package xjson

import "testing"

func TestJsonStr_UnmarshalJSON(t *testing.T) {
	s := `{"name": "key1", value:"value1"}`
	var s1 JsonStr
	err := s1.UnmarshalJSON([]byte(s))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("s1.UnmarshalJSON:", s1)

	bytes, err := s1.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("s1.MarshalJSON bytes:", bytes)
	t.Log("s1.MarshalJSON:", string(bytes))
}
