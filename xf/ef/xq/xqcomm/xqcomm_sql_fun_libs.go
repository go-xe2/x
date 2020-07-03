package xqcomm

import . "github.com/go-xe2/x/xf/ef/xqi"

func SqlFunIfNull(v1 interface{}, v2 interface{}) SqlField {
	fun := NewSqlFunIfNull(v1, v2)
	return NewSqlField(fun, "")
}

func SqlFunIsNull(val interface{}) SqlField {
	fun := NewSqlFunIsNull(val)
	return NewSqlField(fun, "")
}

func SqlFunCase(exp SqlCompiler) SqlCase {
	var fun *TSqlFunCase
	if exp == nil {
		fun = NewSqlFunCase()
	} else {
		fun = NewSqlFunCase(exp)
	}
	return fun
}

func SqlFunIf(exp, v1, v2 interface{}) SqlField {
	fun := NewSqlFunIf(exp, v1, v2)
	return NewSqlField(fun, "")
}

func SqlFunCast(exp interface{}, asType DbType) SqlField {
	fun := NewSqlFunCast(exp, asType)
	return NewSqlField(fun, "")
}

// ***聚合函数****//

func SqlFunCount(exp SqlCompiler) SqlField {
	fun := newAggregate(SFCount, exp)
	return NewSqlField(fun, "")
}

func SqlFunMax(exp SqlCompiler) SqlField {
	fun := newAggregate(SFMax, exp)
	return NewSqlField(fun, "")
}

func SqlFunMin(exp SqlCompiler) SqlField {
	fun := newAggregate(SFMin, exp)
	return NewSqlField(fun, "")
}

func SqlFunAvg(exp SqlCompiler) SqlField {
	return NewSqlField(newAggregate(SFAvg, exp), "")
}

func SqlFunSum(exp SqlCompiler) SqlField {
	return NewSqlField(newAggregate(SFSum, exp), "")
}
