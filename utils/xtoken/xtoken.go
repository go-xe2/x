/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-11 11:22
* Description:
*****************************************************************/

package xtoken

import (
	"encoding/json"
	"fmt"
	"github.com/go-xe2/x/crypto/xhmac"
	"github.com/go-xe2/x/encoding/xbase64"
	"github.com/go-xe2/x/type/t"
	"sort"
	"strings"
)

func mapToString(mp map[string]interface{}) string {
	keys := make([]string, len(mp))
	i := 0
	for k := range mp {
		keys[i] = k
		i++
	}
	// 对字段进行排序
	sort.Slice(keys, func(i, j int) bool {
		return strings.Compare(keys[i], keys[j]) < 0
	})
	str := ""
	for _, k := range keys {
		val := t.String(mp[k])
		str += "&" + fmt.Sprintf("%s=%s", k, val)
	}
	if len(str) > 0 {
		str = str[1:]
	}
	return str
}

// token编码
func EncodeToken(data map[string]interface{}, expire int64, secret string) (string, error) {
	keys := make([]string, len(data))
	i := 0
	for k := range data {
		keys[i] = k
		i++
	}
	// 对字段进行排序
	sort.Slice(keys, func(i, j int) bool {
		return strings.Compare(keys[i], keys[j]) < 0
	})
	// 去除传入的t参数
	delete(data, "t")
	str := mapToString(data)
	str = str + "&t=" + t.String(expire)
	signKey := xhmac.HMAC_SHA1(str, secret)
	paramData := map[string]interface{}{
		"data":   data,
		"expire": expire,
		"sign":   signKey,
	}
	bts, err := json.Marshal(paramData)
	if err != nil {
		return "", err
	}
	res := xbase64.Encode(bts)
	return string(res), nil
}

// token解码
func DecodeToken(token string, secret string) (data map[string]interface{}, expire int64, err error) {
	bts, err := xbase64.DecodeString(token)
	if err != nil {
		return nil, 0, fmt.Errorf("token无效")
	}
	src := make(map[string]interface{})
	err = json.Unmarshal(bts, &src)
	if err != nil {
		return nil, 0, fmt.Errorf("token无效")
	}
	sign, ok := src["sign"]
	if !ok {
		return nil, 0, fmt.Errorf("token无效")
	}
	if m, ok := src["data"].(map[string]interface{}); !ok {
		return nil, 0, fmt.Errorf("token无效")
	} else {
		data = m
	}
	if v, ok := src["expire"]; ok {
		expire = t.Int64(v)
	} else {
		expire = 0
	}
	// 验证token是否被篡改
	str := mapToString(data)
	str = str + "&t=" + t.String(expire)
	verifySign := xhmac.HMAC_SHA1(str, secret)
	if verifySign != sign {
		return nil, 0, fmt.Errorf("token无效")
	}
	return data, expire, nil
}
