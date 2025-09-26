package switchcase

import "fmt"

func MultiCaseExample() {
	// 方法1：多个case值用逗号分隔
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend! 🎉")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Workday 😴")
	default:
		fmt.Println("Unknown day")
	}

	// 数字示例
	score := 85

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80, score >= 85: // 注意：这里85会被80的条件先匹配
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// 更好的数字范围写法
	switch {
	case score >= 90:
		fmt.Println("Excellent!")
	case score >= 80 && score < 90:
		fmt.Println("Good!")
	case score >= 70 && score < 80:
		fmt.Println("Average")
	case score >= 60 && score < 70:
		fmt.Println("Pass")
	default:
		fmt.Println("Fail")
	}
}

func CharacterTypeExample() {
	char := 'A'

	switch char {
	case 'A', 'E', 'I', 'O', 'U':
		fmt.Printf("%c is a vowel\n", char)
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Printf("%c is a lowercase vowel\n", char)
	default:
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			fmt.Printf("%c is a consonant\n", char)
		} else {
			fmt.Printf("%c is not a letter\n", char)
		}
	}
}

func StatusCodeExample() {
	statusCode := 404

	switch statusCode {
	case 200, 201, 202, 204:
		fmt.Println("Success! ✅")
	case 301, 302, 304:
		fmt.Println("Redirection 🔀")
	case 400, 401, 403, 404:
		fmt.Println("Client Error ❌")
	case 500, 502, 503:
		fmt.Println("Server Error 💥")
	default:
		fmt.Println("Unknown status code")
	}
}
