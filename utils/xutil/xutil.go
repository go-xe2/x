package xutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-xe2/x/core/empty"
	"github.com/go-xe2/x/type/t"
	"os"
)

func Dump(i ...interface{}) {
	s := Export(i...)
	if s != "" {
		fmt.Println(s)
	}
}

func Export(i ...interface{}) string {
	buffer := bytes.NewBuffer(nil)
	for _, v := range i {
		if b, ok := v.([]byte); ok {
			buffer.Write(b)
		} else {
			if m := t.Map(v); m != nil {
				v = m
			}
			encoder := json.NewEncoder(buffer)
			encoder.SetEscapeHTML(false)
			encoder.SetIndent("", "\t")
			if err := encoder.Encode(v); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
		}
	}
	return buffer.String()
}

func Throw(exception interface{}) {
	panic(exception)
}

func TryCatch(try func(), catch ...func(exception interface{})) {
	if len(catch) > 0 {
		// If <catch> is given, it's used to handle the exception.
		defer func() {
			if e := recover(); e != nil {
				catch[0](e)
			}
		}()
	} else {
		// If no <catch> function passed, it filters the exception.
		defer func() {
			recover()
		}()
	}
	try()
}

func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}
