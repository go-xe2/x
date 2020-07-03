package xstackQe

type TStackElement struct {
	prev  *TStackElement
	next  *TStackElement
	stack *TStackQueue
	value interface{}
}

func NewElement(v interface{}) *TStackElement {
	return &TStackElement{
		value: v,
	}
}

func (e *TStackElement) Val() interface{} {
	return e.value
}

func (e *TStackElement) Set(v interface{}) *TStackElement {
	e.value = v
	return e
}

func (e *TStackElement) Prev() *TStackElement {
	if p := e.prev; e.stack != nil && p != &e.stack.root {
		return p
	}
	return nil
}

func (e *TStackElement) Next() *TStackElement {
	if p := e.next; e.stack != nil && p != &e.stack.root {
		return p
	}
	return nil
}

func (e *TStackElement) MoveToFront() {
	if e.stack != nil {
		e.stack.MoveToFront(e)
	}
}

func (e *TStackElement) MoveToBack() {
	if e.stack != nil {
		e.stack.MoveToBack(e)
	}
}

func (e *TStackElement) MoveAfter(mark *TStackElement) {
	if e.stack != nil {
		e.stack.MoveAfter(e, mark)
	}
}

func (e *TStackElement) MoveBefore(mark *TStackElement) {
	if e.stack != nil {
		e.stack.MoveBefore(e, mark)
	}
}

func (e *TStackElement) Remove() interface{} {
	if e.stack == nil {
		return e.value
	}
	return e.stack.Remove(e)
}
