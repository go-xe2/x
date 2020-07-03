package xconfig

import "github.com/go-xe2/x/core/cmdenv"

const (
	// 是否打印配置错误信息, 默认true
	mERROR_PRINT_KEY = "x.config.errorPrint"
)

func errorPrint() bool {
	return cmdenv.Get(mERROR_PRINT_KEY, true).Bool()
}
