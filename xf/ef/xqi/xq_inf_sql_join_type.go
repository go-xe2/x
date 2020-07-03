package xqi

type SqlJoinType int

const (
	SqlInnerJoinType SqlJoinType = iota
	SqlLeftJoinType
	SqlRightJoinType
	SqlCrossJoinType
)

var SqlJoinTypeExps = map[SqlJoinType]string{
	SqlInnerJoinType: " JOIN ",
	SqlLeftJoinType:  " LEFT JOIN ",
	SqlRightJoinType: " RIGHT JOIN ",
	SqlCrossJoinType: " CROSS JOIN ",
}

func (sjt SqlJoinType) Exp() string {
	if s, ok := SqlJoinTypeExps[sjt]; ok {
		return s
	}
	return " JOIN "
}

func (sjt SqlJoinType) String() string {
	return sjt.Exp()
}
