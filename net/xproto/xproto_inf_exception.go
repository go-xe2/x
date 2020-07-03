package xproto

type ProtoException interface {
	ProtoSerializer
	ProtoUnSerializer
	ExceptName() string
}
