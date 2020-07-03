package xproto

func RegistDataset(className string, constructor func() ProtoDataset) bool {
	key := "struct:" + className
	if constructFactories.Contains(key) {
		return false
	}
	constructFactories.Set(key, constructor)
	return true
}

func NewMemDataset(className string) ProtoDataset {
	key := "struct:" + className
	if v := constructFactories.Get(key); v != nil {
		if fn, ok := v.(func() ProtoDataset); ok {
			return fn()
		}
	}
	return nil
}
