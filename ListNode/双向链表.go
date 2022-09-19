package main

import "fmt"

// 创建双向链表模型
type twoListNode struct {
	val  int
	Next *twoListNode
	Pre  *twoListNode
}

// 头插法
func headInsert(head *twoListNode, Node *twoListNode) {
	temp := head
	for {
		if temp.Pre == nil {
			break
		}
		temp = temp.Pre
	}
	temp.Pre = Node
	Node.Next = temp

}

// 尾插法
func tailInsert(head *twoListNode, Node *twoListNode) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = Node
	Node.Pre = temp
}

// 打印链表尾插法
func outTwoListNode(head *twoListNode) {
	for head != nil {
		fmt.Println(head.val)
		head = head.Next
	}
}

// 打印头插法
func outTwoListNode02(head *twoListNode) {
	for head != nil {
		fmt.Println(head.val)
		head = head.Pre
	}
}


// 测试
func main() {
	head := &twoListNode{}
	node1 := &twoListNode{val: 2}
	node2 := &twoListNode{val: 4}
	node3 := &twoListNode{val: 7}

	//tailInsert(head, node3)
	//tailInsert(head, node1)
	//tailInsert(head, node2)

	headInsert(head, node1)
	headInsert(head, node3)
	headInsert(head, node2)

	outTwoListNode02(head)
	fmt.Println(*head)
}









































