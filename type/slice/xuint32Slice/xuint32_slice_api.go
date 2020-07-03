package xuint32Slice

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TUint32Array) Contain(item uint32) bool {
	return Contain(iar, item)
}

func (iar TUint32Array) ForEach(fn Uint32ArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TUint32Array) Find(fn Uint32ArrSearchFunc) uint32 {
	return Find(iar, fn)
}

func (iar TUint32Array) FindIndex(fn Uint32ArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TUint32Array) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TUint32Array) Size() int {
	return len(iar)
}

func (iar *TUint32Array) Append(item ...uint32) int {
	*iar = TUint32Array(append(*iar, item...))
	return iar.Size()
}

func (iar *TUint32Array) Prepend(item ...uint32) int {
	old := *iar
	*iar = append(make([]uint32, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TUint32Array) Insert(index int, item uint32) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TUint32Array) Clear() {
	*iar = make([]uint32, 0)
}

func (iar *TUint32Array) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]uint32, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TUint32Array) Concat(arrs ...[]uint32) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TUint32Array) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TUint32Array) String() string {
	return String(iar)
}
