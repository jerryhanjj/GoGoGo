package method

import (
	"fmt"
	"math"
)

/*
你只能为在同一个包中定义的接收者类型声明方法，而不能为其它别的包中定义的类型 （包括 int 之类的内置类型）声明方法。

（译注：就是接收者的类型定义和方法声明必须在同一包内。）
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func MethodContinue() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) // 调用MyFloat类型的方法Abs
}
