package xuint8Slice

type TUint8Slice []uint8

func New(src ...[]uint8) TUint8Slice {
	var inst []uint8
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]uint8, 0)
	}
	return TUint8Slice(inst)
}
