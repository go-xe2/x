package xuintSlice

type TUintArray []uint

func New(src ...[]uint) TUintArray {
	var inst []uint
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]uint, 0)
	}
	return TUintArray(inst)
}
