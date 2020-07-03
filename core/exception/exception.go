package exception

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strings"
)

type IException interface {
	Error() string
	Cause() error
	Format(s fmt.State, verb rune)
	Stack() string
}

type Exception struct {
	error error  // Wrapped error.
	stack stack  // Stack array, which records the stack information when this error is created or wrapped.
	text  string // Error text, which is created by New* functions.
}

const (
	mFILTER_KEY = "/core/core/exception"
)

var (
	// goRootForFilter is used for stack filtering purpose.
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.Replace(goRootForFilter, "\\", "/", -1)
	}
}

// Error implements the interface of Error, it returns the error as string.
func (err *Exception) Error() string {
	if err.text != "" {
		if err.error != nil {
			return err.text + ": " + err.error.Error()
		}
		return err.text
	}
	return err.error.Error()
}

// Cause returns the root cause error.
func (err *Exception) Cause() error {
	loop := err
	for loop != nil {
		if loop.error != nil {
			if e, ok := loop.error.(*Exception); ok {
				loop = e
			} else {
				return loop.error
			}
		} else {
			return loop
		}
	}
	return nil
}

// Format formats the frame according to the fmt.Formatter interface.
//
// %v, %s   : Print the error string;
// %-v, %-s : Print current error string;
// %+s      : Print full stack error list;
// %+v      : Print the error string and full stack error list;
func (err *Exception) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		switch {
		case s.Flag('-'):
			if err.text != "" {
				io.WriteString(s, err.text)
			} else {
				io.WriteString(s, err.Error())
			}
		case s.Flag('+'):
			if verb == 's' {
				io.WriteString(s, err.Stack())
			} else {
				io.WriteString(s, err.Error()+"\n"+err.Stack())
			}
		default:
			io.WriteString(s, err.Error())
		}
	}
}

// Stack returns the stack callers as string.
// It returns an empty string id the <err> does not support stacks.
func (err *Exception) Stack() string {
	if err == nil {
		return ""
	}
	loop := err
	index := 1
	buffer := bytes.NewBuffer(nil)
	for loop != nil {
		buffer.WriteString(fmt.Sprintf("%d. %-v\n", index, loop))
		index++
		formatSubStack(loop.stack, buffer)
		if loop.error != nil {
			if e, ok := loop.error.(*Exception); ok {
				loop = e
			} else {
				buffer.WriteString(fmt.Sprintf("%d. %s\n", index, loop.error.Error()))
				index++
				break
			}
		} else {
			break
		}
	}
	return buffer.String()
}

// formatSubStack formats the stack for error.
func formatSubStack(st stack, buffer *bytes.Buffer) {
	index := 1
	space := "  "
	for _, p := range st {
		if fn := runtime.FuncForPC(p - 1); fn != nil {
			file, line := fn.FileLine(p - 1)
			if strings.Contains(file, mFILTER_KEY) {
				continue
			}
			if goRootForFilter != "" && len(file) >= len(goRootForFilter) && file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}
			if index > 9 {
				space = " "
			}
			buffer.WriteString(fmt.Sprintf("   %d).%s%s\n    \t%s:%d\n", index, space, fn.Name(), file, line))
			index++
		}
	}
}
