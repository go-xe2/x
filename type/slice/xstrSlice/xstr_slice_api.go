package xstrSlice

import "errors"
import . "github.com/go-xe2/x/type/slice/comm"

func (sar TStrSlice) Contain(item string) bool {
	return Contain(sar, item)
}

func (sar TStrSlice) ForEach(fn StrArrForEachFunc) {
	ForEach(sar, fn)
}

func (sar TStrSlice) Find(fn StrArrSearchFunc) string {
	return Find(sar, fn)
}

func (sar TStrSlice) FindIndex(fn StrArrSearchFunc) int {
	return FindIndex(sar, fn)
}

func (sar TStrSlice) Sort(comparer ...ArrSortCompareFunc) {
	Sort(sar, comparer...)
}

func (sar TStrSlice) Size() int {
	return len(sar)
}

func (sar *TStrSlice) Append(item ...string) int {
	*sar = TStrSlice(append(*sar, item...))
	return sar.Size()
}

func (sar *TStrSlice) Prepend(item ...string) int {
	old := *sar
	*sar = append(make([]string, 0), item...)
	*sar = append(*sar, old...)
	return len(*sar)
}

func (sar *TStrSlice) Insert(index int, item string) error {
	if index < 0 || index >= sar.Size() {
		return errors.New("数组下标越界")
	}
	old := *sar
	*sar = append(old[:index], item)
	*sar = append(*sar, old[index:]...)
	return nil
}

func (sar *TStrSlice) Clear() {
	*sar = make([]string, 0)
}

func (sar *TStrSlice) Delete(index int) error {
	if index < 0 || index >= sar.Size() {
		return errors.New("数组下标越界")
	}
	old := *sar
	*sar = append(make([]string, 0), old[:index-1]...)
	*sar = append(*sar, old[index+1:]...)
	return nil
}

func (sar *TStrSlice) Concat(arrs ...[]string) int {
	for _, nar := range arrs {
		*sar = append(*sar, nar...)
	}
	return sar.Size()
}

func (sar TStrSlice) Join(sep string) string {
	return Join(sar, sep)
}

func (sar TStrSlice) String() string {
	return String(sar)
}

func (sar TStrSlice) AsInterface() []interface{} {
	return AsInterface(sar)
}
