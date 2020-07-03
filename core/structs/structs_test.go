package structs

import (
	"fmt"
	"testing"
)

type TestA struct {
	FieldAName string
}

type TestB struct {
	TestA
	FieldBName string
}

func (a *TestA) HelloA() string {
	return fmt.Sprintf("helloA %s", a.FieldAName)
}

func (b *TestB) HelloB(age int) string {
	return fmt.Sprintf("helloB %s my age:%d", b.FieldAName, age)
}

func TestNewStruct(t *testing.T) {
	inst := &TestB{}

	obj, err := New(inst)

	if err != nil {
		t.Error(err)
	}

	t.Log("inst fields:", obj.Fields())

	t.Log("inst methods:", obj.Methods())

	obj.SetFieldVal("FieldBName", "my name is b")
	t.Log("instB:", inst)

	s, err := obj.Call("HelloB", 33)
	if err != nil {
		t.Error(err)
	}
	t.Log("call HelloB result:", s)

	obj.SetFieldVal("FieldAName", "my name is a")
	t.Log("instB:", inst)

	s, err = obj.Call("HelloA")
	if err != nil {
		t.Error(err)
	}
	t.Log("call HelloA result:", s)
}
