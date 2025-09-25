package types

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func FnTypes() {
	fmt.Printf("类型: %T 值: %v\n", ToBe, ToBe)
	fmt.Printf("类型: %T 值: %v\n", MaxInt, MaxInt)
	fmt.Printf("类型: %T 值: %v\n", z, z)
}

func FnZeroValue() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
