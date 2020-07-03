package xyaml

import "github.com/go-xe2/third/github.com/ghodss/yaml"

func Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func Decode(v []byte) (interface{}, error) {
	var result interface{}
	if err := yaml.Unmarshal(v, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func DecodeTo(v []byte, result interface{}) error {
	return yaml.Unmarshal(v, &result)
}

func ToJson(v []byte) ([]byte, error) {
	return yaml.YAMLToJSON(v)
}
