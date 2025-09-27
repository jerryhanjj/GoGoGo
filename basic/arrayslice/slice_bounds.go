package arrayslice

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func RunSliceBounds() {
	s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	// fmt.Println(s)
	// s = s[:2]
	// fmt.Println(s)
	// s = s[:]
	// fmt.Println(s)
	/*
		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/

	printSlice(s)
	//截取长度为0
	s = s[:0]
	printSlice(s)
	//拓展长度
	s = s[:4]
	printSlice(s)
	//舍弃前两个值
	s = s[2:]
	printSlice(s)

	var s1 []int
	fmt.Println(s1, len(s1), cap(s1))
	if s1 == nil {
		fmt.Println("nil")
	}
}
