```go
package interfacego

import (
	"fmt"
	"math"
)

// 定义接口
type Calculator interface {
	Calculate() float64
}

// 使用值接收者的类型
type Circle struct {
	Radius float64
}

// 值接收者方法
func (c Circle) Calculate() float64 {
	fmt.Printf("  计算圆面积，半径=%.1f\n", c.Radius)
	return math.Pi * c.Radius * c.Radius
}

// 使用指针接收者的类型
type Rectangle struct {
	Width, Height float64
}

// 指针接收者方法
func (r *Rectangle) Calculate() float64 {
	fmt.Printf("  计算矩形面积，宽=%.1f，高=%.1f\n", r.Width, r.Height)
	return r.Width * r.Height
}

func ValueReceiverDemo() {
	fmt.Println("=== 值接收者：值和指针都能实现接口 ===\n")

	// Circle 使用值接收者方法
	circle := Circle{Radius: 5}
	circlePtr := &Circle{Radius: 3}

	var calc Calculator

	// 1. 值可以实现接口
	calc = circle // ✅ Circle 实现了 Calculator
	fmt.Printf("1. 值实现接口: circle -> %T\n", calc)
	result1 := calc.Calculate()
	fmt.Printf("   结果: %.2f\n\n", result1)

	// 2. 指针也可以实现接口
	calc = circlePtr // ✅ *Circle 也实现了 Calculator
	fmt.Printf("2. 指针实现接口: circlePtr -> %T\n", calc)
	result2 := calc.Calculate()
	fmt.Printf("   结果: %.2f\n\n", result2)

	fmt.Println("结论：值接收者方法让 值和指针 都能实现接口！")
}

func PointerReceiverDemo() {
	fmt.Println("\n=== 指针接收者：只有指针能实现接口 ===\n")

	// Rectangle 使用指针接收者方法
	rectPtr := &Rectangle{Width: 5, Height: 6}

	var calc Calculator

	// 1. 值不能实现接口
	// rect := Rectangle{Width: 4, Height: 3}
	// calc = rect  // ❌ 编译错误：Rectangle 没有实现 Calculator
	fmt.Println("1. 值不能实现接口: rect -> 编译错误")
	fmt.Println("   错误信息: Rectangle does not implement Calculator")
	fmt.Println("   原因: Calculate method has pointer receiver")

	// 2. 只有指针可以实现接口
	calc = rectPtr // ✅ *Rectangle 实现了 Calculator
	fmt.Printf("\n2. 指针实现接口: rectPtr -> %T\n", calc)
	result := calc.Calculate()
	fmt.Printf("   结果: %.2f\n\n", result)

	fmt.Println("结论：指针接收者方法只让 指针 能实现接口！")
}

func MethodCallVsInterface() {
	fmt.Println("\n=== 方法调用 vs 接口实现的区别 ===\n")

	circle := Circle{Radius: 2}
	rect := Rectangle{Width: 3, Height: 4}

	fmt.Println("方法调用（都可以自动转换）:")

	// 值接收者方法调用
	fmt.Print("  circle.Calculate(): ")
	circle.Calculate() // ✅ 直接调用

	fmt.Print("  (&circle).Calculate(): ")
	(&circle).Calculate() // ✅ 指针调用值接收者方法，自动解引用

	// 指针接收者方法调用
	fmt.Print("  rect.Calculate(): ")
	rect.Calculate() // ✅ 值调用指针接收者方法，自动取地址

	fmt.Print("  (&rect).Calculate(): ")
	(&rect).Calculate() // ✅ 直接调用

	fmt.Println("\n接口赋值（规则更严格）:")
	var calc Calculator

	fmt.Println("  值接收者方法:")
	calc = circle // ✅ 值可以
	fmt.Printf("    calc = circle: 成功 -> %T\n", calc)
	calc = &circle // ✅ 指针也可以
	fmt.Printf("    calc = &circle: 成功 -> %T\n", calc)

	fmt.Println("  指针接收者方法:")
	// calc = rect     // ❌ 值不行
	fmt.Println("    calc = rect: 失败 - 编译错误")
	calc = &rect // ✅ 只有指针可以
	fmt.Printf("    calc = &rect: 成功 -> %T\n", calc)
}

func WhyThisRule() {
	fmt.Println("\n=== 为什么有这样的规则？ ===\n")

	fmt.Println("1. 值接收者方法特点:")
	fmt.Println("   - 接收者是值的拷贝")
	fmt.Println("   - 不会修改原始数据")
	fmt.Println("   - 值和指针调用效果相同")
	fmt.Println("   - 所以值和指针都能安全地实现接口")

	fmt.Println("\n2. 指针接收者方法特点:")
	fmt.Println("   - 接收者是指针，可能修改原始数据")
	fmt.Println("   - 如果接口存储值，方法无法修改原始对象")
	fmt.Println("   - 这会导致不一致的行为")
	fmt.Println("   - 所以只允许指针实现接口")

	fmt.Println("\n3. 设计哲学:")
	fmt.Println("   - 保证接口行为的一致性")
	fmt.Println("   - 避免意外的副作用")
	fmt.Println("   - 明确值语义 vs 指针语义")
}

func PracticalExample() {
	fmt.Println("\n=== 实际应用示例 ===\n")

	// 一个实际的例子：计算器切片
	var calculators []Calculator

	// 值接收者类型：两种方式都可以
	circle1 := Circle{Radius: 1}
	circle2 := &Circle{Radius: 2}

	calculators = append(calculators, circle1) // ✅ 值
	calculators = append(calculators, circle2) // ✅ 指针

	// 指针接收者类型：只能用指针
	// rect1 := Rectangle{Width: 3, Height: 4}
	rect2 := &Rectangle{Width: 5, Height: 6}

	// calculators = append(calculators, rect1)  // ❌ 不能用值
	calculators = append(calculators, rect2) // ✅ 只能用指针

	fmt.Println("计算所有图形的面积:")
	total := 0.0
	for i, calc := range calculators {
		fmt.Printf("图形%d (%T): ", i+1, calc)
		area := calc.Calculate()
		fmt.Printf("面积 = %.2f\n", area)
		total += area
	}
	fmt.Printf("总面积: %.2f\n", total)
}
```