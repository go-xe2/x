package int16Array

type TInt16Slice []int16

func New(src ...[]int16) TInt16Slice {
	var inst []int16
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]int16, 0)
	}
	return TInt16Slice(inst)
}
