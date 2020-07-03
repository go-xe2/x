package xfloat64Slice

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TFloat64Slice) Contain(item float64) bool {
	return Contain(iar, item)
}

func (iar TFloat64Slice) ForEach(fn Float64ArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TFloat64Slice) Find(fn Float64ArrSearchFunc) float64 {
	return Find(iar, fn)
}

func (iar TFloat64Slice) FindIndex(fn Float64ArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TFloat64Slice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TFloat64Slice) Size() int {
	return len(iar)
}

func (iar *TFloat64Slice) Append(item ...float64) int {
	*iar = TFloat64Slice(append(*iar, item...))
	return iar.Size()
}

func (iar *TFloat64Slice) Prepend(item ...float64) int {
	old := *iar
	*iar = append(make([]float64, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TFloat64Slice) Insert(index int, item float64) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TFloat64Slice) Clear() {
	*iar = make([]float64, 0)
}

func (iar *TFloat64Slice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]float64, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TFloat64Slice) Concat(arrs ...[]float64) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TFloat64Slice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TFloat64Slice) String() string {
	return String(iar)
}
