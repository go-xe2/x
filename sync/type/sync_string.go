package _type

import "sync/atomic"

type TString struct {
	value atomic.Value
}

func NewString(v ...string) *TString {
	inst := &TString{}
	if len(v) > 0 {
		inst.value.Store(v[0])
	}
	return inst
}

func (s *TString) Clone() *TString {
	return NewString(s.Val())
}

func (s *TString) Set(v string) string {
	old := s.Val()
	s.value.Store(v)
	return old
}

func (s *TString) Val() string {
	v := s.value.Load()
	if v != nil {
		return v.(string)
	}
	return ""
}
