package xclass

import (
	_type "github.com/go-xe2/x/sync/type"
	"reflect"
)

type Object interface {
	// 类对象构造方法
	// inst:为类实例指定, props参数由实现子类管理参数个数
	Constructor(props ...interface{}) interface{}
	// 析构方法
	Destroy()
	// 类实例指针
	This() interface{}
	// 创建新实例
	NewInstance(props ...interface{}) interface{}
	// 是否继承某类
	Implements(parent Class) bool
	// 转换成父类指针
	ToParent(parent Class) interface{}
	// 类名称
	ClassName() string
	// 类类型
	ClassType() reflect.Type
	// 设置字段值
	Set(fieldName string, value interface{}) error
	// 自动转换数据类型为字段类型后设置字段值
	SafeSet(fieldName string, value interface{}) (err error)
	// 获取字段值
	Get(fieldName string) interface{}
	GetVar(fieldName string) *_type.TVar
	// 设置动态字段值, 如果字段不存在则不设置，返回字段原来值
	DynamicSet(fieldName string, value interface{}) interface{}
	// 设置动态字段值，字段不存在时则创建，返回字段原来值
	DynamicTrySet(fieldName string, value interface{}) interface{}
	// 获取动态字段值
	DynamicGet(fieldName string) interface{}
	DynamicGetVar(fieldName string) *_type.TVar
	// 获取或设置动态字段, 如果不存在时，使使用setValue进行设置，setValue可以是实际值或func() interface{}类型的函数
	DynamicGetOrSet(fieldName string, setValue interface{}) interface{}
	DynamicGetOrSetVar(fieldName string, setValue interface{}) *_type.TVar
	// 调用类方法
	Call(method string, params ...interface{}) (results []interface{}, err error)
	// 检查类是否存在某方法
	HasMethod(method string) bool
	// 检查类是否存在某字段
	HasField(fieldName string) bool
	// 是否存在某动态字段
	HasDynamicField(fieldName string) bool
	// 动态字段值
	DynamicFieldCount() int
	// 遍历动态字段列表，fn返回false时结束遍历
	ForEachDynamicFields(fn func(k string, v interface{}) bool)
	// 类对象定义字段数
	FieldCount() int
	// 遍历对象字段列表
	ForEachFields(fn func(fieldName string, tag ClassTag, fieldType reflect.Type) bool)
	// 类对象定义方法数
	MethodCount() int
	// 类tag注解
	ClassTag() ClassTag
	// 获取类对象所有字段列表
	Fields() []*ClassField
	// 获取类对象方法列表
	Methods() []*ClassMethod
	// 遍历类对象方法
	ForEachMethods(fn func(methodName string, params []reflect.Type, results []reflect.Type) bool)
}

var classType = reflect.TypeOf((*Object)(nil)).Elem()
