package main

import (
	"errors"
	"fmt"
)

// 创建栈模型
type stack struct {
	MaxTop int
	Top int
	arr [5]int
}

// 入栈
func (this *stack) push(val int) (err error) {
	// 判断栈是否已满
	if this.Top == this.MaxTop - 1 {
		return errors.New("stack full")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

// 出栈， 后进先出
func (this *stack) pop() (val int, err error) {
	// 判断栈是否为空
	if this.Top == -1 {
		return 0,  errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}


// 打印栈
func (this *stack) list()  {
	// 判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d", i, this.arr[i])
	}
}

// 测试
func main() {
	stack := stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)

	stack.pop()
	stack.pop()

	stack.list()
}

































