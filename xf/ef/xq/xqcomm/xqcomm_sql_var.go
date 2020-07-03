package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlVar struct {
	paramName string
	val       interface{}
	this      interface{}
}

var _ SqlVarExpr = &TSqlVar{}

func NewSqlVar(inherited interface{}, val interface{}, name ...string) *TSqlVar {
	szName := ""
	if len(name) > 0 {
		szName = name[0]
	}
	inst := &TSqlVar{
		paramName: szName,
		val:       val,
		this:      nil,
	}
	if _, ok := inherited.(SqlVarExpr); ok {
		inst.this = inherited
	} else {
		inst.this = inst
	}
	return inst
}

func (srv *TSqlVar) TokenType() SqlTokenType {
	return SqlVarTokenType
}

func (srv *TSqlVar) Val() interface{} {
	return srv.val
}

func (srv *TSqlVar) SetVal(v interface{}) SqlVarExpr {
	srv.val = v
	return srv
}

func (srv *TSqlVar) This() interface{} {
	return srv.this
}

func (srv *TSqlVar) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	isPrepare := true
	if len(unPrepare) > 0 {
		isPrepare = !unPrepare[0]
	}
	var result = NewSqlToken("", SqlVarTokenType)
	if isPrepare {
		var paramName = ""
		state := cxt.State()
		if srv.paramName != "" {
			paramName = srv.paramName
		} else if state == SCPQrSelectJoinItemState {
			paramName = cxt.MakeParamId()
		}
		cxt.AddParam(paramName, srv.val)
		result.AddParam(paramName, srv.val)
		result.SetVal(builder.PlaceHolder(paramName))
	} else {
		result.SetVal(builder.MakeRealValue(srv.val))
	}
	return result
}
