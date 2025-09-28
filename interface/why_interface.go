package interfacego

import (
	"fmt"
	"math"
)

// 没有接口时的问题演示
func WithoutInterface() {
	fmt.Println("=== 没有接口时的问题 ===")

	// 不同的图形，相同的需求：计算面积
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}
	triangle := Triangle{Base: 6, Height: 4}

	// 每种类型都需要单独处理
	fmt.Printf("圆形面积: %.2f\n", calculateCircleArea(circle))
	fmt.Printf("矩形面积: %.2f\n", calculateRectangleArea(rectangle))
	fmt.Printf("三角形面积: %.2f\n", calculateTriangleArea(triangle))

	// 如果要批量处理？每种类型都需要不同的处理方式
	fmt.Println("\n问题：无法统一处理不同类型")
}

// 为每种类型定义单独的函数（代码重复）
func calculateCircleArea(c Circle) float64 {
	fmt.Printf("处理圆形...")
	return math.Pi * c.Radius * c.Radius
}

func calculateRectangleArea(r Rectangle) float64 {
	fmt.Printf("处理矩形...")
	return r.Width * r.Height
}

func calculateTriangleArea(t Triangle) float64 {
	fmt.Printf("处理三角形...")
	return 0.5 * t.Base * t.Height
}

// 定义类型
type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }
type Triangle struct{ Base, Height float64 }

// 使用接口解决问题
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 每种类型实现相同的接口方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	// 简化计算，假设是等腰三角形
	side := math.Sqrt(t.Base*t.Base/4 + t.Height*t.Height)
	return t.Base + 2*side
}

// 有接口后的优雅解决方案
func WithInterface() {
	fmt.Println("\n=== 使用接口的优雅解决方案 ===")

	// 统一的类型：Shape接口
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 3},
		Triangle{Base: 6, Height: 4},
	}

	// 统一的处理方式
	totalArea := 0.0
	totalPerimeter := 0.0

	for i, shape := range shapes {
		area := shape.Area()
		perimeter := shape.Perimeter()

		fmt.Printf("图形%d (%T): 面积=%.2f, 周长=%.2f\n",
			i+1, shape, area, perimeter)

		totalArea += area
		totalPerimeter += perimeter
	}

	fmt.Printf("总面积: %.2f, 总周长: %.2f\n", totalArea, totalPerimeter)
}
