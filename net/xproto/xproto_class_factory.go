package xproto

func RegisterClass(className string, constructor func() ProtoClass) bool {
	key := "class:" + className
	if constructFactories.Contains(key) {
		return false
	}
	constructFactories.Set(key, constructor)
	return true
}

func NewClass(className string) ProtoClass {
	key := "class:" + className
	if v := constructFactories.Get(key); v != nil {
		if fn, ok := v.(func() ProtoClass); ok {
			return fn()
		}
	}
	return nil
}
