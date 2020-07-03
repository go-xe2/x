package xtcp

import "time"

// 简单协议: (面向短链接)发送消息包
func SendPkg(addr string, data []byte, option ...TPkgOption) error {
	conn, err := NewConn(addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendPkg(data, option...)
}

// 简单协议: (面向短链接)发送数据并等待接收返回数据
func SendRecvPkg(addr string, data []byte, option ...TPkgOption) ([]byte, error) {
	conn, err := NewConn(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvPkg(data, option...)
}

// 简单协议: (面向短链接)带超时时间的数据发送
func SendPkgWithTimeout(addr string, data []byte, timeout time.Duration, option ...TPkgOption) error {
	conn, err := NewConn(addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendPkgWithTimeout(data, timeout, option...)
}

// 简单协议: (面向短链接)发送数据并等待接收返回数据(带返回超时等待时间)
func SendRecvPkgWithTimeout(addr string, data []byte, timeout time.Duration, option ...TPkgOption) ([]byte, error) {
	conn, err := NewConn(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvPkgWithTimeout(data, timeout, option...)
}
