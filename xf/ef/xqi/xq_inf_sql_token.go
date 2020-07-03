package xqi

import "github.com/go-xe2/x/xf/ef/xdriveri"

type SqlToken interface {
	// sql脚本
	Val() string
	// sql节点类型
	TType() SqlTokenType
	// sql节点参数
	Params() []xdriveri.SqlParam
	AddParam(name string, val interface{}) SqlToken
	String() string
}
