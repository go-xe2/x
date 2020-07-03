package xproto

func RegistStruct(className string, constructor func() ProtoStruct) bool {
	key := "struct:" + className
	if constructFactories.Contains(key) {
		return false
	}
	constructFactories.Set(key, constructor)
	return true
}

func NewStruct(className string) ProtoStruct {
	key := "struct:" + className
	if v := constructFactories.Get(key); v != nil {
		if fn, ok := v.(func() ProtoStruct); ok {
			return fn()
		}
	}
	return nil
}
