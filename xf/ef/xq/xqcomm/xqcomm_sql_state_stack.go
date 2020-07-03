package xqcomm

type sqlStack struct {
	prev *sqlStack
	val  interface{}
}

// 编译状态堆栈
type tSqlStateStack struct {
	root sqlStack
}

func newSqlStateStack() *tSqlStateStack {
	result := &tSqlStateStack{}
	result.root.val = 0
	result.root.prev = &result.root
	return result
}

func (ssk *tSqlStateStack) push(state interface{}) {
	p := &sqlStack{
		prev: ssk.root.prev,
		val:  state,
	}
	ssk.root.prev = p
}

func (ssk *tSqlStateStack) pop() interface{} {
	p := ssk.root.prev
	ssk.root.prev = p.prev
	return p.val
}

func (ssk *tSqlStateStack) peek() interface{} {
	return ssk.root.prev.val
}

func (ssk *tSqlStateStack) Clear() {
	ssk.root.prev = &ssk.root
	ssk.root.val = 0
}
