package xproto

import "time"

type XNetProtocol interface {
	ProtocolWriter
	ProtocolReader
}

type ProtocolWriter interface {
	WriteAny(val interface{}) (size int64, err error)
	WriteNull() (size int64, err error)
	WriteStr(val string) (size int64, err error)
	WriteInt8(val int8) (size int64, err error)
	WriteInt16(val int16) (size int64, err error)
	WriteInt32(val int32) (size int64, err error)
	WriteInt64(val int64) (size int64, err error)
	WriteDouble(val float64) (size int64, err error)
	WriteBool(val bool) (size int64, err error)
	WriteDatetime(val time.Time) (size int64, err error)

	WriteMap(val interface{}) (size int64, err error)
	WriteList(val interface{}) (size int64, err error)
	WriteDataset(val ProtoDataset) (size int64, err error)

	WriteEnum(val ProtoEnum) (size int64, err error)
	WriteStruct(val ProtoStruct) (size int64, err error)
	WriteUnion(val ProtoUnion) (size int64, err error)
	WriteClass(val ProtoClass) (size int64, err error)
	WriteException(val ProtoException) (size int64, err error)
}

type ProtocolReader interface {
	ReadAny() (val interface{}, err error)
	ReadType() (dType ProtoDataType)
	ReadStr() (val string, err error)
	ReadInt8() (val int8, err error)
	ReadInt16() (val int16, err error)
	ReadInt32() (val int32, err error)
	ReadInt64() (val int64, err error)
	ReadDouble() (val float64, err error)
	ReadBool() (val bool, err error)
	ReadDatetime() (val time.Time, err error)

	ReadMap() (val interface{}, err error)
	ReadList() (val interface{}, err error)

	ReadDataset() (ds ProtoDataset, err error)
	ReadEnum() (enum ProtoEnum, err error)
	ReadStruct() (st ProtoStruct, err error)
	ReadUnion() (un ProtoUnion, err error)
	ReadClass() (cls ProtoClass, err error)
	ReadException() (exp ProtoException, err error)
}
