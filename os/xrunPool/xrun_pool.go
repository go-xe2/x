// 可重复使用线程池,实现接口goroutine reusable pool

package xrunPool

import (
	"errors"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/sync/xsafeStack"
)

type TRunPool struct {
	limit  int                      // Max goroutine count limit.
	count  *_type.TInt              // Current running goroutine count.
	list   *xsafeStack.TSafeStackQe // Job list for asynchronous job adding purpose.
	closed *_type.TBool             // Is pool closed or not.
}

// 全局线程池实例
var pool = New()

func New(limit ...int) *TRunPool {
	p := &TRunPool{
		limit:  -1,
		count:  _type.NewInt(),
		list:   xsafeStack.New(),
		closed: _type.NewBool(),
	}
	if len(limit) > 0 && limit[0] > 0 {
		p.limit = limit[0]
	}
	return p
}

// 添加工作方法到线程池
func Add(f func()) error {
	return pool.Add(f)
}

// 获取线程池当前线程数
func Size() int {
	return pool.Size()
}

// 获取池线程中所有工作方法数
func Jobs() int {
	return pool.Jobs()
}

// 添加工作方法到线程池
func (p *TRunPool) Add(f func()) error {
	for p.closed.Val() {
		return errors.New("pool closed")
	}
	p.list.PushFront(f)
	var n int
	for {
		n = p.count.Val()
		if p.limit != -1 && n >= p.limit {
			return nil
		}
		if p.count.CompareSwap(n, n+1) {
			break
		}
	}
	p.fork()
	return nil
}

// 获取线程池最大线程数
func (p *TRunPool) Cap() int {
	return p.limit
}

// 获取线程池当前线程数
func (p *TRunPool) Size() int {
	return p.count.Val()
}

// 获取池中当前工作方法数
func (p *TRunPool) Jobs() int {
	return p.list.Size()
}

// 取出工作方法并执行
func (p *TRunPool) fork() {
	go func() {
		defer p.count.Add(-1)
		job := (interface{})(nil)
		for !p.closed.Val() {
			if job = p.list.PopBack(); job != nil {
				job.(func())()
			} else {
				return
			}
		}
	}()
}

// 检查线程池是否已关闭
func (p *TRunPool) IsClosed() bool {
	return p.closed.Val()
}

// 关闭线程池
func (p *TRunPool) Close() {
	p.closed.Set(true)
}
