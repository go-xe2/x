package xdvrlibs

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	. "github.com/go-xe2/x/xdb/xtableInf"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
)

func MakeSqlFieldName(builder DbDriverSqlBuilder, field DbField) string {
	s := ""
	if field.Table().TableAlias() != "" {
		s = field.Table().TableAlias() + "."
	}
	s += builder.QuotesName(field.FieldName())
	return s
}

func GetDateTimeStr(val interface{}) string {
	tv := t.XTime(val)
	return fmt.Sprintf("'%s'", tv.String())
}

func GetDateUnixStr(val interface{}) string {
	tv := t.XTime(val)
	return t.String(tv.Unix())
}
