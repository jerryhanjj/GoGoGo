package method

import "fmt"

func FuncMethodIndirection() {
	v := Vertex{3, 4}
	v.Scale(2)
	Scale(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	Scale(p, 8)

	fmt.Println(v, p)

	v1 := Vertex{3, 4}
	fmt.Println(v1.Abs())
	fmt.Println(Abs(v1))

	p1 := &Vertex{4, 3}
	fmt.Println(p1.Abs())
	fmt.Println(Abs(*p1))
}
