package xq

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xe2/x/xf/ef/xq/xdatabase"
	"testing"
)

type TableDemo struct {
	SqlBaseTable
}

var _ SqlTable = &TableDemo{}

func NewDemoTable(name string, alias ...string) *TableDemo {
	inst := &TableDemo{}
	base := NewBaseTable(inst, name, alias...)
	inst.SqlBaseTable = base
	return inst
}

var config = xdatabase.NewMysqlConfigFromMap(map[string]interface{}{
	"host":        "127.0.1",
	"port":        3306,
	"user":        "root",
	"password":    "123456",
	"db":          "xqdb",
	"maxOpenCons": 0,
	"maxIdleCons": 0,
	"driver":      "mysql",
	"charset":     "utf8",
	"parseTime":   true,
})

func init() {
	db := Database()
	if err := db.Connection(config); err != nil {
		panic(err)
	}
}

func TestQuery(t *testing.T) {

	tb1 := NewDemoTable("tableA", "A")
	tb2 := NewDemoTable("tableB", "B")
	tb1_userId := NewBaseField(nil, tb1, "user_id", "u_id")
	tb1_name := NewBaseField(nil, tb1, "name")
	tb1_sex := NewBaseField(nil, tb1, "sex")
	tb1_age := NewBaseField(nil, tb1, "age")
	tb1_base_salary := NewBaseField(nil, tb1, "base_salary")
	tb2_className := NewBaseField(nil, tb2, "className")
	tb2_project := NewBaseField(nil, tb2, "project")
	tb2_achievement := NewBaseField(nil, tb2, "achievement", "avc")
	tb2_userId := NewBaseField(nil, tb2, "user_id", "u_id")
	tb2_salary := NewBaseField(nil, tb2, "salary")

	//qr2 := Query().Fields(func(tables SqlTables) []SqlField {
	//	return []SqlField{
	//		tb2_project,
	//	}
	//}).From(tb2).Where(func(where SqlCondition, tables SqlTables) SqlCondition {
	//	return where.And(tb2_userId.Gt(10))
	//})

	qry := Query().Fields(func(tables SqlTables) []SqlField {
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
			NewField("test").Alias("row_type"),
			SFCase(tb1_name).When(1, "李三").When(2, "李四").When(3, "王五").End().Alias("name_ext"),
			tb1_sex.CaseGt(0).ThenElse("男", "女").Alias("sex_name"),
			tb1_age.Sub(6).Case().When(30, "青年").When(40, "中年").When(50, "老年").End().Alias("age_name"),
			tb1_sex.CaseGt(0).ThenElse("男", "女").Alias("sex1"),
			tb1_age.Cast(DTVarchar.Size(10)).Alias("age_name"),
			tb1_base_salary.Avg().Alias("avg_salary"),
		}
	}).From(tb1).LeftJoin(tb2, func(joinTable SqlTable, tables SqlTables, on SqlCondition) SqlCondition {
		//_ := tables.Table("A") // 测试
		return on.And(tb1_userId.Eq(tb2_userId))
	}).Where(func(where SqlCondition, tables SqlTables) SqlCondition {
		//_ := tables.Table("A") // 测试
		//_ := tables.Table("B") // 测试
		return where.And(
			Gte(tb1_userId, RealValue(0)),
			//Like(tb2_project, RealValue("math"))).Or(
			//Eq(tb1_sex, RealValue("女")),
			//tb1_age.In([]int{34, 35, 36}), tb2_project.In(qr2),
			//tb2_project.SubEq(1, 3, "yuwen"),
			//tb2_project.Substr(1, 4).Like("yuwen'3"),
			//tb1_sex.Not(),
			//tb1_age.Plus(40).Gt(60),
		)
	})

	//sql, vars := qry.Sql()
	//fmt.Println("===========>>sql:")
	//fmt.Print(sql)
	////s := sql
	////for len(s) > 0 {
	////	n, _ := fmt.Print(sql)
	////	s = s[n:]
	////}
	//fmt.Println("")
	//fmt.Println("===========>>vars:", vars)

	xml, err := qry.Xml()
	fmt.Println("xml:", xml.String(), ", err:", err)

	json, err := qry.Json()
	fmt.Println("json:", json.String(), ", err:", err)

	mp, err := qry.Rows()
	fmt.Println("map:", mp, ", err:", err)

}
