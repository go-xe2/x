package xquery

import "github.com/go-xe2/x/xf/ef/xqi"

var _ xqi.QueryInfo = (*tQueryExp)(nil)

func (qe *tQueryExp) GetTable() xqi.SqlTable {
	if qe.com == nil {
		return nil
	}
	return qe.com.GetFromTable()
}

func (qe *tQueryExp) GetSelectFields() []xqi.SqlField {
	if qe.com == nil {
		return nil
	}
	return qe.com.GetQueryFields()
}

func (qe *tQueryExp) GetWhere() xqi.SqlCondition {
	if qe.com == nil {
		return nil
	}
	return qe.com.GetWhere()
}

func (qe *tQueryExp) GetHaving() xqi.SqlCondition {
	if qe.com == nil {
		return nil
	}
	return qe.com.GetHaving()
}

func (qe *tQueryExp) GetOrders() []xqi.SqlOrderField {
	if qe.com == nil {
		return nil
	}
	return qe.com.GetOrderFields()
}

func (qe *tQueryExp) GetGroups() []xqi.SqlField {
	if qe.com == nil {
		return nil
	}
	return qe.com.GetGroupFields()
}

func (qe *tQueryExp) GetJoins() []xqi.SqlJoin {
	if qe.com == nil {
		return nil
	}
	return qe.com.GetJoins()
}

func (qe *tQueryExp) UseTables() []xqi.SqlTable {
	result := make([]xqi.SqlTable, 1)
	result[0] = qe.GetTable()
	joins := qe.GetJoins()
	for _, item := range joins {
		result = append(result, item.JoinTable())
	}
	return result
}
