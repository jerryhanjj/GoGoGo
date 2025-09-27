package closure

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	first := true
	return func() int {
		if first {
			first = false
			return a
		}
		a, b = b, a+b
		return a
	}
}

func RunFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
