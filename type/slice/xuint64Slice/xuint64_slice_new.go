package xuint64Slice

type TUint64Slice []uint64

func New(src ...[]uint64) TUint64Slice {
	var inst []uint64
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]uint64, 0)
	}
	return TUint64Slice(inst)
}
