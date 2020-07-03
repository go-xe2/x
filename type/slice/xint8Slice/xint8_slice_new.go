package int8Array

type TInt8Slice []int8

func New(src ...[]int8) TInt8Slice {
	var inst []int8
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]int8, 0)
	}
	return TInt8Slice(inst)
}
