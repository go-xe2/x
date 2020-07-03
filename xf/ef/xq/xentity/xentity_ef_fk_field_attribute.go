package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type entityForeignFieldAttribute struct {
	maps       map[string]interface{}
	foreignKey string
	fieldAlias string
	rule       string
	formatter  string
}

var _ xqi.EntityForeignFieldAttribute = (*entityForeignFieldAttribute)(nil)

func newEntityForeignFieldAttribute(foreignKey, fieldAlias, rule string, formatter string) xqi.EntityForeignFieldAttribute {
	return &entityForeignFieldAttribute{
		foreignKey: foreignKey,
		fieldAlias: fieldAlias,
		rule:       rule,
		formatter:  formatter,
	}
}

func (fat *entityForeignFieldAttribute) AttrName() string {
	return "foreign"
}

func (fat *entityForeignFieldAttribute) ForeignKey() string {
	return fat.foreignKey
}

func (fat *entityForeignFieldAttribute) FieldAlias() string {
	return fat.fieldAlias
}

func (fat *entityForeignFieldAttribute) Rule() string {
	return fat.rule
}

func (fat *entityForeignFieldAttribute) Formatter() string {
	return fat.formatter
}

func (fat *entityForeignFieldAttribute) Map() map[string]interface{} {
	if fat.maps == nil {
		fat.maps = make(map[string]interface{})
		fat.maps["fk"] = fat.foreignKey
		fat.maps["alias"] = fat.fieldAlias
		fat.maps["rule"] = fat.rule
	}
	return fat.maps
}

func (fat *entityForeignFieldAttribute) FromMap(mp map[string]interface{}) {
	if mp != nil {
		if v, ok := mp["fk"]; ok {
			fat.foreignKey = t.String(v)
		}
		if v, ok := mp["alias"]; ok {
			fat.fieldAlias = t.String(v)
		}
		if v, ok := mp["rule"]; ok {
			fat.rule = t.String(v)
		}
	}
}

func (fat *entityForeignFieldAttribute) FromTagStr(szTag string) {
	mp := xstring.ParseKeyValue(szTag)
	if v, ok := mp["primary"]; ok {
		fat.foreignKey = t.String(v)
	}
	if v, ok := mp["alias"]; ok {
		fat.fieldAlias = t.String(v)
	}
	if v, ok := mp["rule"]; ok {
		fat.fieldAlias = v
	}
}

func (fat *entityForeignFieldAttribute) Tag() string {
	return fmt.Sprintf("fk:%v; alias:%s; rule:%s", fat.foreignKey, fat.fieldAlias, fat.rule)
}

func (fat *entityForeignFieldAttribute) This() interface{} {
	return fat
}

func (fat *entityForeignFieldAttribute) String() string {
	return fat.Tag()
}
