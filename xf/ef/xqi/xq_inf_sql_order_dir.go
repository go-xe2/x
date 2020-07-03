package xqi

type SqlOrderDirect int

const (
	SqlOrderAscDirect SqlOrderDirect = iota
	SqlOrderDescDirect
)

var SqlOrderDirectExps = map[SqlOrderDirect]string{
	SqlOrderAscDirect:  " ASC ",
	SqlOrderDescDirect: " DESC ",
}

func (sod SqlOrderDirect) Exp() string {
	if s, ok := SqlOrderDirectExps[sod]; ok {
		return s
	}
	return " ASC "
}
