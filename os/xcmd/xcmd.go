package xcmd

import (
	"github.com/go-xe2/x/os/xlog"
	"os"
	"regexp"
)

var Value = &tCmdValue{}
var Option = &tCmdOption{}
var cmdFuncMap = make(map[string]func())

func init() {
	doInit()
}

func doInit() {
	Value.values = Value.values[:0]
	Option.options = make(map[string]string)
	reg := regexp.MustCompile(`^\-{1,2}(\w+)={0,1}(.*)`)
	for i := 0; i < len(os.Args); i++ {
		result := reg.FindStringSubmatch(os.Args[i])
		if len(result) > 1 {
			Option.options[result[1]] = result[2]
		} else {
			Value.values = append(Value.values, os.Args[i])
		}
	}
}

func BindHandle(cmd string, f func()) {
	if _, ok := cmdFuncMap[cmd]; ok {
		xlog.Fatal("duplicated handle for command:" + cmd)
	} else {
		cmdFuncMap[cmd] = f
	}
}

func RunHandle(cmd string) {
	if handle, ok := cmdFuncMap[cmd]; ok {
		handle()
	} else {
		xlog.Fatal("no handle found for command:" + cmd)
	}
}

func AutoRun() {
	if cmd := Value.Get(1); cmd != "" {
		if handle, ok := cmdFuncMap[cmd]; ok {
			handle()
		} else {
			xlog.Fatal("no handle found for command:" + cmd)
		}
	} else {
		xlog.Fatal("no command found")
	}
}
