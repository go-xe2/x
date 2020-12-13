package logger

type LogLevel int

const (
	LEVEL_ALL           = LEVEL_DEBU | LEVEL_INFO | LEVEL_WARN | LEVEL_ERRO
	LEVEL_DEV           = LEVEL_ALL
	LEVEL_PROD          = LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
	LEVEL_DEBU LogLevel = 1 << iota
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERRO
	LEVEL_NOTI
	LEVEL_CRIT
)

var LogLevelNames = map[LogLevel]string{
	LEVEL_DEBU: "调试",
	LEVEL_INFO: "提示",
	LEVEL_WARN: "警告",
	LEVEL_ERRO: "错误",
	LEVEL_NOTI: "通知",
	LEVEL_CRIT: "关键",
	LEVEL_PROD: "生产环境",
	LEVEL_DEV:  "开发环境",
}

func (l LogLevel) String() string {
	if v, ok := LogLevelNames[l]; ok {
		return v
	}
	return "其他级别"
}

type ILogger interface {
	// 输出错误信息
	Error(tag string, v ...interface{})
	// 格式化输出错误信息
	ErrorF(tag string, format string, v ...interface{})
	// 格式化输出错误信息及换行符
	ErrorFLn(tag string, format string, v ...interface{})
	// 输出警告信息
	Warning(tag string, v ...interface{})
	// 格式化输出警告信息
	WarningF(tag string, format string, v ...interface{})
	// 格式化输出警告信息及换行
	WarningFLn(tag string, format string, v ...interface{})
	// 输出调试信息
	Debug(tag string, v ...interface{})
	// 格式化输出调度信息
	DebugF(tag string, format string, v ...interface{})
	// 格式化输出调度信息及换行
	DebugFLn(tag string, format string, v ...interface{})
	// 输出信息
	Info(tag string, v ...interface{})
	// 格式化输出信息
	InfoF(tag string, format string, v ...interface{})
	// 格式输出信息及换行符
	InfoFLn(tag string, format string, v ...interface{})
	// 控制台打印信息
	Print(tag string, level uint8, v ...interface{})
	// 控制台打印格式化信息
	PrintF(tag string, level uint8, format string, v ...interface{})
	// 控制台打印格式化信息及换行符
	PrintFLn(tag string, level uint8, format string, v ...interface{})
	// 输出通知信息
	Notice(tag string, v ...interface{})
	// 格式输出通知信息
	NoticeF(tag string, format string, v ...interface{})
	// 格式输出通知信息及换行
	NoticeFLn(tag string, format string, v ...interface{})
	// 输出关键信息
	Critical(tag string, v ...interface{})
	// 格式化输出关键信息
	CriticalF(tag string, format string, v ...interface{})
	// 格式化输出关键信息及换行
	CriticalFLn(tag string, format string, v ...interface{})
}
