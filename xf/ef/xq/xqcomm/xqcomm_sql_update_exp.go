package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tSqlUpdateExp struct {
	updateTable  xqi.SqlTable
	updateWhere  xqi.SqlCondition
	updateFields []xqi.FieldValue
	joins        []xqi.SqlJoinItem
	joinInfo     []xqi.SqlJoin
	tables       xqi.SqlTables
}

var _ xqi.SqlUpdateExp = (*tSqlUpdateExp)(nil)

func NewSqlUpdateExp(table xqi.SqlTable) xqi.SqlUpdateExp {
	inst := &tSqlUpdateExp{}
	return inst.Table(table)
}

func (exp *tSqlUpdateExp) GetTable() xqi.SqlTable {
	return exp.updateTable
}

func (exp *tSqlUpdateExp) initJoinItems() {
	if exp.joinInfo != nil {
		return
	}
	count := len(exp.joins)
	exp.joinInfo = make([]xqi.SqlJoin, count)
	exp.tables = NewSqlTables()
	// 收集表达式所关联的表
	exp.tables.Add(exp.updateTable)
	for i := 0; i < count; i++ {
		exp.tables.Add(exp.joins[i].JoinTable())
	}
	for i := 0; i < count; i++ {
		joinItem := exp.joins[i]
		var condition xqi.SqlCondition
		if onFn := joinItem.LazyConditionFn(); onFn != nil {
			condition = onFn(joinItem.JoinTable(), exp.useTables())
		}
		exp.joinInfo[i] = NewSqlJoin(joinItem.JoinType(), joinItem.JoinTable(), condition)
	}
}

func (exp *tSqlUpdateExp) useTables() xqi.SqlTables {
	exp.initJoinItems()
	return exp.tables
}

func (exp *tSqlUpdateExp) GetJoins() []xqi.SqlJoin {
	exp.initJoinItems()
	return exp.joinInfo
}

func (exp *tSqlUpdateExp) GetWhere() xqi.SqlCondition {
	return exp.updateWhere
}

func (exp *tSqlUpdateExp) GetFields() []xqi.FieldValue {
	return exp.updateFields
}

func (exp *tSqlUpdateExp) This() interface{} {
	return exp
}

func (exp *tSqlUpdateExp) TokenType() xqi.SqlTokenType {
	return xqi.SqlUpdateTokenType
}

func (exp *tSqlUpdateExp) Compile(builder xdriveri.DbDriverSqlBuilder, cxt xqi.SqlCompileContext, unPrepare ...bool) xqi.SqlToken {
	if exp.updateTable == nil {
		panic(exception.New("没有设置要更新数据的表"))
	}
	if len(exp.updateFields) == 0 || exp.updateFields[0] == nil {
		panic(exception.New("没有设置要更新数据的字段"))
	}
	exp.initJoinItems()

	isPrepare := true
	if len(unPrepare) > 0 {
		isPrepare = !unPrepare[0]
	}

	cxt.Clear()
	cxt.Tables().Add(exp.updateTable)

	// 要更新的表
	szTable := ""
	// 更新字段
	szFieldList := ""
	// 关联关系
	szJoins := ""
	// 更新条件
	szWhere := ""

	result := NewSqlToken("", exp.TokenType())
	// 编译关联更新的join表达式
	joinCount := len(exp.joinInfo)

	// 编译待更新的表名
	if joinCount > 0 {
		cxt.PushState(xqi.SCPBuildUpdateTableWithFromState)
	} else {
		cxt.PushState(xqi.SCPBuildUpdateTableState)
	}
	if tk := exp.updateTable.Compile(builder, cxt, unPrepare...); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
		panic(exception.Newf("更新的表%s表达式错误", exp.updateTable.TableName()))
	} else {
		szTable = tk.Val()
	}
	cxt.PopState()

	// 编译更新字段列表
	fieldCount := len(exp.updateFields)
	for i := 0; i < fieldCount; i++ {
		fdExpr := exp.updateFields[i]
		if fdExpr == nil {
			continue
		}
		fd := exp.updateFields[i].Field()
		fv := exp.updateFields[i].Value()
		if fd == nil {
			// 忽略为nil的字段
			continue
		}
		szFieldName := ""
		if joinCount > 0 {
			cxt.PushState(xqi.SCPBuildUpdateFieldWithFromState)
		} else {
			cxt.PushState(xqi.SCPBuildUpdateFieldState)
		}
		if tk := fd.Compile(builder, cxt, unPrepare...); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
			panic(exception.Newf("表达式中更新的第%d个字段%s表达式错误,编译失败", i, fd.FieldName()))
		} else {
			szFieldName = tk.Val()
		}
		cxt.PopState()

		szFieldValue := ""
		if c, ok := fv.(xqi.SqlCompiler); ok {
			if joinCount > 0 {
				cxt.PushState(xqi.SCPBuildUpdateFieldValueFromState)
			} else {
				cxt.PushState(xqi.SCPBuildUpdateFieldValueState)
			}
			if tk := c.Compile(builder, cxt, unPrepare...); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
				panic(exception.Newf("表达式中更新的第%d个字段%s的数据源表达式错误，编译失败", i, fd.FieldName()))
			} else {
				szFieldValue = tk.Val()
			}
			cxt.PopState()
		} else {
			if isPrepare {
				vn := cxt.MakeParamId()
				result.AddParam(vn, fv)
				cxt.AddParam(vn, fv)
				szFieldValue = builder.PlaceHolder(vn)
			} else {
				szFieldValue = builder.MakeRealValue(fv)
			}
		}
		if szFieldList != "" {
			szFieldList += ","
		}
		szFieldList += fmt.Sprintf("%s = %s", szFieldName, szFieldValue)
	}
	// 检查表达式完整性，所有引用的字段所属表是否已定义关联关系

	// 检查字段表有效性
	tables := cxt.Tables().All()
	for _, table := range tables {
		if !exp.useTables().HasTable(table.TableName()) {
			panic(exception.Newf("表达式缺少定义表%s的关联关系", table.TableName()))
		}
	}

	for i := 0; i < joinCount; i++ {
		item := exp.joinInfo[i]
		szOnCondition := ""
		if c := item.OnCondition(); c == nil {
			panic(exception.Newf("表达式中表%s的关系表达式未定义", item.JoinTable().TableName()))
		} else {
			cxt.PushState(xqi.SCPBuildUpdateJoinOnState)
			if tk := c.Compile(builder, cxt, unPrepare...); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
				panic(exception.Newf("表达式中表%s的关系表达式定义错误", item.JoinTable().TableName()))
			} else {
				szOnCondition = tk.Val()
			}
			cxt.PopState()
		}
		szJoinTable := ""
		cxt.PushState(xqi.SCPBuildUpdateJoinTableState)
		if tk := item.JoinTable().Compile(builder, cxt, unPrepare...); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
			panic(exception.Newf("表达式中表%s的定义错误", item.JoinTable().TableName()))
		} else {
			szJoinTable = tk.Val()
		}
		cxt.PopState()
		szJoins += fmt.Sprintf(" %s %s ON %s", item.JoinType().Exp(), szJoinTable, szOnCondition)
	}

	// 编译where条件表达式
	if exp.updateWhere != nil {
		if joinCount > 0 {
			cxt.PushState(xqi.SCPBuildUpdateWhereFromState)
		} else {
			cxt.PushState(xqi.SCPBuildUpdateWhereState)
		}
		if tk := exp.updateWhere.Compile(builder, cxt, unPrepare...); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
			panic(exception.Newf("条件表达式错误"))
		} else {
			szWhere = tk.Val()
		}
		cxt.PopState()
	}

	// 从其他数据源更新的sql语句差别
	// mysql: update tableA a left join tableB b on a.user_id = b.user_id set a.name = 'new name', a.base_salary = a.base_salary + b.salary where b.user_id > 3 and a.name like '%张三%';
	// mssql: update a set a.name = 'new name', a.base_salary = a.base_salary + b.salary from tableA a left join tableB on a.user_id = b.user_id where b.user_id > 3 and a.name like '%张三%'

	szExp := builder.BuildUpdate(szTable, szFieldList, szJoins, szWhere)
	return result.SetVal(szExp)
}

func (exp *tSqlUpdateExp) Table(table xqi.SqlTable) xqi.SqlUpdateExp {
	exp.updateTable = table
	return exp
}

func (exp *tSqlUpdateExp) Set(fields ...xqi.FieldValue) xqi.SqlUpdateExp {
	if len(fields) == 0 {
		return exp
	}
	exp.updateFields = fields
	return exp
}

func (exp *tSqlUpdateExp) Where(where xqi.SqlCondition) xqi.SqlUpdateExp {
	exp.updateWhere = where
	return exp
}

func (exp *tSqlUpdateExp) Join(joinType xqi.SqlJoinType, joinTable xqi.SqlTable, on func(join xqi.SqlTable, others xqi.SqlTables) xqi.SqlCondition) xqi.SqlUpdateExp {
	item := NewSqlJoinItem(joinType, joinTable, on)
	exp.joins = append(exp.joins, item)
	return exp
}

func (exp *tSqlUpdateExp) InnerJoin(table xqi.SqlTable, on func(join xqi.SqlTable, others xqi.SqlTables) xqi.SqlCondition) xqi.SqlUpdateExp {
	return exp.Join(xqi.SqlInnerJoinType, table, on)
}

func (exp *tSqlUpdateExp) LeftJoin(table xqi.SqlTable, on func(join xqi.SqlTable, others xqi.SqlTables) xqi.SqlCondition) xqi.SqlUpdateExp {
	return exp.Join(xqi.SqlLeftJoinType, table, on)
}

func (exp *tSqlUpdateExp) RightJoin(table xqi.SqlTable, on func(join xqi.SqlTable, others xqi.SqlTables) xqi.SqlCondition) xqi.SqlUpdateExp {
	return exp.Join(xqi.SqlRightJoinType, table, on)
}

func (exp *tSqlUpdateExp) CrossJoin(table xqi.SqlTable, on func(join xqi.SqlTable, others xqi.SqlTables) xqi.SqlCondition) xqi.SqlUpdateExp {
	return exp.Join(xqi.SqlCrossJoinType, table, on)
}
