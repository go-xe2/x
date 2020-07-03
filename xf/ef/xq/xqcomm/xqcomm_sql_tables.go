package xqcomm

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TSqlTables struct {
	items    []SqlTable
	itemMaps map[string]SqlTable
}

var _ SqlTables = &TSqlTables{}

func NewSqlTables() *TSqlTables {
	return &TSqlTables{
		itemMaps: make(map[string]SqlTable),
		items:    make([]SqlTable, 0),
	}
}

func NewSqlTablesFromMap(mp map[string]SqlTable) *TSqlTables {
	inst := NewSqlTables()
	for _, tb := range mp {
		inst.Add(tb)
	}
	return inst
}

func (sts *TSqlTables) Table(name string) SqlTable {
	if v, ok := sts.itemMaps[name]; ok {
		return v
	}
	panic(exception.Newf("表达式中表实体%s不存在", name))
}

func (sts *TSqlTables) HasTable(name string) bool {
	if _, ok := sts.itemMaps[name]; ok {
		return true
	}
	return false
}

func (sts *TSqlTables) All() []SqlTable {
	return sts.items
}

func (sts *TSqlTables) Count() int {
	return len(sts.items)
}

func (sts *TSqlTables) Add(table SqlTable) SqlTables {
	if _, ok := sts.itemMaps[table.TableName()]; !ok {
		sts.itemMaps[table.TableName()] = table
		sts.items = append(sts.items, table)
	}
	if table.TableAlias() != "" {
		if v, ok := sts.itemMaps[table.TableAlias()]; ok {
			if v.TableName() != table.TableName() {
				panic(exception.Newf("表%s别名%s已经存在", v.TableName(), v.TableAlias()))
			}
		} else {
			sts.itemMaps[table.TableAlias()] = table
		}
	}
	return sts
}

func (sts *TSqlTables) AddTable(name string, alias ...string) SqlTable {
	result := NewSqlTable(nil, name, alias...)
	sts.Add(result)
	return result
}

func (sts *TSqlTables) Clone(target SqlTables) {
	items := target.All()
	for _, item := range items {
		sts.Add(item)
	}
}

func (sts *TSqlTables) Clear() SqlTables {
	sts.items = make([]SqlTable, 0)
	sts.itemMaps = make(map[string]SqlTable)
	return sts
}

func (sts *TSqlTables) This() interface{} {
	return sts
}

func (sts *TSqlTables) String() string {
	s := ""
	for k, v := range sts.itemMaps {
		s += fmt.Sprintf(", %s:%s", k, v.String())
	}
	return "{ " + s[1:] + " }"
}
