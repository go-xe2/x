package xqi

type SqlCompareType int

const (
	// ==
	SqlCompareEQType SqlCompareType = iota
	// <>
	SqlCompareNEQType
	// >
	SqlCompareGTType
	// >=
	SqlCompareGTEType
	// <
	SqlCompareLTType
	// <=
	SqlCompareLTEType
	// like
	SqlCompareLKType
	// not like
	SqlCompareNLKType
	// in
	SqlCompareINType
	// not in
	SqlCompareNINType
	// not
	SqlCompareNTType
)

var SqlCompareTypeExps = map[SqlCompareType]string{
	SqlCompareEQType:  " = ",
	SqlCompareNEQType: " <> ",
	SqlCompareGTType:  " > ",
	SqlCompareGTEType: " >= ",
	SqlCompareLTType:  " < ",
	SqlCompareLTEType: " <= ",
	SqlCompareLKType:  " like ",
	SqlCompareNLKType: " not like ",
	SqlCompareINType:  " in ",
	SqlCompareNINType: " not in ",
	SqlCompareNTType:  " not ",
}

func (sct SqlCompareType) Val() string {
	if s, ok := SqlCompareTypeExps[sct]; ok {
		return s
	}
	return " = "
}

func (sct SqlCompareType) String() string {
	return sct.Val()
}

func (sct SqlCompareType) Exp(rExp string) string {
	if s, ok := SqlCompareTypeExps[sct]; ok {
		return s + rExp
	}
	return " = " + rExp
}
