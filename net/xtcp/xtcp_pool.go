package xtcp

import (
	"github.com/go-xe2/x/container/xpool"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/sync/xsafeMap"
	"time"
)

// 链接池链接对象
type TPoolConn struct {
	*TConn              // 继承底层链接接口对象
	pool   *xpool.TPool // 对应的链接池对象
	status int          // 当前对象的状态，主要用于失败重连判断
}

const (
	mDEFAULT_POOL_EXPIRE = 60000 // (毫秒)默认链接对象过期时间
	mCONN_STATUS_UNKNOWN = 0     // 未知，表示未经过连通性操作;
	mCONN_STATUS_ACTIVE  = 1     // 正常，表示已经经过连通性操作
	mCONN_STATUS_ERROR   = 2     // 错误，表示该接口操作产生了错误，不应当被循环使用了

)

var (
	// 连接池对象map，键名为地址端口，键值为对应的连接池对象
	pools = xsafeMap.NewStrAnyMap()
)

// 创建TCP链接池对象
func NewPoolConn(addr string, timeout ...int) (*TPoolConn, error) {
	var pool *xpool.TPool
	if v := pools.Get(addr); v == nil {
		pools.LockFunc(func(m map[string]interface{}) {
			if v, ok := m[addr]; ok {
				pool = v.(*xpool.TPool)
			} else {
				pool = xpool.New(mDEFAULT_POOL_EXPIRE, func() (interface{}, error) {
					if conn, err := NewConn(addr, timeout...); err == nil {
						return &TPoolConn{conn, pool, mCONN_STATUS_ACTIVE}, nil
					} else {
						return nil, err
					}
				})
				m[addr] = pool
			}
		})
	} else {
		pool = v.(*xpool.TPool)
	}

	if v, err := pool.Get(); err == nil {
		return v.(*TPoolConn), nil
	} else {
		return nil, err
	}
}

// (方法覆盖)覆盖底层接口对象的Close方法
func (c *TPoolConn) Close() error {
	if c.pool != nil && c.status == mCONN_STATUS_ACTIVE {
		c.status = mCONN_STATUS_UNKNOWN
		c.pool.Put(c)
	} else {
		return c.Conn.Close()
	}
	return nil
}

// (方法覆盖)发送数据
func (c *TPoolConn) Send(data []byte, retry ...TRetry) error {
	var err error
	if err = c.TConn.Send(data, retry...); err != nil && c.status == mCONN_STATUS_UNKNOWN {
		if v, e := c.pool.NewFunc(); e == nil {
			c.Conn = v.(*TPoolConn).Conn
			err = c.TConn.Send(data, retry...)
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

// (方法覆盖)接收数据
func (c *TPoolConn) Recv(length int, retry ...TRetry) ([]byte, error) {
	data, err := c.TConn.Recv(length, retry...)
	if err != nil {
		c.status = mCONN_STATUS_ERROR
	} else {
		c.status = mCONN_STATUS_ACTIVE
	}
	return data, err
}

// (方法覆盖)按行读取数据，阻塞读取，直到完成一行读取位置(末尾以'\n'结尾，返回数据不包含换行符)
func (c *TPoolConn) RecvLine(retry ...TRetry) ([]byte, error) {
	data, err := c.TConn.RecvLine(retry...)
	if err != nil {
		c.status = mCONN_STATUS_ERROR
	} else {
		c.status = mCONN_STATUS_ACTIVE
	}
	return data, err
}

// (方法覆盖)带超时时间的数据获取
func (c *TPoolConn) RecvWithTimeout(length int, timeout time.Duration, retry ...TRetry) (data []byte, err error) {
	if err := c.SetRecvDeadline(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	defer func() {
		err = exception.Wrap(c.SetRecvDeadline(time.Time{}), "SetRecvDeadline error")
	}()
	data, err = c.Recv(length, retry...)
	return
}

// (方法覆盖)带超时时间的数据发送
func (c *TPoolConn) SendWithTimeout(data []byte, timeout time.Duration, retry ...TRetry) (err error) {
	if err := c.SetSendDeadline(time.Now().Add(timeout)); err != nil {
		return err
	}
	defer func() {
		err = exception.Wrap(c.SetSendDeadline(time.Time{}), "SetSendDeadline error")
	}()
	err = c.Send(data, retry...)
	return
}

// (方法覆盖)发送数据并等待接收返回数据
func (c *TPoolConn) SendRecv(data []byte, receive int, retry ...TRetry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.Recv(receive, retry...)
	} else {
		return nil, err
	}
}

// (方法覆盖)发送数据并等待接收返回数据(带返回超时等待时间)
func (c *TPoolConn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...TRetry) ([]byte, error) {
	if err := c.Send(data, retry...); err == nil {
		return c.RecvWithTimeout(receive, timeout, retry...)
	} else {
		return nil, err
	}
}
