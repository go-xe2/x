package xintSlie

type TIntSlice []int

func New(src ...[]int) TIntSlice {
	var inst []int
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]int, 0)
	}
	return TIntSlice(inst)
}
