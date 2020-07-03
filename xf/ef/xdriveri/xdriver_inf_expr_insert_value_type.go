package xdriveri

type ExprInsertValueType int

const (
	SqlInsertStaticValueType ExprInsertValueType = iota
	SqlInsertFromScriptType
)
