package formodule

import "fmt"

func FnFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func ForClear() {
	sum := 1
	// for ; sum < 1000; {
	// 	sum += sum
	// }
	// Go 语言中，如果 for 循环的初始语句或后置语句为空，则分号可以省略。
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
