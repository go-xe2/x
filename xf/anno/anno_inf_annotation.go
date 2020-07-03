package anno

type Annotation interface {
	// 元注解名称
	AnnotationName() string
	// 元注解创建初始化
	// @param caller 调用创建元注解的实例
	// @param annParams 元注解带的参数, 在元注解字符串中定义的参数
	// @param callParams 实例化元注解时传入的其他参数，根据实际注解类型创建时传入
	AnnCreate(caller interface{}, annParams map[string]interface{}, callParams ...interface{}) interface{}
	// 元注解实例
	Instance() interface{}
}

// 元注解容器
type AnnotationContainer interface {
	// 元注解类型名称
	AnnName() string
	// 元注解实例
	Ann() Annotation
	// 元注解参数
	AnnParams() map[string]interface{}
	// 调用元注解实例的AnnCreate方法, 以创建具体属性类
	Create(caller interface{}, params ...interface{}) interface{}
}
