package gostruct

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	pV = &Vertex{1, 2} // 类型为 *Vertex
)

func RunStructs() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.Y = 1e9
	fmt.Println(v)
	fmt.Println(v1, pV, v2, v3)
}
