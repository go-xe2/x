package defaultDriver

import (
	"fmt"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
)

func (dr *TDbDefaultDriver) IfNull(v1 string, v2 string) string {
	return fmt.Sprintf("IFNULL(%s,%s)", v1, v2)
}

func (dr *TDbDefaultDriver) If(exp string, v1, v2 string) string {
	return fmt.Sprintf("IF(%s,%s,%s)", exp, v1, v2)
}

func (dr *TDbDefaultDriver) ISNULL(exp string) string {
	return fmt.Sprintf("ISNULL(%s)", exp)
}

func (dr *TDbDefaultDriver) CAST(exp string, ddType DbDataType, size int, decimal ...int) string {
	var asType string
	switch ddType {
	case DbDataInt, DbDataTinyint, DbDataSmallint, DbDataBigint:
		asType = "SIGNED"
		break
	case DbDataFloat, DbDataDouble, DbDataDecimal:
		if size > 0 && len(decimal) == 0 {
			asType = fmt.Sprintf("DECIMAL(%d,2)", size)
		} else if size == 0 && len(decimal) > 0 {
			asType = fmt.Sprintf("DECIMAL(%d,%d)", size+decimal[0], decimal[0])
		} else {
			asType = "DECIMAL"
		}
		break
	case DbDataDate:
		asType = "DATE"
		break
	case DbDataTime:
		asType = "TIME"
		break
	case DbDataDateTime:
		asType = "DATETIME"
		break
	case DbDataChar, DbDataVarchar, DbDataTinytext, DbDataText, DbDataLongText:
		if size <= 0 {
			asType = "CHAR"
		} else {
			asType = fmt.Sprintf("CHAR(%d)", size)
		}
		break
	case DbDataBlob, DbDataTinyblob, DbDataLongblob:
		asType = "CHAR"
		break
	case DbDataBinary, DbDataVarbinary:
		asType = "BINARY"
		break
	default:
		asType = "CHAR"
	}
	return fmt.Sprintf("CAST(%s as %s)", exp, asType)
}

func (dr *TDbDefaultDriver) Case(exp string, whenThen [][]string, elseValue ...string) string {
	if len(whenThen) == 0 || len(whenThen[0]) != 2 {
		return ""
	}
	s := "CASE " + exp
	for _, items := range whenThen {
		s += " WHEN " + items[0] + " THEN " + items[1]
	}
	if len(elseValue) > 0 && elseValue[0] != "" {
		s += " ELSE " + elseValue[0]
	}
	s += " END"
	return s
}
