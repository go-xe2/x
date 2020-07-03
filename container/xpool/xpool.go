package xpool

import (
	"errors"
	"github.com/go-xe2/x/os/xtimer"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/sync/xsafeStack"
	"github.com/go-xe2/x/type/xtime"
	"time"
)

// 线程池
type TPool struct {
	list       *xsafeStack.TSafeStackQe
	closed     *_type.TBool
	Expire     int64
	NewFunc    func() (interface{}, error)
	ExpireFunc func(interface{})
}

// 线程池项
type poolItem struct {
	expire int64
	value  interface{}
}

type NewFunc func() (interface{}, error)

type ExpireFunc func(interface{})

func New(expire int, newFunc NewFunc, expireFunc ...ExpireFunc) *TPool {
	r := &TPool{
		list:    xsafeStack.New(),
		closed:  _type.NewBool(),
		Expire:  int64(expire),
		NewFunc: newFunc,
	}
	if len(expireFunc) > 0 {
		r.ExpireFunc = expireFunc[0]
	}
	xtimer.AddSingleton(time.Second, r.checkExpire)
	return r
}

// 对象放入到线程池
func (p *TPool) Put(value interface{}) {
	item := &poolItem{
		value: value,
	}
	if p.Expire == 0 {
		item.expire = 0
	} else {
		item.expire = xtime.Millisecond() + p.Expire
	}
	p.list.PushBack(item)
}

// 清空线程池
func (p *TPool) Clear() {
	p.list.RemoveAll()
}

// 从线程池中获取一项
func (p *TPool) Get() (interface{}, error) {
	for !p.closed.Val() {
		if r := p.list.PopFront(); r != nil {
			f := r.(*poolItem)
			if f.expire == 0 || f.expire > xtime.Millisecond() {
				return f.value, nil
			}
		} else {
			break
		}
	}
	if p.NewFunc != nil {
		return p.NewFunc()
	}
	return nil, errors.New("线程池没有内容")
}

// 获取池大小
func (p *TPool) Size() int {
	return p.list.Size()
}

// 关闭线程池
func (p *TPool) Close() {
	p.closed.Set(true)
}

// 检查超时方法
func (p *TPool) checkExpire() {
	if p.closed.Val() {
		if p.ExpireFunc != nil {
			for {
				if r := p.list.PopFront(); r != nil {
					p.ExpireFunc(r.(*poolItem).value)
				} else {
					break
				}
			}
		}
		xtimer.Exit()
	}
	for {
		if r := p.list.PopFront(); r != nil {
			item := r.(*poolItem)
			if item.expire == 0 || item.expire > xtime.Millisecond() {
				p.list.PushFront(item)
				break
			}
			if p.ExpireFunc != nil {
				p.ExpireFunc(item.value)
			}
		} else {
			break
		}
	}
}
