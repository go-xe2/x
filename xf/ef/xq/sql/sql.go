package sql

import (
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

// 创建sql变量
func Var(varName string, v interface{}) xqi.SqlVarExpr {
	return xqcomm.NewSqlVar(nil, v, varName)
}

// 创建sql常量
func Static(v interface{}) xqi.SqlStaticExpr {
	return xqcomm.NewSqlStatic(v)
}

func Field(expr interface{}) xqi.SqlField {
	return xqcomm.NewSqlField(expr, "")
}
