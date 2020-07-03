package xqcomm

import (
	"github.com/go-xe2/x/core/exception"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
	"strings"
)

type lazyLoadFlags int

const (
	// 加载join连接的表
	lazyLoadTables lazyLoadFlags = 1 << iota
	// 加载where条件
	lazyLoadWhereCond
	// 加载having条件
	lazyLoadHavingCond
	// 加载join条件
	lazyLoadJoinItems
	// 加载fields字段
	LazyLoadFields
	// 加载group fields字段
	lazyLoadGroupFields
	// 加载order fields字段
	lazyLoadOrderFields
)

type TSqlQuery struct {
	// 查询主表
	fromTable SqlTable
	// 查询所用到的表
	qrTables SqlTables
	// 联合查询到的表
	joinTables []SqlJoinItem
	// 查询字段列表
	qrFieldFun LazyGetFieldFun
	// 分组字段列表
	groupFieldFun LazyGetFieldFun
	// 获取排序字段列表
	orderFieldFun LazyGetOrderFieldFun
	// 获取where条件方法
	whereFun LazyGetConditionFun
	// 获having方法
	havingFun LazyGetConditionFun
	// limit offset
	limitOffset int
	// limit rows
	limitRows int
	// 分页大小
	pageSize int
	// 页码
	pageIndex int

	where       SqlCondition
	having      SqlCondition
	fields      []SqlField
	groupFields []SqlField
	orderFields []SqlOrderField
	joins       []SqlJoin

	// 是否已经加载所有表
	loadFlags lazyLoadFlags
}

var _ SqlQuery = &TSqlQuery{}

func NewSqlQuery(table SqlTable) *TSqlQuery {
	return &TSqlQuery{
		fromTable:     table,
		qrTables:      NewSqlTables(),
		joinTables:    make([]SqlJoinItem, 0),
		qrFieldFun:    nil,
		groupFieldFun: nil,
		orderFieldFun: nil,
		whereFun:      nil,
		havingFun:     nil,
		limitOffset:   0,
		limitRows:     0,
		pageSize:      0,
		pageIndex:     0,
		loadFlags:     0,
		where:         nil,
		having:        nil,
		fields:        nil,
		groupFields:   nil,
		orderFields:   nil,
		joins:         nil,
	}
}

func (sq *TSqlQuery) Exp() interface{} {
	return sq
}

func (sq *TSqlQuery) isLoad(flag lazyLoadFlags) bool {
	return sq.loadFlags&flag == flag
}

func (sq *TSqlQuery) includeLoad(flag lazyLoadFlags) {
	sq.loadFlags = sq.loadFlags | flag
}

func (sq *TSqlQuery) excludeLoad(flag lazyLoadFlags) {
	sq.loadFlags = sq.loadFlags & (^flag)
}

func (sq *TSqlQuery) TokenType() SqlTokenType {
	return SqlQueryTokenType
}

func (sq *TSqlQuery) Compile(builder DbDriverSqlBuilder, cxt SqlCompileContext, unPrepare ...bool) SqlToken {
	if sq.fromTable == nil {
		panic(exception.NewText("查询表达式中未选择表."))
	}
	szQueryFields := ""
	szQueryTable := ""
	szQueryJoins := ""
	szQueryWhere := ""
	szQueryGroup := ""
	szQueryHaving := ""
	szQueryOrder := ""

	// 查询字段
	queryFields := sq.GetQueryFields()
	resultFields := make(map[string]SqlField)

	// 收集查询字段 fields
	if len(queryFields) == 0 {
		szQueryFields = "*"
	} else {
		fieldNames := make([]string, 0)
		cxt.PushState(SCPQrSelectFieldsState)
		for _, field := range queryFields {
			fdName := field.AliasName()
			if _, ok := resultFields[fdName]; !ok {
				resultFields[fdName] = field
			}
			if fdToken := field.Compile(builder, cxt, unPrepare...); fdToken != nil && fdToken.TType() != SqlEmptyTokenType {
				fieldNames = append(fieldNames, fdToken.Val())
			}
		}
		cxt.PopState()
		szQueryFields = strings.Join(fieldNames, ",")
	}

	// from
	cxt.PushState(SCPQrSelectFromState)
	fromToken := sq.fromTable.Compile(builder, cxt, unPrepare...)
	cxt.PopState()
	if fromToken != nil && fromToken.TType() != SqlEmptyTokenType {
		szQueryTable = fromToken.Val()
	}

	// 收集join表达式
	joins := sq.GetJoins()
	if len(joins) > 0 {
		joinRows := make([]string, 0)
		// 创建join连接表达式
		for _, joinItem := range joins {
			cxt.PushState(SCPQrSelectJoinTableState)
			joinToken := joinItem.JoinTable().Compile(builder, cxt, unPrepare...)
			cxt.PopState()
			var szJoin = ""
			if joinToken != nil && joinToken.TType() != SqlEmptyTokenType {
				szJoin = joinItem.JoinType().Exp() + joinToken.Val()
			}
			if szJoin == "" {
				continue
			}
			if condition := joinItem.OnCondition(); condition != nil {
				cxt.PushState(SCPQrSelectJoinCondState)
				conditionToken := condition.Compile(builder, cxt, unPrepare...)
				cxt.PopState()
				if conditionToken != nil && conditionToken.TType() != SqlEmptyTokenType {
					szJoin += " ON " + conditionToken.Val()
				}
			}
			joinRows = append(joinRows, szJoin)
		}
		if len(joinRows) > 0 {
			szQueryJoins = strings.Join(joinRows, " ")
		}
	}

	// where condition
	where := sq.GetWhere()
	if where != nil {
		cxt.PushState(SCPQrSelectWhereCondState)
		whereToken := where.Compile(builder, cxt, unPrepare...)
		cxt.PopState()
		if whereToken != nil && whereToken.TType() != SqlEmptyTokenType {
			szQueryWhere = whereToken.Val()
		}
	}

	// group by fields
	groups := sq.GetGroupFields()
	groupFieldCount := len(groups)
	if groupFieldCount == 0 {
		// 检查是否使用聚合函数，如果使用生成分组字段列表
		aggreFields := make([]string, 0)
		tmpFields := make([]SqlField, 0)
		for _, field := range queryFields {
			if field.TokenType() == SqlAggregateFunTokenType {
				aggreFields = append(aggreFields, field.AliasName())
			} else {
				tmpFields = append(tmpFields, field)
			}
		}
		if len(aggreFields) > 0 {
			// 存在使用聚合函数，创建未使用聚合函数的字段列表
			groupFieldCount = len(tmpFields)
			groups = tmpFields
		}
	}
	if groupFieldCount > 0 {
		fieldNames := make([]string, 0)
		for _, field := range groups {
			fdName := field.AliasName()
			if _, ok := resultFields[fdName]; !ok {
				panic(exception.Newf("分组统计字符%s无效，字段不在查询字段列表中", fdName))
			}
			cxt.PushState(SCPQrSelectGroupFieldState)
			fieldToken := field.Compile(builder, cxt, unPrepare...)
			cxt.PopState()
			if fieldToken != nil && fieldToken.TType() != SqlEmptyTokenType {
				fieldNames = append(fieldNames, fieldToken.Val())
			}
		}
		if len(fieldNames) > 0 {
			szQueryGroup = strings.Join(fieldNames, ",")
		}
	}

	// having condition
	having := sq.GetHaving()
	if having != nil {
		cxt.PushState(SCPQrSelectHavingCondState)
		havingToken := having.Compile(builder, cxt, unPrepare...)
		cxt.PopState()
		if havingToken != nil && havingToken.TType() != SqlEmptyTokenType {
			szQueryHaving = havingToken.Val()
		}
	}

	// order fields
	orders := sq.GetOrderFields()
	if len(orders) > 0 {
		orderNames := make([]string, 0)
		for _, field := range orders {
			cxt.PushState(SCPQrSelectOrderFieldState)
			fieldToken := field.Compile(builder, cxt, unPrepare...)
			cxt.PopState()
			if fieldToken != nil && fieldToken.TType() != SqlEmptyTokenType {
				orderNames = append(orderNames, fieldToken.Val())
			}
		}
		if len(orderNames) > 0 {
			szQueryOrder = strings.Join(orderNames, ",")
		}
	}
	sql := builder.BuildQuery(szQueryTable, szQueryFields, szQueryJoins, szQueryWhere, szQueryGroup, szQueryHaving, szQueryOrder, sq.limitRows, sq.limitOffset)
	return NewSqlToken(sql, sq.TokenType())
}

func (sq *TSqlQuery) This() interface{} {
	return sq
}

func (sq *TSqlQuery) GetTables() SqlTables {
	if !sq.isLoad(lazyLoadTables) {
		sq.includeLoad(lazyLoadTables)
		sq.qrTables.Clear()
		sq.qrTables.Add(sq.fromTable)
		for _, join := range sq.joinTables {
			sq.qrTables.Add(join.JoinTable())
		}
	}
	return sq.qrTables
}

func (sq *TSqlQuery) GetFromTable() SqlTable {
	return sq.fromTable
}

func (sq *TSqlQuery) GetQueryFields() []SqlField {
	if !sq.isLoad(LazyLoadFields) {
		sq.includeLoad(LazyLoadFields)
		tbs := sq.GetTables()
		sq.fields = make([]SqlField, 0)
		if fn := sq.qrFieldFun; fn != nil {
			sq.fields = fn(tbs)
		}
	}
	return sq.fields
}

func (sq *TSqlQuery) GetJoins() []SqlJoin {
	if !sq.isLoad(lazyLoadJoinItems) {
		sq.includeLoad(lazyLoadJoinItems)
		tbs := sq.GetTables()
		sq.joins = make([]SqlJoin, 0)
		for _, item := range sq.joinTables {
			var condition SqlCondition = nil
			if fn := item.LazyConditionFn(); fn != nil {
				condition = fn(item.JoinTable(), tbs)
			}
			sq.joins = append(sq.joins, NewSqlJoin(item.JoinType(), item.JoinTable(), condition))
		}
	}
	return sq.joins
}

func (sq *TSqlQuery) GetWhere() SqlCondition {
	if !sq.isLoad(lazyLoadWhereCond) {
		sq.includeLoad(lazyLoadWhereCond)
		if fn := sq.whereFun; fn != nil {
			sq.where = fn(sq.GetTables())
		}
	}
	return sq.where
}

func (sq *TSqlQuery) GetGroupFields() []SqlField {
	if !sq.isLoad(lazyLoadGroupFields) {
		sq.includeLoad(lazyLoadGroupFields)
		if fn := sq.groupFieldFun; fn != nil {
			sq.groupFields = fn(sq.GetTables())
		}
	}
	return sq.groupFields
}

func (sq *TSqlQuery) GetHaving() SqlCondition {
	if !sq.isLoad(lazyLoadHavingCond) {
		sq.includeLoad(lazyLoadHavingCond)
		if fn := sq.havingFun; fn != nil {
			sq.having = fn(sq.GetTables())
		}
	}
	return sq.having
}

func (sq *TSqlQuery) GetOrderFields() []SqlOrderField {
	if !sq.isLoad(lazyLoadOrderFields) {
		sq.includeLoad(lazyLoadOrderFields)
		if fn := sq.orderFieldFun; fn != nil {
			sq.orderFields = fn(sq.GetTables())
		}
	}
	return sq.orderFields
}

func (sq *TSqlQuery) GetLimitRows() int {
	return sq.limitRows
}

func (sq *TSqlQuery) GetLimitOffset() int {
	return sq.limitOffset
}

func (sq *TSqlQuery) Fields(fields func(tables SqlTables) []SqlField) SqlQuery {
	sq.fields = nil
	sq.excludeLoad(LazyLoadFields)
	sq.qrFieldFun = fields
	return sq
}

func (sq *TSqlQuery) Where(condition func(tables SqlTables) SqlCondition) SqlQuery {
	sq.excludeLoad(lazyLoadWhereCond)
	sq.whereFun = condition
	sq.where = nil
	return sq
}

func (sq *TSqlQuery) join(joinType SqlJoinType, table SqlTable, on LazyGetJoinConditionFun) {
	sq.excludeLoad(lazyLoadJoinItems)
	sq.excludeLoad(lazyLoadTables)
	sq.joins = nil
	sq.joinTables = append(sq.joinTables, NewSqlJoinItem(joinType, table, on))
}

func (sq *TSqlQuery) Join(table SqlTable, on func(joinTable SqlTable, tables SqlTables) SqlCondition) SqlQuery {
	sq.join(SqlInnerJoinType, table, on)
	return sq
}

func (sq *TSqlQuery) LeftJoin(table SqlTable, on func(joinTable SqlTable, tables SqlTables) SqlCondition) SqlQuery {
	sq.join(SqlLeftJoinType, table, on)
	return sq
}

func (sq *TSqlQuery) RightJoin(table SqlTable, on func(joinTable SqlTable, tables SqlTables) SqlCondition) SqlQuery {
	sq.join(SqlRightJoinType, table, on)
	return sq
}

func (sq *TSqlQuery) CrossJoin(table SqlTable, on func(joinTable SqlTable, tables SqlTables) SqlCondition) SqlQuery {
	sq.join(SqlCrossJoinType, table, on)
	return sq
}

func (sq *TSqlQuery) Group(fields func(tables SqlTables) []SqlField) SqlQuery {
	sq.excludeLoad(lazyLoadGroupFields)
	sq.groupFields = nil
	sq.groupFieldFun = fields
	return sq
}

func (sq *TSqlQuery) Having(condition func(tables SqlTables) SqlCondition) SqlQuery {
	sq.excludeLoad(lazyLoadHavingCond)
	sq.having = nil
	sq.havingFun = condition
	return sq
}

func (sq *TSqlQuery) Order(fields func(tables SqlTables) []SqlOrderField) SqlQuery {
	sq.orderFields = nil
	sq.excludeLoad(lazyLoadOrderFields)
	sq.orderFieldFun = fields
	return sq
}

func (sq *TSqlQuery) Limit(rows, offset int) SqlQuery {
	sq.limitRows = rows
	sq.limitOffset = offset
	sq.pageSize = rows
	sq.pageIndex = 0
	return sq
}

func (sq *TSqlQuery) Page(size, index int) SqlQuery {
	if size < 1 {
		size = 10
	}
	if index < 1 {
		index = 1
	}
	sq.limitOffset = (index - 1) * size
	sq.limitRows = size
	sq.pageSize = size
	return sq
}

func (sq *TSqlQuery) Alias(alias string) SqlTable {
	return NewSqlQueryTable(sq.This().(SqlQuery), alias)
}
