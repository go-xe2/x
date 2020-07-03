package xclass

import "strings"

func repeatString(s string, l int) string {
	items := make([]string, l)
	for i := 0; i < l; i++ {
		items = append(items, s)
	}
	return strings.Join(items, "")
}
