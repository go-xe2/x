package xproc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-xe2/x/net/xtcp"
	"github.com/go-xe2/x/os/xfileCache"
	"github.com/go-xe2/x/type/t"
	"io"
	"time"
)

const (
	mPROC_COMM_FAILURE_RETRY_COUNT   = 3    // 失败重试次数
	mPROC_COMM_FAILURE_RETRY_TIMEOUT = 1000 // (毫秒)失败重试间隔
	mPROC_COMM_SEND_TIMEOUT          = 5000 // (毫秒)发送超时时间
	mPROC_COMM_DEAFULT_GRUOP_NAME    = ""   // 默认分组名称
)

// 向指定gproc进程发送数据.
func Send(pid int, data []byte, group ...string) error {
	msg := TProcessMsg{
		SendPid: Pid(),
		RecvPid: pid,
		Group:   mPROC_COMM_DEAFULT_GRUOP_NAME,
		Data:    data,
	}
	if len(group) > 0 {
		msg.Group = group[0]
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	var buf []byte
	var conn *xtcp.TPoolConn
	// 循环获取连接TCP对象
	for i := mPROC_COMM_FAILURE_RETRY_COUNT; i > 0; i-- {
		if conn, err = getConnByPid(pid); err == nil {
			break
		}
		time.Sleep(mPROC_COMM_FAILURE_RETRY_TIMEOUT * time.Millisecond)
	}
	if conn == nil {
		return err
	}
	defer conn.Close()
	// 执行数据发送
	buf, err = conn.SendRecvPkgWithTimeout(msgBytes, mPROC_COMM_SEND_TIMEOUT*time.Millisecond)
	if len(buf) > 0 {
		if !bytes.EqualFold(buf, []byte("ok")) {
			err = errors.New(string(buf))
		}
	}
	// EOF不算异常错误
	if err == io.EOF {
		err = nil
	}
	return err
}

// 获取指定进程的TCP通信对象
func getConnByPid(pid int) (*xtcp.TPoolConn, error) {
	port := getPortByPid(pid)
	if port > 0 {
		if conn, err := xtcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", port)); err == nil {
			return conn, nil
		} else {
			return nil, err
		}
	}
	return nil, errors.New(fmt.Sprintf("could not find port for pid: %d", pid))
}

// 获取指定进程监听的端口号
func getPortByPid(pid int) int {
	path := getCommFilePath(pid)
	content := xfileCache.GetContents(path)
	return t.Int(content)
}
