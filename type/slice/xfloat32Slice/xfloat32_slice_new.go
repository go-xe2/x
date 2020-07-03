package xfloat32Slice

type TFloat32Slice []float32

func New(src ...[]float32) TFloat32Slice {
	var inst []float32
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]float32, 0)
	}
	return TFloat32Slice(inst)
}
