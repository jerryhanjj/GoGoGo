package switchcase

import "fmt"

func FallthroughExample() {
	// 方法2：使用fallthrough（谨慎使用）
	grade := 'B'

	fmt.Print("Grade analysis: ")
	switch grade {
	case 'A':
		fmt.Print("Excellent! ")
		fallthrough // 继续执行下一个case
	case 'B':
		fmt.Print("Good work! ")
		fallthrough
	case 'C':
		fmt.Print("You passed. ")
		fallthrough
	case 'D':
		fmt.Print("Better study more. ")
	case 'F':
		fmt.Println("You failed.")
	default:
		fmt.Println("Invalid grade.")
	}

	// 注意：fallthrough有风险，因为它会无条件执行下一个case
	// 即使下一个case的条件不匹配也会执行

	number := 2
	fmt.Printf("\nAnalyzing number %d:\n", number)

	switch number {
	case 1:
		fmt.Println("It's one")
		fallthrough
	case 2:
		fmt.Println("It's small") // 这个会执行
		fallthrough
	case 3:
		fmt.Println("It's prime") // 这个也会执行（即使number不是3）
		// 没有fallthrough，所以停止
	case 4:
		fmt.Println("It's even") // 这个不会执行
	}
}

// 更安全的fallthrough示例
func SafeFallthroughExample() {
	userType := "admin"

	permissions := []string{}

	switch userType {
	case "superadmin":
		permissions = append(permissions, "system_config")
		fallthrough
	case "admin":
		permissions = append(permissions, "user_management")
		fallthrough
	case "moderator":
		permissions = append(permissions, "content_moderation")
		fallthrough
	case "user":
		permissions = append(permissions, "basic_access")
	}

	fmt.Printf("User type: %s\n", userType)
	fmt.Printf("Permissions: %v\n", permissions)
}
