package debug

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

const (
	mMAX_DEPTH  = 1000
	mFILTER_KEY = "/x/core/debug/stack.go"
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

// PrintStack prints to standard error the stack trace returned by runtime.Stack.
func PrintStack(skip ...int) {
	fmt.Print(Stack(skip...))
}

// Stack returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
func Stack(skip ...int) string {
	return StackWithFilter("", skip...)
}

// StackWithFilter returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
//
// The parameter <filter> is used to filter the path of the caller.
func StackWithFilter(filter string, skip ...int) string {
	number := 0
	if len(skip) > 0 {
		number = skip[0]
	}
	name := ""
	space := "  "
	index := 1
	buffer := bytes.NewBuffer(nil)
	for i := callerFromIndex(filter) + number; i < mMAX_DEPTH; i++ {
		if pc, file, line, ok := runtime.Caller(i); ok {
			if goRootForFilter != "" && len(file) >= len(goRootForFilter) && file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}
			if filter != "" && strings.Contains(file, filter) {
				continue
			}
			if strings.Contains(file, mFILTER_KEY) {
				continue
			}
			if fn := runtime.FuncForPC(pc); fn == nil {
				name = "unknown"
			} else {
				name = fn.Name()
			}
			if index > 9 {
				space = " "
			}
			buffer.WriteString(fmt.Sprintf("%d.%s%s\n    %s:%d\n", index, space, name, file, line))
			index++
		} else {
			break
		}
	}
	return buffer.String()
}

// CallerPath returns the absolute file path along with its line number of the caller.
func Caller(skip ...int) string {
	return CallerWithFilter("", skip...)
}

// CallerPathWithFilter returns the absolute file path along with its line number of the caller.
//
// The parameter <filter> is used to filter the path of the caller.
func CallerWithFilter(filter string, skip ...int) string {
	number := 0
	if len(skip) > 0 {
		number = skip[0]
	}
	for i := callerFromIndex(filter) + number; i < mMAX_DEPTH; i++ {
		if _, file, line, ok := runtime.Caller(i); ok {
			if filter != "" && strings.Contains(file, filter) {
				continue
			}
			if strings.Contains(file, mFILTER_KEY) {
				continue
			}
			return fmt.Sprintf(`%s:%d`, file, line)
		} else {
			break
		}
	}
	return ""
}

// callerFromIndex returns the caller position exclusive of the debug package.
func callerFromIndex(filter string) int {
	for i := 0; i < mMAX_DEPTH; i++ {
		if _, file, _, ok := runtime.Caller(i); ok {
			if filter != "" && strings.Contains(file, filter) {
				continue
			}
			if strings.Contains(file, mFILTER_KEY) {
				continue
			}
			// exclude the depth from the function of current package.
			return i - 1
		}
	}
	return 0
}
