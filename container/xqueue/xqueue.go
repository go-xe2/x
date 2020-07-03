package xqueue

import (
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/sync/xsafeStack"
	"math"
)

type TQueue struct {
	limit  int                      // Limit for queue size.
	list   *xsafeStack.TSafeStackQe // Underlying list structure for data maintaining.
	closed *_type.TBool             // Whether queue is closed.
	events chan struct{}            // Events for data writing.
	C      chan interface{}         // Underlying channel for data reading.
}

const (
	mDEFAULT_QUEUE_SIZE     = 10000
	mDEFAULT_MAX_BATCH_SIZE = 10
)

func New(limit ...int) *TQueue {
	q := &TQueue{
		closed: _type.NewBool(),
	}
	if len(limit) > 0 && limit[0] > 0 {
		q.limit = limit[0]
		q.C = make(chan interface{}, limit[0])
	} else {
		q.list = xsafeStack.New()
		q.events = make(chan struct{}, math.MaxInt32)
		q.C = make(chan interface{}, mDEFAULT_QUEUE_SIZE)
		go q.startAsyncLoop()
	}
	return q
}

func (q *TQueue) startAsyncLoop() {
	defer func() {
		if q.closed.Val() {
			_ = recover()
		}
	}()
	for !q.closed.Val() {
		<-q.events
		for !q.closed.Val() {
			if length := q.list.Size(); length > 0 {
				if length > mDEFAULT_MAX_BATCH_SIZE {
					length = mDEFAULT_MAX_BATCH_SIZE
				}
				for _, v := range q.list.PopFronts(length) {
					q.C <- v
				}
			} else {
				break
			}
		}
		for i := 0; i < len(q.events)-1; i++ {
			<-q.events
		}
	}
	close(q.C)
}

func (q *TQueue) Push(v interface{}) {
	if q.limit > 0 {
		q.C <- v
	} else {
		q.list.PushBack(v)
		if len(q.events) < mDEFAULT_QUEUE_SIZE {
			q.events <- struct{}{}
		}
	}
}

func (q *TQueue) Pop() interface{} {
	return <-q.C
}

func (q *TQueue) Close() {
	q.closed.Set(true)
	if q.events != nil {
		close(q.events)
	}
	if q.limit > 0 {
		close(q.C)
	}
	for i := 0; i < mDEFAULT_MAX_BATCH_SIZE; i++ {
		q.Pop()
	}
}

func (q *TQueue) Len() (length int) {
	if q.list != nil {
		length += q.list.Size()
	}
	length += len(q.C)
	return
}

// Len别名
func (q *TQueue) Size() int {
	return q.Len()
}
