package xclass

// 类对象观察器
type ObjectFieldObserver interface {
	// 当设置字段值时触发，仅在调用Set、DynamicSet、DynamicTrySet、DynamicGetOrSet方法设置新值时发生.如果返回false将取消设置
	OnFieldChanged(fieldName string, oldV, newV interface{}) bool
}

type ObjectDynamicFieldObserver interface {
	// 当设置动态字段值触发，仅在调用DynamicSet、DynamicTrySet、DynamicGetOrSet方法设置新值时发生, 如果返回false将取消设置
	OnDynamicFieldChanged(fieldName string, oldV, newV interface{}) bool
}
