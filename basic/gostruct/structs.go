package gostruct

import "fmt"

type Vertex struct {
	X, Y int
}

func RunStructs() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.Y = 1e9
	fmt.Println(v)
}
