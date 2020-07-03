package xfile

import (
	"io"
	"io/ioutil"
	"os"
)

const (
	mREAD_BUFFER = 1024
)

// 获取文件内容，返回字符串，文件不存在时返回空字符串
func GetContents(path string) string {
	return string(GetBinContents(path))
}

// 获取文件内容，文件不存在时返回nil
func GetBinContents(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return data
}

// 将二进制数据写入文件
func putContents(path string, data []byte, flag int, perm int) error {
	dir := Dir(path)
	if !Exists(dir) {
		if err := Mkdir(dir); err != nil {
			return err
		}
	}
	f, err := OpenWithFlagPerm(path, flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	if n, err := f.Write(data); err != nil {
		return err
	} else if n < len(data) {
		return io.ErrShortWrite
	}
	return nil
}

// 清空文件
func Truncate(path string, size int) error {
	return os.Truncate(path, int64(size))
}

// 将字符串写入文件
func PutContents(path string, content string) error {
	return putContents(path, []byte(content), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mDEFAULT_PERM)
}

// 追加字符串到文件，文件不存在时创建文件
func PutContentsAppend(path string, content string) error {
	return putContents(path, []byte(content), os.O_WRONLY|os.O_CREATE|os.O_APPEND, mDEFAULT_PERM)
}

// 写入二进制数据到文件，文件不存在时创建
func PutBinContents(path string, content []byte) error {
	return putContents(path, content, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mDEFAULT_PERM)
}

// 追回二进制数据到文件，文件不存在时创建
func PutBinContentsAppend(path string, content []byte) error {
	return putContents(path, content, os.O_WRONLY|os.O_CREATE|os.O_APPEND, mDEFAULT_PERM)
}

// 从start之后查找char字符，如果找到返回找到位置，否则返回-1
func GetNextCharOffset(reader io.ReaderAt, char byte, start int64) int64 {
	buffer := make([]byte, mREAD_BUFFER)
	offset := start
	for {
		if n, err := reader.ReadAt(buffer, offset); n > 0 {
			for i := 0; i < n; i++ {
				if buffer[i] == char {
					return int64(i) + offset
				}
			}
			offset += int64(n)
		} else if err != nil {
			break
		}
	}
	return -1
}

// 在文件path中从start之后查找字符char,找到则返回所在位置，否则返回-1
func GetNextCharOffsetByPath(path string, char byte, start int64) int64 {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, mDEFAULT_PERM); err == nil {
		defer f.Close()
		return GetNextCharOffset(f, char, start)
	}
	return -1
}

// 在reader中，获取从start到字符char之间二进制数据，包含char在内
func GetBinContentsTilChar(reader io.ReaderAt, char byte, start int64) ([]byte, int64) {
	if offset := GetNextCharOffset(reader, char, start); offset != -1 {
		return GetBinContentsByTwoOffsets(reader, start, offset+1), offset
	}
	return nil, -1
}

// 在文件path中，获取从start位置到字符char之间的二进制数据，包含char在内
func GetBinContentsTilCharByPath(path string, char byte, start int64) ([]byte, int64) {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, mDEFAULT_PERM); err == nil {
		defer f.Close()
		return GetBinContentsTilChar(f, char, start)
	}
	return nil, -1
}

// 从reader中获取从位置start到end之间的二进制数据
func GetBinContentsByTwoOffsets(reader io.ReaderAt, start int64, end int64) []byte {
	buffer := make([]byte, end-start)
	if _, err := reader.ReadAt(buffer, start); err != nil {
		return nil
	}
	return buffer
}

// 从文件path中，获取从位置start到end之间的二进制数据
func GetBinContentsByTwoOffsetsByPath(path string, start int64, end int64) []byte {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, mDEFAULT_PERM); err == nil {
		defer f.Close()
		return GetBinContentsByTwoOffsets(f, start, end)
	}
	return nil
}
