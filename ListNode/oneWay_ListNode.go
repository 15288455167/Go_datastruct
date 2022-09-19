package main

import "fmt"

// ListNode 创建链表模型
type ListNode struct {
	Val  int
	Next *ListNode
}

// InsertNode 头插法
func headInsertNode(head *ListNode, Node *ListNode) {
	Node.Next = head
}

// 尾插法
func tailInsertNode(head *ListNode, Node *ListNode) {
	for head.Next != nil {
		head = head.Next
	}
	head.Next = Node
}

// 输出链表
func outListNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// 测试
func main() {
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 1}
	//node2 := &ListNode{Val: 2}
	//node3 := &ListNode{Val: 3}

	tailInsertNode(head, node1)
	//tailInsertNode(head, node2)
	//tailInsertNode(head, node3)
	outListNode(head)
}




































