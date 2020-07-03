package int8Array

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TInt8Slice) Contain(item int8) bool {
	return Contain(iar, item)
}

func (iar TInt8Slice) ForEach(fn Int8ArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TInt8Slice) Find(fn Int8ArrSearchFunc) int8 {
	return Find(iar, fn)
}

func (iar TInt8Slice) FindIndex(fn Int8ArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TInt8Slice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TInt8Slice) Size() int {
	return len(iar)
}

func (iar *TInt8Slice) Append(item ...int8) int {
	*iar = TInt8Slice(append(*iar, item...))
	return iar.Size()
}

func (iar *TInt8Slice) Prepend(item ...int8) int {
	old := *iar
	*iar = append(make([]int8, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TInt8Slice) Insert(index int, item int8) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TInt8Slice) Clear() {
	*iar = make([]int8, 0)
}

func (iar *TInt8Slice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]int8, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TInt8Slice) Concat(arrs ...[]int8) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TInt8Slice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TInt8Slice) String() string {
	return String(iar)
}
