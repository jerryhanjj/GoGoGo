package interfacego

import (
	"fmt"
	"strings"
)

func RunAllExamples() {
	fmt.Println("=== Go语言隐式接口实现详解 ===")

	// 运行原始示例
	fmt.Println("1. 原始示例:")
	InterfaceImplicitly()

	fmt.Println("\n" + strings.Repeat("=", 50))

	// 运行详细演示
	ImpicitDemo()
	DemonstrateImplicitInterface()
	DemonstrateTypeAssertion()
	DemonstrateEmptyInterface()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("=== 为什么隐式实现如此重要？ ===")
	fmt.Println("1. 解耦: 类型和接口可以在不同包中定义")
	fmt.Println("2. 灵活: 已有类型可以'后天'实现新接口")
	fmt.Println("3. 组合: 支持接口组合和多重实现")
	fmt.Println("4. 测试友好: 容易创建mock对象")
}
