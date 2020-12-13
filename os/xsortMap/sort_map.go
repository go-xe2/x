/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-08-20 16:50
* Description:
*****************************************************************/

package xsortMap

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Smap []*SortMapNode

type SortMapNode struct {
	Key string
	Val interface{}
}

func (c *Smap) Put(key string, val interface{}) {
	index, _, ok := c.get(key)
	if ok {
		(*c)[index].Val = val
	} else {
		node := &SortMapNode{Key: key, Val: val}
		*c = append(*c, node)
	}
}

func (c *Smap) Get(key string) (interface{}, bool) {
	_, val, ok := c.get(key)
	return val, ok
}

func (c *Smap) get(key string) (int, interface{}, bool) {
	for index, node := range *c {
		if node.Key == key {
			return index, node.Val, true
		}
	}
	return -1, nil, false
}

func (c *Smap) MarshalJSON() ([]byte, error) {
	s := toSortedMapJson(c)
	return json.Marshal(s)
}

func toSortedMapJson(smap *Smap) string {
	s := "{"
	for _, node := range *smap {
		v := node.Val
		isSamp := false
		str := ""
		switch v.(type) {
		case *Smap:
			isSamp = true
			str = toSortedMapJson(v.(*Smap))
		}

		if !isSamp {
			b, _ := json.Marshal(node.Val)
			str = string(b)
		}

		s = fmt.Sprintf("%s\"%s\":%s,", s, node.Key, str)
	}
	s = strings.TrimRight(s, ",")
	s = fmt.Sprintf("%s}", s)
	return s
}
