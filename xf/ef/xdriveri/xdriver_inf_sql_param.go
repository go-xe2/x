package xdriveri

type SqlParam interface {
	// 参数名称
	Name() string
	// 参数值
	Val() interface{}
	String() string
}
