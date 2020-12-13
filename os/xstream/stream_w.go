/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-08-20 09:59
* Description:
*****************************************************************/

package xstream

import (
	"github.com/go-xe2/x/type/xbinary"
	"io"
)

type tEncodeMode int8

const (
	leEncode tEncodeMode = iota
	beEncode
)

type tStreamWriter struct {
	mode tEncodeMode
	w    io.Writer
}

var _ StreamWriter = (*tStreamWriter)(nil)

func NewLeStreamWriter(writer io.Writer) StreamWriter {
	return &tStreamWriter{
		mode: leEncode,
		w:    writer,
	}
}

func NewBeStreamWriter(writer io.Writer) StreamWriter {
	return &tStreamWriter{
		mode: beEncode,
		w:    writer,
	}
}

func (p *tStreamWriter) writeNodeType(nt TNodeType) error {
	if _, err := p.w.Write(xbinary.EncodeInt8(int8(nt))); err != nil {
		return err
	}
	return nil
}

func (p *tStreamWriter) WriteStr(str string) error {
	size := len(str)
	if err := p.writeNodeType(STR_NODE); err != nil {
		return err
	}
	if p.mode == leEncode {
		if _, err := p.w.Write(xbinary.LeEncodeInt32(int32(size))); err != nil {
			return err
		}
	} else {
		if _, err := p.w.Write(xbinary.BeEncodeInt32(int32(size))); err != nil {
			return err
		}
	}
	if _, err := p.w.Write([]byte(str)); err != nil {
		return err
	}
	return nil
}

func (p *tStreamWriter) WriteInt8(v int8) error {
	var buf []byte
	if p.mode == leEncode {
		buf = xbinary.LeEncodeInt8(v)
	} else {
		buf = xbinary.BeEncodeInt8(v)
	}
	if err := p.writeNodeType(INT8_NODE); err != nil {
		return err
	}
	if _, err := p.w.Write(buf); err != nil {
		return err
	}
	return nil
}

func (p *tStreamWriter) WriteInt16(v int16) error {
	var buf []byte
	if p.mode == leEncode {
		buf = xbinary.LeEncodeInt16(v)
	} else {
		buf = xbinary.BeEncodeInt16(v)
	}
	if err := p.writeNodeType(INT16_NODE); err != nil {
		return err
	}
	if _, err := p.w.Write(buf); err != nil {
		return err
	}
	return nil
}

func (p *tStreamWriter) WriteInt32(v int32) error {
	var buf []byte
	if p.mode == leEncode {
		buf = xbinary.LeEncodeInt32(v)
	} else {
		buf = xbinary.BeEncodeInt32(v)
	}
	if err := p.writeNodeType(INT32_NODE); err != nil {
		return err
	}
	if _, err := p.w.Write(buf); err != nil {
		return err
	}
	return nil
}

func (p *tStreamWriter) WriteInt64(v int64) error {
	var buf []byte
	if p.mode == leEncode {
		buf = xbinary.LeEncodeInt64(v)
	} else {
		buf = xbinary.BeEncodeInt64(v)
	}
	if err := p.writeNodeType(INT64_NODE); err != nil {
		return err
	}
	if _, err := p.w.Write(buf); err != nil {
		return err
	}
	return nil
}

func (p *tStreamWriter) WriteFloat32(v float32) error {
	var buf []byte
	if p.mode == leEncode {
		buf = xbinary.LeEncodeFloat32(v)
	} else {
		buf = xbinary.BeEncodeFloat32(v)
	}
	if err := p.writeNodeType(FLOAT32_NODE); err != nil {
		return err
	}
	if _, err := p.w.Write(buf); err != nil {
		return err
	}
	return nil
}

func (p *tStreamWriter) WriteFloat64(v float64) error {
	var buf []byte
	if p.mode == leEncode {
		buf = xbinary.LeEncodeFloat64(v)
	} else {
		buf = xbinary.BeEncodeFloat64(v)
	}
	if err := p.writeNodeType(FLOAT64_NODE); err != nil {
		return err
	}
	if _, err := p.w.Write(buf); err != nil {
		return err
	}
	return nil
}

func (p *tStreamWriter) WriteBool(v bool) error {
	if err := p.writeNodeType(BOOL_NODE); err != nil {
		return err
	}
	if v {
		if _, err := p.w.Write([]byte{1}); err != nil {
			return err
		}
	} else {
		if _, err := p.w.Write([]byte{0}); err != nil {
			return err
		}
	}
	return nil
}

func (p *tStreamWriter) WriteBytes(buf []byte) error {
	if err := p.writeNodeType(BYTES_NODE); err != nil {
		return err
	}
	size := len(buf)
	if p.mode == leEncode {
		if _, err := p.w.Write(xbinary.LeEncodeInt64(int64(size))); err != nil {
			return err
		}
	} else {
		if _, err := p.w.Write(xbinary.BeEncodeInt64(int64(size))); err != nil {
			return err
		}
	}
	if _, err := p.w.Write(buf); err != nil {
		return err
	}
	return nil
}
