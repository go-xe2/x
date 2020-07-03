package xuintSlice

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TUintArray) Contain(item uint) bool {
	return Contain(iar, item)
}

func (iar TUintArray) ForEach(fn UintArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TUintArray) Find(fn UintArrSearchFunc) uint {
	return Find(iar, fn)
}

func (iar TUintArray) FindIndex(fn UintArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TUintArray) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TUintArray) Size() int {
	return len(iar)
}

func (iar *TUintArray) Append(item ...uint) int {
	*iar = TUintArray(append(*iar, item...))
	return iar.Size()
}

func (iar *TUintArray) Prepend(item ...uint) int {
	old := *iar
	*iar = append(make([]uint, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TUintArray) Insert(index int, item uint) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TUintArray) Clear() {
	*iar = make([]uint, 0)
}

func (iar *TUintArray) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]uint, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TUintArray) Concat(arrs ...[]uint) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TUintArray) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TUintArray) String() string {
	return String(iar)
}
