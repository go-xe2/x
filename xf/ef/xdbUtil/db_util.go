package xdbUtil

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/t"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// 如果v1为空，则返回v2
func EmptyThen(v1 interface{}, v2 interface{}) interface{} {
	if v1 == nil || v1 == "" || v1 == 0 || v1 == false {
		return v2
	}
	return v1
}

// 如果expr成立，则返回trueV, 否则返回falseV
func IfThen(expr interface{}, trueV, falseV interface{}) interface{} {
	var returnValue = func(val interface{}) interface{} {
		if fn, ok := val.(func() interface{}); ok {
			return fn()
		}
		return val
	}
	if fn, ok := expr.(func() bool); ok {
		if fn() {
			return returnValue(trueV)
		}
		return returnValue(falseV)
	}
	if t.Bool(expr, false) {
		return returnValue(trueV)
	}
	return returnValue(falseV)
}

func MakeRandomInt(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

func WithLockContext(fn func()) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	fn()
}

func WithRunTimeContext(closer func() (result interface{}, error error), callback func(duration time.Duration, err error)) (result interface{}, err error) {
	// 记录开始时间
	start := time.Now()
	defer func() {
		if e := recover(); e != nil {
			if err1, ok := e.(error); ok {
				err = err1
			} else {
				err = exception.New(e)
			}
		}
	}()
	result, err = closer()
	duration := time.Since(start)
	//log.Println("执行完毕,用时:", timeduration.Seconds(),timeduration.Seconds()>1.1)
	callback(duration, err)
	return
}

func AddQuotes(data interface{}, quotes ...string) string {
	sep := "\""
	if len(quotes) > 0 {
		sep = quotes[0]
	}
	ret := t.New(data).String()
	//ret = strings.Replace(ret, `\`, `\\`, -1)
	ret = strings.Replace(ret, `"`, `\"`, -1)
	ret = strings.Replace(ret, `'`, `\'`, -1)
	return fmt.Sprintf("%s%s%s", sep, ret, sep)
}
