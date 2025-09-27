package method

import "fmt"

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MethodsPointers() {
	v := Vertex{3, 4}
	// 这里为什么不是&v.Scale(10)？
	// 因为 go 会自动帮我们取地址
	// 也就是说 v.Scale(10) 实际上是 (&v).Scale(10)
	v.Scale(10)
	fmt.Println(v.Abs())
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
