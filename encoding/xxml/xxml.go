package xxml

import (
	"fmt"
	"github.com/go-xe2/third/github.com/clbanning/mxj"
	"github.com/go-xe2/x/encoding/xcharset"
	"github.com/go-xe2/x/utils/xregex"
	"strings"
)

// 将XML内容解析为map变量
func Decode(content []byte) (map[string]interface{}, error) {
	res, err := convert(content)
	if err != nil {
		return nil, err
	}
	return mxj.NewMapXml(res)
}

// 将map变量解析为XML格式内容
func Encode(v map[string]interface{}, rootTag ...string) ([]byte, error) {
	return mxj.Map(v).Xml(rootTag...)
}

func EncodeWithIndent(v map[string]interface{}, rootTag ...string) ([]byte, error) {
	return mxj.Map(v).XmlIndent("", "\t", rootTag...)
}

// XML格式内容直接转换为JSON格式内容
func ToJson(content []byte) ([]byte, error) {
	res, err := convert(content)
	if err != nil {
		fmt.Println("convert error. ", err)
		return nil, err
	}

	mv, err := mxj.NewMapXml(res)
	if err == nil {
		return mv.Json()
	} else {
		return nil, err
	}
}

// XML字符集预处理
func convert(xml []byte) (res []byte, err error) {
	patten := `<\?xml.*encoding\s*=\s*['|"](.*?)['|"].*\?>`
	matchStr, err := xregex.MatchString(patten, string(xml))
	if err != nil {
		return nil, err
	}
	xmlEncode := "UTF-8"
	if len(matchStr) == 2 {
		xmlEncode = matchStr[1]
	}
	xmlEncode = strings.ToUpper(xmlEncode)
	res, err = xregex.Replace(patten, []byte(""), xml)
	if err != nil {
		return nil, err
	}
	if xmlEncode != "UTF-8" && xmlEncode != "UTF8" {
		dst, err := xcharset.Convert("UTF-8", xmlEncode, string(res))
		if err != nil {
			return nil, err
		}
		res = []byte(dst)
	}
	return res, nil
}
