package xproto

import "io"

type TBinProtocol struct {
	writer io.Writer
	reader io.ReadSeeker
	this   interface{}
}

var _ XNetProtocol = (*TBinProtocol)(nil)

const szDataStreamVerErrorMsg = "数据流格式错误或版本不正确"
const szErrorDataTypeReadMsg = "当前数据类型为%s非%s类型"

func NewBinProtocol(writer io.Writer, reader io.ReadSeeker) *TBinProtocol {
	inst := &TBinProtocol{
		writer: writer,
		reader: reader,
	}
	inst.this = inst
	return inst
}
