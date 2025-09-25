package multiresult

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func Fn_multiresult() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
