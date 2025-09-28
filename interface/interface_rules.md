```go
package interfacego

import (
	"fmt"
	"math"
)

// 接口定义
type AbsInterface interface {
	Abs() float64
}

// 两个不同的类型来演示
type VertexValue struct {
	X, Y float64
}

type VertexPointer struct {
	X, Y float64
}

// 值接收者方法
func (v VertexValue) Abs() float64 {
	fmt.Println("  → 调用值接收者方法")
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 指针接收者方法
func (v *VertexPointer) Abs() float64 {
	fmt.Println("  → 调用指针接收者方法")
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func InterfaceImplementationRules() {
	fmt.Println("=== 接口实现规则详解 ===")

	// 1. 值接收者的情况
	fmt.Println("\n1. 值接收者方法 func (v VertexValue) Abs()")

	v1 := VertexValue{3, 4}
	p1 := &VertexValue{3, 4}

	fmt.Println("   方法调用（都可以）:")
	fmt.Print("   v1.Abs(): ")
	v1.Abs() // ✅ 直接调用
	fmt.Print("   p1.Abs(): ")
	p1.Abs() // ✅ 自动解引用

	fmt.Println("   接口赋值:")
	var a1 AbsInterface
	a1 = v1 // ✅ VertexValue 实现了 AbsInterface
	fmt.Printf("   a1 = v1: 成功，类型是 %T\n", a1)

	a1 = p1 // ✅ *VertexValue 也实现了 AbsInterface（因为值接收者）
	fmt.Printf("   a1 = p1: 成功，类型是 %T\n", a1)

	// 2. 指针接收者的情况
	fmt.Println("\n2. 指针接收者方法 func (v *VertexPointer) Abs()")

	v2 := VertexPointer{3, 4}
	p2 := &VertexPointer{3, 4}

	fmt.Println("   方法调用（都可以）:")
	fmt.Print("   v2.Abs(): ")
	v2.Abs() // ✅ 自动取地址
	fmt.Print("   p2.Abs(): ")
	p2.Abs() // ✅ 直接调用

	fmt.Println("   接口赋值:")
	var a2 AbsInterface

	// a2 = v2  // ❌ 这会编译错误！
	fmt.Println("   a2 = v2: ❌ 编译错误 - VertexPointer 没有实现 AbsInterface")

	a2 = p2 // ✅ *VertexPointer 实现了 AbsInterface
	fmt.Printf("   a2 = p2: 成功，类型是 %T\n", a2)
}

func WhyThisRestriction() {
	fmt.Println("\n=== 为什么有这个限制？ ===")

	fmt.Println("1. 接口值的存储机制:")
	fmt.Println("   - 接口值包含 (值, 类型) 对")
	fmt.Println("   - 如果方法需要修改接收者，必须存储指针")
	fmt.Println("   - 如果存储值，方法无法修改原始数据")

	fmt.Println("\n2. 一致性保证:")
	fmt.Println("   - 值接收者方法：对值和指针都适用")
	fmt.Println("   - 指针接收者方法：只对指针适用")
	fmt.Println("   - 这保证了方法语义的一致性")

	fmt.Println("\n3. 防止意外行为:")
	fmt.Println("   - 避免值拷贝导致的修改丢失")
	fmt.Println("   - 明确区分值语义和引用语义")
}

func MethodVsInterface() {
	fmt.Println("\n=== 方法调用 vs 接口实现 ===")

	type Counter struct {
		count int
	}

	// 指针接收者方法（会修改状态）
	increment := func(c *Counter) {
		c.count++
		fmt.Printf("   计数器值: %d\n", c.count)
	}

	c := Counter{0}

	fmt.Println("方法调用:")
	fmt.Print("  increment(&c): ")
	increment(&c) // 需要显式传指针
	fmt.Printf("  c 的值: %+v\n", c)

	// 如果允许接口存储值，会发生什么？
	fmt.Println("\n如果接口可以存储值会怎样？")
	fmt.Println("  - 接口存储的是 c 的拷贝")
	fmt.Println("  - 调用方法修改的是拷贝，不是原始值")
	fmt.Println("  - 这会导致非常困惑的行为")
	fmt.Println("  - 所以 Go 禁止这种情况")
}

func Solutions() {
	fmt.Println("\n=== 解决方案 ===")

	fmt.Println("1. 改为值接收者（如果不需要修改）:")
	fmt.Println("   func (v Vertex) Abs() float64 { ... }")

	fmt.Println("\n2. 使用指针（如果需要指针接收者）:")
	fmt.Println("   a = &v  // 而不是 a = v")

	fmt.Println("\n3. 提供两个版本的方法（不推荐）:")
	fmt.Println("   func (v Vertex) Abs() float64 { return v.abs() }")
	fmt.Println("   func (v *Vertex) abs() float64 { ... }")
}
```
