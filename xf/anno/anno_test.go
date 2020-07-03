package anno

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	szAnno1 := `@field (name = name, sex = true, age = 33)`
	szAnno2 := `@ExampleAnno(name = testAnno, sex = true)`

	name, anno2 := Create(szAnno2)
	fmt.Println("====> anno2 name:", name, ",ann:", anno2)

	name, anno1 := Create(szAnno1)
	fmt.Println("===> anno1 name:", name, ",ann:", anno1)

}
