package exception

import (
	"fmt"
	"github.com/go-xe2/x/type/xstring"
)

var _ IException = (*Exception)(nil)

func New(value ...interface{}) IException {
	if value == nil {
		return nil
	}
	s := make([]string, len(value))
	for i := 0; i < len(value); i++ {
		s[i] = fmt.Sprintf("%v", value[0])
	}
	return NewText(xstring.Join(s, " "))
}

func Newf(format string, args ...interface{}) IException {
	s := fmt.Sprintf(format, args...)
	return NewText(s)
}

func NewText(text string) IException {
	if text == "" {
		return nil
	}
	return &Exception{
		stack: callers(),
		text:  text,
	}
}

func Wrap(err error, text string) IException {
	if err == nil {
		return nil
	}
	return &Exception{
		error: err,
		stack: callers(),
		text:  text,
	}
}

func Wrapf(err error, format string, args ...interface{}) IException {
	if err == nil {
		return nil
	}
	return &Exception{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
	}
}

func Cause(err error) error {
	if err != nil {
		if e, ok := err.(ICause); ok {
			return e.Cause()
		}
	}
	return err
}

func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(IStack); ok {
		return e.Stack()
	}
	return ""
}
