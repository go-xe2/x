package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tEntityAttribute struct {
	name  string
	alias string
}

var _ xqi.EntityAttribute = (*tEntityAttribute)(nil)

func MakeEntityAttr(tableName string, alias ...string) xqi.EntityAttribute {
	s := ""
	if len(alias) > 0 {
		s = alias[0]
	}
	return &tEntityAttribute{
		name:  tableName,
		alias: s,
	}
}

func (tg *tEntityAttribute) AttrName() string {
	return "EntityAttribute"
}

func (tg *tEntityAttribute) TableName() string {
	return tg.name
}

func (tg *tEntityAttribute) TableAlias() string {
	return tg.alias
}

func (tg *tEntityAttribute) Map() map[string]interface{} {
	result := make(map[string]interface{})
	result["tableName"] = tg.name
	result["alias"] = tg.alias
	return result
}

func (tg *tEntityAttribute) FromMap(mp map[string]interface{}) {
	if mp != nil {
		if v, ok := mp["tableName"]; ok {
			tg.name = t.String(v)
		}
		if v, ok := mp["alias"]; ok {
			tg.alias = t.String(v)
		}
	}
}

func (tg *tEntityAttribute) FromTagStr(szTag string) {
	mp := xstring.ParseKeyValue(szTag)
	if v, ok := mp["tableName"]; ok {
		tg.name = v
	}
	if v, ok := mp["alias"]; ok {
		tg.alias = v
	}
}

func (tg *tEntityAttribute) Tag() string {
	return fmt.Sprintf("tableName:%v;alias:%s", tg.name, tg.alias)
}

func (tg *tEntityAttribute) This() interface{} {
	return tg
}

func (tg *tEntityAttribute) String() string {
	return tg.Tag()
}
