package xdataset

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xqi"
	"io"
)

type tDatasetFields struct {
	names      []string
	ds         *tDataset
	fieldMap   map[string]*tDsField
	fields     []*tDsField
	fieldCount int
}

var _ xqi.DSFields = (*tDatasetFields)(nil)

func newDataSetFields(ds *tDataset) *tDatasetFields {
	return &tDatasetFields{
		ds:         ds,
		fieldMap:   make(map[string]*tDsField),
		fieldCount: 0,
		fields:     make([]*tDsField, 0),
	}
}

func (dfs *tDatasetFields) Count() int {
	return dfs.fieldCount
}

func (dfs *tDatasetFields) Clear() xqi.DSFields {
	dfs.fieldCount = 0
	dfs.fields = make([]*tDsField, 0)
	dfs.fieldMap = make(map[string]*tDsField)
	dfs.names = nil
	dfs.ds.Clear()
	return dfs
}

func (dfs *tDatasetFields) Add(fieldName string, fieldType xqi.FieldDataType) xqi.DSFields {
	fd := newDsField(dfs.ds, fieldName, dfs.fieldCount, fieldType)
	dfs.fieldCount += 1
	dfs.fieldMap[fieldName] = fd
	dfs.fields = append(dfs.fields, fd)
	dfs.names = nil
	return dfs
}

func (dfs *tDatasetFields) AddDef(fieldDef *xqi.TQueryColDef) xqi.DSFields {
	fd := newDsField(dfs.ds, fieldDef.ColName, dfs.fieldCount, fieldDef.ColType)
	dfs.fieldCount += 1
	dfs.fieldMap[fieldDef.ColName] = fd
	dfs.fields = append(dfs.fields, fd)
	dfs.names = nil
	return dfs
}

func (dfs *tDatasetFields) fieldByIndex(index int) *tDsField {
	return dfs.fields[index]
}

func (dfs *tDatasetFields) Field(index int) xqi.DSField {
	return dfs.fieldByIndex(index)
}

func (dfs *tDatasetFields) fieldByName(name string) *tDsField {
	return dfs.fieldMap[name]
}

func (dfs *tDatasetFields) FieldByName(name string) xqi.DSField {
	return dfs.fieldByName(name)
}

func (dfs *tDatasetFields) Names() []string {
	if dfs.names == nil || len(dfs.names) != dfs.fieldCount {
		dfs.names = make([]string, dfs.fieldCount)
		for i := 0; i < dfs.fieldCount; i++ {
			dfs.names[i] = dfs.Field(i).FieldName()
		}
	}
	return dfs.names
}

func (dfs *tDatasetFields) WriteJson(writer io.Writer) (n int, err error) {
	size := 0
	if n, err := writer.Write([]byte("{ ")); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := writer.Write([]byte(fmt.Sprintf(`"count": %d`, dfs.fieldCount))); err != nil {
		return n, err
	} else {
		size += n
	}
	for i := 0; i < dfs.fieldCount; i++ {
		fd := dfs.fieldByIndex(i)
		if n, err := writer.Write([]byte{','}); err != nil {
			return size, err
		} else {
			size += n
		}
		s := fmt.Sprintf(`"%d":`, i)
		if n, err := writer.Write([]byte(s)); err != nil {
			return size, err
		} else {
			size += n
		}
		if n, err := fd.writeDefineJson(writer); err != nil {
			return size, err
		} else {
			size += n
		}
	}
	if n, err := writer.Write([]byte{' ', '}'}); err != nil {
		return size, err
	} else {
		size += n
	}
	return size, nil
}
