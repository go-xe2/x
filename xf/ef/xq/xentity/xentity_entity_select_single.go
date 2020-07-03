package xentity

import (
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/encoding/xxml"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type tEntitySelectSingle struct {
	sel *tEntitySelect
}

var _ EntitySelectSingle = (*tEntitySelectSingle)(nil)

func newEntitySelectSingle(entSelect *tEntitySelect) *tEntitySelectSingle {
	return &tEntitySelectSingle{
		sel: entSelect,
	}
}

func (ess *tEntitySelectSingle) Map() (data map[string]interface{}, err error) {
	rows, err := ess.sel.Top(1).List().Rows()
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	return rows[0], nil
}

func (ess *tEntitySelectSingle) Xml() (data xxml.XmlStr, err error) {
	rows, err := ess.sel.Top(1).List().Xml()
	if err != nil {
		return "", err
	}
	return rows, nil
}

func (ess *tEntitySelectSingle) Json() (data xjson.JsonStr, err error) {
	rows, err := ess.sel.Top(1).List().Json()
	if err != nil {
		return "", err
	}
	return rows, nil
}

func (ess *tEntitySelectSingle) Dataset() (data Dataset, err error) {
	rows, err := ess.sel.Top(1).List().Dataset()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (ess *tEntitySelectSingle) Slice() (data []interface{}, err error) {
	rows, err := ess.sel.Top(1).List().Slice()
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	return rows[0], nil
}

func (ess *tEntitySelectSingle) Visit(visitor QueryBinderVisit) (data interface{}, err error) {
	v, err := ess.sel.Top(1).List().Visit(visitor)
	if err != nil {
		return nil, err
	}
	if rows, ok := v.([]interface{}); ok {
		data = rows[0]
	} else {
		data = v
	}
	return
}

func (ess *tEntitySelectSingle) Convert() EntitySelectConvert {
	return ess.sel
}
