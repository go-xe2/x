package xbytes

import "github.com/go-xe2/x/core/exception"

// 拆分byte数组为多个数组
func Split(buf []byte, lim int) [][]byte {
	if lim == 0 {
		panic(exception.New("数组拆分大小不能为0"))
	}
	var chunk []byte
	bufLen := len(buf)
	chunkLen := bufLen/lim + 1
	chunks := make([][]byte, 0, chunkLen)
	for bufLen >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
		bufLen = len(buf)
	}
	if bufLen > 0 {
		chunks = append(chunks, buf[:bufLen])
	}
	return chunks
}
