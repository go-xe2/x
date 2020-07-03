package xdataset

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xqi"
	"io"
)

type tDatasetRow struct {
	ds       *tDataset
	rowIndex int
}

var _ xqi.DatasetRow = (*tDatasetRow)(nil)

func newDatasetRow(ds *tDataset, rowIndex int) *tDatasetRow {
	return &tDatasetRow{
		ds:       ds,
		rowIndex: rowIndex,
	}
}

func (dsr *tDatasetRow) RowIndex() int {
	return dsr.ds.cursor
}

func (dsr *tDatasetRow) Field(index int) xqi.DSField {
	return dsr.fieldByIndex(index)
}

func (dsr *tDatasetRow) fieldByIndex(index int) *tDsField {
	return dsr.ds.fieldByColRow(index, dsr.rowIndex)
}

func (dsr *tDatasetRow) fieldByName(name string) *tDsField {
	return dsr.ds.fieldByNameRow(name, dsr.rowIndex)
}

func (dsr *tDatasetRow) FieldByName(name string) xqi.DSField {
	return dsr.fieldByName(name)
}

func (dsr *tDatasetRow) Values() []interface{} {
	return dsr.ds.rowData(dsr.rowIndex)
}

func (dsr *tDatasetRow) Names() []string {
	return dsr.ds.fields.Names()
}

func (dsr *tDatasetRow) Delete() {
	dsr.ds.DeleteByRow(dsr.rowIndex)
}

func (dsr *tDatasetRow) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	for i := 0; i < dsr.ds.FieldCount(); i++ {
		result[dsr.Field(i).FieldName()] = dsr.Field(i).Value()
	}
	return result
}

func (dsr *tDatasetRow) WriteJson(writer io.Writer) (n int, err error) {
	size := 0
	if n, err := writer.Write([]byte("{ ")); err != nil {
		return size, err
	} else {
		size += n
	}
	isEmpty := true
	for i := 0; i < dsr.ds.FieldCount(); i++ {
		fd := dsr.fieldByIndex(i)
		if fd.Value() == nil {
			continue
		}
		if !isEmpty {
			if n, err := writer.Write([]byte{','}); err != nil {
				return size, err
			} else {
				size += n
			}
		}
		isEmpty = false
		if n, err := writer.Write([]byte(fmt.Sprintf(`"%d":`, i))); err != nil {
			return size, err
		} else {
			size += n
		}
		if n, err := fd.writeDataJson(writer); err != nil {
			return size, err
		} else {
			size += n
		}
	}
	if n, err := writer.Write([]byte(" }")); err != nil {
		return size, err
	} else {
		size += n
	}
	return size, nil
}
