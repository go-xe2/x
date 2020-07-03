package xhash

import "testing"

func TestMap(t *testing.T) {
	var map1 = THashMap{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"name1": "张三",
				"name2": "李四",
			},
			"age": 33,
		},
		"age": 88,
	}

	s1 := GetPathInterface(map1, "a/b/name1")
	t.Log("name1:", s1)
	age1 := GetPathInterface(map1, "a/age")
	t.Log("age1:", age1)
	age2 := GetPathInt(map1, "a/age")
	t.Log("age2:", age2)

	name2 := GetPathString(map1, "a/b/name2")
	t.Log("name2:", name2)
	m2 := GetPathMap(map1, "a/b")
	t.Log("m2:", m2)
}
