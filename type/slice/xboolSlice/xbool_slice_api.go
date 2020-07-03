package xboolSlice

import . "github.com/go-xe2/x/type/slice/comm"

func (arr TBoolSlice) Contain(item bool) bool {
	return Contain(arr, item)
}

func (arr TBoolSlice) Map(fn BoolArrMapFunc) {
	Map(arr, fn)
}

func (arr TBoolSlice) ForEach(fn BoolArrForEachFunc) {
	ForEach(arr, fn)
}

func (arr TBoolSlice) Join(sep string) string {
	return Join(arr, sep)
}

func (arr TBoolSlice) String() string {
	return String(arr)
}

func (arr TBoolSlice) AsInterface() []interface{} {
	return AsInterface(arr)
}
