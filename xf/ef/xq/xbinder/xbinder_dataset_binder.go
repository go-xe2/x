package xbinder

import (
	"github.com/go-xe2/x/xf/ef/xq/xdataset"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TDbQueryDatasetBinder struct {
	dataset  MemDataset
	colCount int
	// 字段名映射
	fieldMaps map[string]string
	// 字段数据转换
	fieldConvert    map[string]func(in interface{}) (out interface{})
	builder         func(row int, colInfos *[]QueryColValue) interface{}
	fieldNameBuf    []string
	fieldConvertBuf []func(old interface{}) interface{}
}

var _ DbQueryBinder = (*TDbQueryDatasetBinder)(nil)

func NewQryDatasetBinder(fieldMaps map[string]string, converts ...map[string]func(interface{}) interface{}) DbQueryBinder {
	inst := &TDbQueryDatasetBinder{fieldMaps: fieldMaps}
	if converts != nil {
		inst.fieldConvert = converts[0]
	}
	return inst
}

func (bd *TDbQueryDatasetBinder) NewInstance(options ...map[string]interface{}) DbQueryBinder {
	if len(options) > 0 && options[0] != nil {
		option := options[0]
		return NewQryDatasetBinder(nil).SetOptions(option)
	}
	return bd
}

func (bd *TDbQueryDatasetBinder) SetOptions(options map[string]interface{}) DbQueryBinder {
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

func (bd *TDbQueryDatasetBinder) FieldName(colIndex int, qryName string) string {
	if bd.fieldNameBuf == nil {
		return qryName
	}
	return bd.fieldNameBuf[colIndex]
}

func (bd *TDbQueryDatasetBinder) FieldConvert(colIndex int, qryName string, val interface{}) interface{} {
	if bd.fieldConvertBuf == nil {
		return val
	}
	if fn := bd.fieldConvertBuf[colIndex]; fn != nil {
		return fn(val)
	}
	return val
}

func (bd *TDbQueryDatasetBinder) buildConvertBuf(colInfos ...*TQueryColInfo) {
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

// 开始绑定，返回false时结束绑定
func (bd *TDbQueryDatasetBinder) StartBuild(colInfos ...*TQueryColInfo) bool {
	bd.dataset = xdataset.NewMemDataset()
	bd.colCount = len(colInfos)
	bd.buildConvertBuf(colInfos...)
	fieldDefs := make([]*TQueryColDef, bd.colCount)
	for i := 0; i < bd.colCount; i++ {
		fieldDefs[i] = NewQueryColDef(bd.FieldName(colInfos[i].ColIndex, colInfos[i].ColName), colInfos[i].ColType)
	}
	bd.dataset.CreateDataSet(fieldDefs...)
	return bd.dataset != nil
}

// 开始创建行,返回false忽略该行
func (bd *TDbQueryDatasetBinder) StartBuildRow(rowIndex int, colCount int) bool {
	return true
}

// 创建行
// row:当前行号
// colInfos行中的列信息
// 返回:result行数据结构，exit: 返回true结束绑定
func (bd *TDbQueryDatasetBinder) BuildRow(row int, colInfos *[]QueryColValue) (result interface{}, exit bool) {
	rowData := make([]interface{}, bd.colCount)
	for i := 0; i < bd.colCount; i++ {
		col := (*colInfos)[i]
		rowData[i] = bd.FieldConvert(col.ColIndex, col.ColName, col.ColValue)
	}
	return rowData, false
}

// 行创建完成
// rowData为BuildRow所创建的行数据
func (bd *TDbQueryDatasetBinder) EndBuildRow(rowData interface{}) {
	if items, ok := rowData.([]interface{}); ok {
		bd.dataset.Append(items...)
	}
}

// 绑定完成,返回所有行数据
func (bd *TDbQueryDatasetBinder) EndBuild() interface{} {
	// 绑定完成后清空之前的配置，以防止之前配置影响绑定结果
	bd.fieldMaps = nil
	bd.fieldConvert = nil
	bd.fieldNameBuf = nil
	bd.fieldConvertBuf = nil
	return bd.dataset
}
