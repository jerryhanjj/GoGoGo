package interfacego

import "math"

type Abser interface {
	Abs() float64
}

func InterfaceGo() {
	var a Abser
	f := MyFloat(-math.Sqrt2)

	a = f // a MyFloat 实现了 Abser
	println(a.Abs())
	v := Vertex{3, 4}
	a = &v // a *Vertex 实现了 Abser
	println(a.Abs())
	// a = v
	// cannot use v (type Vertex) as type Abser in assignment:
	// Vertex does not implement Abser (Abs method has pointer receiver)
	// println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
