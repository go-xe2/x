package xdriveri

type DatePart int

const (
	DateUnknownPart DatePart = iota
	// 年
	DateYearPart
	// 月
	DateMonthPart
	// 天
	DateDayPart
	// 小时
	DateHourPart
	// 分
	DateMinutePart
	// 秒
	DateSecondPart
	// 毫秒
	DateMicrosecondPart
	// 季节
	DateQuarterPart
	// 周
	DateWeekPart
)

var DatePartNames = map[DatePart]string{
	DateYearPart:        "year",
	DateMonthPart:       "month",
	DateDayPart:         "day",
	DateHourPart:        "hour",
	DateMinutePart:      "minute",
	DateSecondPart:      "second",
	DateMicrosecondPart: "millisecond",
	DateQuarterPart:     "quarter",
	DateWeekPart:        "week",
}
