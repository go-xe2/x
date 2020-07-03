package xentity

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/slice/xstrSlice"
	"github.com/go-xe2/x/utils/xvalid"
	"github.com/go-xe2/x/xf/ef/xqi"
	"reflect"
)

type foreignItem struct {
	table          xqi.SqlTable
	joinType       xqi.SqlJoinType
	relationTables []xqi.SqlTable
	onCondition    func(joinTable xqi.SqlTable, otherTables xqi.SqlTables, on xqi.SqlCondition) xqi.SqlCondition
}

func (fi *foreignItem) String() string {
	s := ""
	for _, table := range fi.relationTables {
		if s != "" {
			s += ","
		}
		s += table.TableName()
	}
	return fmt.Sprintf("{item(%s %s),relation(%s)}", fi.joinType, fi.table.TableName(), s)
}

type tEntityEditState int

const (
	// 空闲状态
	entityFreeState tEntityEditState = iota
	// 插入状态
	entityInsertState
	// 更新状态
	entityUpdateState
)

type entityFieldValidRule struct {
	field xqi.EntField
	valid xqi.FieldValid
}

type fieldFormatterStore struct {
	options   map[string]interface{}
	formatter func(old interface{}, options ...map[string]interface{}) interface{}
}

func newfieldFormatterStore(options map[string]interface{}, formatter func(old interface{}, options ...map[string]interface{}) interface{}) *fieldFormatterStore {
	return &fieldFormatterStore{
		options:   options,
		formatter: formatter,
	}
}

type TEntity struct {
	*baseEntity
	lastSql  string
	keyField xqi.EntField
	// 是否已经创建实体之间的关系
	// 外联实体关系字典，如果为空表达未创建实体关系
	joinConditions map[string]*foreignItem
	// 字段所关联到的实体
	fieldConditions map[xqi.SqlField][]*foreignItem
	dataset         xqi.Dataset
	// 当前数据游标
	dataCursor int
	// 查询返回的数据
	dataRows [][]interface{}
	// 实体当前编译状态
	editState tEntityEditState
	// 当前编译的数据
	editFieldValues []xqi.FieldValue
	// 主键
	editKeyValue xqi.FieldValue
	// 当前实体使用的数据库
	dbName []string

	updateFieldRuleIndex map[string]int
	insertFieldRuleIndex map[string]int
	// 更新时要使用的字段规则过滤,表字段名 => 规则
	updateFieldValidRule []*entityFieldValidRule
	// 表字段名 => 规则
	insertFieldValidRule []*entityFieldValidRule
	// 字段数据格式化映射词典
	fieldFormatter map[string]func(old interface{}) interface{}
}

var _ xqi.Entity = (*TEntity)(nil)

var entityType = reflect.TypeOf((*TEntity)(nil)).Elem()

func newEntity(inherited interface{}, tableName string, alias ...string) *TEntity {
	inst := &TEntity{
		baseEntity: nil,
		keyField:   nil,
		// 此处不能初始化，用于检查是否已建立关联关系
		//joinConditions:  make(map[string]*foreignItem),
		fieldConditions:      make(map[xqi.SqlField][]*foreignItem),
		updateFieldValidRule: make([]*entityFieldValidRule, 0),
		insertFieldValidRule: make([]*entityFieldValidRule, 0),
		updateFieldRuleIndex: make(map[string]int),
		insertFieldRuleIndex: make(map[string]int),
		fieldFormatter:       make(map[string]func(old interface{}) interface{}),
	}
	var this interface{} = inst
	if _, ok := inherited.(xqi.SqlTable); ok {
		this = inherited
	}
	base := newBaseEntity(this, tableName, alias...)
	inst.baseEntity = base
	return inst
}

func (ent *TEntity) Supper() xqi.Entity {
	return nil
}

func (ent *TEntity) Implement(supper interface{}) {
	if v, ok := supper.(*baseEntity); ok {
		ent.baseEntity = v
	}
}

func (ent *TEntity) Constructor(attrs []xqi.XqAttribute, inherited ...interface{}) interface{} {
	ent.baseEntity.Constructor(attrs, inherited...)
	return ent
}

func (ent *TEntity) isBuildRelation() bool {
	return ent.joinConditions != nil && ent.fieldConditions != nil
}

func (ent *TEntity) String() string {
	return fmt.Sprintf("entity{ baseEntity: %v }", ent.baseEntity)
}

func (ent *TEntity) Database(dbName string) xqi.Entity {
	ent.dbName = []string{dbName}
	return ent
}

func (ent *TEntity) FieldFormat(defineName string) func(old interface{}) interface{} {
	if fn, ok := ent.fieldFormatter[defineName]; ok {
		return fn
	}
	return nil
}

func (ent *TEntity) GetUpdateRules(cate ...string) []string {
	if ruleInf, ok := ent.This().(xqi.EntityUpdateRule); ok {
		return ruleInf.UpdateRule(cate...)
	}
	result := make([]string, 0)
	szCate := ""
	if len(cate) > 0 {
		szCate = cate[0]
	}
	for _, row := range ent.updateFieldValidRule {
		cates := row.valid.Cate()
		if (len(cates) == 0 || xstrSlice.Contain(cates, szCate)) && (row.valid.Rule() != "") {
			result = append(result, row.valid.MakeValidString(row.field.FieldName()))
		}
	}
	return result
}

func (ent *TEntity) GetInsertRules(cate ...string) []string {
	if ruleInf, ok := ent.This().(xqi.EntityInsertRule); ok {
		return ruleInf.InsertRule(cate...)
	}
	result := make([]string, 0)
	szCate := ""
	if len(cate) > 0 {
		szCate = cate[0]
	}
	for _, row := range ent.insertFieldValidRule {
		cates := row.valid.Cate()
		if (len(cates) == 0 || xstrSlice.Contain(cates, szCate)) && (row.valid.Rule() != "") {
			result = append(result, row.valid.MakeValidString(row.field.FieldName()))
		}
	}
	return result
}

// 使用定义更新规则过滤参数
// @param validParams 可选参数，默认为false, 为true时，检查数据有效性
func (ent *TEntity) FilterUpdateParams(params map[string]interface{}, validParams ...bool) (map[string]interface{}, error) {
	var result = make(map[string]interface{})
	bValid := false
	if len(validParams) > 0 {
		bValid = validParams[0]
	}
	for k, v := range params {
		if validIdx, ok := ent.updateFieldRuleIndex[k]; ok {
			vf := ent.updateFieldValidRule[validIdx]
			if bValid && vf.valid.Rule() != "" {
				if errs := xvalid.Check(v, vf.valid.Rule(), vf.valid.Msg()); errs != nil {
					return params, exception.NewText(errs.FirstString())
				}
			}
			result[k] = v
		}
	}
	return result, nil
}

// 使用定义插入规则过滤参数
// @param validParams 可选参数，默认为false, 为true时，检查数据有效性
func (ent *TEntity) FilterInsertParams(params map[string]interface{}, validParams ...bool) (map[string]interface{}, error) {
	var result = make(map[string]interface{})
	bValid := false
	if len(validParams) > 0 {
		bValid = validParams[0]
	}
	for k, v := range params {
		if validIdx, ok := ent.insertFieldRuleIndex[k]; ok {
			vf := ent.insertFieldValidRule[validIdx]
			if bValid && vf.valid.Rule() != "" {
				if errs := xvalid.Check(v, vf.valid.Rule(), vf.valid.Msg()); errs != nil {
					return params, exception.NewText(errs.FirstString())
				}
			}
			result[k] = v
		}
	}
	return result, nil
}

func (ent *TEntity) GetInsertValuesFromParams(params map[string]interface{}, validParams ...bool) ([]xqi.FieldValue, error) {
	var result = make([]xqi.FieldValue, 0)
	bValid := false
	if len(validParams) > 0 {
		bValid = validParams[0]
	}
	for k, v := range params {
		if validIdx, ok := ent.insertFieldRuleIndex[k]; ok {
			vf := ent.insertFieldValidRule[validIdx]
			if bValid && vf.valid.Rule() != "" {
				if errs := xvalid.Check(v, vf.valid.Rule(), vf.valid.Msg()); errs != nil {
					return nil, exception.NewText(errs.FirstString())
				}
			}
			result = append(result, vf.field.Set(v))
		}
	}
	return result, nil
}

func (ent *TEntity) GetUpdateValuesFromParams(params map[string]interface{}, validParams ...bool) ([]xqi.FieldValue, error) {
	var result = make([]xqi.FieldValue, 0)
	bValid := false
	if len(validParams) > 0 {
		bValid = validParams[0]
	}
	for k, v := range params {
		if validIdx, ok := ent.updateFieldRuleIndex[k]; ok {
			vf := ent.updateFieldValidRule[validIdx]
			if bValid && vf.valid.Rule() != "" {
				if errs := xvalid.Check(v, vf.valid.Rule(), vf.valid.Msg()); errs != nil {
					return nil, exception.NewText(errs.FirstString())
				}
			}
			result = append(result, vf.field.Set(v))
		}
	}
	return result, nil
}
