package xstackQe

type TStackQueue struct {
	// 根节点
	root TStackElement
	nLen int
}

func New() *TStackQueue {
	return new(TStackQueue).Init()
}

func (sq *TStackQueue) Init() *TStackQueue {
	sq.root.prev = &sq.root
	sq.root.next = &sq.root
	sq.nLen = 0
	return sq
}

func (sq *TStackQueue) lazyInit() {
	if sq.root.next == nil {
		sq.Init()
	}
}

func (sq *TStackQueue) insert(e, at *TStackElement) *TStackElement {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.stack = sq
	sq.nLen++
	return e
}

func (sq *TStackQueue) remove(e *TStackElement) *TStackElement {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.stack = nil
	sq.nLen--
	return e
}

func (sq *TStackQueue) move(e, at *TStackElement) *TStackElement {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	return e
}

func (sq *TStackQueue) insertValue(v interface{}, at *TStackElement) *TStackElement {
	return sq.insert(NewElement(v), at)
}

// 删除元素
func (sq *TStackQueue) Remove(e *TStackElement) interface{} {
	if e.stack == sq {
		sq.remove(e)
	}
	return e.Val()
}

// 加入栈顶
func (sq *TStackQueue) PushFront(v interface{}) *TStackElement {
	sq.lazyInit()
	return sq.insertValue(v, &sq.root)
}

func (sq *TStackQueue) PushBack(v interface{}) *TStackElement {
	sq.lazyInit()
	return sq.insertValue(v, sq.root.prev)
}

// 取栈顶元素并删除,push -> pop 实现LIFO后进选出堆栈操作
func (sq *TStackQueue) PopFront() interface{} {
	e := sq.PeekFront()
	if e == nil {
		return nil
	}
	return sq.Remove(e)
}

func (sq *TStackQueue) PopBack() interface{} {
	e := sq.PeekBack()
	if e == nil {
		return nil
	}
	return sq.Remove(e)
}

// 获取栈顶元素不删除
func (sq *TStackQueue) PeekFront() *TStackElement {
	if sq.nLen == 0 {
		return nil
	}
	return sq.root.next
}

func (sq *TStackQueue) PeekBack() *TStackElement {
	if sq.nLen == 0 {
		return nil
	}
	return sq.root.prev
}

// 栈长度
func (sq *TStackQueue) Size() int {
	return sq.nLen
}

// 取前一个元素
func (sq *TStackQueue) Prev(e *TStackElement) *TStackElement {
	if e != nil {
		return e.Prev()
	}
	return nil
}

// 取后一个元素
func (sq *TStackQueue) Next(e *TStackElement) *TStackElement {
	if e != nil {
		return e.Next()
	}
	return nil
}

func (sq *TStackQueue) InsertBefore(v interface{}, mark *TStackElement) *TStackElement {
	sq.lazyInit()
	if mark.stack != sq {
		return nil
	}
	return sq.insertValue(v, mark.prev)
}

func (sq *TStackQueue) InsertAfter(v interface{}, mark *TStackElement) *TStackElement {
	sq.lazyInit()
	if mark.stack != sq {
		return nil
	}
	return sq.insertValue(v, mark)
}

// 移动到栈顶
func (sq *TStackQueue) MoveToFront(e *TStackElement) {
	if e.stack != sq || sq.root.next == e {
		return
	}
	sq.move(e, &sq.root)
}

// 移动到栈底
func (sq *TStackQueue) MoveToBack(e *TStackElement) {
	if e.stack != sq || sq.root.prev == e {
		return
	}
	sq.move(e, sq.root.prev)
}

// 移动e到mark之前
func (sq *TStackQueue) MoveBefore(e, mark *TStackElement) {
	if e.stack != sq || e == mark || mark.stack != sq {
		return
	}
	sq.move(e, mark.prev)
}

// 移动e到mark之后
func (sq *TStackQueue) MoveAfter(e, mark *TStackElement) {
	if e.stack != sq || e == mark || mark.stack != sq {
		return
	}
	sq.move(e, mark)
}

// 复制堆栈队列到other顶部
func (sq *TStackQueue) AssignFront(src *TStackQueue) {
	sq.lazyInit()
	for i, e := src.Size(), src.PeekFront(); i > 0; i, e = i-1, e.Next() {
		sq.insertValue(e.Val(), sq.root.prev)
	}
}

// 复制堆栈队列到other底部
func (sq *TStackQueue) AssignBack(src *TStackQueue) {
	sq.lazyInit()
	for i, e := src.Size(), src.PeekBack(); i > 0; i, e = i-1, e.Prev() {
		sq.insertValue(e.Val(), &sq.root)
	}
}
