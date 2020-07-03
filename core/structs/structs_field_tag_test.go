package structs

import "testing"

func TestFieldTag_Parse(t *testing.T) {
	fTag := NewFieldTag()
	fTag.Parse(`json:"json_json_json22" c:"cccccc" e:"eeeeee" d:"33" f:"43"`)
	t.Log("tag:", fTag.items)
	t.Log("tag str:", fTag)
}
