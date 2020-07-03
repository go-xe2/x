package xsafeStack

import (
	"github.com/go-xe2/x/container/xstackQe"
	"github.com/go-xe2/x/core/rwmutex"
)

type TSafeStackQe struct {
	mu   *rwmutex.RWMutex
	data *xstackQe.TStackQueue
}

func New(unsafe ...bool) *TSafeStackQe {
	return &TSafeStackQe{
		mu:   rwmutex.New(unsafe...),
		data: xstackQe.New(),
	}
}

func (sq *TSafeStackQe) Init() *TSafeStackQe {
	sq.mu.Lock()
	sq.data.Init()
	sq.mu.Unlock()
	return sq
}

// 删除元素
func (sq *TSafeStackQe) Remove(e *xstackQe.TStackElement) interface{} {
	sq.mu.Lock()
	v := sq.data.Remove(e)
	sq.mu.Unlock()
	return v
}

func (sq *TSafeStackQe) RemoveAll() {
	sq.mu.Lock()
	sq.data = xstackQe.New()
	sq.mu.Unlock()
}

func (sq *TSafeStackQe) Clear() {
	sq.RemoveAll()
}

// 加入栈顶
func (sq *TSafeStackQe) PushFront(v interface{}) *xstackQe.TStackElement {
	sq.mu.Lock()
	e := sq.data.PushFront(v)
	sq.mu.Unlock()
	return e
}

func (sq *TSafeStackQe) PushBack(v interface{}) *xstackQe.TStackElement {
	sq.mu.Lock()
	e := sq.data.PushBack(v)
	sq.mu.Unlock()
	return e
}

// 取栈顶元素并删除,push -> pop 实现LIFO后进选出堆栈操作
func (sq *TSafeStackQe) PopFront() interface{} {
	sq.mu.Lock()
	v := sq.data.PopFront()
	sq.mu.Unlock()
	return v
}

func (sq *TSafeStackQe) PopFronts(max int) (values []interface{}) {
	sq.mu.Lock()
	length := sq.data.Size()
	if length > 0 {
		if max > 0 && max < length {
			length = max
		}
		values = make([]interface{}, length)
		for i := 0; i < length; i++ {
			values[i] = sq.data.Remove(sq.data.PeekFront())
		}
	}
	sq.mu.Unlock()
	return
}

func (sq *TSafeStackQe) PopFrontAll() []interface{} {
	return sq.PopFronts(-1)
}

func (sq *TSafeStackQe) PopBack() interface{} {
	sq.mu.Lock()
	v := sq.data.PopBack()
	sq.mu.Unlock()
	return v
}

func (sq *TSafeStackQe) PopBacks(max int) (values []interface{}) {
	sq.mu.Lock()
	length := sq.data.Size()
	if length > 0 {
		if max > 0 && max < length {
			length = max
		}
		values = make([]interface{}, length)
		for i := 0; i < length; i++ {
			values[i] = sq.data.Remove(sq.data.PeekBack())
		}
	}
	sq.mu.Unlock()
	return
}

func (sq *TSafeStackQe) PopBackAll() []interface{} {
	return sq.PopBacks(-1)
}

// 获取栈顶元素不删除
func (sq *TSafeStackQe) PeekFront() *xstackQe.TStackElement {
	sq.mu.RLock()
	v := sq.data.PeekFront()
	sq.mu.RUnlock()
	return v
}

func (sq *TSafeStackQe) FrontAll() (values []interface{}) {
	sq.mu.RLock()
	length := sq.data.Size()
	if length > 0 {
		values = make([]interface{}, length)
		for i, e := 0, sq.data.PeekFront(); i < length; i, e = i+1, e.Next() {
			values[i] = e.Val()
		}
	}
	sq.mu.RUnlock()
	return
}

func (sq *TSafeStackQe) PeekBack() *xstackQe.TStackElement {
	sq.mu.RLock()
	v := sq.data.PeekBack()
	sq.mu.RUnlock()
	return v
}

func (sq *TSafeStackQe) BackAll() (values []interface{}) {
	sq.mu.RLock()
	length := sq.data.Size()
	if length > 0 {
		values = make([]interface{}, length)
		for i, e := 0, sq.data.PeekBack(); i < length; i, e = i+1, e.Prev() {
			values[i] = e.Val()
		}
	}
	sq.mu.RUnlock()
	return
}

// 栈长度
func (sq *TSafeStackQe) Size() int {
	sq.mu.RLock()
	n := sq.data.Size()
	sq.mu.RUnlock()
	return n
}

// 取前一个元素
func (sq *TSafeStackQe) Prev(e *xstackQe.TStackElement) *xstackQe.TStackElement {
	sq.mu.RLock()
	en := sq.data.Prev(e)
	sq.mu.RUnlock()
	return en
}

// 取后一个元素
func (sq *TSafeStackQe) Next(e *xstackQe.TStackElement) *xstackQe.TStackElement {
	sq.mu.RLock()
	en := sq.data.Next(e)
	sq.mu.RUnlock()
	return en
}

func (sq *TSafeStackQe) InsertBefore(v interface{}, mark *xstackQe.TStackElement) *xstackQe.TStackElement {
	sq.mu.Lock()
	e := sq.data.InsertBefore(v, mark)
	sq.mu.Unlock()
	return e
}

func (sq *TSafeStackQe) InsertAfter(v interface{}, mark *xstackQe.TStackElement) *xstackQe.TStackElement {
	sq.mu.Lock()
	e := sq.data.InsertAfter(v, mark)
	sq.mu.Unlock()
	return e
}

// 移动到栈顶
func (sq *TSafeStackQe) MoveToFront(e *xstackQe.TStackElement) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.data.MoveToFront(e)
}

// 移动到栈底
func (sq *TSafeStackQe) MoveToBack(e *xstackQe.TStackElement) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.data.MoveToBack(e)
}

// 移动e到mark之前
func (sq *TSafeStackQe) MoveBefore(e, mark *xstackQe.TStackElement) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.data.MoveBefore(e, mark)
}

// 移动e到mark之后
func (sq *TSafeStackQe) MoveAfter(e, mark *xstackQe.TStackElement) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.data.MoveAfter(e, mark)
}

// 复制堆栈队列到other顶部
func (sq *TSafeStackQe) AssignFront(src *TSafeStackQe) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.data.AssignFront(src.data)
}

// 复制堆栈队列到other底部
func (sq *TSafeStackQe) AssignBack(src *TSafeStackQe) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.data.AssignBack(src.data)
}

func (sq *TSafeStackQe) AssignFrontSQ(src *xstackQe.TStackQueue) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.data.AssignFront(src)
}

func (sq *TSafeStackQe) AssignBackSQ(src *xstackQe.TStackQueue) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.data.AssignBack(src)
}
