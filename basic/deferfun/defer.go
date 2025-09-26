package deferfun

import "fmt"

func DeferExample() {
	fmt.Println("Starting DeferExample")

	// 使用defer确保资源释放
	fmt.Println("Opening resource...")
	defer fmt.Println("Closing resource...") // 程序结束时执行

	// 模拟一些操作
	for i := 1; i <= 3; i++ {
		// 推迟调用的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的调用会按照后进先出的顺序调用。
		defer fmt.Printf("Processing step %d\n", i)
	}

	// 模拟一个可能的错误
	if true {
		fmt.Println("An error occurred!")
		return // 即使这里返回，defer仍会执行
	}

	fmt.Println("Finished processing")
}
