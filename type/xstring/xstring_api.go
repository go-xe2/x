package xstring

import (
	"strings"
)

func FirstCharUpperCase(s string) string {
	if s == "" {
		return s
	}
	str := []rune(s)
	if str[0] >= 'a' && str[0] <= 'z' {
		str[0] = str[0] + ('A' - 'a')
	}
	return string(str)
}

func FirstCharLowerCase(s string) string {
	if s == "" {
		return s
	}
	str := []rune(s)
	if str[0] >= 'A' && str[0] <= 'Z' {
		str[0] = str[0] - ('A' - 'a')
	}
	return string(str)
}

func UnderScore2Camel(s string, score string) string {
	str := []rune(s)
	nLen := len(str)
	result := ""
	var tmp = ""
	for i := 0; i < nLen; i++ {
		if string(str[i]) == score {
			if tmp != "" {
				result = result + FirstCharUpperCase(strings.ToLower(tmp))
				tmp = ""
			}
		} else {
			tmp += string(str[i])
		}
	}
	if tmp != "" {
		result = result + FirstCharUpperCase(tmp)
	}
	return result
}

func Camel2UnderScore(s string, score string) string {
	str := []rune(s)
	nLen := len(str)
	result := ""
	tmp := ""
	for i := 0; i < nLen; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			if tmp != "" {
				if result == "" {
					result += strings.ToLower(tmp)
				} else {
					result += score + strings.ToLower(tmp)
				}
			}
			tmp = string(str[i])
		} else {
			tmp += string(str[i])
		}
	}
	if tmp != "" {
		if result == "" {
			result += strings.ToLower(tmp)
		} else {
			result += score + strings.ToLower(tmp)
		}
	}
	return result
}

// 检查字节是否是大写字母
func IsLetterUpper(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}

// 检查节字是否是小写字线
func IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

// 检查字符串是否是纯数字
func IsNumeric(s string) bool {
	length := len(s)
	if length == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if (s[i] < byte('0') || s[i] > byte('9')) && s[i] != '.' {
			return false
		}
	}
	return true
}

// 将字符串首字母转换成大写
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterLower(s[0]) {
		return string(s[0]-32) + s[1:]
	}
	return s
}

// 将字符串首字母转换成小写
func LcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterUpper(s[0]) {
		return string(s[0]+32) + s[1:]
	}
	return s
}

func Replace(origin, search, replace string, count ...int) string {
	n := -1
	if len(count) > 0 {
		n = count[0]
	}
	return strings.Replace(origin, search, replace, n)
}

// 替换字符串
func ReplaceByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = strings.Replace(origin, k, v, -1)
	}
	return origin
}

func Left(s string, l int) string {
	str := []rune(s)
	nLen := len(str)
	if l > nLen {
		l = nLen
	} else if l < 0 {
		l = 0
	}
	return string(s[:l])
}

func BLeft(s string, l int) string {
	nLen := len(s)
	if l > nLen {
		l = nLen
	} else if l < 0 {
		l = 0
	}
	return s[:l]
}

func Right(s string, l int) string {
	str := []rune(s)
	nLen := len(str)
	if l > nLen {
		l = nLen
	} else if l < 0 {
		l = 0
	}
	return string(str[nLen-l:])
}

func BRight(s string, l int) string {
	nLen := len(s)
	if l > nLen {
		l = nLen
	} else if l < 0 {
		l = 0
	}
	return s[nLen-l:]
}

func Len(s string) int {
	runes := []rune(s)
	return len(runes)
}

func SubStr(s string, fromIdx int, slen ...int) string {
	runes := []rune(s)
	nLen := len(runes)
	l := nLen - fromIdx
	if len(slen) > 0 {
		l = slen[0]
	}
	isFromLeft := fromIdx >= 0
	startIdx := 0
	endIdx := 0
	if isFromLeft {
		if fromIdx > nLen {
			startIdx = nLen
		} else {
			startIdx = fromIdx
		}
	} else {
		// from right
		if -fromIdx > nLen {
			startIdx = 0
		} else {
			startIdx = nLen + fromIdx
		}
	}
	endIdx = startIdx + l
	if endIdx > nLen {
		endIdx = nLen
	}
	return string(runes[startIdx:endIdx])
}

func Trim(s string) string {
	return strings.Trim(s, " ")
}

func Splits(s, sep string, more ...string) []string {
	result := Split(s, sep)
	for _, se := range more {
		result = append(result, Split(s, se)...)
	}
	return result
}

func Split(s, sep string, needCount ...int) []string {
	sepRunes := []rune(sep)

	nLen := len(sepRunes)

	srcRunes := []rune(s)
	srcLen := len(srcRunes)

	lastIdx := 0
	result := make([]string, 0)
	nCount := -1
	if len(needCount) > 0 {
		nCount = needCount[0]
	}
	curCount := 0
	for i := 0; i < srcLen && lastIdx < srcLen; i++ {
		szCur := srcRunes[i : i+nLen] //  SubStr(s, i, nLen)
		if string(szCur) == sep {
			curCount++
			item := string(srcRunes[lastIdx:i])
			result = append(result, item)
			lastIdx = i + nLen
			if nCount > 0 && nCount >= curCount {
				break
			}
		}
	}
	if lastIdx < srcLen {
		lastStr := string(srcRunes[lastIdx:])
		if lastStr != sep {
			result = append(result, lastStr)
		}
	}
	return result
}

func MuSplit(s string, sep uint8, more ...uint8) []string {
	srcRunes := []rune(s)
	nLen := len(srcRunes)
	nStart := 0
	seps := make(map[rune]bool)
	seps[rune(sep)] = true
	for i := 0; i < len(more); i++ {
		seps[rune(more[i])] = true
	}
	result := make([]string, 0)
	for i := 0; i < nLen; i++ {
		c := srcRunes[i]
		if _, ok := seps[c]; ok {
			result = append(result, string(srcRunes[nStart:i]))
			nStart = i + 1
		}
	}
	if nStart < nLen-1 {
		result = append(result, string(srcRunes[nStart:nLen-1]))
	}
	return result
}

func IndexOf(s, sub string) int {
	return strings.Index(s, sub)
}

func LastIndexOf(s, sub string) int {
	return strings.LastIndex(s, sub)
}

func Foreach(s string, fn func(i int, s string) bool) {
	bytes := []rune(s)
	nLen := len(bytes)
	for i := 0; i < nLen; i++ {
		if !fn(i, string(bytes[i])) {
			break
		}
	}
}

func Contain(s, sub string) bool {
	return IndexOf(s, sub) >= 0
}

func Join(items []string, sep string) string {
	return strings.Join(items, sep)
}

func IsFirstLetterUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	return IsLetterUpper(s[0])
}

func IsFirstLetterLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	return IsLetterLower(s[0])
}

// 转换在大写
func UpperCase(s string) string {
	return strings.ToUpper(s)
}

// 转换成小写
func LowerCase(s string) string {
	return strings.ToLower(s)
}

// 解析键值对字符串为hashMap
// 格式:
//		1、键1:值1;键2:值2
//		2、键1:值1,键2:值2
//		3、键1=值1&键2=值2
//		4、键1=值1;键2=值2
//		4、键1=值1,键2=值2
func ParseKeyValue(s string) map[string]string {
	result := make(map[string]string)
	if s == "" {
		return result
	}
	cols := MuSplit(s, ',', ';', '&')
	for _, col := range cols {
		kv := MuSplit(col, ':', '=')
		k := kv[0]
		v := ""
		if len(kv) > 1 {
			v = kv[1]
		}
		result[k] = v
	}
	return result
}

// 是否包含前缀
func StartWith(s string, search string) bool {
	return strings.HasPrefix(s, search)
}

// 是否包含后缀
func EndWith(s string, search string) bool {
	return strings.HasSuffix(s, search)
}

func MakeAndFillStrSlice(size int, fillChar string) []string {
	result := make([]string, size)
	l := len(result)
	for i := 0; i < l; i++ {
		result[i] = fillChar
	}
	return result
}

func MakeStrAndFill(count int, repeatChar string, sep string) string {
	if count == 0 {
		return ""
	}
	items := MakeAndFillStrSlice(count, repeatChar)
	return strings.Join(items, sep)
}
