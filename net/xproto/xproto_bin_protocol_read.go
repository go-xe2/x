package xproto

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/xbinary"
	"io"
	"reflect"
	"time"
)

func (pro *TBinProtocol) ReadType() ProtoDataType {
	bytes, err := read(pro.reader, 1)
	if err != nil {
		return PDTUnknown
	}
	n := xbinary.BeDecodeToInt8(bytes)
	return ProtoDataType(n)
}

func (pro *TBinProtocol) readByType(pdt ProtoDataType) (val interface{}, err error) {
	switch pdt {
	case PDTNull:
		return nil, nil
	case PDTString:
		return pro.readStr()
	case PDTInt8:
		return pro.readInt8()
	case PDTInt16:
		return pro.readInt16()
	case PDTInt32:
		return pro.readInt32()
	case PDTInt64:
		return pro.readInt64()
	case PDTDouble:
		return pro.readDouble()
	case PDTBool:
		return pro.readBool()
	case PDTDatetime:
		return pro.readDatetime()
	case PDTList:
		return pro.readList()
	case PDTMap:
		return pro.readMap()
	case PDTDataset:
		return pro.readDataset()
	case PDTEnum:
		return pro.readEnum()
	case PDTStruct:
		return pro.readStruct()
	case PDTUnion:
		return pro.readUnion()
	case PDTClass:
		return pro.readClass()
	case PDTException:
		return pro.readException()
	}
	return nil, exception.Newf("未知数据类型，反序列化失败")
}

func (pro *TBinProtocol) ReadAny() (val interface{}, err error) {
	dType := pro.ReadType()
	if dType == PDTUnknown {
		return nil, exception.Newf("未知数据类型")
	}
	return pro.readByType(dType)
}

func (pro *TBinProtocol) readStr() (val string, err error) {
	bytes, err := read(pro.reader, 8)
	if err != nil {
		return "", err
	}
	nLen := xbinary.BeDecodeToInt64(bytes)
	data, err := read(pro.reader, nLen)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (pro *TBinProtocol) readInt8() (val int8, err error) {
	bytes, err := read(pro.reader, 1)
	if err != nil {
		return 0, err
	}
	return xbinary.BeDecodeToInt8(bytes), nil
}

func (pro *TBinProtocol) readInt16() (val int16, err error) {
	bytes, err := read(pro.reader, 2)
	if err != nil {
		return 0, err
	}
	return xbinary.BeDecodeToInt16(bytes), nil
}

func (pro *TBinProtocol) readInt32() (val int32, err error) {
	bytes, err := read(pro.reader, 4)
	if err != nil {
		return 0, err
	}
	return xbinary.BeDecodeToInt32(bytes), nil
}

func (pro *TBinProtocol) readInt64() (val int64, err error) {
	bytes, err := read(pro.reader, 8)
	if err != nil {
		return 0, err
	}
	return xbinary.BeDecodeToInt64(bytes), nil
}

func (pro *TBinProtocol) readDouble() (val float64, err error) {
	bytes, err := read(pro.reader, 8)
	if err != nil {
		return 0, err
	}
	return xbinary.BeDecodeToFloat64(bytes), nil
}

func (pro *TBinProtocol) readBool() (val bool, err error) {
	bytes, err := read(pro.reader, 1)
	if err != nil {
		return false, err
	}
	return xbinary.BeDecodeToBool(bytes), nil
}

func (pro *TBinProtocol) readDatetime() (val time.Time, err error) {
	bytes, err := read(pro.reader, 8)
	if err != nil {
		return time.Unix(0, 0), err
	}
	n := xbinary.BeDecodeToInt64(bytes)
	return time.Unix(n, 0), nil
}

func (pro *TBinProtocol) readList() (val interface{}, err error) {
	bytes, err := read(pro.reader, 4)
	if err != nil {
		return nil, err
	}
	// 读取list项类型
	nLen := int(xbinary.BeDecodeToInt32(bytes))
	proTypes := make([]ProtoDataType, nLen)
	for i := 0; i < nLen; i++ {
		if t := pro.ReadType(); t == PDTUnknown {
			return nil, exception.NewText(szDataStreamVerErrorMsg)
		} else {
			proTypes[i] = t
		}
	}
	// 创建list列表实例
	st, err := NewTypeByProtoDataTypes(proTypes)
	if err != nil {
		return nil, err
	}
	bytes, err = read(pro.reader, 8)
	if err != nil {
		return nil, err
	}

	sliceLen := int(xbinary.BeDecodeToInt64(bytes))
	listT := reflect.SliceOf(st)
	listV := reflect.MakeSlice(listT, sliceLen, sliceLen)

	// 读取列表项数据
	for i := 0; i < sliceLen; i++ {
		if v, err := pro.ReadAny(); err != nil {
			return nil, err
		} else {
			listV.Index(i).Set(reflect.ValueOf(v))
		}
	}
	return listV.Interface(), nil
}

func (pro *TBinProtocol) readMap() (val interface{}, err error) {
	// 读取key类型
	bytes, err := read(pro.reader, 4)
	if err != nil {
		return nil, err
	}
	keyTypeCount := int(xbinary.BeDecodeToInt32(bytes))
	keyProtoTypes := make([]ProtoDataType, keyTypeCount)
	for i := 0; i < keyTypeCount; i++ {
		if t := pro.ReadType(); t == PDTUnknown {
			return nil, exception.NewText(szDataStreamVerErrorMsg)
		} else {
			keyProtoTypes[i] = t
		}
	}
	// 读取val类型
	bytes, err = read(pro.reader, 4)
	if err != nil {
		return 0, err
	}
	valTypeCount := int(xbinary.BeDecodeToInt32(bytes))
	valProtoTypes := make([]ProtoDataType, valTypeCount)
	for i := 0; i < valTypeCount; i++ {
		if t := pro.ReadType(); t == PDTUnknown {
			return nil, exception.NewText(szDataStreamVerErrorMsg)
		} else {
			valProtoTypes[i] = t
		}
	}
	// 读取key-value键值对数
	bytes, err = read(pro.reader, 8)
	if err != nil {
		return nil, err
	}
	mapKeyCount := int(xbinary.BeDecodeToInt64(bytes))
	// 获取键值类型
	keyType, err := NewTypeByProtoDataTypes(keyProtoTypes)
	if err != nil {
		return nil, err
	}
	valType, err := NewTypeByProtoDataTypes(valProtoTypes)
	if err != nil {
		return nil, err
	}
	// 创建 map实例
	mapT := reflect.MapOf(keyType, valType)
	mapV := reflect.MakeMap(mapT)

	// 读取键值对
	for i := 0; i < mapKeyCount; i++ {
		// 读键名
		key, err := pro.ReadAny()
		if err != nil {
			return nil, err
		}
		val, err := pro.ReadAny()
		if err != nil {
			return nil, err
		}
		mapV.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	}
	return mapV.Interface(), nil
}

func (pro *TBinProtocol) readDataset() (val ProtoDataset, err error) {
	clsName, err := pro.readStr()
	if err != nil {
		return nil, err
	}
	inst := NewMemDataset(clsName)
	if inst == nil {
		return nil, exception.Newf("未注册dataset数据类型：%s", clsName)
	}
	if err = inst.UnSerialize(pro); err != nil {
		return
	}
	return inst, nil
}

func (pro *TBinProtocol) readEnum() (val ProtoEnum, err error) {
	clsName, err := pro.readStr()
	if err != nil {
		return nil, err
	}
	inst := NewEnum(clsName)
	if inst == nil {
		return nil, exception.Newf("未注册enum数据类型：%s", clsName)
	}
	if err = inst.UnSerialize(pro); err != nil {
		return
	}
	return inst, nil
}

func (pro *TBinProtocol) readStruct() (val ProtoStruct, err error) {
	clsName, err := pro.readStr()
	if err != nil {
		return nil, err
	}
	inst := NewStruct(clsName)
	if inst == nil {
		return nil, exception.Newf("未注册struct数据类型：%s", clsName)
	}
	if err = inst.UnSerialize(pro); err != nil {
		return
	}
	return inst, nil
}

func (pro *TBinProtocol) readUnion() (val ProtoUnion, err error) {
	clsName, err := pro.readStr()
	if err != nil {
		return nil, err
	}
	inst := NewUnionStruct(clsName)
	if inst == nil {
		return nil, exception.Newf("未注册union数据类型：%s", clsName)
	}
	if err = inst.UnSerialize(pro); err != nil {
		return
	}
	return inst, nil
}

func (pro *TBinProtocol) readClass() (val ProtoClass, err error) {
	clsName, err := pro.readStr()
	if err != nil {
		return nil, err
	}
	inst := NewClass(clsName)
	if inst == nil {
		return nil, exception.Newf("未注册class数据类型：%s", clsName)
	}
	if err = inst.UnSerialize(pro); err != nil {
		return
	}
	return inst, nil
}

func (pro *TBinProtocol) readException() (val ProtoException, err error) {
	clsName, err := pro.readStr()
	if err != nil {
		return nil, err
	}
	inst := NewException(clsName)
	if inst == nil {
		return nil, exception.Newf("未注册exception数据类型：%s", clsName)
	}
	if err = inst.UnSerialize(pro); err != nil {
		return
	}
	return inst, nil
}

func (pro *TBinProtocol) ReadStr() (val string, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return "", exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTString {
		return pro.readStr()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return "", err
		}
		return "", exception.Newf(szErrorDataTypeReadMsg, vt, PDTString)
	}
}

func (pro *TBinProtocol) ReadInt8() (val int8, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return 0, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTInt8 {
		return pro.readInt8()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return 0, err
		}
		return 0, exception.Newf(szErrorDataTypeReadMsg, vt, PDTInt8)
	}
}

func (pro *TBinProtocol) ReadInt16() (val int16, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return 0, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTInt16 {
		return pro.readInt16()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return 0, err
		}
		return 0, exception.Newf(szErrorDataTypeReadMsg, vt, PDTInt16)
	}
}

func (pro *TBinProtocol) ReadInt32() (val int32, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return 0, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTInt32 {
		return pro.readInt32()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return 0, err
		}
		return 0, exception.Newf(szErrorDataTypeReadMsg, vt, PDTInt32)
	}
}

func (pro *TBinProtocol) ReadInt64() (val int64, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return 0, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTInt64 {
		return pro.readInt64()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return 0, err
		}
		return 0, exception.Newf(szErrorDataTypeReadMsg, vt, PDTInt64)
	}
}

func (pro *TBinProtocol) ReadDouble() (val float64, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return 0, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTDouble {
		return pro.readDouble()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return 0, err
		}
		return 0, exception.Newf(szErrorDataTypeReadMsg, vt, PDTDouble)
	}
}

func (pro *TBinProtocol) ReadBool() (val bool, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return false, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTBool {
		return pro.readBool()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return false, err
		}
		return false, exception.Newf(szErrorDataTypeReadMsg, vt, PDTBool)
	}
}

func (pro *TBinProtocol) ReadDatetime() (val time.Time, err error) {
	defTime := time.Unix(0, 0)
	if vt := pro.ReadType(); vt == PDTUnknown {
		return defTime, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTDatetime {
		return pro.readDatetime()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return defTime, err
		}
		return defTime, exception.Newf(szErrorDataTypeReadMsg, vt, PDTDatetime)
	}
}

func (pro *TBinProtocol) ReadMap() (val interface{}, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return 0, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTMap {
		return pro.readMap()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return nil, err
		}
		return nil, exception.Newf(szErrorDataTypeReadMsg, vt, PDTMap)
	}
}

func (pro *TBinProtocol) ReadList() (val interface{}, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return 0, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTList {
		return pro.readList()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return nil, err
		}
		return nil, exception.Newf(szErrorDataTypeReadMsg, vt, PDTList)
	}
}

func (pro *TBinProtocol) ReadDataset() (ds ProtoDataset, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return nil, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTDataset {
		return pro.readDataset()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return nil, err
		}
		return nil, exception.Newf(szErrorDataTypeReadMsg, vt, PDTDataset)
	}
}

func (pro *TBinProtocol) ReadEnum() (enum ProtoEnum, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return nil, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTEnum {
		return pro.readEnum()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return nil, err
		}
		return nil, exception.Newf(szErrorDataTypeReadMsg, vt, PDTEnum)
	}
}

func (pro *TBinProtocol) ReadStruct() (st ProtoStruct, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return nil, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTStruct {
		return pro.readStruct()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return nil, err
		}
		return nil, exception.Newf(szErrorDataTypeReadMsg, vt, PDTStruct)
	}
}

func (pro *TBinProtocol) ReadUnion() (un ProtoUnion, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return nil, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTUnion {
		return pro.readUnion()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return nil, err
		}
		return nil, exception.Newf(szErrorDataTypeReadMsg, vt, PDTUnion)
	}
}

func (pro *TBinProtocol) ReadClass() (cls ProtoClass, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return nil, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTClass {
		return pro.readClass()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return nil, err
		}
		return nil, exception.Newf(szErrorDataTypeReadMsg, vt, PDTClass)
	}
}

func (pro *TBinProtocol) ReadException() (exp ProtoException, err error) {
	if vt := pro.ReadType(); vt == PDTUnknown {
		return nil, exception.NewText(szDataStreamVerErrorMsg)
	} else if vt == PDTException {
		return pro.readException()
	} else {
		if _, err := pro.reader.Seek(-1, io.SeekCurrent); err != nil {
			return nil, err
		}
		return nil, exception.Newf(szErrorDataTypeReadMsg, vt, PDTException)
	}
}
