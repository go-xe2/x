package xproto

func RegistUnionStruct(className string, constructor func() ProtoUnion) bool {
	key := "union:" + className
	if constructFactories.Contains(key) {
		return false
	}
	constructFactories.Set(key, constructor)
	return true
}

func NewUnionStruct(className string) ProtoUnion {
	key := "union:" + className
	if v := constructFactories.Get(key); v != nil {
		if fn, ok := v.(func() ProtoUnion); ok {
			return fn()
		}
	}
	return nil
}
