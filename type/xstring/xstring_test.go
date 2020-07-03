package xstring

import (
	"fmt"
	"testing"
)

func TestString_Split(t *testing.T) {
	var s = "123:456:794"
	items := Split(s, ":", 1)
	fmt.Println("split items:", Join(items, "\n"))
}
