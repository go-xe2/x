package xtcp

import (
	"github.com/go-xe2/x/core/exception"
	"time"
)

// 简单协议: (方法覆盖)发送数据
func (c *TPoolConn) SendPkg(data []byte, option ...TPkgOption) (err error) {
	if err = c.TConn.SendPkg(data, option...); err != nil && c.status == mCONN_STATUS_UNKNOWN {
		if v, e := c.pool.NewFunc(); e == nil {
			c.Conn = v.(*TPoolConn).Conn
			err = c.TConn.SendPkg(data, option...)
		} else {
			err = e
		}
	}
	if err != nil {
		c.status = mCONN_STATUS_ERROR
	} else {
		c.status = mCONN_STATUS_ACTIVE
	}
	return err
}

// 简单协议: (方法覆盖)接收数据
func (c *TPoolConn) RecvPkg(option ...TPkgOption) ([]byte, error) {
	data, err := c.TConn.RecvPkg(option...)
	if err != nil {
		c.status = mCONN_STATUS_ERROR
	} else {
		c.status = mCONN_STATUS_ACTIVE
	}
	return data, err
}

// 简单协议: (方法覆盖)带超时时间的数据获取
func (c *TPoolConn) RecvPkgWithTimeout(timeout time.Duration, option ...TPkgOption) (data []byte, err error) {
	if err := c.SetRecvDeadline(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	defer func() {
		err = exception.Wrap(c.SetRecvDeadline(time.Time{}), "SetRecvDeadline error")
	}()
	data, err = c.RecvPkg(option...)
	return
}

// 简单协议: (方法覆盖)带超时时间的数据发送
func (c *TPoolConn) SendPkgWithTimeout(data []byte, timeout time.Duration, option ...TPkgOption) (err error) {
	if err := c.SetSendDeadline(time.Now().Add(timeout)); err != nil {
		return err
	}
	defer func() {
		err = exception.Wrap(c.SetSendDeadline(time.Time{}), "SetSendDeadline error")
	}()
	err = c.SendPkg(data, option...)
	return
}

// 简单协议: (方法覆盖)发送数据并等待接收返回数据
func (c *TPoolConn) SendRecvPkg(data []byte, option ...TPkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkg(option...)
	} else {
		return nil, err
	}
}

// 简单协议: (方法覆盖)发送数据并等待接收返回数据(带返回超时等待时间)
func (c *TPoolConn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...TPkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkgWithTimeout(timeout, option...)
	} else {
		return nil, err
	}
}
