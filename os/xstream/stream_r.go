/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-08-20 09:58
* Description:
*****************************************************************/

package xstream

import (
	"errors"
	"github.com/go-xe2/x/type/xbinary"
	"io"
)

type tStreamReader struct {
	m tEncodeMode
	r io.Reader
}

var _ StreamReader = (*tStreamReader)(nil)

func NewLeStreamReader(r io.Reader) StreamReader {
	return &tStreamReader{
		m: leEncode,
		r: r,
	}
}

func NewBeStreamReader(r io.Reader) StreamReader {
	return &tStreamReader{
		m: beEncode,
		r: r,
	}
}

func (p *tStreamReader) readNodeType() (TNodeType, error) {
	buf := make([]byte, 1)
	_, err := p.r.Read(buf)
	if err != nil {
		return UNKNOWN_NODE, err
	}
	n := xbinary.DecodeToInt8(buf)
	return TNodeType(n), nil
}

func (p *tStreamReader) ReadNode() (val interface{}, t TNodeType, err error) {
	t, err = p.readNodeType()
	if err != nil {
		return nil, t, err
	}
	switch t {
	case STR_NODE:
		buf := make([]byte, 4)
		if _, err := p.r.Read(buf); err != nil {
			return "", t, err
		}
		var size int32 = 0
		if p.m == leEncode {
			size = xbinary.LeDecodeToInt32(buf)
		} else {
			size = xbinary.BeDecodeToInt32(buf)
		}
		strBuf := make([]byte, size)
		if _, err := p.r.Read(strBuf); err != nil {
			return "", t, err
		}
		return string(strBuf), t, nil
	case BOOL_NODE:
		buf := make([]byte, 1)
		if _, err := p.r.Read(buf); err != nil {
			return false, t, err
		}
		return buf[0] == 1, t, nil
	case INT8_NODE:
		buf := make([]byte, 1)
		if _, err := p.r.Read(buf); err != nil {
			return 0, t, err
		}
		return int8(buf[0]), t, nil
	case INT16_NODE:
		buf := make([]byte, 2)
		if _, err := p.r.Read(buf); err != nil {
			return 0, t, err
		}
		if p.m == leEncode {
			return xbinary.LeDecodeToInt16(buf), t, nil
		} else {
			return xbinary.BeDecodeToInt16(buf), t, nil
		}
	case INT32_NODE:
		buf := make([]byte, 4)
		if _, err := p.r.Read(buf); err != nil {
			return 0, t, err
		}
		if p.m == leEncode {
			return xbinary.LeDecodeToInt32(buf), t, nil
		} else {
			return xbinary.BeDecodeToInt32(buf), t, nil
		}
		break
	case INT64_NODE:
		buf := make([]byte, 8)
		if _, err := p.r.Read(buf); err != nil {
			return 0, t, err
		}
		if p.m == leEncode {
			return xbinary.LeDecodeToInt64(buf), t, nil
		} else {
			return xbinary.BeDecodeToInt64(buf), t, nil
		}
	case FLOAT32_NODE:
		buf := make([]byte, 4)
		if _, err := p.r.Read(buf); err != nil {
			return 0, t, err
		}
		if p.m == leEncode {
			return xbinary.LeDecodeToFloat32(buf), t, nil
		} else {
			return xbinary.BeDecodeToFloat32(buf), t, nil
		}
	case FLOAT64_NODE:
		buf := make([]byte, 8)
		if _, err := p.r.Read(buf); err != nil {
			return 0, t, err
		}
		if p.m == leEncode {
			return xbinary.LeDecodeToFloat64(buf), t, nil
		} else {
			return xbinary.BeDecodeToFloat64(buf), t, nil
		}
	case BYTES_NODE:
		buf := make([]byte, 8)
		if _, err := p.r.Read(buf); err != nil {
			return 0, t, err
		}
		var size int64 = 0
		if p.m == leEncode {
			size = xbinary.LeDecodeToInt64(buf)
		} else {
			size = xbinary.BeDecodeToInt64(buf)
		}
		data := make([]byte, size)
		if n, err := p.r.Read(data); err != nil {
			return []byte{}, t, err
		} else {
			return data[:n], t, nil
		}
	}
	return nil, t, nil
}

func (p *tStreamReader) ReadBytes() ([]byte, error) {
	b, t, err := p.ReadNode()
	if err != nil {
		return nil, err
	}
	if t != BYTES_NODE {
		return []byte{}, errors.New("不是[]byte类型")
	}
	return b.([]byte), nil
}

func (p *tStreamReader) ReadStr() (string, error) {
	s, t, err := p.ReadNode()
	if err != nil {
		return "", err
	}
	if t == STR_NODE {
		return s.(string), nil
	}
	return "", errors.New("不是string类型")
}

func (p *tStreamReader) ReadInt8() (int8, error) {
	n, t, err := p.ReadNode()
	if err != nil {
		return 0, err
	}
	if t == INT8_NODE {
		return n.(int8), nil
	}
	return 0, errors.New("不是int8类型")
}

func (p *tStreamReader) ReadInt16() (int16, error) {
	n, t, err := p.ReadNode()
	if err != nil {
		return 0, err
	}
	if t == INT16_NODE {
		return n.(int16), nil
	}
	return 0, errors.New("不是int16类型")
}

func (p *tStreamReader) ReadInt32() (int32, error) {
	n, t, err := p.ReadNode()
	if err != nil {
		return 0, err
	}
	if t == INT32_NODE {
		return n.(int32), nil
	}
	return 0, errors.New("不是int32类型")
}

func (p *tStreamReader) ReadInt64() (int64, error) {
	n, t, err := p.ReadNode()
	if err != nil {
		return 0, err
	}
	if t == INT64_NODE {
		return n.(int64), nil
	}
	return 0, errors.New("不是int64类型")
}

func (p *tStreamReader) ReadFloat32() (float32, error) {
	n, t, err := p.ReadNode()
	if err != nil {
		return 0, err
	}
	if t == FLOAT32_NODE {
		return n.(float32), nil
	}
	return 0, errors.New("不是float32类型")
}

func (p *tStreamReader) ReadFloat64() (float64, error) {
	n, t, err := p.ReadNode()
	if err != nil {
		return 0, err
	}
	if t == FLOAT64_NODE {
		return n.(float64), nil
	}
	return 0, errors.New("不是float64类型")
}

func (p *tStreamReader) ReadBool() (bool, error) {
	n, t, err := p.ReadNode()
	if err != nil {
		return false, err
	}
	if t == BOOL_NODE {
		return n.(bool), nil
	}
	return false, errors.New("不是bool类型")
}
