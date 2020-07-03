package xquery

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/encoding/xxml"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xq/xbinder"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

func (qe *tQueryExp) expressToken() (sqlToken SqlToken, vars []interface{}) {
	cxt := xqcomm.NewSqlCompileContext()
	cxt.SetDatabase(qe.db)
	builder := xdriver.GetSqlBuilderByName(cxt.Driver())
	sqlToken = qe.Compile(builder, cxt, false)
	vars = builder.MakeQueryParams(cxt.Params())
	return
}

func (qe *tQueryExp) Sql() (sql string, vars []interface{}) {
	var token SqlToken
	token, vars = qe.expressToken()
	sql = token.Val()
	return
}

func (qe *tQueryExp) Xml() (xxml.XmlStr, error) {
	sqlToken, vars := qe.expressToken()
	if sqlToken.TType() != SqlQueryTokenType {
		return "", exception.Newf("不是查询表达式:%s", sqlToken.Val())
	}
	if v, err := qe.db.Query(XmlBinder, sqlToken.Val(), vars...); err != nil {
		return "", err
	} else {
		if x, ok := v.(xxml.XmlStr); ok {
			return x, nil
		} else {
			return xxml.XmlStr(v.(string)), nil
		}
	}
}

func (qe *tQueryExp) Json() (xjson.JsonStr, error) {
	sqlToken, vars := qe.expressToken()
	if sqlToken.TType() != SqlQueryTokenType {
		return "", exception.Newf("不是查询表达式:%s", sqlToken.Val())
	}
	if v, err := qe.db.Query(JsonBinder, sqlToken.Val(), vars...); err != nil {
		return "", err
	} else {
		if json, ok := v.(xjson.JsonStr); ok {
			return json, nil
		} else {
			return xjson.JsonStr(v.(string)), nil
		}
	}
}

func (qe *tQueryExp) Rows() ([]map[string]interface{}, error) {
	sqlToken, vars := qe.expressToken()
	if sqlToken.TType() != SqlQueryTokenType {
		return nil, exception.Newf("不是查询表达式:%s", sqlToken.Val())
	}
	if v, err := qe.db.Query(MapBinder, sqlToken.Val(), vars...); err != nil {
		return nil, err
	} else {
		if mp, ok := v.([]map[string]interface{}); ok {
			return mp, nil
		} else if v != nil {
			return t.New(v).SliceMapStrAny(), nil
		} else {
			return nil, nil
		}
	}
}

func (qe *tQueryExp) Bind(binder DbQueryBinder) (interface{}, error) {
	sqlToken, vars := qe.expressToken()
	if sqlToken.TType() != SqlQueryTokenType {
		return "", exception.Newf("不是查询表达式:%s", sqlToken.Val())
	}
	if v, err := qe.db.QueryBind(binder, sqlToken.Val(), vars...); err != nil {
		return "", err
	} else {
		return v, nil
	}
}

func (qe *tQueryExp) Slices() ([][]interface{}, error) {
	sqlToken, vars := qe.expressToken()
	if sqlToken.TType() != SqlQueryTokenType {
		return nil, exception.Newf("不是查询表达式:%s", sqlToken.Val())
	}
	if v, err := qe.db.Query(SliceBinder, sqlToken.Val(), vars...); err != nil {
		return nil, err
	} else {
		if sl, ok := v.([][]interface{}); ok {
			return sl, nil
		} else {
			return nil, exception.NewText("数据绑定错误")
		}
	}
}

func (qe *tQueryExp) Dataset() (Dataset, error) {
	sqlToken, vars := qe.expressToken()
	if sqlToken.TType() != SqlQueryTokenType {
		return nil, exception.Newf("不是查询表达式:%s", sqlToken.Val())
	}
	if v, err := qe.db.Query(DatasetBinder, sqlToken.Val(), vars...); err != nil {
		return nil, err
	} else {
		if ds, ok := v.(Dataset); ok {
			return ds, nil
		} else {
			return nil, exception.NewText("数据绑定错误")
		}
	}
}

func (qe *tQueryExp) Visitor(visitor func(row int, values ...interface{}) (interface{}, bool)) ([]interface{}, error) {
	sqlToken, vars := qe.expressToken()
	if sqlToken.TType() != SqlQueryTokenType {
		return nil, exception.Newf("不是查询表达式:%s", sqlToken.Val())
	}
	if v, err := qe.db.QueryBind(xbinder.VisitorBinder(visitor), sqlToken.Val(), vars...); err != nil {
		return nil, err
	} else {
		if sm, ok := v.([]interface{}); ok {
			return sm, nil
		} else {
			return nil, exception.NewText("数据绑定错误")
		}
	}
}
