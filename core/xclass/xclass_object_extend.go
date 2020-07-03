package xclass

// 检查是否继承自某类
func (o *TObject) Implements(parent Class) bool {
	if v, ok := o.extends[parent.Type()]; ok && v != nil {
		return true
	}
	return false
}

// 赋值给父类指针
func (o *TObject) ToParent(parent Class) interface{} {
	if v, ok := o.extends[parent.Type()]; ok && v != nil {
		return v.Instance()
	}
	return nil
}
