package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tFieldValid struct {
	cate []string
	rule string
	msg  string
	op   []string
}

var _ xqi.FieldValid = (*tFieldValid)(nil)

func NewFieldValid(rule, msg string, op string, category ...string) xqi.FieldValid {
	inst := new(tFieldValid)
	inst.msg = msg
	inst.rule = rule
	inst.op = make([]string, 0)
	tmp := xstring.Split(op, "|")
	for _, s := range tmp {
		s = xstring.Trim(s)
		if s == "" {
			continue
		}
		inst.op = append(inst.op, s)
	}
	inst.cate = make([]string, 0)
	if len(category) > 0 {
		tmp := xstring.Split(category[0], "|")
		for _, s := range tmp {
			s = xstring.Trim(s)
			if s != "" {
				inst.cate = append(inst.cate, s)
			}
		}
	}
	return inst
}

func (fv *tFieldValid) Cate() []string {
	return fv.cate
}

func (fv *tFieldValid) Rule() string {
	return fv.rule
}

func (fv *tFieldValid) Msg() string {
	return fv.msg
}

func (fv *tFieldValid) Operation() []string {
	return fv.op
}

func (fv *tFieldValid) AttrName() string {
	return "fieldValid"
}

func (fv *tFieldValid) MakeValidString(fieldName string) string {
	s := ""
	if len(fieldName) > 0 {
		s = fmt.Sprintf("%s@%s", fieldName, fv.rule)
	} else {
		s = fv.rule
	}
	if fv.msg != "" {
		s += "#" + fv.msg
	}
	return s
}

func (fv *tFieldValid) Map() map[string]interface{} {
	return map[string]interface{}{
		"cate": xstring.Join(fv.cate, "|"),
		"rule": fv.rule,
		"msg":  fv.msg,
		"op":   xstring.Join(fv.op, "|"),
	}
}

func (fv *tFieldValid) This() interface{} {
	return fv
}

func (fv *tFieldValid) String() string {
	return fmt.Sprintf("cate:%s,rule:%s,msg:%s,op:%s", xstring.Join(fv.cate, "|"), fv.rule, fv.msg, xstring.Join(fv.op, "|"))
}
