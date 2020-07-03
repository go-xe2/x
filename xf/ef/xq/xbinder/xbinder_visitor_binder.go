package xbinder

import . "github.com/go-xe2/x/xf/ef/xqi"

type TDbQueryVisitorBinder struct {
	data     []interface{}
	colCount int
	visitor  func(row int, values ...interface{}) (interface{}, bool)
}

var _ DbQueryBinder = (*TDbQueryVisitorBinder)(nil)

func VisitorBinder(visitor QueryBinderVisit) DbQueryBinder {
	inst := &TDbQueryVisitorBinder{visitor: visitor}
	return inst
}

func (bd *TDbQueryVisitorBinder) NewInstance(options ...map[string]interface{}) DbQueryBinder {
	if len(options) > 0 && options[0] != nil {
		if v, ok := options[0]["visitor"].(func(row int, values ...interface{}) (interface{}, bool)); ok {
			return VisitorBinder(v)
		}
	}
	return bd
}

func (bd *TDbQueryVisitorBinder) SetOptions(options map[string]interface{}) DbQueryBinder {
	if v, ok := options["visitor"].(func(row int, values ...interface{}) (interface{}, bool)); ok {
		bd.visitor = v
	}
	return bd
}

func (bd *TDbQueryVisitorBinder) FieldName(colIndex int, qryName string) string {
	return qryName
}

func (bd *TDbQueryVisitorBinder) FieldConvert(colIndex int, qryName string, val interface{}) interface{} {
	return val
}

// 开始绑定，返回false时结束绑定
func (bd *TDbQueryVisitorBinder) StartBuild(colInfos ...*TQueryColInfo) bool {
	return true
}

// 开始创建行,返回false忽略该行
func (bd *TDbQueryVisitorBinder) StartBuildRow(rowIndex int, colCount int) bool {
	bd.colCount = colCount
	return true
}

// 创建行
// row:当前行号
// colInfos行中的列信息
// 返回:result行数据结构，exit: 返回true结束绑定
func (bd *TDbQueryVisitorBinder) BuildRow(row int, colInfos *[]QueryColValue) (result interface{}, exit bool) {
	if bd.visitor == nil {
		return nil, true
	}
	values := make([]interface{}, bd.colCount)
	for i := 0; i < bd.colCount; i++ {
		values[i] = (*colInfos)[i].ColValue
	}
	return bd.visitor(row, values...)
}

// 行创建完成
// rowData为BuildRow所创建的行数据
func (bd *TDbQueryVisitorBinder) EndBuildRow(rowData interface{}) {
	if item, ok := rowData.(interface{}); ok {
		bd.data = append(bd.data, item)
	}
}

// 绑定完成,返回所有行数据
func (bd *TDbQueryVisitorBinder) EndBuild() interface{} {
	// 绑定完成后清空之前的配置，以防止之前配置影响绑定结果
	return bd.data
}
