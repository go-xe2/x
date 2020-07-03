package xlog

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-xe2/x/core/debug"
	"github.com/go-xe2/x/os/xfile"
	"github.com/go-xe2/x/os/xfilePool"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xtime"
	"github.com/go-xe2/x/utils/xregex"
	"io"
	"os"
	"strings"
	"time"
)

type TLogger struct {
	parent      *TLogger  // 父日志器
	writer      io.Writer // 输出接口io.Writer
	flags       int       // 日志扩展设置
	path        string    // 日志存放路径
	file        string    // 日志文件名格式字符串
	level       int       // 输出日志级别
	prefix      string    // 日志输出前缀
	stSkip      int       // 日志栈输出层级数
	stStatus    int       // 日志栈启用状态(1: 启有 - 默认; 0: 禁用)
	headerPrint bool      // 是否打印日志头(默认true)
	stdoutPrint bool      // 是否输出到stdout, 默认true
}

const (
	mDEFAULT_FILE_FORMAT     = `{Y-m-d}.log`
	mDEFAULT_FILE_POOL_FLAGS = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	mDEFAULT_FPOOL_PERM      = os.FileMode(0666)
	mDEFAULT_FPOOL_EXPIRE    = 60000
	mPATH_FILTER_KEY         = "/x/os/xlog/xlog"
)

const (
	F_ASYNC      = 1 << iota // Print logging content asynchronously。
	F_FILE_LONG              // Print full file name and line number: /a/b/c/d.go:23.
	F_FILE_SHORT             // Print final file name element and line number: d.go:23. overrides F_FILE_LONG.
	F_TIME_DATE              // Print the date in the local time zone: 2009-01-23.
	F_TIME_TIME              // Print the time in the local time zone: 01:23:23.
	F_TIME_MILLI             // Print the time with milliseconds in the local time zone: 01:23:23.675.
	F_TIME_STD   = F_TIME_DATE | F_TIME_MILLI
)

// 创建日志器实例
func New() *TLogger {
	logger := &TLogger{
		file:        mDEFAULT_FILE_FORMAT,
		flags:       F_TIME_STD,
		level:       LEVEL_ALL,
		stStatus:    1,
		headerPrint: true,
		stdoutPrint: true,
	}
	return logger
}

// 刻隆日志器
func (l *TLogger) Clone() *TLogger {
	logger := TLogger{}
	logger = *l
	logger.parent = l
	return &logger
}

// 设置日志级别
func (l *TLogger) SetLevel(level int) {
	l.level = level
}

// 获取日志级别
func (l *TLogger) GetLevel() int {
	return l.level
}

// 设置是否为调试模式
func (l *TLogger) SetDebug(debug bool) {
	if debug {
		l.level = l.level | LEVEL_DEBU
	} else {
		l.level = l.level & ^LEVEL_DEBU
	}
}

func (l *TLogger) GetDebug() bool {
	return l.level&LEVEL_DEBU == LEVEL_DEBU
}

// 设置是否启用异步输出日志
func (l *TLogger) SetAsync(enabled bool) {
	if enabled {
		l.flags = l.flags | F_ASYNC
	} else {
		l.flags = l.flags & ^F_ASYNC
	}
}

func (l *TLogger) GetAsync() bool {
	return l.flags&F_ASYNC == F_ASYNC
}

// 设置日扩展标识
func (l *TLogger) SetFlags(flags int) {
	l.flags = flags
}

// 获取日志扩散标识
func (l *TLogger) GetFlags() int {
	return l.flags
}

// 设置是否启用输出日志栈
func (l *TLogger) SetStack(enabled bool) {
	if enabled {
		l.stStatus = 1
	} else {
		l.stStatus = 0
	}
}

func (l *TLogger) GetStackStatus() bool {
	return l.stStatus == 1
}

// 设置输出日志栈层级数
func (l *TLogger) SetStackSkip(skip int) {
	l.stSkip = skip
}

func (l *TLogger) GetStackSkip() int {
	return l.stSkip
}

// 设置日志输出接口
func (l *TLogger) SetWriter(writer io.Writer) {
	l.writer = writer
}

// 获取日志输出接口
func (l *TLogger) GetWriter() io.Writer {
	return l.writer
}

// 获取日志存储文件指针
func (l *TLogger) getFilePointer() *xfilePool.File {
	if path := l.path; path != "" {
		// Content containing "{}" in the file name is formatted using gtime
		file, _ := xregex.ReplaceStringFunc(`{.+?}`, l.file, func(s string) string {
			return xtime.Now().Format(strings.Trim(s, "{}"))
		})
		// Create path if it does not exist。
		if !xfile.Exists(path) {
			if err := xfile.Mkdir(path); err != nil {
				fmt.Fprintln(os.Stderr, fmt.Sprintf(`[glog] mkdir "%s" failed: %s`, path, err.Error()))
				return nil
			}
		}
		if fp, err := xfilePool.Open(
			path+xfile.Separator+file,
			mDEFAULT_FILE_POOL_FLAGS,
			mDEFAULT_FPOOL_PERM,
			mDEFAULT_FPOOL_EXPIRE); err == nil {
			return fp
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	return nil
}

// 设置日志存放目录
func (l *TLogger) SetPath(path string) error {
	if path == "" {
		return errors.New("path is empty")
	}
	if !xfile.Exists(path) {
		if err := xfile.Mkdir(path); err != nil {
			fmt.Fprintln(os.Stderr, fmt.Sprintf(`[glog] mkdir "%s" failed: %s`, path, err.Error()))
			return err
		}
	}
	l.path = strings.TrimRight(path, xfile.Separator)
	return nil
}

// 获取日志存放目录
func (l *TLogger) GetPath() string {
	return l.path
}

// 设置日志文件名格式，默认:Y-m-d.log, eg: 2018-01-01.log
func (l *TLogger) SetFile(pattern string) {
	l.file = pattern
}

// 设置日志是否输出到stdout, 默认为true
func (l *TLogger) SetStdoutPrint(enabled bool) {
	l.stdoutPrint = enabled
}

// 设置是否打印日志头,默认true
func (l *TLogger) SetHeaderPrint(enabled bool) {
	l.headerPrint = enabled
}

// 设置日志前掇
func (l *TLogger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *TLogger) GetPrefix() string {
	return l.prefix
}

// 输出内容到接口
func (l *TLogger) print(std io.Writer, lead string, value ...interface{}) {
	buffer := bytes.NewBuffer(nil)
	if l.headerPrint {
		// Time.
		timeFormat := ""
		if l.flags&F_TIME_DATE > 0 {
			timeFormat += "2006-01-02 "
		}
		if l.flags&F_TIME_TIME > 0 {
			timeFormat += "15:04:05 "
		}
		if l.flags&F_TIME_MILLI > 0 {
			timeFormat += "15:04:05.000 "
		}
		if len(timeFormat) > 0 {
			buffer.WriteString(time.Now().Format(timeFormat))
		}
		// Lead string.
		if len(lead) > 0 {
			buffer.WriteString(lead)
			if len(value) > 0 {
				buffer.WriteByte(' ')
			}
		}
		// Caller path.
		callerPath := ""
		if l.flags&F_FILE_LONG > 0 {
			callerPath = debug.CallerWithFilter(mPATH_FILTER_KEY, l.stSkip) + ": "
		}
		if l.flags&F_FILE_SHORT > 0 {
			callerPath = xfile.Basename(debug.CallerWithFilter(mPATH_FILTER_KEY, l.stSkip)) + ": "
		}
		if len(callerPath) > 0 {
			buffer.WriteString(callerPath)
		}
		// Prefix.
		if len(l.prefix) > 0 {
			buffer.WriteString(l.prefix + " ")
		}
	}
	// Convert value to string.
	tempStr := ""
	valueStr := ""
	for _, v := range value {
		if err, ok := v.(error); ok {
			tempStr = fmt.Sprintf("%+v", err)
		} else {
			tempStr = t.String(v)
		}
		if len(valueStr) > 0 {
			if valueStr[len(valueStr)-1] == '\n' {
				// Remove one blank line(\n\n).
				if tempStr[0] == '\n' {
					valueStr += tempStr[1:]
				} else {
					valueStr += tempStr
				}
			} else {
				valueStr += " " + tempStr
			}
		} else {
			valueStr = tempStr
		}
	}
	buffer.WriteString(valueStr + "\n")
	if l.flags&F_ASYNC > 0 {
		asyncPool.Add(func() {
			l.printToWriter(std, buffer)
		})
	} else {
		l.printToWriter(std, buffer)
	}
}

// 输出二进制数据到接口
func (l *TLogger) printToWriter(std io.Writer, buffer *bytes.Buffer) {
	if l.writer == nil {
		if f := l.getFilePointer(); f != nil {
			defer f.Close()
			if _, err := io.WriteString(f, buffer.String()); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
		}
		// Allow output to stdout?
		if l.stdoutPrint {
			if _, err := std.Write(buffer.Bytes()); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
		}
	} else {
		if _, err := l.writer.Write(buffer.Bytes()); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}

// 输出日志到stdout
func (l *TLogger) printStd(lead string, value ...interface{}) {
	l.print(os.Stdout, lead, value...)
}

// 输出日志及错误栈到stdout
func (l *TLogger) printErr(lead string, value ...interface{}) {
	if l.stStatus == 1 {
		if s := l.GetStack(); s != "" {
			value = append(value, "\nStack:\n"+s)
		}
	}
	// In matter of sequence, do not use stderr here, but use the same stdout.
	l.print(os.Stdout, lead, value...)
}

// 格式化日志输出内容
func (l *TLogger) format(format string, value ...interface{}) string {
	return fmt.Sprintf(format, value...)
}

// 打印指定层级数日志栈到输出
func (l *TLogger) PrintStack(skip ...int) {
	if s := l.GetStack(skip...); s != "" {
		l.Println("Stack:\n" + s)
	} else {
		l.Println()
	}
}

// 获取指定层级数日志栈字符串
func (l *TLogger) GetStack(skip ...int) string {
	number := 1
	if len(skip) > 0 {
		number = skip[0] + 1
	}
	return debug.StackWithFilter(mPATH_FILTER_KEY, number)
}
