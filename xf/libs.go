/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-10 16:27
* Description:
*****************************************************************/

package xf

import "fmt"

func TryEF(try func(ctx ...interface{}) error, ex func(err error, ctx ...interface{}) error, fi func(err error, ctx ...interface{}) error, ctx ...interface{}) (err error) {
	defer func() {
		e := recover()
		if fi != nil {
			var e1 error = nil
			if e != nil {
				e1 = fmt.Errorf("%s", e)
			}
			err = fi(e1, ctx...)
		}
	}()
	err = try(ctx...)
	if err != nil {
		if ex != nil {
			e1 := ex(err, ctx...)
			if e1 != nil {
				panic(err)
			}
		}
		panic(err)
	}
	return
}
