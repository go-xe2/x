package xproto

type ProtoSerializer interface {
	Serialize(pro ProtocolWriter) (size int64, err error)
}

type ProtoUnSerializer interface {
	UnSerialize(pro ProtocolReader) (err error)
}
