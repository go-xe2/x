package xuint8Slice

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TUint8Slice) Contain(item uint8) bool {
	return Contain(iar, item)
}

func (iar TUint8Slice) ForEach(fn Uint8ArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TUint8Slice) Find(fn Uint8ArrSearchFunc) uint8 {
	return Find(iar, fn)
}

func (iar TUint8Slice) FindIndex(fn Uint8ArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TUint8Slice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TUint8Slice) Size() int {
	return len(iar)
}

func (iar *TUint8Slice) Append(item ...uint8) int {
	*iar = TUint8Slice(append(*iar, item...))
	return iar.Size()
}

func (iar *TUint8Slice) Prepend(item ...uint8) int {
	old := *iar
	*iar = append(make([]uint8, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TUint8Slice) Insert(index int, item uint8) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TUint8Slice) Clear() {
	*iar = make([]uint8, 0)
}

func (iar *TUint8Slice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]uint8, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TUint8Slice) Concat(arrs ...[]uint8) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TUint8Slice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TUint8Slice) String() string {
	return String(iar)
}
