package xproto

type ProtoEnum interface {
	ProtoSerializer
	ProtoUnSerializer
	EnumType() string
	Items() ProtoEnum
	Value() int64
}
