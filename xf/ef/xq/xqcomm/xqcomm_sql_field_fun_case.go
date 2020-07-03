package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func (sf *TSqlField) Case() SqlCase {
	return NewSqlFunCase(sf.This().(SqlField))
}

// case when field == v then else end
func (sf *TSqlField) CaseEq(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(sf.Eq(v))
}

// case when field != v then else end
func (sf *TSqlField) CaseNeq(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(sf.Neq(v))
}

// case when field > v then else end
func (sf *TSqlField) CaseGt(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(sf.Gt(v))
}

// case when field >= v then else end
func (sf *TSqlField) CaseGte(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(sf.Gte(v))
}

// case when field < v then else end
func (sf *TSqlField) CaseLt(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(sf.Lt(v))
}

// case when field <= v then else end
func (sf *TSqlField) CaseLte(v interface{}) SqlCaseThenElse {
	return NewSqlFunCase(sf.Lte(v))
}
