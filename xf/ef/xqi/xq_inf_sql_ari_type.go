package xqi

type SqlAriType int

const (
	// 加运算 +
	SqlAriPlusType SqlAriType = iota
	// 减运算 -
	SqlAriSubType
	// 乘运算 *
	SqlAriMulType
	// 除运算 /
	SqlAriDivType
)

func (sat SqlAriType) Val() string {
	switch sat {
	case SqlAriPlusType:
		return " + "
	case SqlAriSubType:
		return " - "
	case SqlAriMulType:
		return " * "
	case SqlAriDivType:
		return " / "
	}
	return ""
}
