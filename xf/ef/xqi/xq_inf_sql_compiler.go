package xqi

import (
	"github.com/go-xe2/x/xf/ef/xdriveri"
)

type SqlCompiler interface {
	// sql编译接口
	// unPrepare默认为false, 为true时不建立参数化语句
	// varReceiver 参数化查询时实参收集器
	Compile(builder xdriveri.DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken
	TokenType() SqlTokenType
	This() interface{}
}
