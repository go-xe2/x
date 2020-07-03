package t

func New(o interface{}) Type {
	switch o.(type) {
	case T:
		return Type{o.(T).Any()}
	default:
		return Type{val: o}
	}
}
