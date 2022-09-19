package main

import "fmt"

// Queue 定义队列模型，此处用两个栈模拟一个队列，一个输入栈，一个输出栈。
type Queue struct {
	inStack, outStack []int
}

// Constructor 初始化队列
func Constructor() Queue {
	return Queue{}
}

// Push 入队列操作
func (q *Queue) Push(val int) {
	q.inStack = append(q.inStack, val)
}

// 输入到outStack栈操作
func (q *Queue) allin(){
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack) - 1])
		q.inStack = q.inStack[:len(q.inStack) - 1]
	}
}

// Pop 出队列操作
func (q *Queue) Pop() int {
	if len(q.outStack) == 0 {
		q.allin()
	}
	x := q.outStack[len(q.outStack) - 1]
	return x
}

// Peek 返回队列顶部元素
func (q *Queue) Peek() int {
	if len(q.outStack) == 0 {
		q.allin()
	}
	x := q.outStack[len(q.outStack) - 1]
	return x
}

// Empty 判断队列是否为空
func (q *Queue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}


// test
func main() {
	q := Constructor()
	q.Push(5)
	q.Push(4)
	q.Push(7)
	q.Push(8)

	fmt.Println(q.Pop())
	fmt.Println(q)
}































