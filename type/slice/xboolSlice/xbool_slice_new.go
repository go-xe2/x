package xboolSlice

type TBoolSlice []bool

func New(src ...[]bool) TBoolSlice {
	var inst []bool
	if len(src) > 0 {
		inst = src[0]
	} else {
		inst = make([]bool, 0)
	}
	return TBoolSlice(inst)
}
