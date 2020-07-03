package xentity

import (
	"github.com/go-xe2/x/xf/ef/xqi"
	"reflect"
)

var (
	efStringType  = reflect.TypeOf([]xqi.EFString{}).Elem()
	efIntType     = reflect.TypeOf([]xqi.EFInt{}).Elem()
	efInt8Type    = reflect.TypeOf([]xqi.EFInt8{}).Elem()
	efInt16Type   = reflect.TypeOf([]xqi.EFInt16{}).Elem()
	efInt32Type   = reflect.TypeOf([]xqi.EFInt32{}).Elem()
	efInt64Type   = reflect.TypeOf([]xqi.EFInt64{}).Elem()
	efUintType    = reflect.TypeOf([]xqi.EFUint{}).Elem()
	efUint8Type   = reflect.TypeOf([]xqi.EFUint8{}).Elem()
	efUint16Type  = reflect.TypeOf([]xqi.EFUint16{}).Elem()
	efUint32Type  = reflect.TypeOf([]xqi.EFUint32{}).Elem()
	efUint64Type  = reflect.TypeOf([]xqi.EFUint64{}).Elem()
	efFloatType   = reflect.TypeOf([]xqi.EFFloat{}).Elem()
	efDoubleType  = reflect.TypeOf([]xqi.EFDouble{}).Elem()
	efBoolType    = reflect.TypeOf([]xqi.EFBool{}).Elem()
	efDateType    = reflect.TypeOf([]xqi.EFDate{}).Elem()
	efByteType    = reflect.TypeOf([]xqi.EFByte{}).Elem()
	efBinaryType  = reflect.TypeOf([]xqi.EFBinary{}).Elem()
	efForeignType = reflect.TypeOf([]xqi.EFForeign{}).Elem()
	efExprType    = reflect.TypeOf([]xqi.EFExpr{}).Elem()
)
