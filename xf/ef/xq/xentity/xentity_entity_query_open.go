package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xq"
	"github.com/go-xe2/x/xf/ef/xq/sql"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type tEntityQueryOpen struct {
	entity        *TEntity
	query         Query
	fieldNameMaps map[string]string
}

var _ EntityQueryOpen = (*tEntityQueryOpen)(nil)

func newEntityQueryOpen(entity *TEntity, query Query, fieldNameMaps map[string]string) EntityQueryOpen {
	return &tEntityQueryOpen{
		entity:        entity,
		query:         query,
		fieldNameMaps: fieldNameMaps,
	}
}

func (enr *tEntityQueryOpen) All() (int, error) {
	// 初始化字段序号
	for i := 0; i < enr.entity.FieldCount(); i++ {
		if fIdx, ok := enr.entity.Field(i).(EntFieldIndex); ok {
			fIdx.SetIndex(-1)
		}
	}
	var err error
	enr.entity.dataset, err = enr.query.Dataset()
	enr.setLastSql(enr.query.DB().LastSql())
	if err != nil {
		return 0, err
	}
	// 设置字段序号
	qryInfo := enr.query.Info()
	fields := qryInfo.GetSelectFields()
	entFields := enr.entity.fieldMaps
	for i, field := range fields {
		if entFd, ok := field.This().(EntField); ok && entFd.Table() == enr.entity.This() {
			fd, ok := entFields[entFd.DefineName()].(EntFieldIndex)
			if ok {
				fd.SetIndex(i)
			}
		} else if fd, ok := entFields[field.AliasName()]; ok {
			if fIdx, idxOk := fd.This().(EntFieldIndex); idxOk {
				fIdx.SetIndex(i)
			}
		}
	}
	return enr.entity.dataset.RowCount(), nil
}

// 分页获取数据
func (enr *tEntityQueryOpen) Page(pageIndex int, pageSize int) (pageInfo QueryPageInfo, err error) {
	CheckPageIndex(&pageIndex)
	CheckPageSize(&pageSize)
	enr.query = enr.query.Page(pageSize, pageIndex)
	qryInfo := enr.query.Info()
	if qryInfo == nil {
		return nil, exception.Newf("查询表达式不完整")
	}
	// 统计记录数，获取查询条件,等相关信息
	tableExpr := qryInfo.GetTable()
	whereExpr := qryInfo.GetWhere()
	joinExpr := qryInfo.GetJoins()
	totalQry := xq.QueryByDb(enr.query.DB()).Fields(func(SqlTables) []SqlField {
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
		enr.setLastSql(totalQry.DB().LastSql())
		return nil, err
	}
	enr.setLastSql(totalQry.DB().LastSql())
	pageCount := totalRow / pageSize
	if totalRow%pageSize > 0 {
		pageCount += 1
	}
	if _, err := enr.All(); err != nil {
		return nil, err
	}
	return NewQueryPageInfo(pageIndex, pageSize, pageCount, totalRow), nil
}

func (enr *tEntityQueryOpen) Top(count int) (int, error) {
	enr.query = enr.query.Limit(count)
	return enr.All()
}

func (enr *tEntityQueryOpen) setLastSql(sql string) {
	enr.entity.lastSql = sql
}

func (enr *tEntityQueryOpen) Where(where ...SqlCondition) EntityQueryOpen {
	if len(where) > 0 {
		enr.query = enr.query.Where(func(w SqlCondition, tables SqlTables) SqlCondition {
			w.And((func(items []SqlCondition) []SqlLogic {
				result := make([]SqlLogic, len(items))
				for i, v := range items {
					result[i] = v
				}
				return result
			})(where)...)
			return w
		})
	}
	return enr
}

func (enr *tEntityQueryOpen) Having(having ...SqlCondition) EntityQueryOpen {
	if len(having) > 0 {
		enr.query = enr.query.Having(func(w SqlCondition, tables SqlTables) SqlCondition {
			w.And((func(items []SqlCondition) []SqlLogic {
				result := make([]SqlLogic, len(items))
				for i, v := range items {
					result[i] = v
				}
				return result
			})(having)...)
			return w
		})
	}
	return enr
}

func (enr *tEntityQueryOpen) Sql() (sql string, vars []interface{}) {
	return enr.query.Sql()
}

func (enr *tEntityQueryOpen) Order(fields ...SqlOrderField) EntityQueryOpen {
	enr.query = enr.query.Order(func(tables SqlTables) []SqlOrderField {
		return fields
	})
	return enr
}

func (enr *tEntityQueryOpen) Group(fields ...SqlField) EntityQueryOpen {
	enr.query = enr.query.Group(func(tables SqlTables) []SqlField {
		return fields
	})
	return enr
}

func (enr *tEntityQueryOpen) Limit(rows int, offset ...int) EntityQueryOpen {
	enr.query = enr.query.Limit(rows, offset...)
	return enr
}
