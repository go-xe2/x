package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xdriveri"
)

type TSqlParam struct {
	name string
	val  interface{}
}

var _ xdriveri.SqlParam = &TSqlParam{}

// 创建参数化查询参数
func NewSqlParam(name string, val interface{}) *TSqlParam {
	return &TSqlParam{
		name: name,
		val:  val,
	}
}

func (sp *TSqlParam) Name() string {
	return sp.name
}

func (sp *TSqlParam) Val() interface{} {
	return sp.val
}

func (sp *TSqlParam) String() string {
	return fmt.Sprintf("{ name:%s, value:%v }", sp.name, sp.val)
}
