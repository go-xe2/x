package xentity

import "github.com/go-xe2/x/xf/ef/xqi"

type tEntitySelectBuilder struct {
	sel *tEntitySelect
}

var _ xqi.EntitySelectBuilder = (*tEntitySelectBuilder)(nil)

func newEntitySelectBuilder(entSelect *tEntitySelect) *tEntitySelectBuilder {
	return &tEntitySelectBuilder{
		sel: entSelect,
	}
}

func (esb *tEntitySelectBuilder) Sql() (sql string, vars []interface{}) {
	return esb.sel.query.Sql()
}

func (esb *tEntitySelectBuilder) Convert() xqi.EntitySelectConvert {
	return esb.sel
}
