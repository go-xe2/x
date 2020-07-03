package xparser

// 设置子层访问符，默认为"."
func (p *TParser) SetSplitChar(char byte) {
	p.json.SetSplitChar(char)
}

// 是否启用访问子层, 默认false
func (p *TParser) SetViolenceCheck(check bool) {
	p.json.SetViolenceCheck(check)
}
