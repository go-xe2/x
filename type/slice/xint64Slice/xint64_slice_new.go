package int64Array

type TInt64Slice []int64

func New(src ...[]int64) TInt64Slice {
	var inst []int64
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]int64, 0)
	}
	return TInt64Slice(inst)
}
