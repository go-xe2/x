package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/encoding/xxml"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xq"
	"github.com/go-xe2/x/xf/ef/xq/sql"
	"github.com/go-xe2/x/xf/ef/xq/xbinder"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type tEntitySelectPage struct {
	pageIndex int
	pageSize  int
	sel       *tEntitySelect
}

var _ EntitySelectPage = (*tEntitySelectPage)(nil)

func newEntitySelectPage(entSelect *tEntitySelect, pageIndex, pageSize int) *tEntitySelectPage {
	inst := &tEntitySelectPage{
		sel:       entSelect,
		pageIndex: pageIndex,
		pageSize:  pageSize,
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return inst
}

// map数组列表
func (esp *tEntitySelectPage) Rows() (data []map[string]interface{}, info QueryPageInfo, err error) {
	binder := xbinder.GetQueryBinder(MapBinder)
	var v interface{}
	v, info, err = esp.Bind(binder)
	if err != nil {
		return
	}
	data = v.([]map[string]interface{})
	return
}

// xml 字符串
func (esp *tEntitySelectPage) Xml() (data xxml.XmlStr, info QueryPageInfo, err error) {
	binder := xbinder.GetQueryBinder(XmlBinder)
	var v interface{}
	v, info, err = esp.Bind(binder)
	if err != nil {
		return
	}
	data = v.(xxml.XmlStr)
	return
}

// json字符串
func (esp *tEntitySelectPage) Json() (data xjson.JsonStr, info QueryPageInfo, err error) {
	binder := xbinder.GetQueryBinder(JsonBinder)
	var v interface{}
	v, info, err = esp.Bind(binder)
	if err != nil {
		return
	}
	data = v.(xjson.JsonStr)
	return
}

// 返回数据集
func (esp *tEntitySelectPage) Dataset() (data Dataset, info QueryPageInfo, err error) {
	binder := xbinder.GetQueryBinder(DatasetBinder)
	var v interface{}
	v, info, err = esp.Bind(binder)
	if err != nil {
		return
	}
	data = v.(Dataset)
	return
}

// 返回数组列表
func (esp *tEntitySelectPage) Slice() (data [][]interface{}, info QueryPageInfo, err error) {
	binder := xbinder.GetQueryBinder(SliceBinder)
	var v interface{}
	v, info, err = esp.Bind(binder)
	if err != nil {
		return
	}
	data = v.([][]interface{})
	return
}

//  访问返回数据并构造返回数据
func (esp *tEntitySelectPage) Visit(visitor QueryBinderVisit) (data interface{}, info QueryPageInfo, err error) {
	binder := xbinder.VisitorBinder(visitor)
	return esp.Bind(binder)
}

// 自定议绑定器绑定返回数据方法
func (esp *tEntitySelectPage) Bind(binder DbQueryBinder) (data interface{}, info QueryPageInfo, err error) {
	esp.sel.query = esp.sel.query.Page(esp.pageSize, esp.pageIndex)
	qryInfo := esp.sel.query.Info()
	if qryInfo == nil {
		return nil, nil, exception.Newf("查询表达式不完整")
	}
	// 统计记录数，获取查询条件,等相关信息
	tableExpr := qryInfo.GetTable()
	whereExpr := qryInfo.GetWhere()
	joinExpr := qryInfo.GetJoins()
	totalQry := xq.QueryByDb(esp.sel.query.DB()).Fields(func(SqlTables) []SqlField {
		return []SqlField{sql.Count("*").Alias("count")}
	}).From(tableExpr)
	var makeJoinCondition = func(condition func() SqlCondition) func(SqlTable, SqlTables, SqlCondition) SqlCondition {
		var info = struct {
			fn func() SqlCondition
		}{
			fn: condition,
		}
		return func(SqlTable, SqlTables, SqlCondition) SqlCondition {
			return info.fn()
		}
	}
	if len(joinExpr) > 0 {
		for _, item := range joinExpr {
			switch item.JoinType() {
			case SqlInnerJoinType:
				totalQry = totalQry.Join(item.JoinTable(), makeJoinCondition(item.OnCondition))
				break
			case SqlLeftJoinType:
				totalQry = totalQry.LeftJoin(item.JoinTable(), makeJoinCondition(item.OnCondition))
				break
			case SqlRightJoinType:
				totalQry = totalQry.RightJoin(item.JoinTable(), makeJoinCondition(item.OnCondition))
				break
			case SqlCrossJoinType:
				totalQry = totalQry.CrossJoin(item.JoinTable(), makeJoinCondition(item.OnCondition))
				break
			}
		}

	}
	if whereExpr != nil {
		totalQry = totalQry.Where(func(SqlCondition, SqlTables) SqlCondition {
			return whereExpr
		})
	}
	var totalRow = 0
	if _, err = totalQry.Limit(1).Visitor(func(row int, values ...interface{}) (i interface{}, b bool) {
		totalRow = t.Int(values[0])
		return totalRow, true
	}); err != nil {
		esp.sel.setLastSql(totalQry.DB().LastSql())
		return nil, nil, err
	}
	esp.sel.setLastSql(totalQry.DB().LastSql())
	pageCount := totalRow / esp.pageSize
	if totalRow%esp.pageSize > 0 {
		pageCount += 1
	}

	var convertBinder = esp.sel.initBinder(binder)

	v, err := esp.sel.query.Page(esp.pageSize, esp.pageIndex).Bind(convertBinder)
	if err != nil {
		return nil, nil, err
	}
	esp.sel.setLastSql(esp.sel.query.DB().LastSql())

	return v, NewQueryPageInfo(esp.pageIndex, esp.pageSize, pageCount, totalRow), nil
}

func (esp *tEntitySelectPage) Convert() EntitySelectConvert {
	return esp.sel
}
