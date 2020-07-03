package xparser

func (p *TParser) MarshalJSON() ([]byte, error) {
	return p.json.MarshalJSON()
}

func (p *TParser) ToXml(rootTag ...string) ([]byte, error) {
	return p.json.ToXml(rootTag...)
}

func (p *TParser) ToXmlIndent(rootTag ...string) ([]byte, error) {
	return p.json.ToXmlIndent(rootTag...)
}

func (p *TParser) ToJson() ([]byte, error) {
	return p.json.ToJson()
}

func (p *TParser) ToJsonString() (string, error) {
	return p.json.ToJsonString()
}

func (p *TParser) ToJsonIndent() ([]byte, error) {
	return p.json.ToJsonIndent()
}

func (p *TParser) ToJsonIndentString() (string, error) {
	return p.json.ToJsonIndentString()
}

func (p *TParser) ToYaml() ([]byte, error) {
	return p.json.ToYaml()
}

func (p *TParser) ToToml() ([]byte, error) {
	return p.json.ToToml()
}

func VarToXml(value interface{}, rootTag ...string) ([]byte, error) {
	return New(value).ToXml(rootTag...)
}

func VarToXmlIndent(value interface{}, rootTag ...string) ([]byte, error) {
	return New(value).ToXmlIndent(rootTag...)
}

func VarToJson(value interface{}) ([]byte, error) {
	return New(value).ToJson()
}

func VarToJsonString(value interface{}) (string, error) {
	return New(value).ToJsonString()
}

func VarToJsonIndent(value interface{}) ([]byte, error) {
	return New(value).ToJsonIndent()
}

func VarToJsonIndentString(value interface{}) (string, error) {
	return New(value).ToJsonIndentString()
}

func VarToYaml(value interface{}) ([]byte, error) {
	return New(value).ToYaml()
}

func VarToToml(value interface{}) ([]byte, error) {
	return New(value).ToToml()
}

func VarToStruct(value interface{}, obj interface{}) error {
	return New(value).ToStruct(obj)
}
