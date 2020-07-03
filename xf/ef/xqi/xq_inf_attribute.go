package xqi

type XqAttribute interface {
	// 属性名称
	AttrName() string
	String() string
	This() interface{}
}

type XqAttributeMap interface {
	AttrName() string
	Map() map[string]interface{}
	FromMap(mp map[string]interface{})
}
