package xcmd

import (
	_type "github.com/go-xe2/x/sync/type"
)

type tCmdOption struct {
	options map[string]string
}

func (c *tCmdOption) GetAll() map[string]string {
	return c.options
}

func (c *tCmdOption) Get(key string, def ...string) string {
	if option, ok := c.options[key]; ok {
		return option
	} else if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (c *tCmdOption) GetVar(key string, def ...string) *_type.TVar {
	return _type.NewVar(c.Get(key, def...), true)
}
