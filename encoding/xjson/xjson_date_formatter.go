package xjson

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xtime"
	"reflect"
	"time"
)

type TJsonDateFormatter struct {
}

var dateFormatter JsonDataFormatter = (*TJsonDateFormatter)(nil)

func (jf *TJsonDateFormatter) Format(val interface{}) interface{} {
	if t, ok := val.(time.Time); ok {
		xt := xtime.New(t)
		return fmt.Sprintf("%s%s%s", JsonQuotes(), xt.String(), JsonQuotes())
	} else if t, ok := val.(*time.Time); ok {
		xt := xtime.New(*t)
		return fmt.Sprintf("%s%s%s", JsonQuotes(), xt.String(), JsonQuotes())
	}
	return fmt.Sprintf("%s%s%s", JsonQuotes(), t.XTime(val).String(), JsonQuotes())
}

func init() {
	RegJsonFormatter(reflect.TypeOf(time.Time{}), dateFormatter)
	RegJsonFormatter(reflect.TypeOf((*time.Time)(nil)), dateFormatter)
}
