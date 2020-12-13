/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-08-20 10:04
* Description:
*****************************************************************/

package xstream

type TNodeType int8

const (
	STR_NODE TNodeType = iota
	BOOL_NODE
	INT8_NODE
	INT16_NODE
	INT32_NODE
	INT64_NODE
	FLOAT32_NODE
	FLOAT64_NODE
	BYTES_NODE
	BEGIN_NODE
	END_NODE
	UNKNOWN_NODE
)

func (nt TNodeType) String() string {
	switch nt {
	case STR_NODE:
		return "str_node"
	case BOOL_NODE:
		return "bool_node"
	case INT8_NODE:
		return "int8_node"
	case INT16_NODE:
		return "int16_node"
	case INT32_NODE:
		return "int32_node"
	case INT64_NODE:
		return "int64_node"
	case FLOAT32_NODE:
		return "float32_node"
	case FLOAT64_NODE:
		return "float64_node"
	case BYTES_NODE:
		return "bytes_node"
	case BEGIN_NODE:
		return "begin_node"
	case END_NODE:
		return "end_node"
	default:
		return "unknown_node"
	}
}
