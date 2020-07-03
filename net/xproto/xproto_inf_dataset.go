package xproto

type DSField interface {
	FieldName() string
	FieldType() ProtoDataType
	Value() interface{}
	Index() int
}

type ProtoDataset interface {
	ProtoSerializer
	ProtoUnSerializer
	DatasetName() string
	FieldNames() []string
	Rowcount() int64
	IsBof() bool
	IsEof() bool
	First() bool
	Next() bool
	Last() bool
	Prior() bool
	FieldValue(fieldIndex int) interface{}
	Field(index int) DSField
	FieldByName(name string) DSField
}
