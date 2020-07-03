package xuint16Slice

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TUint16Slice) Contain(item uint16) bool {
	return Contain(iar, item)
}

func (iar TUint16Slice) ForEach(fn Uint16ArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TUint16Slice) Find(fn Uint16ArrSearchFunc) uint16 {
	return Find(iar, fn)
}

func (iar TUint16Slice) FindIndex(fn Uint16ArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TUint16Slice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TUint16Slice) Size() int {
	return len(iar)
}

func (iar *TUint16Slice) Append(item ...uint16) int {
	*iar = TUint16Slice(append(*iar, item...))
	return iar.Size()
}

func (iar *TUint16Slice) Prepend(item ...uint16) int {
	old := *iar
	*iar = append(make([]uint16, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TUint16Slice) Insert(index int, item uint16) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TUint16Slice) Clear() {
	*iar = make([]uint16, 0)
}

func (iar *TUint16Slice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]uint16, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TUint16Slice) Concat(arrs ...[]uint16) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TUint16Slice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TUint16Slice) String() string {
	return String(iar)
}
