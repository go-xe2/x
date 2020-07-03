package xquery

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	. "github.com/go-xe2/x/xf/ef/xqi"
	"testing"
)

type TableDemo struct {
	*xqcomm.TSqlTable
}

var _ SqlTable = &TableDemo{}

func NewDemoTable(name string, alias ...string) *TableDemo {
	inst := &TableDemo{}
	base := xqcomm.NewSqlTable(inst, name, alias...)
	inst.TSqlTable = base
	return inst
}

func TestNewQueryExp(t *testing.T) {
	tb1 := NewDemoTable("tableA", "A")
	tb2 := NewDemoTable("tableB", "B")
	tb1_userId := xqcomm.NewSqlTableField(nil, tb1, "user_id", "u_id")
	tb1_name := xqcomm.NewSqlTableField(nil, tb1, "name")
	tb1_sex := xqcomm.NewSqlTableField(nil, tb1, "sex")
	tb1_age := xqcomm.NewSqlTableField(nil, tb1, "age")
	tb1_base_salary := xqcomm.NewSqlTableField(nil, tb1, "base_salary")
	tb2_className := xqcomm.NewSqlTableField(nil, tb2, "className")
	tb2_project := xqcomm.NewSqlTableField(nil, tb2, "project")
	tb2_achievement := xqcomm.NewSqlTableField(nil, tb2, "achievement", "avc")
	tb2_userId := xqcomm.NewSqlTableField(nil, tb2, "user_id", "u_id")
	tb2_salary := xqcomm.NewSqlTableField(nil, tb2, "salary")

	qr2 := NewQueryExp(nil).Fields(func(tables SqlTables) []SqlField {
		return []SqlField{
			tb2_project,
		}
	}).From(tb2).Where(func(where SqlCondition, tables SqlTables) SqlCondition {
		return where.And(tb2_userId.Gt(10))
	})

	qry := NewQueryExp(nil).Fields(func(tables SqlTables) []SqlField {
		//a := tables.Table("A")
		//b := tables.Table("B")
		return []SqlField{
			tb1_userId,
			tb1_name.Alias("tb1_name"),
			tb1_sex,
			tb1_age,
			tb1_age.Plus(30).Sub(5).Alias("max_age"),
			tb1_base_salary.Plus(tb2_salary).Alias("user_salary"),
			tb2_className.Substr(0, 10).Alias("cls_name"),
			tb2_project,
			tb2_achievement,
			xqcomm.NewSqlField("test", "row_type"),
			xqcomm.SqlFunCase(tb1_name).When(1, "李三").When(2, "李四").When(3, "王五").End().Alias("name_ext"),
			tb1_sex.CaseGt(0).ThenElse("男", "女").Alias("sex_name"),
			tb1_age.Sub(6).Case().When(30, "青年").When(40, "中年").When(50, "老年").End().Alias("age_name"),
			tb1_sex.CaseGt(0).ThenElse("男", "女").Alias("sex1"),
			tb1_age.Cast(xqcomm.DTVarchar.Size(10)).Alias("age_name"),
		}
	}).From(tb1).LeftJoin(tb2, func(joinTable SqlTable, tables SqlTables, on SqlCondition) SqlCondition {
		//_ := tables.Table("A") // 测试
		return on.And(tb1_userId.Eq(tb2_userId))
	}).Where(func(where SqlCondition, tables SqlTables) SqlCondition {
		//_ := tables.Table("A") // 测试
		//_ := tables.Table("B") // 测试
		return where.And(xqcomm.NewSqlConditionItem(tb1_userId, xqcomm.NewSqlVar(nil, 3), SqlCompareGTEType),
			xqcomm.NewSqlConditionItem(tb2_project, xqcomm.NewSqlVar(nil, "math"), SqlCompareLKType)).Or(
			xqcomm.NewSqlConditionItem(tb1_sex, xqcomm.NewSqlVar(nil, "女"), SqlCompareEQType),
			tb1_age.In([]int{34, 55, 66}), tb2_project.In(qr2),
			tb2_project.SubEq(1, 3, "yuwen"),
			tb2_project.Substr(1, 4).Like("yuwen3"),
			tb1_sex.Not(),
			tb1_age.Plus(40).Gt(60),
		)
	})

	builder := xdriver.GetSqlBuilderByName("mysql")
	var params = xqcomm.NewSqlCompileContext()
	tk := qry.Compile(builder, params, false)
	fmt.Println("prepare make:\n", tk, ", params:", params)

	params1 := xqcomm.NewSqlCompileContext()
	tk = qry.Compile(builder, params1, true)
	fmt.Println("unPrepare make:\n", tk, ", params:", params1)
}
