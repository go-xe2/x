/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-08-20 09:56
* Description:
*****************************************************************/

package xstream

type StreamWriter interface {
	WriteStr(str string) error
	WriteInt8(v int8) error
	WriteInt16(v int16) error
	WriteInt32(v int32) error
	WriteInt64(v int64) error
	WriteFloat32(v float32) error
	WriteFloat64(v float64) error
	WriteBool(v bool) error
	WriteBytes(buf []byte) error
}
