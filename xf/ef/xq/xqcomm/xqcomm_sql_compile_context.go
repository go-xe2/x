package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

// 参数化查询实参收集器
type TSqlCompileContext struct {
	items []xdriveri.SqlParam
	// 编译状态
	states      SCPState
	stateStack  *tSqlStateStack
	lTokenStack *tSqlStateStack
	maxParamId  int
	maxIndentId int
	tables      SqlTables
	database    Database
}

var _ SqlCompileContext = &TSqlCompileContext{}

func NewSqlCompileContext() *TSqlCompileContext {
	return &TSqlCompileContext{
		items:       make([]xdriveri.SqlParam, 0),
		states:      0,
		stateStack:  newSqlStateStack(),
		lTokenStack: newSqlStateStack(),
		maxParamId:  0,
		maxIndentId: 0,
		tables:      NewSqlTables(),
	}
}

func (pms *TSqlCompileContext) SetDatabase(db Database) {
	pms.database = db
}

func (pms *TSqlCompileContext) Database() Database {
	return pms.database
}

func (pms *TSqlCompileContext) TablePrefix() string {
	if pms.database == nil {
		return ""
	}
	return pms.database.Prefix()
}

func (pms *TSqlCompileContext) Driver() string {
	if pms.database == nil {
		return "mysql"
	}
	return pms.database.Driver()
}

func (pms *TSqlCompileContext) MakeParamId() string {
	pms.maxParamId++
	return fmt.Sprintf("P%d", pms.maxParamId)
}

func (pms *TSqlCompileContext) MakeIndentId() string {
	pms.maxIndentId++
	return fmt.Sprintf("T%d", pms.maxIndentId)
}

func (pms *TSqlCompileContext) AddParam(name string, val interface{}) xdriveri.SqlParam {
	result := NewSqlParam(name, val)
	pms.Add(result)
	return result
}

// 添加参数
func (pms *TSqlCompileContext) Add(val ...xdriveri.SqlParam) SqlCompileContext {
	pms.items = append(pms.items, val...)
	return pms
}

// 所有已被收到的实参
func (pms *TSqlCompileContext) Params() []xdriveri.SqlParam {
	return pms.items
}

// 清空实参
func (pms *TSqlCompileContext) ClearParams() SqlCompileContext {
	pms.items = make([]xdriveri.SqlParam, 0)
	return pms
}

func (pms *TSqlCompileContext) Clear() SqlCompileContext {
	pms.ClearParams()
	pms.tables.Clear()
	pms.stateStack.Clear()
	pms.lTokenStack.Clear()
	pms.maxIndentId = 0
	pms.maxParamId = 0
	return pms
}

// 拷贝实参
func (pms *TSqlCompileContext) Copy(src SqlCompileContext) SqlCompileContext {
	srcItems := src.Params()
	pms.items = append(pms.items, srcItems...)
	return pms
}

// 参数个数
func (pms *TSqlCompileContext) ParamCount() int {
	return len(pms.items)
}

func (pms *TSqlCompileContext) ParamByIndex(index int) xdriveri.SqlParam {
	for k, v := range pms.items {
		if k == index {
			return v
		}
	}
	return nil
}

func (pms *TSqlCompileContext) DeleteParam(index int) xdriveri.SqlParam {
	if index == 0 {
		value := pms.items[0]
		pms.items = pms.items[1:]
		return value
	} else if index == len(pms.items)-1 {
		value := pms.items[index]
		pms.items = pms.items[:index]
		return value
	}
	value := pms.items[index]
	pms.items = append(pms.items[:index], pms.items[index+1:]...)
	return value
}

func (pms *TSqlCompileContext) DeleteParamByName(name string) xdriveri.SqlParam {
	for k, v := range pms.items {
		if v.Name() == name {
			return pms.DeleteParam(k)
		}
	}
	return nil
}

func (pms *TSqlCompileContext) String() string {
	return fmt.Sprintf("%v", pms.items)
}

func (pms *TSqlCompileContext) PushState(state SCPState) SqlCompileContext {
	pms.stateStack.push(state)
	return pms
}

func (pms *TSqlCompileContext) PopState() SCPState {
	if n, ok := pms.stateStack.pop().(SCPState); ok {
		return n
	}
	return SCPUnknown
}

func (pms *TSqlCompileContext) State() SCPState {
	if n, ok := pms.stateStack.peek().(SCPState); ok {
		return n
	}
	return SCPUnknown
}

func (pms *TSqlCompileContext) PushLExp(exp SqlCompiler) SqlCompileContext {
	pms.lTokenStack.push(exp)
	return pms
}

func (pms *TSqlCompileContext) PopLExp() SqlCompiler {
	v := pms.lTokenStack.pop()
	if tk, ok := v.(SqlCompiler); ok {
		return tk
	}
	return nil
}

func (pms *TSqlCompileContext) LExp() SqlCompiler {
	v := pms.lTokenStack.peek()
	if tk, ok := v.(SqlCompiler); ok {
		return tk
	}
	return nil
}

func (pms *TSqlCompileContext) AssignTables(src SqlTables) {
	pms.tables = src
}

func (pms *TSqlCompileContext) UseTable(table SqlTable) {
	pms.tables.Add(table)
}

func (pms *TSqlCompileContext) Tables() SqlTables {
	return pms.tables
}
