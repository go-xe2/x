package int16Array

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TInt16Slice) Contain(item int16) bool {
	return Contain(iar, item)
}

func (iar TInt16Slice) ForEach(fn Int16ArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TInt16Slice) Find(fn Int16ArrSearchFunc) int16 {
	return Find(iar, fn)
}

func (iar TInt16Slice) FindIndex(fn Int16ArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TInt16Slice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TInt16Slice) Size() int {
	return len(iar)
}

func (iar *TInt16Slice) Append(item ...int16) int {
	*iar = TInt16Slice(append(*iar, item...))
	return iar.Size()
}

func (iar *TInt16Slice) Prepend(item ...int16) int {
	old := *iar
	*iar = append(make([]int16, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TInt16Slice) Insert(index int, item int16) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TInt16Slice) Clear() {
	*iar = make([]int16, 0)
}

func (iar *TInt16Slice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]int16, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TInt16Slice) Concat(arrs ...[]int16) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TInt16Slice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TInt16Slice) String() string {
	return String(iar)
}
