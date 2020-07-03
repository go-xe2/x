package xparser

import (
	"github.com/go-xe2/x/encoding/xjson"
)

func New(value interface{}, unsafe ...bool) *TParser {
	return &TParser{xjson.New(value, unsafe...)}
}

func NewUnsafe(value ...interface{}) *TParser {
	if len(value) > 0 {
		return &TParser{xjson.New(value[0], false)}
	}
	return &TParser{xjson.New(nil, false)}
}

func Load(path string, unsafe ...bool) (*TParser, error) {
	if j, e := xjson.Load(path, unsafe...); e == nil {
		return &TParser{j}, nil
	} else {
		return nil, e
	}
}

func LoadContent(data interface{}, unsafe ...bool) (*TParser, error) {
	if j, e := xjson.LoadContent(data, unsafe...); e == nil {
		return &TParser{j}, nil
	} else {
		return nil, e
	}
}
