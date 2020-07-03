package xbinder

import . "github.com/go-xe2/x/xf/ef/xqi"

type TDbQuerySliceBinder struct {
	data     [][]interface{}
	colCount int
	// 字段名映射
	fieldMaps map[string]string
	// 字段数据转换
	fieldConvert    map[string]func(in interface{}) (out interface{})
	builder         func(row int, colInfos *[]QueryColValue) interface{}
	fieldConvertBuf []func(old interface{}) interface{}
	fieldNameBuf    []string
}

var _ DbQueryBinder = (*TDbQuerySliceBinder)(nil)

func NewQrySliceBinder(fieldMaps map[string]string, converts ...map[string]func(interface{}) interface{}) DbQueryBinder {
	inst := &TDbQuerySliceBinder{fieldMaps: fieldMaps}
	if converts != nil {
		inst.fieldConvert = converts[0]
	}
	return inst
}

func (bd *TDbQuerySliceBinder) NewInstance(options ...map[string]interface{}) DbQueryBinder {
	if len(options) > 0 && options[0] != nil {
		option := options[0]
		return NewQryMapBinder(nil).SetOptions(option)
	}
	return bd
}

func (bd *TDbQuerySliceBinder) SetOptions(options map[string]interface{}) DbQueryBinder {
	fieldMap := bd.fieldMaps
	converts := bd.fieldConvert
	if v, ok := options["fieldMaps"].(map[string]string); ok {
		fieldMap = v
	}
	if v, ok := options["converts"].(map[string]func(interface{}) interface{}); ok {
		converts = v
	}
	bd.fieldMaps = fieldMap
	bd.fieldConvert = converts
	return bd
}

func (bd *TDbQuerySliceBinder) buildConvertBuf(colInfos ...*TQueryColInfo) {
	colCount := len(colInfos)
	bd.fieldNameBuf = make([]string, colCount)
	bd.fieldConvertBuf = make([]func(old interface{}) interface{}, colCount)
	for i, col := range colInfos {
		if bd.fieldMaps != nil {
			if s, ok := bd.fieldMaps[col.ColName]; ok {
				bd.fieldNameBuf[i] = s
			} else {
				bd.fieldNameBuf[i] = col.ColName
			}
		} else {
			bd.fieldNameBuf[i] = col.ColName
		}
		if bd.fieldConvert != nil {
			if fn, ok := bd.fieldConvert[col.ColName]; ok {
				bd.fieldConvertBuf[i] = fn
			} else {
				bd.fieldConvertBuf[i] = nil
			}
		} else {
			bd.fieldConvertBuf[i] = nil
		}
	}
}

func (bd *TDbQuerySliceBinder) FieldName(colIndex int, qryName string) string {
	if bd.fieldNameBuf == nil {
		return qryName
	}
	return bd.fieldNameBuf[colIndex]
}

func (bd *TDbQuerySliceBinder) FieldConvert(colIndex int, qryName string, val interface{}) interface{} {
	if bd.fieldConvertBuf == nil {
		return val
	}
	if fn := bd.fieldConvertBuf[colIndex]; fn != nil {
		return fn(val)
	}
	return val
}

// 开始绑定，返回false时结束绑定
func (bd *TDbQuerySliceBinder) StartBuild(colInfos ...*TQueryColInfo) bool {
	bd.buildConvertBuf(colInfos...)
	bd.data = make([][]interface{}, 0)
	bd.colCount = len(colInfos)
	return bd.data != nil
}

// 开始创建行,返回false忽略该行
func (bd *TDbQuerySliceBinder) StartBuildRow(rowIndex int, colCount int) bool {
	bd.colCount = colCount
	return true
}

// 创建行
// row:当前行号
// colInfos行中的列信息
// 返回:result行数据结构，exit: 返回true结束绑定
func (bd *TDbQuerySliceBinder) BuildRow(row int, colInfos *[]QueryColValue) (result interface{}, exit bool) {
	rowData := make([]interface{}, bd.colCount)
	for i := 0; i < bd.colCount; i++ {
		col := (*colInfos)[i]
		var v = bd.FieldConvert(col.ColIndex, col.ColName, col.ColValue)
		rowData[i] = v
	}
	return rowData, false
}

// 行创建完成
// rowData为BuildRow所创建的行数据
func (bd *TDbQuerySliceBinder) EndBuildRow(rowData interface{}) {
	if item, ok := rowData.([]interface{}); ok {
		bd.data = append(bd.data, item)
	}
}

// 绑定完成,返回所有行数据
func (bd *TDbQuerySliceBinder) EndBuild() interface{} {
	// 绑定完成后清空之前的配置，以防止之前配置影响绑定结果
	bd.fieldMaps = nil
	bd.fieldConvert = nil
	bd.fieldNameBuf = nil
	bd.fieldConvertBuf = nil
	return bd.data
}
