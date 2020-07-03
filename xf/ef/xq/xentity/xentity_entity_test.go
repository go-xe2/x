package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	. "github.com/go-xe2/x/xf/ef/xqi"
	"testing"
)

type SalaryEntity struct {
	*TEntity
	UserId     EFInt    `ef:"@field(name=user_id);@dbType(type=int,increment=true,allowNull=false)"`
	BaseSalary EFFloat  `ef:"@field(name=base_salary);@dbType(type=decimal,size=18,decimal=2,allowNull=false,default=0)"`
	Salary     EFFloat  `ef:"@field(name=salary)"`
	ClassId    EFInt    `ef:"@field(name=class_id)"`
	DeductRate EFDouble `ef:"@field(name=deduct_rate)"`
}

var SalaryEntityClass = ClassOfEntity(func() Entity {
	return &SalaryEntity{}
}, []XqAttribute{MakeEntityAttr("salaryTable", "st")})

// 设置实体继承类实例, 供构架内部调用
func (de *SalaryEntity) Implement(supper interface{}) {
	if v, ok := supper.(*TEntity); ok {
		de.TEntity = v
	}
}

// 继承的父类
func (de *SalaryEntity) Supper() Entity {
	return de.TEntity
}

// 实体构造方法
func (de *SalaryEntity) Constructor(attrs []XqAttribute, inherited ...interface{}) interface{} {
	de.Supper().Constructor(attrs, inherited...)
	return de
}

type UserClassName struct {
	*TEntity
	ClassId EFInt    `ef:"@field(name=class_id,primary=true)"`
	Name    EFString `ef:"@field(name=class_name)"`
}

var UserClassNameClass = ClassOfEntity(func() Entity {
	return &UserClassName{}
}, []XqAttribute{MakeEntityAttr("userClass", "uc")})

// 设置实体继承类实例, 供构架内部调用
func (de *UserClassName) Implement(supper interface{}) {
	if v, ok := supper.(*TEntity); ok {
		de.TEntity = v
	}
}

// 继承的父类
func (de *UserClassName) Supper() Entity {
	return de.TEntity
}

// 实体构造方法
func (de *UserClassName) Constructor(attrs []XqAttribute, inherited ...interface{}) interface{} {
	de.Supper().Constructor(attrs, inherited...)
	return de
}

type DemoEntity struct {
	*TEntity
	UserId    EFInt     `ef:"@field(name=user_id,primary=true,rule=M|U|G|S);@dbType(type=int,increment=true,allowNull=false)"`
	Name      EFString  `ef:"@field(name=name);@valid(cate=U|M, rule='required|length:2,10', msg='产品ID不能为空|名称长度只能输入2到10个字符之间', op=I)"`
	Sex       EFBool    `ef:"@field(name=sex)"`
	Birthday  EFDate    `ef:"@field(name=birthday)"`
	Weight    EFInt     `ef:"@field(name=weight);@valid(cate = U|M, op=I|U, rule='required|min:100|max:10000',msg='重量必填|重量不能小于100|重量不能大小10000')"`
	Salary    EFForeign `ef:"@foreign(fk=fk_demo_salary, alias=salary);@dbType(type=decimal,size=18,decimal=2,allowNull=false,default=0)"`
	ClassName EFForeign `ef:"@foreign(fk=fk_salary_class_name, alias=class_name)"`
	SexName   EFExpr    `ef:"@field(alias=sex_name)"` // 定义表达式定义在constructor构造方法中
}

var DemoEntityClass = ClassOfEntity(func() Entity {
	var inst Entity = new(DemoEntity)
	return inst
}, []XqAttribute{MakeEntityAttr("demoTable", "dt")},

	// 以下注册外联字段类型
	// 定义外联字段Salary
	TypeLeftJoinField("fk_demo_salary", SalaryEntityClass, func(ent interface{}) SqlField {
		salaryEnt := ent.(*SalaryEntity)
		return salaryEnt.BaseSalary.Plus(salaryEnt.Salary) // BaseSalary + salary

	}, func(on SqlCondition, joinEnt interface{}, leftEntities SqlTables) SqlCondition {
		salaryEnt := joinEnt.(*SalaryEntity)
		demoEnt := leftEntities.Table("demoTable").This().(*DemoEntity)
		return on.And(salaryEnt.UserId.Eq(demoEnt.UserId))
	}),

	TypeLeftJoinField("fk_salary_class_name", UserClassNameClass, func(ent interface{}) SqlField {
		this := ent.(*UserClassName)
		return this.Name.Alias("class_name1")
	}, func(on SqlCondition, joinEnt interface{}, leftEntities SqlTables) SqlCondition {
		this := joinEnt.(*UserClassName)
		us := leftEntities.Table(SalaryEntityClass.TableName()).(*SalaryEntity)
		return on.And(this.ClassId.Eq(us.ClassId))
	}),
)

// 设置实体继承类实例, 供构架内部调用
func (de *DemoEntity) Implement(supper interface{}) {
	if v, ok := supper.(*TEntity); ok {
		de.TEntity = v
	}
}

// 继承的父类
func (de *DemoEntity) Supper() Entity {
	return de.TEntity
}

// 实体构造方法
func (de *DemoEntity) Constructor(attrs []XqAttribute, inherited ...interface{}) interface{} {
	de.Supper().Constructor(attrs, inherited...)
	de.SexName.Expression(func(ent Entity) SqlField {
		return de.Sex.Case().When(0, "未知").When(1, "男").When(2, "女").End()
	})
	return de
}

func (de *DemoEntity) String() string {
	return "DemoEntity"
}

type DemoEntityB struct {
	*DemoEntity
	Weight EFDouble `ef:"@field(name=weight)"`
}

func (deb *DemoEntityB) Implement(supper interface{}) {
	if v, ok := supper.(*DemoEntity); ok {
		deb.DemoEntity = v
	}
}

func (deb *DemoEntityB) Supper() Entity {
	return deb.DemoEntity
}

func (deb *DemoEntityB) Constructor(attrs []XqAttribute, inherited ...interface{}) interface{} {
	deb.Supper().Constructor(attrs, inherited...)
	return deb
}

var DemoEntityBClass = ClassOfEntity(func() Entity { return &DemoEntityB{} },
	[]XqAttribute{MakeEntityAttr("demoBTable", "dtb")})

func TestEntity(t *testing.T) {
	//var entCls = ClassOfEntity(func() Entity {
	//	return &baseEntity{}
	//}, nil)
	//fmt.Println("ent.TSqlTable type:", entCls)

	//fmt.Println("DemoEntityClass:", DemoEntityClass)
	//fmt.Println("DemoEntityClass.Annotations:", DemoEntityClass.Annotations())

	demo := DemoEntityClass.Create().(*DemoEntity)
	//fmt.Println("demo:", demo)
	//s, _ := xjson.Encode(demo)
	//fmt.Println("demo json:", string(s))
	//fmt.Println("demo.TableName:", demo.TableName())
	//fmt.Println("demo.TableAlias:", demo.TableAlias())
	//fmt.Println("demo.this type:", reflect.TypeOf(demo.This()))
	//fmt.Println("demo.fields:", demo.AllField())
	//
	//demoB := NewEntity(DemoEntityBClass).(*DemoEntityB)
	//fmt.Println("demoB:", demoB)
	//s, _ = xjson.Encode(demo)
	//fmt.Println("demoB json:", string(s))
	//fmt.Println("demoB.TableName:", demoB.TableName())
	//fmt.Println("demoB.TableAlias:", demoB.TableAlias())
	//fmt.Println("demoB.this type:", reflect.TypeOf(demoB.This()))

	qry := demo.Select(
		demo.UserId,
		demo.Name,
		//demo.Salary,
		//demo.ClassName,
		demo.SexName,
		demo.Sex).Where(xqcomm.NewSqlCondition().And(demo.UserId.Gt(0), demo.Name.Like("dddd")).And(demo.Sex.Not()).Or(demo.Salary.Gt(1000), demo.Birthday.Eq("1900-01-01"))).Order(
		demo.UserId.Desc(),
		demo.Name.Asc())

	sql, vars := qry.Build().Sql()
	fmt.Println("demo query sql:", sql, ", vars:", vars)

	// 获取元备注
	ann := demo.Salary.GetAnnotation(DbTypeAnnName)
	if ann != nil {
		fmt.Println("ann:", ann)
	}

	fmt.Println("field userId isPrimary:", demo.UserId.IsPrimary())
	fmt.Println("field userId.alias() isPrimary:", demo.UserId.Alias("u_id").(EntField).IsPrimary())
	fmt.Println("field userId.Set isPrimary:", demo.UserId.Set("123456").Field().(EntField).IsPrimary())

	dbAttribute := demo.UserId.GetAnnotation(DbTypeAnnName).(EntityDbTypeAttribute)
	fmt.Println("dbAttribute:", dbAttribute)

}

func TestGetValidRules(t *testing.T) {
	demo := DemoEntityClass.Create().(*DemoEntity)

	rules := demo.GetInsertRules("U")
	fmt.Println("insert rules:", xstring.Join(rules, "\n"))

	rules = demo.GetUpdateRules("M")
	fmt.Println("update rules:", xstring.Join(rules, "\n"))

}

func TestGetAllEntities(t *testing.T) {
	entities := GetAllEntities()
	fmt.Println("all entities:", entities)
}

func TestFormatTime(tg *testing.T) {
	var tr = t.XTime(0)
	fmt.Println("time:", tr.String())
}
