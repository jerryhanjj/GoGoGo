package variables

import "fmt"

var c, python, java bool

func Fn_variables() {
	var i int // 默认值为0
	fmt.Println(i, c, python, java)
}

var i, j int = 1, 2

func FnVariablesWithInit() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func FnShortVar() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
