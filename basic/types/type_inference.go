package types

import "fmt"

func TypeInference() {
	v := 42                                 // v is inferred as int
	f := 3.14                               // f is inferred as float64
	s := "hello"                            // s is inferred as string
	b := true                               // b is inferred as bool
	g := 0.867 + 0.5i                       // g is inferred as complex128
	arr := []int{1, 2, 3}                   // arr is inferred as []int
	m := map[string]int{"one": 1, "two": 2} // m is inferred as map[string]int

	fmt.Printf("v: %T, f: %T, s: %T, b: %T, arr: %T, m: %T, g: %T\n",
		v, f, s, b, arr, m, g)
}
