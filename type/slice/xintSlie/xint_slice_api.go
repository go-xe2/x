package xintSlie

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TIntSlice) Contain(item int) bool {
	return Contain(iar, item)
}

func (iar TIntSlice) ForEach(fn IntArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TIntSlice) Find(fn IntArrSearchFunc) int {
	return Find(iar, fn)
}

func (iar TIntSlice) FindIndex(fn IntArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TIntSlice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TIntSlice) Size() int {
	return len(iar)
}

func (iar *TIntSlice) Append(item ...int) int {
	*iar = TIntSlice(append(*iar, item...))
	return iar.Size()
}

func (iar *TIntSlice) Prepend(item ...int) int {
	old := *iar
	*iar = append(make([]int, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TIntSlice) Insert(index int, item int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TIntSlice) Clear() {
	*iar = make([]int, 0)
}

func (iar *TIntSlice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]int, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TIntSlice) Concat(arrs ...[]int) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TIntSlice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TIntSlice) String() string {
	return String(iar)
}
