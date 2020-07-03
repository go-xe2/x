package cmdenv

import (
	_type "github.com/go-xe2/x/sync/type"
	"os"
	"regexp"
	"strings"
)

var (
	cmdOptions = make(map[string]string)
)

func init() {
	doInit()
}

func doInit() {
	reg := regexp.MustCompile(`\-\-{0,1}(.+?)=(.+)`)
	for i := 0; i < len(os.Args); i++ {
		result := reg.FindStringSubmatch(os.Args[i])
		if len(result) > 1 {
			cmdOptions[result[1]] = result[2]
		}
	}
}

func Get(key string, def ...interface{}) *_type.TVar {
	value := interface{}(nil)
	if len(def) > 0 {
		value = def[0]
	}
	if v, ok := cmdOptions[key]; ok {
		value = v
	} else {
		key = strings.ToUpper(strings.Replace(key, ".", "_", -1))
		if v := os.Getenv(key); v != "" {
			value = v
		}
	}
	return _type.NewVar(value, true)
}
