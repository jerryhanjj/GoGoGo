package interfacego

import "fmt"

func TypeAssertion() {
	var i interface{} = "hello"

	s, ok := i.(string)
	if ok {
		println("字符串值:", s)
	} else {
		println("类型断言失败")
	}

	n, ok := i.(int)
	fmt.Println(n, ok) // 0 false

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false
}
