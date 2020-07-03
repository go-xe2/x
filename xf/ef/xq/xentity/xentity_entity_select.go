package xentity

import (
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type tEntitySelect struct {
	entity        *TEntity
	fieldNameMaps map[string]string
	query         Query
	builder       *tEntitySelectBuilder
	list          *tEntitySelectList
	page          *tEntitySelectPage
	single        *tEntitySelectSingle
}

var _ EntitySelect = (*tEntitySelect)(nil)

func newEntitySelect(ent *TEntity, query Query, fieldNameMaps map[string]string) *tEntitySelect {
	return &tEntitySelect{
		entity:        ent,
		fieldNameMaps: fieldNameMaps,
		query:         query,
	}
}

func (eqs *tEntitySelect) Top(count int) EntitySelect {
	eqs.query = eqs.query.Limit(count, 0)
	return eqs
}

func (eqs *tEntitySelect) Limit(rows int, offset ...int) EntitySelect {
	eqs.query = eqs.query.Limit(rows, offset...)
	return eqs
}

func (eqs *tEntitySelect) Where(where ...SqlCondition) EntitySelect {
	eqs.query = eqs.query.Where(func(w SqlCondition, tables SqlTables) SqlCondition {
		w = w.And((func(items []SqlCondition) []SqlLogic {
			result := make([]SqlLogic, len(items))
			for i, v := range items {
				result[i] = v
			}
			return result
		})(where)...)
		return w
	})
	return eqs
}

func (eqs *tEntitySelect) Having(having ...SqlCondition) EntitySelect {
	eqs.query = eqs.query.Having(func(w SqlCondition, tables SqlTables) SqlCondition {
		w.And((func(items []SqlCondition) []SqlLogic {
			result := make([]SqlLogic, len(items))
			for i, v := range items {
				result[i] = v
			}
			return result
		})(having)...)
		return w
	})
	return eqs
}

func (eqs *tEntitySelect) Order(fields ...SqlOrderField) EntitySelect {
	eqs.query = eqs.query.Order(func(tables SqlTables) []SqlOrderField {
		return fields
	})
	return eqs
}

func (eqs *tEntitySelect) Group(fields ...SqlField) EntitySelect {
	eqs.query = eqs.query.Group(func(tables SqlTables) []SqlField {
		return fields
	})
	return eqs
}

func (eqs *tEntitySelect) Alias(alias string) SqlTable {
	return eqs.query.Alias(alias)
}

func (eqs *tEntitySelect) Page(index, size int) EntitySelectPage {
	if eqs.page == nil {
		eqs.page = newEntitySelectPage(eqs, index, size)
	}
	return eqs.page
}

// 返回数据类型相关方法
func (eqs *tEntitySelect) List() EntitySelectList {
	if eqs.list == nil {
		eqs.list = newEntitySelectList(eqs)
	}
	return eqs.list
}

func (eqs *tEntitySelect) First() EntitySelectSingle {
	if eqs.single == nil {
		eqs.single = newEntitySelectSingle(eqs)
	}
	return eqs.single
}

// 生成sql脚本相关
func (eqs *tEntitySelect) Build() EntitySelectBuilder {
	if eqs.builder == nil {
		eqs.builder = newEntitySelectBuilder(eqs)
	}
	return eqs.builder
}
