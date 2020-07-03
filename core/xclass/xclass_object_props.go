package xclass

import "reflect"

func (o *TObject) ClassTag() ClassTag {
	return o.classTag
}

// 获取类名称
func (o *TObject) ClassName() string {
	return classAliasName(o.classValue.Type().String())
}

// 获取类类型
func (o *TObject) ClassType() reflect.Type {
	return o.classValue.Type()
}
