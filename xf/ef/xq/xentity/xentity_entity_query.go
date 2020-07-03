package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xq"
	"github.com/go-xe2/x/xf/ef/xq/sql"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

var _ EntityQuery = (*TEntity)(nil)

func (ent *TEntity) Select(fields ...SqlField) EntitySelect {
	ent.buildRelation()

	fieldList := ent.fields
	if len(fields) > 0 {
		fieldList = fields
	}

	joinTables, fieldNameMaps := ent.buildJoinCondition(fieldList)

	query := xq.Query(ent.dbName...).Fields(func(tables SqlTables) []SqlField {
		return fieldList
	}).From(ent)

	if len(joinTables) > 0 {
		for _, item := range joinTables {
			if item.table == ent {
				// 不使用join连接本身表实体
				continue
			}
			switch item.joinType {
			case SqlInnerJoinType:
				query = query.Join(item.table, item.onCondition)
				break
			case SqlLeftJoinType:
				query = query.LeftJoin(item.table, item.onCondition)
				break
			case SqlRightJoinType:
				query = query.RightJoin(item.table, item.onCondition)
				break
			case SqlCrossJoinType:
				query = query.CrossJoin(item.table, item.onCondition)
				break
			}
		}
	}
	return newEntitySelect(ent, query, fieldNameMaps)
}

func (ent *TEntity) SelectR(rule string) EntitySelect {
	ent.buildRelation()
	fieldList := ent.getFieldListByRule(rule)
	return ent.Select(fieldList...)
}

// 获取第一条数据，如果不存在，如果无记录，则返回nil
func (ent *TEntity) First(where SqlCondition, orders []SqlOrderField, fields ...SqlField) (map[string]interface{}, error) {
	return ent.Select(fields...).Where(where).Order(orders...).First().Map()
}

func (ent *TEntity) Find(where SqlCondition, orders []SqlOrderField, fields ...SqlField) ([]map[string]interface{}, error) {
	return ent.Select(fields...).Where(where).Order(orders...).List().Rows()
}

func (ent *TEntity) GetById(id interface{}, fields ...SqlField) map[string]interface{} {
	if ent.KeyField() == nil {
		panic(exception.NewText("实体未定义主键字段"))
	}
	result, err := ent.First(sql.Where(ent.KeyField().Eq(id)), nil, fields...)
	if err != nil {
		panic(exception.Wrap(err, "获取数据出错"))
	}
	return result
}

// 获取指定数量的数据
func (ent *TEntity) GetTop(count int, where SqlCondition, orders []SqlOrderField, fields ...SqlField) []map[string]interface{} {
	rows, err := ent.Select(fields...).Where(where).Order(orders...).Top(count).List().Rows()
	if err != nil {
		panic(exception.Wrap(err, "查询出错"))
	}
	return rows
}

func (ent *TEntity) Exists(where func(where SqlCondition, this interface{}) SqlCondition) (bool, error) {
	count := 0
	_, err := xq.Query(ent.dbName...).Fields(func(tables SqlTables) []SqlField {
		return []SqlField{sql.Count("*")}
	}).From(ent).Where(func(w SqlCondition, tables SqlTables) SqlCondition {
		return where(w, ent.This())
	}).Visitor(func(row int, values ...interface{}) (i interface{}, b bool) {
		count = t.Int(values[0], 0)
		return count, true
	})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ent *TEntity) CheckExists(where ...SqlCondition) bool {
	b, err := ent.Exists(func(w SqlCondition, this interface{}) SqlCondition {
		w = w.And((func(items []SqlCondition) (result []SqlLogic) {
			result = make([]SqlLogic, len(items))
			for i, v := range items {
				result[i] = v
			}
			return
		})(where)...)
		return w
	})
	if err != nil {
		panic(exception.Wrap(err, "查询表达式错误"))
	}
	return b
}

func (ent *TEntity) Open(fields ...SqlField) EntityQueryOpen {
	ent.buildRelation()
	fieldList := ent.fields
	if len(fields) > 0 {
		fieldList = fields
	}
	joinTables, fieldNameMaps := ent.buildJoinCondition(fieldList)

	query := xq.Query(ent.dbName...).Fields(func(tables SqlTables) []SqlField {
		return fieldList
	}).From(ent)

	if len(joinTables) > 0 {
		for _, item := range joinTables {
			if item.table == ent {
				// 不使用join连接本身表实体
				continue
			}
			switch item.joinType {
			case SqlInnerJoinType:
				query = query.Join(item.table, item.onCondition)
				break
			case SqlLeftJoinType:
				query = query.LeftJoin(item.table, item.onCondition)
				break
			case SqlRightJoinType:
				query = query.RightJoin(item.table, item.onCondition)
				break
			case SqlCrossJoinType:
				query = query.CrossJoin(item.table, item.onCondition)
				break
			}
		}
	}
	return newEntityQueryOpen(ent, query, fieldNameMaps)
}

// 打开数据集，以实体定义字段方式访问
func (ent *TEntity) OpenR(rule string) EntityQueryOpen {
	fieldList := ent.getFieldListByRule(rule)
	return ent.Open(fieldList...)
}

func (ent *TEntity) OpenTop(count int, where SqlCondition, orders []SqlOrderField, fields ...SqlField) int {
	n, err := ent.Open(fields...).Where(where).Order(orders...).Top(count)
	if err != nil {
		panic(exception.Wrap(err, "打开数据出错"))
	}
	return n
}

func (ent *TEntity) OpenFirst(where SqlCondition, orders []SqlOrderField, fields ...SqlField) bool {
	n := ent.OpenTop(1, where, orders, fields...)
	return n == 1
}

func (ent *TEntity) OpenById(id interface{}, fields ...SqlField) bool {
	if ent.KeyField() == nil {
		panic(exception.NewText("实体未定义主键字段"))
	}
	return ent.OpenFirst(sql.Where(ent.KeyField().Eq(id)), nil, fields...)
}
