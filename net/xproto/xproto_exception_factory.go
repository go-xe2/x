package xproto

func RegistException(className string, constructor func() ProtoException) bool {
	key := "exception:" + className
	if constructFactories.Contains(key) {
		return false
	}
	constructFactories.Set(key, constructor)
	return true
}

func NewException(className string) ProtoException {
	key := "exception:" + className
	if v := constructFactories.Get(key); v != nil {
		if fn, ok := v.(func() ProtoException); ok {
			return fn()
		}
	}
	return nil
}
