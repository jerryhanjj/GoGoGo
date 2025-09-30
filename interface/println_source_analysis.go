package interfacego

import (
	"fmt"
	"reflect"
)

/*
fmt.Println 调用 String() 方法的源码分析

调用链路：
fmt.Println() → fmt.Fprintln() → fmt.Fprint() → pp.doPrint() → pp.printArg() → pp.handleMethods()

让我们逐层分析：
*/

// 模拟fmt包的核心逻辑
type printer struct {
	// 简化版的printer结构
}

// 1. fmt.Println的实现（简化版）
func simulatePrintln(a ...interface{}) {
	fmt.Printf("=== 模拟 fmt.Println 的执行过程 ===\n")

	for i, arg := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		// 这里调用核心的格式化逻辑
		simulatePrintArg(arg)
	}
	fmt.Println()
}

// 2. 核心的参数处理函数（模拟fmt包的printArg）
func simulatePrintArg(arg interface{}) {
	fmt.Printf("开始处理参数: %T\n", arg)

	// 检查是否实现了fmt.Stringer接口
	if handleStringer(arg) {
		return
	}

	// 如果没有实现Stringer，使用默认格式
	fmt.Printf("使用默认格式: %v", arg)
}

// 3. 处理Stringer接口的函数（这是关键部分）
func handleStringer(arg interface{}) bool {
	// 这相当于fmt包中的类型断言检查
	if stringer, ok := arg.(fmt.Stringer); ok {
		fmt.Printf("检测到实现了fmt.Stringer接口\n")
		result := stringer.String()
		fmt.Printf("调用String()方法，结果: %s", result)
		return true
	}
	fmt.Printf("未实现fmt.Stringer接口\n")
	return false
}

// 演示类型
type DemoPerson struct {
	Name string
	Age  int
}

func (p DemoPerson) String() string {
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

type DemoAnimal struct {
	Species string
	Name    string
}

// DemoAnimal没有String()方法

func SourceCodeAnalysis() {
	fmt.Println("=== Go源码中fmt.Println的实现分析 ===")

	fmt.Println("\n1. 调用链路:")
	fmt.Println("   fmt.Println() →")
	fmt.Println("   fmt.Fprintln(os.Stdout, a...) →")
	fmt.Println("   fmt.Fprint(w, a...) →")
	fmt.Println("   p.doPrint(a) →")
	fmt.Println("   p.printArg(arg) →")
	fmt.Println("   p.handleMethods()")

	fmt.Println("\n2. 关键代码片段分析:")
	fmt.Println(`
   // 在fmt包的print.go中，handleMethods函数:
   func (p *pp) handleMethods(verb rune) (handled bool) {
       // 检查是否实现了Stringer接口
       if verb != 'T' && verb != 'p' {
           if stringer, ok := p.arg.(Stringer); ok {
               p.fmt.fmtS(stringer.String())
               return true
           }
       }
       return false
   }`)

	fmt.Println("\n3. 实际演示:")
	person := DemoPerson{Name: "Alice", Age: 30}
	animal := DemoAnimal{Species: "Cat", Name: "Whiskers"}

	fmt.Println("\n--- 使用我们的模拟函数 ---")
	simulatePrintln(person)
	fmt.Println()
	simulatePrintln(animal)

	fmt.Println("\n--- 使用真正的fmt.Println ---")
	fmt.Println(person)
	fmt.Println(animal)
}

// 详细的类型检查演示
func TypeCheckDemo() {
	fmt.Println("\n=== 类型检查演示 ===")

	person := DemoPerson{Name: "Bob", Age: 25}

	// 1. 使用反射检查接口实现
	personType := reflect.TypeOf(person)
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	fmt.Printf("DemoPerson类型: %v\n", personType)
	fmt.Printf("fmt.Stringer接口: %v\n", stringerType)
	fmt.Printf("DemoPerson是否实现了Stringer: %v\n", personType.Implements(stringerType))

	// 2. 直接类型断言
	if stringer, ok := interface{}(person).(fmt.Stringer); ok {
		fmt.Printf("类型断言成功，调用String(): %s\n", stringer.String())
	} else {
		fmt.Println("类型断言失败")
	}

	// 3. 演示接口值
	var any interface{} = person
	if stringer, ok := any.(fmt.Stringer); ok {
		fmt.Printf("从interface{}转换成功: %s\n", stringer.String())
	}
}

// 性能对比演示
func PerformanceDemo() {
	fmt.Println("\n=== 性能对比演示 ===")

	person := DemoPerson{Name: "Charlie", Age: 35}

	// 直接调用String()方法
	fmt.Printf("直接调用String(): %s\n", person.String())

	// 通过fmt.Sprintf调用（内部会检查Stringer）
	fmt.Printf("通过fmt.Sprintf: %s\n", fmt.Sprintf("%v", person))

	// 通过fmt.Println调用（内部会检查Stringer）
	fmt.Print("通过fmt.Print: ")
	fmt.Println(person)
}
