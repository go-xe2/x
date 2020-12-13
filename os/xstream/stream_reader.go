/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-08-20 09:55
* Description:
*****************************************************************/

package xstream

type StreamReader interface {
	ReadStr() (string, error)
	ReadInt8() (int8, error)
	ReadInt16() (int16, error)
	ReadInt32() (int32, error)
	ReadInt64() (int64, error)
	ReadFloat32() (float32, error)
	ReadFloat64() (float64, error)
	ReadBool() (bool, error)
	ReadBytes() ([]byte, error)
	ReadNode() (val interface{}, t TNodeType, err error)
}
