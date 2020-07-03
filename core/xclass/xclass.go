package xclass

func Create(class Class, props ...interface{}) interface{} {
	vt := classToClassVT(class.Type())
	return classAlloc(vt, props...)
}

func Free(obj *Object) {
	o := *obj
	o.Destroy()
	*obj = nil
}
