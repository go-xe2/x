package xuint16Slice

type TUint16Slice []uint16

func New(src ...[]uint16) TUint16Slice {
	var inst []uint16
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]uint16, 0)
	}
	return TUint16Slice(inst)
}
