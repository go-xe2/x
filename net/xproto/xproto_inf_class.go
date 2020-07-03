package xproto

type ProtoClass interface {
	ProtoSerializer
	ProtoUnSerializer
	ClassName() string
}
