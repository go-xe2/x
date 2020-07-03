package xjson

type JsonStr string

func (js JsonStr) String() string {
	return string(js)
}

func (js JsonStr) MarshalJSON() (data []byte, err error) {
	return []byte(js), nil
}

func (js *JsonStr) UnmarshalJSON(data []byte) (err error) {
	*js = JsonStr(string(data))
	return nil
}
