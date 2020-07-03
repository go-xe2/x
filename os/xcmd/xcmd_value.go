package xcmd

import (
	_type "github.com/go-xe2/x/sync/type"
)

type tCmdValue struct {
	values []string
}

func (c *tCmdValue) GetAll() []string {
	return c.values
}

func (c *tCmdValue) Get(index int, def ...string) string {
	if index < len(c.values) {
		return c.values[index]
	} else if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (c *tCmdValue) GetVar(index int, def ...string) *_type.TVar {
	return _type.NewVar(c.Get(index, def...), true)
}
