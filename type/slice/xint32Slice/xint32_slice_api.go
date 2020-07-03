package xint32Slice

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TInt32Slice) Contain(item int32) bool {
	return Contain(iar, item)
}

func (iar TInt32Slice) ForEach(fn Int32ArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TInt32Slice) Find(fn Int32ArrSearchFunc) int32 {
	return Find(iar, fn)
}

func (iar TInt32Slice) FindIndex(fn Int32ArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TInt32Slice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TInt32Slice) Size() int {
	return len(iar)
}

func (iar *TInt32Slice) Append(item ...int32) int {
	*iar = TInt32Slice(append(*iar, item...))
	return iar.Size()
}

func (iar *TInt32Slice) Prepend(item ...int32) int {
	old := *iar
	*iar = append(make([]int32, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TInt32Slice) Insert(index int, item int32) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TInt32Slice) Clear() {
	*iar = make([]int32, 0)
}

func (iar *TInt32Slice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]int32, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TInt32Slice) Concat(arrs ...[]int32) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TInt32Slice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TInt32Slice) String() string {
	return String(iar)
}
