package generics

import "fmt"

// Node 链表节点
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList 链表结构
type LinkedList[T any] struct {
	head *Node[T]
	size int
}

// NewLinkedList 创建新链表
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		size: 0,
	}
}

// Append 在链表末尾添加元素
func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value, Next: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.size++
}

// Prepend 在链表开头添加元素
func (l *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value, Next: l.head}
	l.head = newNode
	l.size++
}

// InsertAt 在指定位置插入元素
func (l *LinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("索引越界: %d", index)
	}

	if index == 0 {
		l.Prepend(value)
		return nil
	}

	newNode := &Node[T]{Value: value}
	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	l.size++
	return nil
}

// Delete 删除第一个匹配的元素（需要comparable约束）
func (l *LinkedList[T]) DeleteValue(value T, equals func(a, b T) bool) bool {
	if l.head == nil {
		return false
	}

	// 如果头节点就是要删除的
	if equals(l.head.Value, value) {
		l.head = l.head.Next
		l.size--
		return true
	}

	current := l.head
	for current.Next != nil {
		if equals(current.Next.Value, value) {
			current.Next = current.Next.Next
			l.size--
			return true
		}
		current = current.Next
	}

	return false
}

// DeleteAt 删除指定位置的元素
func (l *LinkedList[T]) DeleteAt(index int) error {
	if index < 0 || index >= l.size {
		return fmt.Errorf("索引越界: %d", index)
	}

	if index == 0 {
		l.head = l.head.Next
		l.size--
		return nil
	}

	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	l.size--
	return nil
}

// Get 获取指定位置的元素
func (l *LinkedList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, fmt.Errorf("索引越界: %d", index)
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, nil
}

// Set 修改指定位置的元素
func (l *LinkedList[T]) Set(index int, value T) error {
	if index < 0 || index >= l.size {
		return fmt.Errorf("索引越界: %d", index)
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	current.Value = value
	return nil
}

// IndexOf 查找元素第一次出现的位置
func (l *LinkedList[T]) IndexOf(value T, equals func(a, b T) bool) int {
	current := l.head
	index := 0

	for current != nil {
		if equals(current.Value, value) {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

// Contains 检查链表是否包含某个元素
func (l *LinkedList[T]) Contains(value T, equals func(a, b T) bool) bool {
	return l.IndexOf(value, equals) != -1
}

// Length 获取链表长度
func (l *LinkedList[T]) Length() int {
	return l.size
}

// IsEmpty 检查链表是否为空
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear 清空链表
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.size = 0
}

// ToSlice 转换为切片
func (l *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, l.size)
	current := l.head

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// Print 打印链表内容
func (l *LinkedList[T]) Print() {
	if l.IsEmpty() {
		fmt.Println("链表为空")
		return
	}

	current := l.head
	fmt.Print("链表: ")
	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

// Reverse 反转链表
func (l *LinkedList[T]) Reverse() {
	if l.head == nil || l.head.Next == nil {
		return
	}

	var prev *Node[T]
	current := l.head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.head = prev
}

// 旧的 List 类型（保持兼容性）
type List[T any] struct {
	next  *List[T]
	value T
}
