package xproto

import (
	"bytes"
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"reflect"
	"testing"
	"time"
)

func TestNewBinProtocol(t *testing.T) {
	var map1 = map[int]string{
		1: "map value1",
		2: "map value2",
		3: "map value3",
	}

	var map2 = map[string]interface{}{
		"key1": "map str value",
		"key2": 123,
		"key3": int64(44444),
		"key4": 123.33,
		"key5": true,
		"key6": []int{1, 2, 3},
	}

	var slice1 = []string{
		"slice value1",
		"slice value2",
		"slice value3",
	}
	var slice2 = []int{
		10,
		11,
		12,
		13,
		14,
	}
	var slice3 = []interface{}{
		"any slice str value",
		101,
		true,
		101.33,
		time.Now(),
	}

	var slice4 = []ProtoClass{nil}

	fmt.Println("data:", map1, slice1, slice2, slice3)

	slice3Type := reflect.TypeOf(slice3).Elem()
	slice4Type := reflect.TypeOf(slice4).Elem()

	fmt.Println("slice3 elem type:", slice3Type, ", kind:", slice3Type.Kind(), ", name:", slice3Type.Name() == "")
	fmt.Println("slice4 elem type:", slice4Type, ", kind:", slice4Type.Kind(), ", name:", slice4Type.Name())

	slice3ProTypes := GetProtoDataTypes(reflect.TypeOf(slice3))
	fmt.Println("slice3 proto types:", slice3ProTypes)

	slice4ProTypes := GetProtoDataTypes(reflect.TypeOf(slice4))
	fmt.Println("slice4 proto types:", slice4ProTypes)

	slice3TypeRestore, err := NewTypeByProtoDataTypes(slice3ProTypes)
	fmt.Println("restore slice3Type:", slice3TypeRestore, ", err:", err)

	slice4TypeRestore, err := NewTypeByProtoDataTypes(slice4ProTypes)
	fmt.Println("restore slice4Type:", slice4TypeRestore, ", err:", err)

	map2ProTypes := GetProtoDataTypes(reflect.TypeOf(map2))
	fmt.Println("map2ProTypes:", map2ProTypes)

	map2TypeRestore, err := NewTypeByProtoDataTypes(map2ProTypes)
	fmt.Println("map2TypeRestore:", map2TypeRestore, ", err:", err)

	fmt.Println("================测试序列化数据为二进制流=============")
	writeBuf := make([]byte, 0)
	readBuf := make([]byte, 0)

	reader := bytes.NewReader(readBuf)
	writer := bytes.NewBuffer(writeBuf)

	proto := NewBinProtocol(writer, reader)

	nSize, err := proto.WriteAny(map2)

	if n, err := proto.WriteAny(map1); err != nil {
		t.Fatal(err)
	} else {
		nSize += n
	}

	if n, err := proto.WriteAny(slice1); err != nil {
		t.Fatal(err)
	} else {
		nSize += n
	}
	if n, err := proto.WriteAny(slice2); err != nil {
		t.Fatal(err)
	} else {
		nSize += n
	}

	if n, err := proto.WriteAny(slice3); err != nil {
		t.Fatal(err)
	} else {
		nSize += n
	}

	fmt.Println("proto write result size:", nSize)
	fmt.Println("buf binary data:", writer.Bytes(), ", size:", len(writer.Bytes()))

	fmt.Println("================读取数据==============")

	reader.Reset(writer.Bytes())

	restoreMap2, err := proto.ReadAny()
	fmt.Println("restoreMap2:", restoreMap2, ", err:", err)
	if e, ok := (interface{})(err).(*exception.Exception); ok {
		fmt.Println(e.Stack())
	}
	map1Restore, err := proto.ReadAny()
	fmt.Println("map1Restore:", map1Restore, ", err:", err)

	slice1Restore, err := proto.ReadAny()
	fmt.Println("slice1Restore:", slice1Restore, ", err:", err)

	slice2Restore, err := proto.ReadAny()
	fmt.Println("slice2Restore:", slice2Restore, ", err:", err)

	slice3Restore, err := proto.ReadAny()
	fmt.Println("slice3Restore:", slice3Restore, ", err:", err)

}
