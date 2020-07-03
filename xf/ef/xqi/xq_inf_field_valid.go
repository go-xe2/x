package xqi

type FieldValid interface {
	XqAttribute
	// 检查规则分组名，默认为空，表示所有分组中都使用该检查规则
	Cate() []string
	// 规则名称，如果是多个规则使用|分割
	Rule() string
	// 提交消息，多个规则对应消息时使用|分割, 如果为空，使用规则的默认提示
	Msg() string
	// 支持的操作类型，插入:I, 更新U, 多种操作使用|连接，如I|U, 默认为空表示同时支持插入和更新
	Operation() []string
	MakeValidString(fieldName string) string
}
