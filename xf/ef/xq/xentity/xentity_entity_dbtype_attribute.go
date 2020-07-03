package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type EntityDbTypeAttribute interface {
	xqi.XqAttribute
	Type() xdriveri.DbDataType
	Size() int
	Decimal() int
	AllowNull() bool
	Default() interface{}
	Increment() bool
}

// 实体字段数据库定义属性, 用于创建数据库表
type tEntityDbTypeAttribute struct {
	// 数据类型
	dataType xdriveri.DbDataType
	// 字段长度,非必填
	dataSize int
	// 小数据点位数,只对DbDataDecimal数据类型有效, 非必填
	decimal int
	// 是否允许为空
	allowNull bool
	// 字段默认值
	defaultValue interface{}
	// 是否自动增长,只对int字段类型有效
	increment bool
	maps      map[string]interface{}
}

func (et *tEntityDbTypeAttribute) Type() xdriveri.DbDataType {
	return et.dataType
}

func (et *tEntityDbTypeAttribute) Size() int {
	return et.dataSize
}

func (et *tEntityDbTypeAttribute) Decimal() int {
	return et.decimal
}

func (et *tEntityDbTypeAttribute) AllowNull() bool {
	return et.allowNull
}

func (et *tEntityDbTypeAttribute) Default() interface{} {
	return et.defaultValue
}

func (et *tEntityDbTypeAttribute) Increment() bool {
	return et.increment
}

func (et *tEntityDbTypeAttribute) AttrName() string {
	return "db"
}

func (et *tEntityDbTypeAttribute) Map() map[string]interface{} {
	if et.maps == nil {
		et.maps = make(map[string]interface{})
		et.maps["type"] = et.dataType
		et.maps["size"] = et.dataSize
		et.maps["decimal"] = et.decimal
		et.maps["allowNull"] = et.allowNull
		et.maps["default"] = et.defaultValue
		et.maps["increment"] = et.increment
	}
	return et.maps
}

func (et *tEntityDbTypeAttribute) FromMap(mp map[string]interface{}) {
	if mp != nil {
		if v, ok := mp["type"]; ok {
			var dt xdriveri.DbDataType
			et.dataType = dt.Parse(t.String(v))
		}
		if v, ok := mp["size"]; ok {
			et.dataSize = t.Int(v)
		}
		if v, ok := mp["decimal"]; ok {
			et.decimal = t.Int(v)
		}
		if v, ok := mp["allowNull"]; ok {
			et.allowNull = t.Bool(v)
		}
		if v, ok := mp["default"]; ok {
			et.defaultValue = v
		}
		if v, ok := mp["increment"]; ok {
			et.increment = t.Bool(v)
		}
	}
}

func (et *tEntityDbTypeAttribute) FromTagStr(szTag string) {
	mp := xstring.ParseKeyValue(szTag)
	if v, ok := mp["type"]; ok {
		var dt xdriveri.DbDataType
		et.dataType = dt.Parse(t.String(v))
	}
	if v, ok := mp["size"]; ok {
		et.dataSize = t.Int(v)
	}
	if v, ok := mp["decimal"]; ok {
		et.decimal = t.Int(v)
	}
	if v, ok := mp["allowNull"]; ok {
		et.allowNull = t.Bool(v)
	}
	if v, ok := mp["increment"]; ok {
		et.increment = t.Bool(v)
	}
}

func (et *tEntityDbTypeAttribute) Tag() string {
	return fmt.Sprintf("type:%v;size:%d; decimal:%d;allowNull:%v, increment:%v", et.dataType, et.dataSize, et.decimal, et.allowNull, et.increment)
}

func (et *tEntityDbTypeAttribute) This() interface{} {
	return et
}

func (et *tEntityDbTypeAttribute) String() string {
	return et.Tag()
}
