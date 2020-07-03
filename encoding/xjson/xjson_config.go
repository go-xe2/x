package xjson

// 设置访问子级连接符，默认为"."
func (j *TJson) SetSplitChar(char byte) {
	j.mu.Lock()
	j.c = char
	j.mu.Unlock()
}

// 设置是否允许通过带"."的键名访问子级,默认false
func (j *TJson) SetViolenceCheck(enabled bool) {
	j.mu.Lock()
	j.vc = enabled
	j.mu.Unlock()
}

// json键值使用的引号
func JsonQuotes() string {
	return mJSON_VALUE_QUOTES
}
