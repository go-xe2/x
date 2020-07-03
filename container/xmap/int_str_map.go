package xmap

type TIntStrMap map[int]string

func NewIntStrMap(mp ...map[int]string) TIntStrMap {
	var def map[int]string
	if len(mp) > 0 {
		def = mp[0]
	} else {
		def = make(map[int]string)
	}
	return TIntStrMap(def)
}

func (mp TIntStrMap) Map() map[int]string {
	return mp
}

func (mp TIntStrMap) ToArray() [][]interface{} {
	result := make([][]interface{}, 0)
	for k, v := range mp {
		result = append(result, []interface{}{k, v})
	}
	return result
}

func (mp TIntStrMap) GetString(key int, def ...string) string {
	s := ""
	if len(def) > 0 {
		s = def[0]
	}
	if s1, ok := mp[key]; ok {
		return s1
	}
	return s
}
