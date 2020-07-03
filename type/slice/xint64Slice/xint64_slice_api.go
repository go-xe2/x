package int64Array

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TInt64Slice) Contain(item int64) bool {
	return Contain(iar, item)
}

func (iar TInt64Slice) ForEach(fn Int64ArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TInt64Slice) Find(fn Int64ArrSearchFunc) int64 {
	return Find(iar, fn)
}

func (iar TInt64Slice) FindIndex(fn Int64ArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TInt64Slice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TInt64Slice) Size() int {
	return len(iar)
}

func (iar *TInt64Slice) Append(item ...int64) int {
	*iar = TInt64Slice(append(*iar, item...))
	return iar.Size()
}

func (iar *TInt64Slice) Prepend(item ...int64) int {
	old := *iar
	*iar = append(make([]int64, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TInt64Slice) Insert(index int, item int64) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TInt64Slice) Clear() {
	*iar = make([]int64, 0)
}

func (iar *TInt64Slice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]int64, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TInt64Slice) Concat(arrs ...[]int64) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TInt64Slice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TInt64Slice) String() string {
	return String(iar)
}
