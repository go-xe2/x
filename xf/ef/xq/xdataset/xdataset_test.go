package xdataset

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xqi"
	"testing"
)

func TestNewDataset(t *testing.T) {
	ds := NewMemDataset()
	ds.CreateDataSet(xqi.NewQueryColDef("name", xqi.FDTString),
		xqi.NewQueryColDef("sex", xqi.FDTBool),
		xqi.NewQueryColDef("age", xqi.FDTInt))
	ds.Append("name1", true, 32)
	ds.Append("name2", false, 33)
	ds.Append("name3", true, 34)
	ds.Append("name4", true, 35)
	ds.Append("name5", true, 36)
	ds.Append(nil, true, 36)

	ds.Iterator(func(row xqi.DatasetRow) bool {
		fmt.Printf("row(%d):field1:%v,field2:%v,field3:%v\n", row.RowIndex(), row.Field(0).Value(), row.Field(1).Value(), row.Field(2).Value())
		return true
	})
	fmt.Println("==========>> dataset.next")
	ds.MoveFirst()
	for ds.Next() {
		row := ds.Current()
		fmt.Printf("row(%d):field1:%v,field2:%v,field3:%v\n", row.RowIndex(), row.Field(0).Value(), row.Field(1).Value(), row.Field(2).Value())
	}

	fmt.Println("==========>> dataset.prior")
	ds.MoveLast()
	for ds.Prior() {
		row := ds.Current()
		fmt.Printf("row(%d):field1:%v,field2:%v,field3:%v\n", row.RowIndex(), row.Field(0).Value(), row.Field(1).Value(), row.Field(2).Value())
	}

	mps := ds.ToMap()
	fmt.Println("row map:", mps)

	bJson, err := ds.MarshalJSON()
	fmt.Println("json:", string(bJson), ", err:", err)

	fmt.Println("=============> unmarshal:")
	ds2 := NewMemDataset()
	err = ds2.UnmarshalJSON(bJson)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("unmarshalJson ds2:", ds2.ToMap())
	fmt.Println("============= 获取字段值开始:==============")
	ds.MoveFirst()
	for ds.Next() {
		fmt.Println("ds.field0:", ds.FieldValue(0), "ds.field1:", ds.DSField(1).Value())
	}

}
