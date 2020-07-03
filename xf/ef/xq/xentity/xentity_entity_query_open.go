package xentity

import . "github.com/go-xe2/x/xf/ef/xqi"

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
