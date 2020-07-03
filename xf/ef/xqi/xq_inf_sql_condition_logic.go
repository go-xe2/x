package xqi

type SqlConditionLogic int

const (
	SqlConditionEqLogic SqlConditionLogic = iota
	SqlConditionNeqLogic
	SqlConditionGtLogic
	SqlConditionGteLogic
	SqlConditionLtLogic
	SqlConditionLteLogic
	SqlConditionLkLogic
	SqlConditionNlkLogic
	SqlConditionInLogic
	SqlConditionNinLogic
	// 逻辑与 and
	SqlConditionAndLogic
	// 逻辑颧 or
	SqlConditionOrLogic
	// 逻辑异或 xor = and (not value)
	SqlConditionXorLogic
)

var SqlConditionLogicNames = map[SqlConditionLogic]string{
	SqlConditionEqLogic:  " = ",
	SqlConditionNeqLogic: " <> ",
	SqlConditionGtLogic:  " > ",
	SqlConditionGteLogic: " >= ",
	SqlConditionLtLogic:  " < ",
	SqlConditionLteLogic: " <= ",
	SqlConditionLkLogic:  " LIKE ",
	SqlConditionNlkLogic: " NOT LIKE ",
	SqlConditionInLogic:  " IN ",
	SqlConditionNinLogic: " NOT IN ",
	SqlConditionAndLogic: " AND ",
	SqlConditionOrLogic:  " OR ",
	SqlConditionXorLogic: " XOR ",
}

func (sl SqlConditionLogic) Exp(rExpr string) string {
	if s, ok := SqlConditionLogicNames[sl]; ok {
		if sl == SqlConditionXorLogic {
			return " AND ( NOT " + rExpr + ")"
		}
		return s + rExpr
	}
	return " AND " + rExpr
}

func (sl SqlConditionLogic) String() string {
	if s, ok := SqlConditionLogicNames[sl]; ok {
		return s
	}
	return " AND "
}
