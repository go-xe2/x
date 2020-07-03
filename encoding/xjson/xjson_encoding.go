package xjson

import (
	"encoding/json"
	"github.com/go-xe2/x/encoding/xtoml"
	"github.com/go-xe2/x/encoding/xxml"
	"github.com/go-xe2/x/encoding/xyaml"
)

func (j *TJson) ToXml(rootTag ...string) ([]byte, error) {
	return xxml.Encode(j.ToMap(), rootTag...)
}

func (j *TJson) ToXmlString(rootTag ...string) (string, error) {
	b, e := j.ToXml(rootTag...)
	return string(b), e
}

func (j *TJson) ToXmlIndent(rootTag ...string) ([]byte, error) {
	return xxml.EncodeWithIndent(j.ToMap(), rootTag...)
}

func (j *TJson) ToXmlIndentString(rootTag ...string) (string, error) {
	b, e := j.ToXmlIndent(rootTag...)
	return string(b), e
}

func (j *TJson) ToJson() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return Encode(*(j.p))
}

func (j *TJson) ToJsonString() (string, error) {
	b, e := j.ToJson()
	return string(b), e
}

func (j *TJson) ToJsonIndent() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return json.MarshalIndent(*(j.p), "", "\t")
}

func (j *TJson) ToJsonIndentString() (string, error) {
	b, e := j.ToJsonIndent()
	return string(b), e
}

func (j *TJson) ToYaml() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return xyaml.Encode(*(j.p))
}

func (j *TJson) ToYamlString() (string, error) {
	b, e := j.ToYaml()
	return string(b), e
}

func (j *TJson) ToToml() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return xtoml.Encode(*(j.p))
}

func (j *TJson) ToTomlString() (string, error) {
	b, e := j.ToToml()
	return string(b), e
}
