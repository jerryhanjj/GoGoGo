package generics

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v 和 x 的类型为 T，它拥有 comparable 可比较的约束，
		// 因此我们可以使用 ==
		if v == x {
			return i
		}
	}
	return -1
}

func indexDemo() {
	si := []int{10, 15, 20, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bob", "jerry"}
	fmt.Println(Index(ss, "hello"))
}

// 链表演示函数
func listDemo() {
	fmt.Println("\n=== 链表演示 ===")

	// 创建整数链表
	intList := NewLinkedList[int]()

	// 添加元素
	fmt.Println("\n1. 添加元素:")
	intList.Append(10)
	intList.Append(20)
	intList.Append(30)
	intList.Prepend(5)
	intList.Print() // 5 -> 10 -> 20 -> 30

	// 在指定位置插入
	fmt.Println("\n2. 在索引2插入15:")
	intList.InsertAt(2, 15)
	intList.Print() // 5 -> 10 -> 15 -> 20 -> 30

	// 获取元素
	fmt.Println("\n3. 获取元素:")
	if val, err := intList.Get(2); err == nil {
		fmt.Printf("索引2的元素: %d\n", val)
	}

	// 修改元素
	fmt.Println("\n4. 修改索引2的元素为100:")
	intList.Set(2, 100)
	intList.Print() // 5 -> 10 -> 100 -> 20 -> 30

	// 查找元素索引
	fmt.Println("\n5. 查找元素:")
	equals := func(a, b int) bool { return a == b }
	idx := intList.IndexOf(20, equals)
	fmt.Printf("元素20的索引: %d\n", idx)
	fmt.Printf("是否包含100: %v\n", intList.Contains(100, equals))
	fmt.Printf("是否包含999: %v\n", intList.Contains(999, equals))

	// 获取长度
	fmt.Println("\n6. 链表长度:")
	fmt.Printf("当前长度: %d\n", intList.Length())

	// 删除元素
	fmt.Println("\n7. 删除元素:")
	intList.DeleteValue(100, equals)
	intList.Print() // 5 -> 10 -> 20 -> 30

	fmt.Println("\n8. 删除索引1的元素:")
	intList.DeleteAt(1)
	intList.Print() // 5 -> 20 -> 30

	// 反转链表
	fmt.Println("\n9. 反转链表:")
	intList.Reverse()
	intList.Print() // 30 -> 20 -> 5

	// 转换为切片
	fmt.Println("\n10. 转换为切片:")
	slice := intList.ToSlice()
	fmt.Printf("切片: %v\n", slice)

	// 字符串链表示例
	fmt.Println("\n=== 字符串链表演示 ===")
	strList := NewLinkedList[string]()
	strList.Append("Hello")
	strList.Append("World")
	strList.Append("Go")
	strList.Prepend("!")
	strList.Print()

	strEquals := func(a, b string) bool { return a == b }
	idx2 := strList.IndexOf("Go", strEquals)
	fmt.Printf("\"Go\"的索引: %d\n", idx2)

	// 清空链表
	fmt.Println("\n11. 清空链表:")
	strList.Clear()
	fmt.Printf("清空后长度: %d, 是否为空: %v\n", strList.Length(), strList.IsEmpty())
}

func RunGenerics() {
	indexDemo()
	listDemo()
}
