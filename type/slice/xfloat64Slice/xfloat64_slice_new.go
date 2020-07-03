package xfloat64Slice

type TFloat64Slice []float64

func New(src ...[]float64) TFloat64Slice {
	var inst []float64
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]float64, 0)
	}
	return TFloat64Slice(inst)
}
