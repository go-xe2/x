package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlToken struct {
	val    string
	tType  xqi.SqlTokenType
	params []xdriveri.SqlParam
}

var _ xqi.SqlToken = &TSqlToken{}

// 空节点
var EmptySqlToken xqi.SqlToken = &TSqlToken{
	tType: xqi.SqlEmptyTokenType,
	val:   "",
}

func NewSqlToken(val string, tType xqi.SqlTokenType) *TSqlToken {
	return &TSqlToken{
		val:    val,
		tType:  tType,
		params: make([]xdriveri.SqlParam, 0),
	}
}

func (stt *TSqlToken) Val() string {
	return stt.val
}

func (stt *TSqlToken) SetVal(val string) *TSqlToken {
	stt.val = val
	return stt
}

func (stt *TSqlToken) TType() xqi.SqlTokenType {
	return stt.tType
}

func (stt *TSqlToken) Params() []xdriveri.SqlParam {
	return stt.params
}

func (stt *TSqlToken) AddParam(name string, val interface{}) xqi.SqlToken {
	stt.params = append(stt.params, NewSqlParam(name, val))
	return stt
}

func (stt *TSqlToken) String() string {
	return fmt.Sprintf("[%s]%s", stt.tType.Name(), stt.Val())
}
