package xstrSlice

type TStrSlice []string

func New(src ...[]string) TStrSlice {
	var inst []string
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]string, 0)
	}
	return TStrSlice(inst)
}
