package xproc

import (
	"github.com/go-xe2/x/os/xfile"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/type/t"
	"os"
)

// 进程通信数据结构
type TProcessMsg struct {
	SendPid int    `json:"spid"`  // 发送进程ID
	RecvPid int    `json:"rpid"`  // 接收进程ID
	Group   string `json:"group"` // 分组名称
	Data    []byte `json:"data"`  // 原始数据
}

// 本地进程通信接收消息队列(按照分组进行构建的map，键值为*gqueue.Queue对象)
var commReceiveQueues = xsafeMap.NewStrAnyMap()

// (用于发送)已建立的PID对应的Conn通信对象，键值为一个Pool，防止并行使用同一个通信对象造成数据重叠
var commPidConnMap = xsafeMap.NewIntAnyMap()

// 获取指定进程的通信文件地址
func getCommFilePath(pid int) string {
	return getCommDirPath() + xfile.Separator + t.String(pid)
}

// 获取进程间通信目录地址
func getCommDirPath() string {
	tempDir := os.Getenv(mPROC_TEMP_DIR_ENV_KEY)
	if tempDir == "" {
		tempDir = xfile.TempDir()
	}
	return tempDir + xfile.Separator + "xproc"
}
