package method

import (
	"fmt"
	"math"
)

/*
Go 没有类。不过你可以为类型定义方法。

方法就是一类带特殊的 接收者 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
*/

type Vertex struct {
	X, Y float64
}

// 方法：与特定类型关联的函数
// 方法与函数的区别在于，方法作用于特定的类型（这里是 Vertex），而函数没有这个限制。
func (v Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

// 普通函数
func Abs(v Vertex) float64 {
	return math.Abs(v.X*v.X + v.Y*v.Y)
}

func MethodExample() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // 调用方法Abs
	fmt.Println(Abs(v))  // 调用函数Abs
}
