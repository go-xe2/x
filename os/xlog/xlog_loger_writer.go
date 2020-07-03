package xlog

import "bytes"

func (l *TLogger) Write(p []byte) (n int, err error) {
	l.Header(false).Print(string(bytes.TrimRight(p, "\r\n")))
	return len(p), nil
}
