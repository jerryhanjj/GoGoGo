package interfacego

import "fmt"

type Robot struct {
	model string
}

// Robot实现Write方法，隐式实现Writer接口
func (r Robot) Write(data []byte) (int, error) {
	fmt.Printf("机器人 %s 写入: %s\n", r.model, string(data))
	return len(data), nil
}

// Robot实现Speak方法，隐式实现Speaker接口
func (r Robot) Speak() string {
	return fmt.Sprintf("我是机器人 %s", r.model)
}

// ====== 实际应用示例 ======

// 假设这是第三方库的类型
type FileSystem struct {
	root string
}

func (fs FileSystem) Write(data []byte) (int, error) {
	fmt.Printf("写入文件系统 %s: %s\n", fs.root, string(data))
	return len(data), nil
}

// 我们定义的接口
type DataStore interface {
	Write([]byte) (int, error)
}

func ImpicitDemo() {
	fmt.Println("=== Go语言隐式接口实现详解 ===")

	// 概念解释
	fmt.Println("🤔 问题：如何理解'不需要显式声明T实现了I'？")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("📝 传统语言（如Java）的做法：")
	fmt.Println("   class T implements I { ... }")
	fmt.Println("   需要显式声明T实现了接口I")
	fmt.Println()

	fmt.Println("🎯 Go语言的做法：")
	fmt.Println("   只要T有I要求的所有方法，T就自动实现了I")
	fmt.Println("   这叫做'鸭子类型'：如果它像鸭子一样叫，它就是鸭子")
	fmt.Println()

	// 基础示例
	fmt.Println("🔥 基础示例：")
	fmt.Println("────────────")
	var i I = T{"Hello, 隐式接口!"} // T自动实现了I，无需声明
	i.M()
	fmt.Println()

	// 复杂示例
	fmt.Println("🚀 复杂示例：")
	fmt.Println("────────────")
	robot := Robot{model: "T-800"}

	// 同一个robot可以赋值给不同接口
	var writer Writer = robot   // robot隐式实现了Writer
	var speaker Speaker = robot // robot隐式实现了Speaker

	writer.Write([]byte("机器人数据"))
	fmt.Println(speaker.Speak())
	fmt.Println()

	// 实际应用场景
	fmt.Println("🌍 实际应用场景：")
	fmt.Println("───────────────")

	// FileSystem自动实现了我们的DataStore接口！
	var store DataStore = FileSystem{root: "/tmp"}
	store.Write([]byte("这展示了隐式实现的威力"))
	fmt.Println("💡 注意：FileSystem可能来自第三方库，但自动实现了我们的接口！")
	fmt.Println()

	// 关键优势
	fmt.Println("✨ 隐式接口的关键优势：")
	fmt.Println("─────────────────────")
	fmt.Println("1. 🔗 解耦性：接口定义者和实现者可以完全独立")
	fmt.Println("2. 🔄 灵活性：已存在的类型可以'后天'实现新接口")
	fmt.Println("3. 📦 组合性：一个类型可以同时实现多个接口")
	fmt.Println("4. 🧪 可测试：容易创建测试用的Mock对象")
	fmt.Println()

	// 核心原理
	fmt.Println("⚡ 核心原理总结：")
	fmt.Println("─────────────")
	fmt.Println("• 结构化类型系统：只看结构，不看声明")
	fmt.Println("• 鸭子类型：如果它叫起来像鸭子，那它就是鸭子")
	fmt.Println("• 编译时检查：Go编译器会自动检查类型是否满足接口")
	fmt.Println("• 运行时多态：通过接口可以实现多态行为")
}
