package xentity

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
	"regexp"
)

type tEntForeignData struct {
	lookEnt   interface{}
	lookField xqi.SqlField
}

type TEFForeign struct {
	entity xqi.Entity
	*xqcomm.TSqlField
	data    *tEntForeignData
	joinEnt xqi.EntityClass // 关联的实体
	// entity为lookEnt类型的实体实例
	lookField   func(entity interface{}) xqi.SqlField
	onCondition func(on xqi.SqlCondition, lookEnt interface{}, leftTables xqi.SqlTables) xqi.SqlCondition
	joinType    xqi.SqlJoinType
	rule        string
	annotations map[string]interface{}
	alias       string
	foreignKey  string
	expr        string
	defineName  string
	// 查询时的序号
	index int
}

var _ xqi.EFForeign = (*TEFForeign)(nil)
var _ xqi.DSField = (*TEFForeign)(nil)

func newEFForeignField(entity xqi.Entity, defineName string, attrs []xqi.XqAttribute, annotations map[string]interface{}, params ...interface{}) xqi.EFForeign {
	var foreignAttr xqi.EntityForeignFieldAttribute
	for _, attr := range attrs {
		if tmp, ok := attr.(xqi.EntityForeignFieldAttribute); ok {
			foreignAttr = tmp
			break
		}
	}
	if foreignAttr == nil {
		panic(exception.Newf("外联查字段%s关未设置元注解@foreign配置属性", defineName))
	}
	foreignKeyDefine := getFieldDefine(foreignAttr.ForeignKey())
	if foreignKeyDefine == nil {
		panic(exception.Newf("外联查字段%s定义的外联字段类型%s不存在或未注册", defineName, foreignAttr.ForeignKey()))
	}
	alias := defineName
	foreignKey := foreignAttr.ForeignKey()
	if foreignAttr.FieldAlias() != "" {
		alias = foreignAttr.FieldAlias()
	}
	inst := &TEFForeign{
		entity:      entity,
		annotations: annotations,
		defineName:  defineName,
		joinEnt:     foreignKeyDefine.JoinEnt,
		joinType:    foreignKeyDefine.JoinType,
		lookField:   foreignKeyDefine.Look,
		onCondition: foreignKeyDefine.OnCondition,
		foreignKey:  foreignKey,
		alias:       alias,
		rule:        foreignAttr.Rule(),
	}
	base := xqcomm.NewSqlField(nil, alias, inst)
	inst.TSqlField = base
	if entity != nil {
		if fields, ok := entity.This().(xqi.SqlTableFields); ok {
			fields.AddField(inst)
		}
	}
	return inst
}

func (ef *TEFForeign) NewInstance(alias string, inherited ...interface{}) xqi.EntField {
	inst := &TEFForeign{
		entity:      ef.entity,
		annotations: ef.annotations,
		defineName:  ef.defineName,
		joinEnt:     ef.joinEnt,
		joinType:    ef.joinType,
		lookField:   ef.lookField,
		onCondition: ef.onCondition,
		foreignKey:  ef.foreignKey,
		alias:       alias,
		rule:        ef.rule,
	}
	inst.TSqlField = xqcomm.NewSqlField(nil, alias, inst)
	return inst
}

func (ef *TEFForeign) Supper() xqi.EntField {
	return nil
}

func (ef *TEFForeign) IsPrimary() bool {
	return false
}

func (ef *TEFForeign) IsForeign() bool {
	return true
}

func (ef *TEFForeign) Table() xqi.SqlTable {
	return ef.entity
}

func (ef *TEFForeign) Entity() xqi.Entity {
	return ef.entity
}

func (ef *TEFForeign) FieldName() string {
	return ef.defineName
}

func (ef *TEFForeign) Value() interface{} {
	if ef.index < 0 {
		panic(exception.Newf("查询结果中无%s字段", ef.AliasName()))
	}
	if ds, ok := ef.entity.This().(xqi.Dataset); ok {
		return ds.FieldValue(ef.index)
	}
	return nil
}

func (ef *TEFForeign) IsOpen() bool {
	return ef.index >= 0
}

func (ef *TEFForeign) TryGetVal() interface{} {
	if ef.index < 0 {
		return nil
	} else {
		return ef.Value()
	}
}

func (ef *TEFForeign) Set(val interface{}) xqi.FieldValue {
	// 不设置外联字段
	return nil
}

func (ef *TEFForeign) Rule() string {
	return ef.rule
}

func (ef *TEFForeign) ForeignKey() string {
	return ef.foreignKey
}

func (ef *TEFForeign) GetAnnotation(annName string) interface{} {
	if ef.annotations == nil {
		return nil
	}
	if v, ok := ef.annotations[annName]; ok {
		return v
	}
	return nil
}

func (ef *TEFForeign) AliasName() string {
	return ef.alias
}

func (ef *TEFForeign) Alias(alias string) xqi.SqlField {
	ef.alias = alias
	return ef
}

func (ef *TEFForeign) JoinType() xqi.SqlJoinType {
	return xqi.SqlLeftJoinType
}

func (ef *TEFForeign) JoinTable() xqi.SqlTable {
	ef.build()
	return ef.data.lookEnt.(xqi.Entity)
}

func (ef *TEFForeign) On() func(on xqi.SqlCondition, joinTable interface{}, leftTables xqi.SqlTables) xqi.SqlCondition {
	return ef.onCondition
}

func (ef *TEFForeign) LookField() xqi.SqlField {
	ef.build()
	return ef.data.lookField
}

func (ef *TEFForeign) DefineName() string {
	return ef.defineName
}

func (ef *TEFForeign) Desc() xqi.SqlOrderField {
	return xqcomm.NewSqlOrderField(ef, xqi.SqlOrderDescDirect)
}

func (ef *TEFForeign) Asc() xqi.SqlOrderField {
	return xqcomm.NewSqlOrderField(ef, xqi.SqlOrderAscDirect)
}

func (ef *TEFForeign) String() string {
	fdName := ef.FieldName()
	if ef.AliasName() != "" {
		fdName = ef.AliasName()
	}
	if ef.expr == "" {
		cxt := xqcomm.NewSqlCompileContext()
		builder := xdriver.GetSqlBuilderByName(cxt.Driver())
		cxt.PushState(xqi.SCPQrSelectFieldsState)
		token := ef.Compile(builder, cxt, true)
		cxt.PopState()
		ef.expr = token.Val()
	}
	return fmt.Sprintf("%s:[%s]", fdName, ef.expr)
}

func (ef *TEFForeign) TokenType() xqi.SqlTokenType {
	return xqi.SqlExpressTokenType
}

func (ef *TEFForeign) Compile(builder xdriveri.DbDriverSqlBuilder, cxt xqi.SqlCompileContext, unPrepare ...bool) xqi.SqlToken {
	field := ef.LookField()
	if field == nil {
		return xqcomm.EmptySqlToken
	}
	if tk := field.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != xqi.SqlEmptyTokenType {
		state := cxt.State()
		if state == xqi.SCPQrSelectWhereCondState || state == xqi.SCPQrSelectJoinCondState || state == xqi.SCPBuildUpdateWhereState ||
			state == xqi.SCPBuildUpdateWhereFromState {
			return tk
		}
		if tk.TType() == xqi.SqlFieldTokenType {
			if ef.alias != "" {
				reg := regexp.MustCompile(`^(.+?)(\s+as\s+|\s+)(.+?)$`)
				fdName := tk.Val()
				if reg.MatchString(fdName) {
					matchRows := reg.FindStringSubmatch(fdName)
					fdName = matchRows[1]
					//return tk
				}
				return xqcomm.NewSqlToken(fmt.Sprintf("%s %s", fdName, ef.alias), tk.TType())
			}
			return tk
		}
		if ef.alias != "" {
			return xqcomm.NewSqlToken(fmt.Sprintf("%s %s", tk.Val(), ef.alias), tk.TType())
		}
		return tk
	} else {
	}
	return xqcomm.EmptySqlToken
}

func (ef *TEFForeign) build() {
	if ef.data == nil {
		ef.data = &tEntForeignData{}
		ef.data.lookEnt = ef.joinEnt.Create()
		ef.data.lookField = ef.lookField(ef.data.lookEnt)
	}
}

// 实现EntFieldIndex接口

func (ef *TEFForeign) SetIndex(index int) {
	ef.index = index
}

func (ef *TEFForeign) GetIndex() int {
	return ef.index
}

func (ef *TEFForeign) FieldType() xqi.FieldDataType {
	return xqi.FDTUnknown
}

func (ef *TEFForeign) FieldIndex() int {
	return ef.index
}

// 更新时字段运算, 外联字段不支持更新，以下方法为实例EntField接口用
func (ef *TEFForeign) Inc(step ...int) xqi.FieldValue {
	return nil
}

// 字段自减
func (ef *TEFForeign) Dec(step ...int) xqi.FieldValue {
	return nil
}

// 字段自乘
func (ef *TEFForeign) UnaryMul(val interface{}) xqi.FieldValue {
	return nil
}

// 字段自乘
func (ef *TEFForeign) UnaryDiv(val interface{}) xqi.FieldValue {
	return nil
}

func (ef *TEFForeign) Formatter() string {
	return ""
}
