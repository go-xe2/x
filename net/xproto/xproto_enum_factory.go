package xproto

// 注册enum类型
func RegisterEnum(enumType string, constructor func() ProtoEnum) bool {
	key := "enum:" + enumType
	if constructFactories.Contains(key) {
		return false
	}
	constructFactories.Set(key, constructor)
	return true
}

func NewEnum(enumType string) ProtoEnum {
	key := "enum:" + enumType
	if v := constructFactories.Get(key); v != nil {
		if fn, ok := v.(func() ProtoEnum); ok {
			return fn()
		}
	}
	return nil
}
