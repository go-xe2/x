package xuint64Slice

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TUint64Slice) Contain(item uint64) bool {
	return Contain(iar, item)
}

func (iar TUint64Slice) ForEach(fn Uint64ArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TUint64Slice) Find(fn Uint64ArrSearchFunc) uint64 {
	return Find(iar, fn)
}

func (iar TUint64Slice) FindIndex(fn Uint64ArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TUint64Slice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TUint64Slice) Size() int {
	return len(iar)
}

func (iar *TUint64Slice) Append(item ...uint64) int {
	*iar = TUint64Slice(append(*iar, item...))
	return iar.Size()
}

func (iar *TUint64Slice) Prepend(item ...uint64) int {
	old := *iar
	*iar = append(make([]uint64, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TUint64Slice) Insert(index int, item uint64) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TUint64Slice) Clear() {
	*iar = make([]uint64, 0)
}

func (iar *TUint64Slice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]uint64, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TUint64Slice) Concat(arrs ...[]uint64) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TUint64Slice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TUint64Slice) String() string {
	return String(iar)
}
