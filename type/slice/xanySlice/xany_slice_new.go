package xanySlice

type TAnySlice []interface{}

func New(src ...[]interface{}) TAnySlice {
	var inst []interface{}
	if len(src) > 0 && src[0] != nil {
		inst = src[0]
	} else {
		inst = make([]interface{}, 0)
	}
	return TAnySlice(inst)
}
