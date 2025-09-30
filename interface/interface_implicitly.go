package interfacego

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，不过我们并不需要显式声明这一点。
func (t T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func InterfaceImplicitly() {
	var i I = T{"hello"}
	i.M()
}

func InterfaceValue() {
	var i I
	i = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func EmptyInterface() {
	var i interface{}
	describeEmpty(i)
	i = 42
	describeEmpty(i)
	i = "hello"
	describeEmpty(i)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
