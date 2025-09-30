package interfacego

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s\n", cases.Title(language.English).String(v))
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}

func InterfaceTypeSwitch() {
	do(21)
	do("hello")
	do(true)
}
