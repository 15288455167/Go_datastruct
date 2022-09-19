package main

import (
	"fmt"
)

// Tree 创建二叉树模型
type Tree struct {
	val int
	left *Tree
	right *Tree
}

// PreOrder 前序遍历
func (T *Tree) PreOrder() (nums []int) {
	var pre func(node *Tree)
	pre = func(node *Tree) {
		if node == nil {
			return
		}
		nums = append(nums, node.val)
		pre(node.left)
		pre(node.right)
	}
	pre(T)
	return
}

// 中序遍历
func (T *Tree) inOrder() (nums []int) {
	var ino func(node *Tree)
	ino = func(node *Tree) {
		if node == nil {
			return
		}
		ino(node.left)
		nums = append(nums, node.val)
		ino(node.right)
	}
	ino(T)
	return
}

// 后序遍历
func (T *Tree) backOrder() (nums []int) {
	var bo func(node *Tree)
	bo = func(node *Tree) {
		if node == nil {
			return
		}
		bo(node.left)
		bo(node.right)
		nums = append(nums, node.val)
	}
	bo(T)
	return
}

// 层序遍历
func (T *Tree) levelOrder() [][]int {
	var ret [][]int
	if T == nil {
		return ret
	}

	q := []*Tree{T}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*Tree
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.val)
			if node.left != nil {
				p = append(p, node.left)
			}
			if node.right != nil {
				p = append(p, node.right)
			}
		}
		q = p
	}
	return ret
}

// 二叉树的最大深度
func (T *Tree) maxDepth() {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	var depth func(node *Tree) int
	depth = func(node *Tree) int {
		if node == nil {
			return 0
		}
		return max(depth(node.left), depth(node.right)) + 1
	}
	depth(T)

	fmt.Println(depth(T))
}

// 判断是否为二叉平衡树
func isSymmetric(root *Tree) bool {
	return check(root, root)
}

func check(p, q *Tree) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.val == q.val && check(p.left, q.right) && check(p.right, q.left)
}

// 反转二叉树
func invertTree(root *Tree) *Tree {
	 if root == nil {
		 return root
	 }
	 left := invertTree(root.left)
	 right := invertTree(root.right)
	 root.left = right
	 root.right = left
	 return root
}

// test
func main() {
	T := &Tree{}
	T.right = &Tree{val: 5}
	T.left = &Tree{val: 6}
	T.right.left = &Tree{val: 2}
	T.right.right = &Tree{val: 3}
	T.right.left.left = &Tree{val: 4}
	T.right.right.right = &Tree{val: 9}

	fmt.Println(T.PreOrder())
	fmt.Println(T.inOrder())
	fmt.Println(T.backOrder())
	fmt.Println(T.levelOrder())
	T.maxDepth()
	fmt.Println(isSymmetric(T))
}


























