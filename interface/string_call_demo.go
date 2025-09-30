package interfacego

import (
	"fmt"
)

// 追踪String()方法的调用
type TrackedPerson struct {
	Name string
	Age  int
}

func (p TrackedPerson) String() string {
	fmt.Printf("🔍 String()方法被调用了！参数: Name=%s, Age=%d\n", p.Name, p.Age)
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

// 没有String()方法的类型
type SimplePerson struct {
	Name string
	Age  int
}

func DemonstrateStringCall() {
	fmt.Println("=== 演示 fmt.Println 如何调用 String() ===")

	// 1. 有String()方法的类型
	fmt.Println("\n1. 使用实现了String()方法的类型:")
	tracked := TrackedPerson{Name: "Alice", Age: 30}
	fmt.Println("调用 fmt.Println(tracked):")
	fmt.Println(tracked) // 这里会打印调用追踪信息

	// 2. 没有String()方法的类型
	fmt.Println("\n2. 使用没有String()方法的类型:")
	simple := SimplePerson{Name: "Bob", Age: 25}
	fmt.Println("调用 fmt.Println(simple):")
	fmt.Println(simple) // 使用默认格式

	// 3. 不同的fmt函数都会检查String()方法
	fmt.Println("\n3. 其他fmt函数也会调用String():")
	fmt.Printf("fmt.Printf(\"%%v\", tracked): ")
	fmt.Printf("%v\n", tracked)

	fmt.Printf("fmt.Printf(\"%%s\", tracked): ")
	fmt.Printf("%s\n", tracked)

	fmt.Printf("fmt.Sprintf(\"%%v\", tracked): ")
	result := fmt.Sprintf("%v", tracked)
	fmt.Printf("%s\n", result)
}

// 演示类型断言的工作原理
func DemonstrateStringTypeAssertion() {
	fmt.Println("\n=== 演示类型断言检查过程 ===")

	tracked := TrackedPerson{Name: "Charlie", Age: 35}
	simple := SimplePerson{Name: "David", Age: 40}

	// 模拟fmt包的类型检查过程
	fmt.Println("\n检查 TrackedPerson:")
	checkStringer(tracked)

	fmt.Println("\n检查 SimplePerson:")
	checkStringer(simple)
}

// 模拟fmt包内部的类型检查
func checkStringer(arg interface{}) {
	fmt.Printf("正在检查类型: %T\n", arg)

	// 这就是fmt包内部做的事情
	if stringer, ok := arg.(fmt.Stringer); ok {
		fmt.Println("✅ 实现了fmt.Stringer接口")
		fmt.Printf("调用String()方法: %s\n", stringer.String())
	} else {
		fmt.Println("❌ 没有实现fmt.Stringer接口")
		fmt.Printf("使用默认格式: %v\n", arg)
	}
}
