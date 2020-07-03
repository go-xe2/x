package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xqi"
)

var _ xqi.Dataset = (*TEntity)(nil)

func (ent *TEntity) FieldValue(index int) interface{} {
	if !ent.IsOpen() {
		return nil
	}
	return ent.dataset.FieldValue(index)
}

func (ent *TEntity) IsEmpty() bool {
	if !ent.IsOpen() {
		return true
	}
	return ent.dataset.IsEmpty()
}

func (ent *TEntity) IsBof() bool {
	if !ent.IsOpen() {
		return true
	}
	return ent.dataset.IsBof()
}

func (ent *TEntity) IsEof() bool {
	if !ent.IsOpen() {
		return true
	}
	return ent.dataset.IsEof()
}

func (ent *TEntity) IsOpen() bool {
	return ent.dataset != nil
}

func (ent *TEntity) MoveFirst() bool {
	if !ent.IsOpen() {
		panic(exception.NewText("未打开数据集"))
	}
	return ent.dataset.MoveFirst()
}

func (ent *TEntity) MoveLast() bool {
	if !ent.IsOpen() {
		panic(exception.NewText("未打开数据集"))
	}
	return ent.dataset.MoveLast()
}

func (ent *TEntity) Close() bool {
	ent.dataset = nil
	return true
}

func (ent *TEntity) Next() bool {
	if !ent.IsOpen() {
		panic(exception.NewText("未打开数据集"))
	}
	return ent.dataset.Next()
}

func (ent *TEntity) Prior() bool {
	if !ent.IsOpen() {
		panic(exception.NewText("未打开数据集"))
	}
	return ent.dataset.Prior()
}

func (ent *TEntity) RowCount() int {
	if !ent.IsOpen() {
		return 0
	}
	return ent.dataset.RowCount()
}

func (ent *TEntity) ToMap() []map[string]interface{} {
	if !ent.IsOpen() {
		return []map[string]interface{}{}
	}
	return ent.dataset.ToMap()
}

func (ent *TEntity) MarshalJSON() (data []byte, err error) {
	if !ent.IsOpen() {
		return []byte{}, nil
	}
	return ent.dataset.MarshalJSON()
}

func (ent *TEntity) UnmarshalJSON(data []byte) error {
	if !ent.IsOpen() {
		return nil
	}
	return ent.dataset.UnmarshalJSON(data)
}

func (ent *TEntity) DSFieldByName(name string) xqi.DSField {
	if fd := ent.FieldByName(name); fd != nil {
		return fd.(xqi.DSField)
	}
	return nil
}

func (ent *TEntity) DSField(index int) xqi.DSField {
	return ent.fields[index].(xqi.DSField)
}
