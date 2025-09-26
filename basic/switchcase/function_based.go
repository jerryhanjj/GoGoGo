package switchcase

import "fmt"

// 方法3：将公共逻辑提取到函数中
func handleWeekend() {
	fmt.Println("It's weekend! Time to relax 🌴")
	fmt.Println("Sleep in, have fun!")
}

func handleWorkday() {
	fmt.Println("It's a workday 💼")
	fmt.Println("Time to be productive!")
}

func handleSpecialDay(day string) {
	fmt.Printf("Special handling for %s\n", day)
}

func FunctionBasedSwitch() {
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		handleWeekend()
	case "Monday", "Tuesday", "Wednesday", "Thursday":
		handleWorkday()
	case "Friday":
		handleWorkday()
		fmt.Println("But TGIF! 🎉")
	default:
		handleSpecialDay(day)
	}
}

// 更复杂的示例：HTTP状态码处理
func logSuccess(code int) {
	fmt.Printf("✅ Success: HTTP %d\n", code)
}

func logError(code int, errorType string) {
	fmt.Printf("❌ %s Error: HTTP %d\n", errorType, code)
}

func logInfo(code int, infoType string) {
	fmt.Printf("ℹ️  %s: HTTP %d\n", infoType, code)
}

func HandleHTTPStatus(statusCode int) {
	switch statusCode {
	case 200, 201, 202, 204:
		logSuccess(statusCode)
	case 400, 401, 403, 404, 409:
		logError(statusCode, "Client")
	case 500, 502, 503, 504:
		logError(statusCode, "Server")
	case 301, 302, 304:
		logInfo(statusCode, "Redirection")
	default:
		fmt.Printf("🤔 Unknown status: HTTP %d\n", statusCode)
	}
}

// 使用接口的高级示例
type ResponseHandler interface {
	Handle(code int)
}

type SuccessHandler struct{}

func (h SuccessHandler) Handle(code int) {
	fmt.Printf("🎉 Request successful with code %d\n", code)
}

type ErrorHandler struct{}

func (h ErrorHandler) Handle(code int) {
	fmt.Printf("💥 Request failed with code %d\n", code)
}

func AdvancedSwitchWithInterface(statusCode int) {
	var handler ResponseHandler

	switch {
	case statusCode >= 200 && statusCode < 300:
		handler = SuccessHandler{}
	case statusCode >= 400:
		handler = ErrorHandler{}
	default:
		fmt.Printf("ℹ️ Info response: %d\n", statusCode)
		return
	}

	handler.Handle(statusCode)
}
