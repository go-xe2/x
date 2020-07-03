package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/ef/xqi"
)

// 实体字段属性
type tEntityFieldAttribute struct {
	maps map[string]interface{}
	// 是否主键
	isPrimary bool
	// 数据库字段名
	field string
	alias string
	// 字段关联的实体名称, 非关联字段为空
	// 字段规则，用于查询
	rule string
	// 数据格式器名称, 格式为: 格式化器名称:参数1,参数2,参数3
	formatter string
}

var _ xqi.EntityFieldAttribute = (*tEntityFieldAttribute)(nil)

func NewFieldAttributeDef() xqi.EntityFieldAttribute {
	return &tEntityFieldAttribute{}
}

func NewFieldAttribute(fieldName string, fieldAlias string, rule string, formatter string, isPrimary ...bool) xqi.EntityFieldAttribute {
	b := false
	if len(isPrimary) > 0 {
		b = isPrimary[0]
	}
	return &tEntityFieldAttribute{
		field:     fieldName,
		alias:     fieldAlias,
		rule:      rule,
		isPrimary: b,
		formatter: formatter,
	}
}

func (et *tEntityFieldAttribute) AttrName() string {
	return "field"
}

func (et *tEntityFieldAttribute) Map() map[string]interface{} {
	if et.maps == nil {
		et.maps = make(map[string]interface{})
		et.maps["primary"] = et.IsPrimary()
		et.maps["field"] = et.FieldName()
		et.maps["rule"] = et.Rule()
		et.maps["alias"] = et.FieldAlias()
		et.maps["format"] = et.formatter
	}
	return et.maps
}

func (et *tEntityFieldAttribute) FromMap(mp map[string]interface{}) {
	if mp != nil {
		if v, ok := mp["primary"]; ok {
			et.isPrimary = t.Bool(v)
		}
		if v, ok := mp["field"]; ok {
			et.field = t.String(v)
		}
		if v, ok := mp["rule"]; ok {
			et.rule = t.String(v)
		}
		if v, ok := mp["alias"]; ok {
			et.alias = t.String(v)
		}
		if v, ok := mp["format"]; ok {
			et.formatter = t.String(v)
		}
	}
}

func (et *tEntityFieldAttribute) FromTagStr(szTag string) {
	mp := xstring.ParseKeyValue(szTag)
	if v, ok := mp["primary"]; ok {
		et.isPrimary = t.Bool(v)
	}
	if v, ok := mp["field"]; ok {
		et.field = v
	}
	if v, ok := mp["rule"]; ok {
		et.rule = v
	}
	if v, ok := mp["alias"]; ok {
		et.alias = v
	}
	if v, ok := mp["format"]; ok {
		et.formatter = v
	}
}

func (et *tEntityFieldAttribute) Tag() string {
	return fmt.Sprintf("primary:%v;field:%s; alias:%s;rule:%s;format:%s", et.IsPrimary(), et.FieldName(), et.FieldAlias(), et.Rule(), et.formatter)
}

func (et *tEntityFieldAttribute) This() interface{} {
	return et
}

func (et *tEntityFieldAttribute) String() string {
	return et.Tag()
}

func (et *tEntityFieldAttribute) IsPrimary() bool {
	return et.isPrimary
}

func (et *tEntityFieldAttribute) FieldName() string {
	return et.field
}

func (et *tEntityFieldAttribute) FieldAlias() string {
	return et.alias
}

func (et *tEntityFieldAttribute) Rule() string {
	return et.rule
}

func (et *tEntityFieldAttribute) Formatter() string {
	return et.formatter
}
