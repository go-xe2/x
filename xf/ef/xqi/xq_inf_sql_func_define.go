package xqi

type SqlFunId uint

const (
	SFUnDefine SqlFunId = iota
	SFCase
	SFSubString
	SFFromBase64
	SFToBase64
	// 字符串连接函数
	SFConcat
	SFDateAdd
	SFDateSub
	SFDateFormat
	SFDateDiff
	SFDateToUnix
	SFUnixToDate

	// 聚合函数
	SFCount
	SFMax
	SFMin
	SFAvg
	SFSum
)

var SqlFunIdNames = map[SqlFunId]string{
	SFCase:       "case",
	SFSubString:  "Substring",
	SFFromBase64: "FromBase64",
	SFToBase64:   "ToBase64",
	SFConcat:     "Concat",
	SFDateAdd:    "DateAdd",
	SFDateSub:    "DateSub",
	SFDateFormat: "DateFormat",
	SFDateDiff:   "DateDiff",
	SFDateToUnix: "DateToUnix",
	SFUnixToDate: "UnixToDate",

	SFCount: "count",
	SFMax:   "max",
	SFMin:   "min",
	SFAvg:   "avg",
	SFSum:   "sum",
}

var FuncNameMapSqlFunId = map[string]SqlFunId{
	"case":       SFCase,
	"Substring":  SFSubString,
	"FromBase64": SFFromBase64,
	"ToBase64":   SFToBase64,
	"Concat":     SFConcat,
	"substring":  SFSubString,
	"DateAdd":    SFDateAdd,
	"dateAdd":    SFDateAdd,
	"DateSub":    SFDateSub,
	"dateSub":    SFDateSub,
	"DateFormat": SFDateFormat,
	"dateFormat": SFDateFormat,
	"DateDiff":   SFDateDiff,
	"dateDiff":   SFDateDiff,
	"DateToUnix": SFDateToUnix,
	"dateToUnix": SFDateToUnix,
	"UnixToDate": SFUnixToDate,
	"unixToDate": SFUnixToDate,
	"count":      SFCount,
	"Count":      SFCount,
	"max":        SFMax,
	"Max":        SFMax,
	"min":        SFMin,
	"Min":        SFMin,
	"avg":        SFAvg,
	"Avg":        SFAvg,
	"sum":        SFSum,
	"Sum":        SFSum,
}

func (sf SqlFunId) String() string {
	if s, ok := SqlFunIdNames[sf]; ok {
		return s
	}
	return "sql func un define"
}

// 是否是聚合函数
func (sf SqlFunId) IsAggregation() bool {
	return sf == SFCount || sf == SFMax || sf == SFMin || sf == SFSum || sf == SFAvg
}

func HasSqlFunId(funcName string) (bool, SqlFunId) {
	if v, ok := FuncNameMapSqlFunId[funcName]; ok {
		return true, v
	}
	return false, SFUnDefine
}
