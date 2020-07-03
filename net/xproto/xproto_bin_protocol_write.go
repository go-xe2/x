package xproto

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/xbinary"
	"github.com/go-xe2/x/type/xtime"
	"reflect"
	"time"
)

func (pro *TBinProtocol) WriteAny(val interface{}) (size int64, err error) {
	if val == nil {
		return pro.WriteNull()
	}
	switch v := val.(type) {
	case string:
		return pro.WriteStr(v)
	case int:
		return pro.WriteInt64(int64(v))
	case int8:
		return pro.WriteInt8(v)
	case int16:
		return pro.WriteInt16(v)
	case int32:
		return pro.WriteInt32(v)
	case int64:
		return pro.WriteInt64(v)
	case uint:
		return pro.WriteInt64(int64(v))
	case uint8:
		return pro.WriteInt8(int8(v))
	case uint16:
		return pro.WriteInt16(int16(v))
	case uint32:
		return pro.WriteInt32(int32(v))
	case uint64:
		return pro.WriteInt64(int64(v))
	case float32:
		return pro.WriteDouble(float64(v))
	case float64:
		return pro.WriteDouble(v)
	case bool:
		return pro.WriteBool(v)
	default:
		if tm, ok := val.(time.Time); ok {
			return pro.WriteDatetime(tm)
		} else if tm, ok := val.(*time.Time); ok {
			return pro.WriteDatetime(*tm)
		} else if tm, ok := val.(xtime.Time); ok {
			return pro.WriteDatetime(tm.Time)
		} else if tm, ok := val.(*xtime.Time); ok {
			return pro.WriteDatetime(tm.Time)
		} else if ds, ok := val.(ProtoDataset); ok {
			return pro.WriteDataset(ds)
		} else if em, ok := val.(ProtoEnum); ok {
			return pro.WriteEnum(em)
		} else if st, ok := val.(ProtoStruct); ok {
			return pro.WriteStruct(st)
		} else if st, ok := val.(ProtoClass); ok {
			return pro.WriteClass(st)
		} else if st, ok := val.(ProtoException); ok {
			return pro.WriteException(st)
		} else {
			vv := reflect.ValueOf(val)
			kind := vv.Kind()
			for kind == reflect.Ptr {
				vv = vv.Elem()
				kind = vv.Kind()
			}
			if kind == reflect.Map {
				return pro.WriteMap(val)
			} else if kind == reflect.Slice || kind == reflect.Array {
				return pro.WriteList(val)
			} else {
				return 0, exception.Newf("不能序列化数据类型：%s", reflect.TypeOf(val))
			}
		}
	}
}

func (pro *TBinProtocol) writeType(dtType ProtoDataType) (size int64, err error) {
	return write(pro.writer, xbinary.BeEncodeInt8(int8(dtType)))
}

func (pro *TBinProtocol) WriteNull() (size int64, err error) {
	return pro.writeType(PDTNull)
}

func (pro *TBinProtocol) WriteStr(val string) (size int64, err error) {
	size = 0
	if n, err := pro.writeType(PDTString); err != nil {
		return size, err
	} else {
		size += n
	}
	nLen := len(val)
	if n, err := write(pro.writer, xbinary.BeEncodeInt64(int64(nLen))); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := write(pro.writer, xbinary.BeEncodeString(val)); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteInt8(val int8) (size int64, err error) {
	size = 0
	if n, err := pro.writeType(PDTInt8); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := write(pro.writer, xbinary.BeEncodeInt8(val)); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteInt16(val int16) (size int64, err error) {
	size = 0
	if n, err := pro.writeType(PDTInt16); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := write(pro.writer, xbinary.BeEncodeInt16(val)); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteInt32(val int32) (size int64, err error) {
	size = 0
	if n, err := pro.writeType(PDTInt32); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := write(pro.writer, xbinary.BeEncodeInt32(val)); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteInt64(val int64) (size int64, err error) {
	size = 0
	if n, err := pro.writeType(PDTInt64); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := write(pro.writer, xbinary.BeEncodeInt64(val)); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteDouble(val float64) (size int64, err error) {
	size = 0
	if n, err := pro.writeType(PDTDouble); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := write(pro.writer, xbinary.BeEncodeFloat64(val)); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteBool(val bool) (size int64, err error) {
	size = 0
	if n, err := pro.writeType(PDTBool); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := write(pro.writer, xbinary.BeEncodeBool(val)); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteDatetime(val time.Time) (size int64, err error) {
	size = 0
	if n, err := pro.writeType(PDTDatetime); err != nil {
		return size, err
	} else {
		size += n
	}
	tm := val.Unix()
	if n, err := write(pro.writer, xbinary.BeEncodeInt64(tm)); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteMap(val interface{}) (size int64, err error) {
	size = 0
	if val == nil {
		return 0, nil
	}
	vv := reflect.ValueOf(val)
	kind := vv.Kind()
	for kind == reflect.Ptr {
		vv = vv.Elem()
		kind = vv.Kind()
	}
	if kind != reflect.Map {
		return 0, exception.Newf("数据类型:%s不是map类型", vv.Type())
	}
	keyType := vv.Type().Key()
	valType := vv.Type().Elem()
	keyProtoTypes := GetProtoDataTypes(keyType)
	valProtoTypes := GetProtoDataTypes(valType)
	keyTypeCount := len(keyProtoTypes)
	valTypeCount := len(valProtoTypes)

	if keyTypeCount == 0 || keyProtoTypes[0] == PDTUnknown {
		return size, exception.Newf("map键数据类型%s不能序列化", keyType)
	}

	if valTypeCount == 0 || valProtoTypes[0] == PDTUnknown {
		return size, exception.Newf("map值数据类型%s不能序列化", valType)
	}

	// 数据类型
	if n, err := pro.writeType(PDTMap); err != nil {
		return size, err
	} else {
		size += n
	}

	// 键类型
	if n, err := write(pro.writer, xbinary.BeEncodeInt32(int32(keyTypeCount))); err != nil {
		return size, err
	} else {
		size += n
	}
	for i := 0; i < keyTypeCount; i++ {
		if n, err := pro.writeType(keyProtoTypes[i]); err != nil {
			return size, err
		} else {
			size += n
		}
	}

	// 值类型
	if n, err := write(pro.writer, xbinary.BeEncodeInt32(int32(valTypeCount))); err != nil {
		return size, err
	} else {
		size += n
	}
	for i := 0; i < valTypeCount; i++ {
		if n, err := pro.writeType(valProtoTypes[i]); err != nil {
			return size, err
		} else {
			size += n
		}
	}

	keys := vv.MapKeys()
	// map键名数
	nLen := len(keys)
	if n, err := write(pro.writer, xbinary.BeEncodeInt64(int64(nLen))); err != nil {
		return size, err
	} else {
		size += n
	}
	// 键值对列表
	for _, k := range keys {
		key := k.Interface()
		value := vv.MapIndex(k).Interface()
		if n, err := pro.WriteAny(key); err != nil {
			return size, err
		} else {
			size += n
		}
		if n, err := pro.WriteAny(value); err != nil {
			return size, err
		} else {
			size += n
		}
	}
	return
}

func (pro *TBinProtocol) WriteList(val interface{}) (size int64, err error) {
	size = 0
	if val == nil {
		return 0, nil
	}
	vv := reflect.ValueOf(val)
	kind := vv.Kind()
	for kind == reflect.Ptr {
		vv = vv.Elem()
		kind = vv.Kind()
	}
	if kind != reflect.Slice && kind != reflect.Array {
		return size, exception.Newf("数据类型:%s不是slice或array类型", vv.Type())
	}
	elemType := vv.Type().Elem()
	elemProtoTypes := GetProtoDataTypes(elemType)
	elemTypeCount := len(elemProtoTypes)

	if elemTypeCount == 0 || elemProtoTypes[0] == PDTUnknown {
		return size, exception.Newf("slice数据项类型%s不支持序列化", elemType)
	}
	// 数据类型
	if n, err := pro.writeType(PDTList); err != nil {
		return size, err
	} else {
		size += n
	}

	// 保存list项数据类型
	if n, err := write(pro.writer, xbinary.BeEncodeInt32(int32(elemTypeCount))); err != nil {
		return size, err
	} else {
		size += n
	}
	for i := 0; i < elemTypeCount; i++ {
		if n, err := pro.writeType(elemProtoTypes[i]); err != nil {
			return size, err
		} else {
			size += n
		}
	}

	// 列表项行数
	nLen := vv.Len()
	if n, err := write(pro.writer, xbinary.BeEncodeInt64(int64(nLen))); err != nil {
		return size, err
	} else {
		size += n
	}
	for i := 0; i < nLen; i++ {
		value := vv.Index(i).Interface()
		if n, err := pro.WriteAny(value); err != nil {
			return size, err
		} else {
			size += n
		}
	}
	return
}

func (pro *TBinProtocol) WriteDataset(val ProtoDataset) (size int64, err error) {
	if val == nil {
		return 0, nil
	}
	size = 0
	if n, err := pro.writeType(PDTDataset); err != nil {
		return size, err
	} else {
		size += n
	}
	dsName := val.DatasetName()
	if n, err := pro.WriteStr(dsName); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := val.Serialize(pro); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteEnum(val ProtoEnum) (size int64, err error) {
	if val == nil {
		return 0, nil
	}
	clsName := val.EnumType()
	size = 0
	if n, err := pro.writeType(PDTEnum); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := pro.WriteStr(clsName); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := val.Serialize(pro); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteStruct(val ProtoStruct) (size int64, err error) {
	if val == nil {
		return 0, nil
	}
	clsName := val.StructName()
	size = 0
	if n, err := pro.writeType(PDTStruct); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := pro.WriteStr(clsName); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := val.Serialize(pro); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteUnion(val ProtoUnion) (size int64, err error) {
	if val == nil {
		return 0, nil
	}
	clsName := val.UnionName()
	size = 0
	if n, err := pro.writeType(PDTUnion); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := pro.WriteStr(clsName); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := val.Serialize(pro); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteClass(val ProtoClass) (size int64, err error) {
	if val == nil {
		return 0, nil
	}
	clsName := val.ClassName()
	size = 0
	if n, err := pro.writeType(PDTClass); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := pro.WriteStr(clsName); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := val.Serialize(pro); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}

func (pro *TBinProtocol) WriteException(val ProtoException) (size int64, err error) {
	if val == nil {
		return 0, nil
	}
	clsName := val.ExceptName()
	size = 0
	if n, err := pro.writeType(PDTException); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := pro.WriteStr(clsName); err != nil {
		return size, err
	} else {
		size += n
	}
	if n, err := val.Serialize(pro); err != nil {
		return size, err
	} else {
		size += n
	}
	return
}
