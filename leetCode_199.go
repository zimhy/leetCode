package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	spliter := TreeNode{}
	queue := NewQueue()
	queue.Push(root)
	queue.Push(&spliter)

	var lastOne *TreeNode
	for queue.Any() {
		node := queue.Pop().(*TreeNode)
		if node == &spliter && lastOne != nil {
			res = append(res, lastOne.Val)
			if queue.Any() {
				queue.Push(&spliter)
			}

		} else {
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		lastOne = node
	}
	return res
}

func main() {
	l1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	views := rightSideView(l1)
	fmt.Println(views)
}
