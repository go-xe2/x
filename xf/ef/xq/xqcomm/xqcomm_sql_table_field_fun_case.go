package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (stf *TSqlTableField) Case() SqlCase {
	return NewSqlFunCase(stf.This().(SqlTableField))
}

// case field == v
func (stf *TSqlTableField) CaseEq(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(stf.Eq(v))
}

// case field <> v
func (stf *TSqlTableField) CaseNeq(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(stf.Neq(v))
}

// case field > v
func (stf *TSqlTableField) CaseGt(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(stf.Gt(v))
}

// case field >= v
func (stf *TSqlTableField) CaseGte(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(stf.Gte(v))
}

// case field < v
func (stf *TSqlTableField) CaseLt(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(stf.Lt(v))
}

// case field <= v
func (stf *TSqlTableField) CaseLte(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(stf.Lte(v))
}
