package xentity

import (
	"github.com/go-xe2/x/xf/ef/xqi"
)

func registerInternalType(entry *tEFTypeConstructorEntry) {
	// 注册字符串类型字段
	entry.register(xqi.FDTString, newEFString)

	entry.register(xqi.FDTBool, newEFBool)

	entry.register(xqi.FDTInt, newEFInt)

	entry.register(xqi.FDTInt8, newEFInt8)

	entry.register(xqi.FDTInt16, newEFInt16)

	entry.register(xqi.FDTInt32, newEFInt32)

	entry.register(xqi.FDTInt64, newEFInt64)

	entry.register(xqi.FDTUint, newEFUint)

	entry.register(xqi.FDTUint8, newEFUint8)

	entry.register(xqi.FDTUint16, newEFUint16)

	entry.register(xqi.FDTUint32, newEFUint32)

	entry.register(xqi.FDTUint64, newEFUint64)

	entry.register(xqi.FDTFloat, newEFFloat)

	entry.register(xqi.FDTDouble, newEFDouble)

	entry.register(xqi.FDTDatetime, newEFDatetime)

	entry.register(xqi.FDTByte, newEFByte)

	entry.register(xqi.FDTBinary, newEFBinary)
}
