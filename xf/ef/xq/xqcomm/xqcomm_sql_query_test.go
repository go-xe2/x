package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xqi"
	"testing"
)

type TableDemo struct {
	*TSqlTable
}

var _ xqi.SqlTable = &TableDemo{}

var (
	tb1             = NewDemoTable("tableA", "A")
	tb2             = NewDemoTable("tableB", "B")
	tb1_userId      = NewSqlTableField(nil, tb1, "user_id", "u_id")
	tb1_name        = NewSqlTableField(nil, tb1, "name")
	tb1_sex         = NewSqlTableField(nil, tb1, "sex")
	tb1_age         = NewSqlTableField(nil, tb1, "age")
	tb1_base_salary = NewSqlTableField(nil, tb1, "base_salary")
	tb2_className   = NewSqlTableField(nil, tb2, "className")
	tb2_project     = NewSqlTableField(nil, tb2, "project")
	tb2_achievement = NewSqlTableField(nil, tb2, "achievement", "avc")
	tb2_userId      = NewSqlTableField(nil, tb2, "user_id", "u_id")
	tb2_salary      = NewSqlTableField(nil, tb2, "salary")
)

func NewDemoTable(name string, alias ...string) *TableDemo {
	inst := &TableDemo{}
	base := NewSqlTable(inst, name, alias...)
	inst.TSqlTable = base
	return inst
}

func TestNewUpdateExp(t *testing.T) {
	// 简单插入
	cxt := NewSqlCompileContext()
	builder := xdriver.GetSqlBuilderByName(cxt.Driver())

	// 简单更新
	fmt.Println("简单更新:")
	expr := NewSqlUpdateExp(tb1)
	expr = expr.Set(NewFieldValue(tb1_name, "李四"), NewFieldValue(tb1_age, 3))
	expr = expr.Where(NewSqlCondition().And(tb1_userId.Gt(3)))
	tk1 := expr.Compile(builder, cxt, false)
	fmt.Println("tk1:", tk1.Val(), ", params:", cxt)

	fmt.Println("复杂更新:")
	// sql:
	// UPDATE `tableA` A   LEFT JOIN  `tableB` B ON A.`user_id`  =  B.`user_id` SET A.`base_salary` = A.`base_salary` + B.`salary`,A.`sex` = 1 WHERE A.`user_id`  >  1

	expr2 := NewSqlUpdateExp(tb1)
	// 设置字段
	expr2 = expr2.Set(NewFieldValue(tb1_base_salary, tb1_base_salary.Plus(tb2_salary)), NewFieldValue(tb1_sex, 1))
	// 设置关联关系
	expr2 = expr2.LeftJoin(tb2, func(join xqi.SqlTable, others xqi.SqlTables) xqi.SqlCondition {
		return NewSqlCondition().And(tb1_userId.Eq(tb2_userId))
	})
	// 设置条件
	expr2 = expr2.Where(NewSqlCondition().And(tb1_userId.Gt(1)))
	tk2 := expr2.Compile(builder, cxt, true)
	fmt.Println("tk2:", tk2.Val(), ", params:", cxt)
}

func TestNewInsertExp(t *testing.T) {
	// 简单插入
	cxt := NewSqlCompileContext()
	builder := xdriver.GetSqlBuilderByName(cxt.Driver())

	// 简单插入表达式
	fmt.Println("简单插入表达式：")
	expr := NewSqlInsertExp(tb1)
	expr = expr.Values(NewFieldValue(tb1_name, "张三"), NewFieldValue(tb1_sex, 1), NewFieldValue(tb1_base_salary, 1000))

	tk := expr.Compile(builder, cxt, false)
	fmt.Println("insert sql:", tk.Val(), ", params:", cxt)

	// 复杂插入表达式
	fmt.Println("复杂插入表达式:")
	fromExp := NewSqlQuery(tb1).Fields(func(tables xqi.SqlTables) []xqi.SqlField {
		return []xqi.SqlField{
			tb1_userId,
			tb1_name,
			tb1_sex,
			tb1_age,
			tb1_base_salary,
		}
	}).Where(func(tables xqi.SqlTables) xqi.SqlCondition {
		return NewSqlCondition().And(tb1_name.Like("张三"), tb1_userId.Gt(3))
	}).Alias("A1")

	expr2 := NewSqlInsertExp(tb1)
	expr2 = expr2.Values(
		NewFieldValue(tb1_name, fromExp.FieldByName("name")),
		NewFieldValue(tb1_sex, fromExp.FieldByName("sex")),
		NewFieldValue(tb1_age, fromExp.FieldByName("age")),
		NewFieldValue(tb1_base_salary, fromExp.FieldByName("base_salary")))
	expr2 = expr2.From(fromExp)

	cxt.Clear()
	tk2 := expr2.Compile(builder, cxt, false)
	fmt.Println("insert sql2:", tk2.Val(), ", params:", cxt)

}

func TestNewSqlQuery(t *testing.T) {

	fmt.Println("tb1.fields:", tb1.AllField())
	fmt.Println("tb2.fields:", tb2.AllField())

	qr2 := NewSqlQuery(tb2).Fields(func(tables xqi.SqlTables) []xqi.SqlField {
		return []xqi.SqlField{
			tb2_project,
		}
	}).Where(func(tables xqi.SqlTables) xqi.SqlCondition {
		return NewSqlCondition().And(tb2_userId.Gt(10))
	})

	qry := NewSqlQuery(tb1).Fields(func(tables xqi.SqlTables) []xqi.SqlField {
		//a := tables.Table("A")
		//b := tables.Table("B")
		return []xqi.SqlField{
			tb1_userId,
			tb1_name.Alias("tb1_name"),
			tb1_sex,
			tb1_age,
			tb1_age.Plus(30).Sub(5).Alias("max_age"),
			tb1_base_salary.Plus(tb2_salary).Alias("user_salary"),
			tb2_className.Substr(0, 10).Alias("cls_name"),
			tb2_project,
			tb2_achievement,
			NewSqlField(3, "row_type"),
			SqlFunCase(tb1_name).When(1, "李三").When(2, "李四").When(3, "王五").End().Alias("name_ext"),
			tb1_sex.CaseGt(0).ThenElse("男", "女").Alias("sex_name"),
			tb1_age.Sub(6).Case().When(30, "青年").When(40, "中年").When(50, "老年").End().Alias("age_name"),
			tb1_sex.CaseGt(0).ThenElse("男", "女").Alias("sex1"),
			tb1_age.Cast(DTVarchar.Size(10)).Alias("age_name"),
		}
	}).LeftJoin(tb2, func(joinTable xqi.SqlTable, tables xqi.SqlTables) xqi.SqlCondition {
		//_ := tables.Table("A") // 测试
		return NewSqlCondition().And(tb1_userId.Eq(tb2_userId))
	}).Where(func(tables xqi.SqlTables) xqi.SqlCondition {
		//_ := tables.Table("A") // 测试
		//_ := tables.Table("B") // 测试
		return NewSqlCondition().And(NewSqlConditionItem(tb1_userId, NewSqlVar(nil, 3), xqi.SqlCompareGTEType),
			NewSqlConditionItem(tb2_project, NewSqlVar(nil, "math"), xqi.SqlCompareLKType)).Or(
			NewSqlConditionItem(tb1_sex, NewSqlVar(nil, "女"), xqi.SqlCompareEQType),
			tb1_age.In([]int{34, 35, 36}), tb2_project.In(qr2),
			tb2_project.SubEq(1, 3, "yuwen"),
			tb2_project.Substr(1, 4).Like("yuwen3"),
			tb1_sex.Not(),
			tb1_age.Plus(40).Gt(60),
		)
	})

	builder := xdriver.GetSqlBuilderByName("mysql")
	var cxt = NewSqlCompileContext()
	tk := qry.Compile(builder, cxt, false)
	fmt.Println("prepare make:\n", tk, ", params:", cxt)

	cxt.Clear()
	tk = qry.Compile(builder, cxt, true)
	fmt.Println("unPrepare make:\n", tk, ", params:", cxt)

	fmt.Println("嵌套查询测试")
	qry2 := NewSqlQuery(qry.Alias("tb1")).Fields(func(tables xqi.SqlTables) []xqi.SqlField {
		tb1 := tables.Table("tb1")
		return []xqi.SqlField{
			tb1.FieldByName("tb1_name"),
			tb1.FieldByName("user_salary"),
		}
	})

	cxt.Clear()
	tk2 := qry2.Compile(builder, cxt, false)
	fmt.Println("tk2:", tk2.Val(), ", params:", cxt)

}
