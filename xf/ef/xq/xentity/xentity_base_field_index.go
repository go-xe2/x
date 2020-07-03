package xentity

import "github.com/go-xe2/x/xf/ef/xqi"

var _ xqi.EntFieldIndex = (*baseField)(nil)

func (ef *baseField) GetIndex() int {
	return ef.index
}

func (ef *baseField) SetIndex(index int) {
	ef.index = index
}
