package xsafeMap

import "testing"

func TestNewAnyAnyMapFrom(t *testing.T) {
	m1 := NewAnyAnyMapFrom(map[interface{}]interface{}{
		"a": "1",
		"b": 2,
	})
	m2 := NewAnyAnyMapFrom(map[interface{}]interface{}{
		"c": 3,
		"d": 4,
	})
	m1.Merge(m2)
	t.Log("m1:", m1)
	t.Log("m2:", m2)
}
