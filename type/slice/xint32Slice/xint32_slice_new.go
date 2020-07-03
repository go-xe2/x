package xint32Slice

type TInt32Slice []int32

func New(src ...[]int32) TInt32Slice {
	var inst []int32
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]int32, 0)
	}
	return TInt32Slice(inst)
}
