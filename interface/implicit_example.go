package interfacego

import "fmt"

// 定义一个写入器接口
type Writer interface {
	Write([]byte) (int, error)
}

// 定义一个说话接口
type Speaker interface {
	Speak() string
}

// 定义一个多功能接口
type WriterSpeaker interface {
	Writer
	Speaker
}

// 文件类型
type File struct {
	name string
}

// File实现了Write方法，因此隐式实现了Writer接口
func (f File) Write(data []byte) (int, error) {
	fmt.Printf("写入文件 %s: %s\n", f.name, string(data))
	return len(data), nil
}

// 人类类型
type Person struct {
	name string
}

// Person实现了Speak方法，因此隐式实现了Speaker接口
func (p Person) Speak() string {
	return fmt.Sprintf("我是 %s", p.name)
}

// 演示隐式接口实现
func DemonstrateImplicitInterface() {
	fmt.Println("=== Go语言隐式接口实现演示 ===")

	// 1. File实现了Writer接口
	var w Writer = File{name: "test.txt"}
	w.Write([]byte("Hello World"))

	// 2. Person实现了Speaker接口
	var s Speaker = Person{name: "张三"}
	fmt.Println(s.Speak())

	// 3. Robot实现了多个接口
	robot := Robot{model: "T-800"}

	// Robot可以作为Writer使用
	var writer Writer = robot
	writer.Write([]byte("机器人数据"))

	// Robot可以作为Speaker使用
	var speaker Speaker = robot
	fmt.Println(speaker.Speak())

	// Robot可以作为WriterSpeaker使用
	var ws WriterSpeaker = robot
	ws.Write([]byte("多功能数据"))
	fmt.Println(ws.Speak())

	fmt.Println("\n=== 关键特性 ===")
	fmt.Println("1. 无需使用 'implements' 关键字")
	fmt.Println("2. 只要有对应的方法，就自动实现接口")
	fmt.Println("3. 可以同时实现多个接口")
	fmt.Println("4. 支持接口组合")
}

// 演示类型断言
func DemonstrateTypeAssertion() {
	fmt.Println("\n=== 类型断言演示 ===")

	var i interface{} = Robot{model: "R2D2"}

	// 类型断言检查具体类型
	if robot, ok := i.(Robot); ok {
		fmt.Printf("这是一个机器人: %s\n", robot.model)
	}

	// 检查是否实现了特定接口
	if writer, ok := i.(Writer); ok {
		fmt.Println("这个对象可以写入数据")
		writer.Write([]byte("类型断言测试"))
	}
}

// 演示空接口的威力
func DemonstrateEmptyInterface() {
	fmt.Println("\n=== 空接口演示 ===")

	// 空接口可以持有任何类型
	var anything interface{}

	anything = 42
	fmt.Printf("数字: %v\n", anything)

	anything = "Hello"
	fmt.Printf("字符串: %v\n", anything)

	anything = Person{name: "李四"}
	fmt.Printf("结构体: %v\n", anything)

	// 类型切换
	switch v := anything.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s\n", v)
	case Person:
		fmt.Printf("人类: %s\n", v.Speak())
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}
