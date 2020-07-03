package defaultDriver

import (
	"fmt"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
)

var _ DbDriverSqlBuilderDateLibs = &TDbDefaultDriver{}

var mysqlDatePartMaps = map[DatePart]string{
	DateYearPart:        "YEAR",
	DateMonthPart:       "MONTH",
	DateDayPart:         "DAY",
	DateHourPart:        "HOUR",
	DateMinutePart:      "MINUTE",
	DateSecondPart:      "SECOND",
	DateMicrosecondPart: "MICROSECOND",
	DateQuarterPart:     "QUARTER",
	DateWeekPart:        "WEEK",
}

func mysqlDatePartName(part DatePart) string {
	if s, ok := mysqlDatePartMaps[part]; ok {
		return s
	}
	return "MINUTE" // 默认为分钟
}

// 数据库时间增加函数映射
func (dr *TDbDefaultDriver) DateAdd(field string, interval int, part DatePart) string {
	return fmt.Sprintf("DATE_ADD(%s, INTERVAL %d %s)", field, interval, mysqlDatePartName(part))
}

// 数据库时间减少函数映射
func (dr *TDbDefaultDriver) DateSub(field string, interval int, part DatePart) string {
	return fmt.Sprintf("DATE_SUB(%s, INTERVAL %d %s)", field, interval, mysqlDatePartName(part))
}

// 格式化时间函数映射, field为DbField可以换转成日期的值
func (dr *TDbDefaultDriver) DateFormat(field string, format string) string {
	return fmt.Sprintf("DATE_FORMAT(%s, '%s')", field, format)
}

// 计算时间差函数映射
func (dr *TDbDefaultDriver) DateDiff(field1 string, field2 string, part DatePart) string {
	return fmt.Sprintf("TIMESTAMPDIFF(%s, %s, %s)", mysqlDatePartName(part), field1, field2)
}

// 时间转时间戳函数映射
func (dr *TDbDefaultDriver) DateToUnix(field string) string {
	return fmt.Sprintf("UNIX_TIMESTAMP(%s)", field)
}

// 时间戳转时间函数映射，field为DbField或实参变量
func (dr *TDbDefaultDriver) UnixToDate(field string) string {
	return fmt.Sprintf("from_unixtime(%s)", field)
}
