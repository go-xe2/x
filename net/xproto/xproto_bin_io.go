package xproto

import (
	"github.com/go-xe2/x/core/exception"
	"io"
)

func write(writer io.Writer, data []byte) (size int64, err error) {
	dSize := len(data)
	if n, err := writer.Write(data); err != nil {
		return 0, err
	} else if n != dSize {
		return 0, exception.NewText("写数据失败")
	}
	return int64(dSize), nil
}

func read(reader io.Reader, size int64) (data []byte, err error) {
	data = make([]byte, size)
	if n, err := reader.Read(data); err != nil {
		return nil, err
	} else if int64(n) != size {
		return nil, exception.NewText("读取数据失败")
	}
	return data, nil
}
