package xstackQe

import "testing"

func TestNew(t *testing.T) {
	stack := New()
	t.Log("stack:\n")
	for i := 0; i < 10; i++ {
		stack.PushFront(i)
	}
	for i := 0; i < 10; i++ {
		v := stack.PopFront()
		t.Logf("i%d=%v\n", i, v)
		if 9-i != v {
			t.Errorf("i%d stack error.", i)
		}
	}
	t.Log("stack size:", stack.Size())
	t.Log("queue:\n")
	for i := 0; i < 10; i++ {
		stack.PushFront(i)
	}
	for i := 0; i < 10; i++ {
		v := stack.PopBack()
		t.Logf("i%d=%v", i, v)
		if i != v {
			t.Errorf("i%d queue error.", i)
		}
	}
	t.Log("stack size:", stack.Size())
	if stack.Size() != 0 {
		t.Error("stack error size not zero.")
	}
}
