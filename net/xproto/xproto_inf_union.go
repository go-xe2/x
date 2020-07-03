package xproto

type ProtoUnion interface {
	ProtoSerializer
	ProtoUnSerializer
	UnionName() string
}
