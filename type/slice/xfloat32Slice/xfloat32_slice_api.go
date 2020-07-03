package xfloat32Slice

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (iar TFloat32Slice) Contain(item float32) bool {
	return Contain(iar, item)
}

func (iar TFloat32Slice) ForEach(fn FloatArrForEachFunc) {
	ForEach(iar, fn)
}

func (iar TFloat32Slice) Find(fn FloatArrSearchFunc) float32 {
	return Find(iar, fn)
}

func (iar TFloat32Slice) FindIndex(fn FloatArrSearchFunc) int {
	return FindIndex(iar, fn)
}

func (iar TFloat32Slice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(iar, comparer...)
}

func (iar TFloat32Slice) Size() int {
	return len(iar)
}

func (iar *TFloat32Slice) Append(item ...float32) int {
	*iar = TFloat32Slice(append(*iar, item...))
	return iar.Size()
}

func (iar *TFloat32Slice) Prepend(item ...float32) int {
	old := *iar
	*iar = append(make([]float32, 0), item...)
	*iar = append(*iar, old...)
	return len(*iar)
}

func (iar *TFloat32Slice) Insert(index int, item float32) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(old[:index], item)
	*iar = append(*iar, old[index:]...)
	return nil
}

func (iar *TFloat32Slice) Clear() {
	*iar = make([]float32, 0)
}

func (iar *TFloat32Slice) Delete(index int) error {
	if index < 0 || index >= iar.Size() {
		return errors.New("数组下标越界")
	}
	old := *iar
	*iar = append(make([]float32, 0), old[:index-1]...)
	*iar = append(*iar, old[index+1:]...)
	return nil
}

func (iar *TFloat32Slice) Concat(arrs ...[]float32) int {
	for _, nar := range arrs {
		*iar = append(*iar, nar...)
	}
	return iar.Size()
}

func (iar TFloat32Slice) Join(sep string) string {
	return Join(iar, sep)
}

func (iar TFloat32Slice) String() string {
	return String(iar)
}
