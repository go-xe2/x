/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-23 09:57
* Description:
*****************************************************************/

package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlFunToBase64 struct {
	*TSqlFun
	exp interface{}
}

var _ SqlFun = &TSqlFunToBase64{}

func NewSqlFunToBase64(exp interface{}) SqlFun {
	inst := &TSqlFunToBase64{
		exp: exp,
	}
	base := newSqlFun(SFToBase64, inst)
	inst.TSqlFun = base
	return inst
}

func (sf *TSqlFunToBase64) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sf.exp == nil {
		return EmptySqlToken
	}
	field := ""
	result := NewSqlToken("", SqlFunExpressTokenType)
	if cv, ok := sf.exp.(SqlCompiler); ok {
		cxt.PushState(SCPQrSelectFunParamState)
		tk := cv.Compile(builder, cxt, unPrepare...)
		cxt.PopState()
		field = tk.Val()
	} else {
		prepare := true
		if len(unPrepare) > 0 {
			prepare = !unPrepare[0]
		}
		if prepare {
			vn := cxt.MakeParamId()
			cxt.AddParam(vn, sf.exp)
			result.AddParam(vn, sf.exp)
			field = builder.PlaceHolder(vn)
		} else {
			field = builder.MakeRealValue(sf.exp)
		}
	}
	result.SetVal(builder.ToBase64(field))
	return result
}
