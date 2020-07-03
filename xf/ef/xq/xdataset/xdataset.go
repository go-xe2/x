package xdataset

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xqi"
	"reflect"
)

type tDataset struct {
	fields *tDatasetFields
	data   [][]interface{}
	// 当前游标
	cursor int
}

var _ xqi.MemDataset = (*tDataset)(nil)

// 创建Dataset
func NewMemDataset() xqi.MemDataset {
	inst := &tDataset{
		data:   make([][]interface{}, 0),
		cursor: -1,
	}
	inst.fields = newDataSetFields(inst)
	return inst
}

// 从map数据列表复制数据创建
func DatasetFromSliceMap(maps []map[string]interface{}) xqi.Dataset {
	return NewMemDataset().CreateFromSliceMap(maps)
}

// 从字段定义dataset
func DatasetDefine(fields ...*xqi.TQueryColDef) xqi.Dataset {
	return NewMemDataset().CreateDataSet(fields...)
}

func (ds *tDataset) Fields() xqi.DSFields {
	return ds.fields
}

func (ds *tDataset) IsOpen() bool {
	return ds.data != nil
}

func (ds *tDataset) Close() bool {
	ds.data = nil
	return true
}

func (ds *tDataset) CreateDataSet(fields ...*xqi.TQueryColDef) xqi.Dataset {
	ds.cursor = -1
	for _, field := range fields {
		ds.fields.Add(field.ColName, field.ColType)
	}
	ds.data = make([][]interface{}, 0)
	return ds
}

func (ds *tDataset) CreateFromSliceMap(maps []map[string]interface{}) xqi.Dataset {
	if len(maps) < 0 || maps[0] == nil {
		return ds
	}
	ds.cursor = -1
	cols := maps[0]
	var i = 0
	for k, v := range cols {
		fieldDef := xqi.NewQueryColDefByType(k, reflect.TypeOf(v))
		ds.fields.AddDef(fieldDef)
		i++
	}
	colCount := ds.fields.Count()
	// 复制数据
	rowCount := len(maps)
	ds.data = make([][]interface{}, rowCount)
	for i := 0; i < rowCount; i++ {
		rowData := make([]interface{}, colCount)
		for j := 0; j < colCount; j++ {
			key := ds.fields.Field(j).FieldName()
			val := maps[i][key]
			rowData[j] = val
		}
		ds.data[i] = rowData
	}
	return ds
}

func (ds *tDataset) FieldCount() int {
	return ds.fields.Count()
}

// 获取字段
func (ds *tDataset) DSField(index int) xqi.DSField {
	return ds.fieldByColRow(index, ds.cursor)
}

func (ds *tDataset) DSFieldByName(name string) xqi.DSField {
	return ds.fieldByNameRow(name, ds.cursor)
}

// 字段值
func (ds *tDataset) FieldValue(index int) interface{} {
	col := ds.fieldByColRow(index, ds.cursor)
	if col == nil {
		return nil
	}
	return col.Value()
}

// 添加数据
func (ds *tDataset) Append(values ...interface{}) xqi.Dataset {
	if len(values) != ds.FieldCount() {
		panic(exception.Newf("需要%d列数据实例传入%d列数据", ds.FieldCount(), len(values)))
	}
	ds.data = append(ds.data, values)
	ds.cursor = 0
	return ds
}

func (ds *tDataset) Row(rowIndex int) xqi.DatasetRow {
	return newDatasetRow(ds, rowIndex)
}

// 遍历所有记录
func (ds *tDataset) Iterator(iter func(row xqi.DatasetRow) bool) {
	dsRow := newDatasetRow(ds, 0)
	for ds.cursor = 0; ds.cursor < ds.RowCount(); ds.cursor++ {
		dsRow.rowIndex = ds.cursor
		if !iter(dsRow) {
			break
		}
	}
}

func (ds *tDataset) Current() xqi.DatasetRow {
	if ds.RowCount() == 0 {
		return nil
	}
	if ds.cursor < 0 {
		ds.cursor = 0
	} else if ds.cursor >= ds.RowCount() {
		ds.cursor = ds.RowCount() - 1
	}
	return newDatasetRow(ds, ds.cursor)
}

func (ds *tDataset) DeleteByRow(row int) bool {
	if row >= 0 && row < ds.RowCount() {
		// 删除
		if row == 0 {
			ds.data = ds.data[1:]
			if ds.cursor > 0 {
				ds.cursor = ds.cursor - 1
			}
		} else if row == ds.RowCount()-1 {
			ds.data = ds.data[:row]
			if ds.cursor == row {
				// 游标被删除行时，光标向前移动
				ds.cursor = ds.cursor - 1
			}
		} else {
			tmp := ds.data[:row]
			ds.data = append(tmp, ds.data[row+1:]...)
			if ds.cursor > row {
				ds.cursor = ds.cursor - 1
			}
		}
		return true
	}
	return false
}

// 删除
func (ds *tDataset) Delete() bool {
	return ds.DeleteByRow(ds.cursor)
}

func (ds *tDataset) RowCount() int {
	return len(ds.data)
}

func (ds *tDataset) MoveFirst() bool {
	if ds.RowCount() > 0 {
		ds.cursor = -1
		return true
	}
	return false
}

func (ds *tDataset) MoveLast() bool {
	if ds.RowCount() > 0 {
		ds.cursor = ds.RowCount()
		return true
	}
	return false
}

func (ds *tDataset) Next() bool {
	if ds.cursor < ds.RowCount()-1 && ds.RowCount() > 0 {
		if ds.cursor < -1 {
			ds.cursor = -1
		}
		ds.cursor++
		return true
	}
	return false
}

func (ds *tDataset) Prior() bool {
	if ds.cursor > 0 && ds.RowCount() > 0 {
		if ds.cursor > ds.RowCount() {
			ds.cursor = ds.RowCount()
		}
		ds.cursor--
		return true
	}
	return false
}

func (ds *tDataset) IsEmpty() bool {
	return ds.RowCount() == 0
}

func (ds *tDataset) IsEof() bool {
	return ds.cursor >= ds.RowCount()
}

func (ds *tDataset) IsBof() bool {
	return ds.cursor < 0
}

func (ds *tDataset) Clear() xqi.Dataset {
	ds.data = make([][]interface{}, 0)
	return ds
}
