package arrayslice

import "fmt"

func RunMakeSlice() {
	a := make([]int, 5)
	printSliceS("a", a)

	b := make([]int, 0, 5)
	printSliceS("b", b)

	c := b[:2]
	printSliceS("c", c)

	d := c[2:5]
	printSliceS("d", d)
}

func printSliceS(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
