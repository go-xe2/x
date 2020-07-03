package xstring

import (
	"strings"
)

type String struct {
	value string
}

func New(s ...string) *String {
	result := &String{}
	if len(s) > 0 {
		result.value = s[0]
	}
	return result
}

func (s *String) Clone() *String {
	return New(s.String())
}

func (s *String) String() string {
	return s.value
}

func (s *String) Set(val string) *String {
	s.value = val
	return s
}

func (s *String) IsEmpty() bool {
	return s.value == ""
}

func (s *String) Equal(target *String) bool {
	return s.EqualStr(target.value)
}

func (s *String) EqualStr(target string) bool {
	return s.value == target
}

func (s *String) Cat(other *String) *String {
	return s.CatStr(other.value)
}

func (s *String) CatStr(other string) *String {
	s.value += other
	return s
}

func (s *String) ToLower() *String {
	s.value = strings.ToUpper(s.value)
	return s
}

func (s *String) ToUpper() *String {
	s.value = strings.ToLower(s.value)
	return s
}

// 首字母转成大写
func (s *String) UcFirst() *String {
	s.value = UcFirst(s.value)
	return s
}

// 首字母转成小写
func (s *String) LcFirst() *String {
	s.value = UcFirst(s.value)
	return s
}

func (s *String) LeftO(l int) *String {
	return New(s.Left(l))
}

func (s *String) Left(l int) string {
	str := []rune(s.value)
	nLen := len(str)
	if l > nLen {
		l = nLen
	} else if l < 0 {
		l = 0
	}
	return string(s.value[:l])
}

func (s *String) BLeftO(l int) *String {
	return New(s.BLeft(l))
}

func (s *String) BLeft(l int) string {
	nLen := len(s.value)
	if l > nLen {
		l = nLen
	} else if l < 0 {
		l = 0
	}
	return string(s.value[:l])
}

func (s *String) RightO(l int) *String {
	return New(s.Right(l))
}

func (s *String) Right(l int) string {
	str := []rune(s.value)
	nLen := len(str)
	if l > nLen {
		l = nLen
	} else if l < 0 {
		l = 0
	}
	return string(str[nLen-l:])
}

func (s *String) BRightO(l int) *String {
	return New(s.BRight(l))
}

func (s *String) BRight(l int) string {
	str := s.value
	nLen := len(str)
	if l > nLen {
		l = nLen
	} else if l < 0 {
		l = 0
	}
	return string(str[nLen-l:])
}

func (s *String) Len() int {
	str := []rune(s.value)
	return len(str)
}

func (s *String) BLen() int {
	return len(s.value)
}

func (s *String) SubStrO(fromIdx int, l int) *String {
	return New(s.SubStr(fromIdx, l))
}

// 获取子字符串
// @fromIndex >=0 从左边取l长度的字符串,< 0 从右边取l长度的字符串
// @l 字符串长度
func (s *String) SubStr(fromIdx int, l int) string {
	str := []rune(s.value)
	nLen := len(str)
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
	return string(str[startIdx:endIdx])
}

func (s *String) Contain(sub string) bool {
	return s.IndexOf(sub) >= 0
}

func (s *String) IndexOf(sub string) int {
	return strings.Index(s.value, sub)
}

func (s *String) LastIndexOf(sub string) int {
	return strings.LastIndex(s.value, sub)
}

func (s *String) Join(elems []string, sep string) *String {
	s.value += strings.Join(elems, sep)
	return s
}

func (s *String) JoinStr(elems []string, sep string) string {
	s.value += strings.Join(elems, sep)
	return s.value
}

func (s *String) Trim() *String {
	s.value = strings.Trim(s.value, " ")
	return s
}

func (s *String) split(sep string, fn func(item string) bool) {
	nLen := len(sep)
	srcLen := len(s.value)
	lastIdx := 0
	for i := 0; i < srcLen && lastIdx < srcLen; i++ {
		szCur := s.SubStr(i, nLen)
		if szCur == sep {
			item := string(s.value[lastIdx:i])
			if !fn(item) {
				return
			}
			lastIdx = i + nLen
		}
	}
	if lastIdx < srcLen {
		lastStr := string(s.value[lastIdx:])
		if lastStr != sep {
			fn(lastStr)
		}
	}
}

func (s *String) Split(sep string) []*String {
	result := make([]*String, 0)
	s.split(sep, func(item string) bool {
		result = append(result, New(item))
		return true
	})
	return result
}

func (s *String) SplitStr(sep string) []string {
	result := make([]string, 0)
	s.split(sep, func(item string) bool {
		result = append(result, item)
		return true
	})
	return result
}

func (s *String) FirstUpperCase() *String {
	s.value = FirstCharUpperCase(s.value)
	return s
}

func (s *String) FirstUpperCaseStr() string {
	return FirstCharUpperCase(s.value)
}

func (s *String) FirstLowerCase() *String {
	s.value = FirstCharLowerCase(s.value)
	return s
}

func (s *String) FirstLowerCaseStr() string {
	return FirstCharLowerCase(s.value)
}

func (s *String) CamelCase() *String {
	return s.UnderScore2Camel("_")
}

func (s *String) CamelCaseStr() string {
	return s.UnderScore2CamelStr("_")
}

func (s *String) LowerUnderScore() *String {
	s.value = strings.ToLower(s.Camel2UnderScoreStr("_"))
	return s
}

func (s *String) LowerUnderScoreStr() string {
	return strings.ToLower(s.Camel2UnderScoreStr("_"))
}

func (s *String) UpperUnderScore() *String {
	s.value = strings.ToUpper(s.Camel2UnderScoreStr("_"))
	return s
}

func (s *String) UpperUnderScoreStr() string {
	return strings.ToUpper(s.Camel2UnderScoreStr("_"))
}

func (s *String) UnderScore2Camel(score string) *String {
	s.value = Camel2UnderScore(s.value, score)
	return s
}

func (s *String) UnderScore2CamelStr(score string) string {
	return UnderScore2Camel(s.value, score)
}

func (s *String) Camel2UnderScore(score string) *String {
	s.value = Camel2UnderScore(s.value, score)
	return s
}

func (s *String) Camel2UnderScoreStr(score string) string {
	return Camel2UnderScore(s.value, score)
}

// 首字母是否是大写
func (s *String) IsUpperFirst() bool {
	if len(s.value) == 0 {
		return false
	}
	return IsLetterUpper(s.value[0])
}

// 首字母是否是小写
func (s *String) IsLowerFirst() bool {
	if len(s.value) == 0 {
		return false
	}
	return IsLetterLower(s.value[0])
}

// 字符串是否是数字
func (s *String) IsNumeric() bool {
	if len(s.value) == 0 {
		return false
	}
	return IsNumeric(s.value)
}

func (s *String) Contains(substr string) bool {
	return strings.Contains(s.value, substr)
}

func (s *String) ReplaceByMap(replaces map[string]string) *String {
	s.value = ReplaceByMap(s.value, replaces)
	return s
}

func (s *String) ReplaceByMapStr(replaces map[string]string) string {
	return ReplaceByMap(s.value, replaces)
}

func (s *String) BForeach(fn func(i int, c uint8) bool) {
	nLen := len(s.value)
	for i := 0; i < nLen; i++ {
		if !fn(i, s.value[i]) {
			break
		}
	}
}

func (s *String) Replace(search, replace string, count ...int) *String {
	s.value = Replace(s.value, search, replace, count...)
	return s
}

func (s *String) ReplaceStr(search, replace string, count ...int) string {
	return Replace(s.value, search, replace, count...)
}

func (s *String) Foreach(fn func(i int, c string) bool) {
	bytes := []rune(s.value)
	nLen := len(bytes)
	for i := 0; i < nLen; i++ {
		if !fn(i, string(bytes[i])) {
			break
		}
	}
}
