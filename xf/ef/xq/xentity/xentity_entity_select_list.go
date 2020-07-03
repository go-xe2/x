package xentity

import (
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/encoding/xxml"
	"github.com/go-xe2/x/xf/ef/xq/xbinder"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type tEntitySelectList struct {
	sel *tEntitySelect
}

var _ EntitySelectList = (*tEntitySelectList)(nil)

func newEntitySelectList(entSelect *tEntitySelect) *tEntitySelectList {
	return &tEntitySelectList{
		sel: entSelect,
	}
}

// 返回[]map[string]interface数据方法
func (esl *tEntitySelectList) Rows() (data []map[string]interface{}, err error) {
	binder := xbinder.NewQryMapBinder(nil, esl.sel.getFieldConvert())
	v, err := esl.sel.query.Bind(binder)
	if err != nil {
		return nil, err
	}
	data = v.([]map[string]interface{})
	esl.sel.setLastSql(esl.sel.query.DB().LastSql())
	return
}

// 返回xml字符串
func (esl *tEntitySelectList) Xml() (data xxml.XmlStr, err error) {
	binder := xbinder.NewQryXmlBinder(nil, esl.sel.getFieldConvert())
	v, err := esl.sel.query.Bind(binder)
	if err != nil {
		return "", err
	}
	data = v.(xxml.XmlStr)
	esl.sel.setLastSql(esl.sel.query.DB().LastSql())
	return
}

// 返回json字符串
func (esl *tEntitySelectList) Json() (data xjson.JsonStr, err error) {
	binder := xbinder.NewQryJsonBinder(nil, esl.sel.getFieldConvert())
	v, err := esl.sel.query.Bind(binder)
	if err != nil {
		return "", err
	}
	data = v.(xjson.JsonStr)
	esl.sel.setLastSql(esl.sel.query.DB().LastSql())
	return
}

// 返回数据集
func (esl *tEntitySelectList) Dataset() (data Dataset, err error) {
	data, err = esl.sel.query.Dataset()
	return
}

// 返回slice数组
func (esl *tEntitySelectList) Slice() (data [][]interface{}, err error) {
	binder := xbinder.NewQrySliceBinder(nil, esl.sel.getFieldConvert())
	v, err := esl.sel.query.Bind(binder)
	if err != nil {
		return nil, err
	}
	data = v.([][]interface{})
	esl.sel.setLastSql(esl.sel.query.DB().LastSql())
	return
}

// 访问返回数据并构造返回数据
func (esl *tEntitySelectList) Visit(visitor QueryBinderVisit) (data interface{}, err error) {
	binder := xbinder.VisitorBinder(visitor)
	return esl.sel.query.Bind(binder)
}

// 自定议绑定器绑定返回数据方法
func (esl *tEntitySelectList) Bind(binder DbQueryBinder) (data interface{}, err error) {
	convertBinder := esl.sel.initBinder(binder)
	result, err := esl.sel.query.Bind(convertBinder)
	esl.sel.setLastSql(esl.sel.query.DB().LastSql())
	return result, err
}

func (esl *tEntitySelectList) Convert() EntitySelectConvert {
	return esl.sel
}
