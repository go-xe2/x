package xproto

type ProtoStruct interface {
	ProtoSerializer
	ProtoUnSerializer
	StructName() string
}
