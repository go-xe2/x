package xuint32Slice

type TUint32Array []uint32

func New(src ...[]uint32) TUint32Array {
	var inst []uint32
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]uint32, 0)
	}
	return TUint32Array(inst)
}
