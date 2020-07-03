package xdataset

import "github.com/go-xe2/x/xf/ef/xqi"

func (ds *tDataset) fieldByNameRow(name string, row int) *tDsField {
	field := ds.fields.fieldByName(name)
	if field != nil {
		field.rowIndex = row
	}
	return field
}

func (ds *tDataset) fieldByColRow(col, row int) *tDsField {
	field := ds.fields.fieldByIndex(col)
	if field != nil {
		field.rowIndex = row
	}
	return field
}

func (ds *tDataset) rowData(row int) []interface{} {
	if row == -1 {
		row = 0
	}
	if row == ds.RowCount() {
		row = ds.RowCount() - 1
	}
	if row < 0 || row > ds.RowCount()-1 {
		return nil
	}
	return ds.data[row]
}

func (ds *tDataset) colData(col int, row int) interface{} {
	rowValues := ds.rowData(row)
	if rowValues == nil {
		return nil
	}
	if col < 0 || col >= len(rowValues) {
		return nil
	}
	return rowValues[col]
}

func (ds *tDataset) ToMap() []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	ds.Iterator(func(row xqi.DatasetRow) bool {
		item := row.ToMap()
		result = append(result, item)
		return true
	})
	return result
}
