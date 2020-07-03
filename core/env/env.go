package env

import (
	_type "github.com/go-xe2/x/sync/type"
	"os"
	"regexp"
	"strings"
)

var (
	// Console options.
	cmdOptions = make(map[string]string)
)

func init() {
	doInit()
}

// doInit does the initialization for this package.
func doInit() {
	reg := regexp.MustCompile(`\-\-{0,1}(.+?)=(.+)`)
	for i := 0; i < len(os.Args); i++ {
		result := reg.FindStringSubmatch(os.Args[i])
		if len(result) > 1 {
			cmdOptions[result[1]] = result[2]
		}
	}
}

// Get returns the command line argument of the specified <key>.
// If the argument does not exist, then it returns the environment variable with specified <key>.
// It returns the default value <def> if none of them exists.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, eg: gf.<package name>.<variable name>;
// 2. Environment arguments are in uppercase format, eg: GF_<package name>_<variable name>；
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
