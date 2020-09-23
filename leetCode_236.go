package main

import "fmt"

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ansestors := make([]*TreeNode, 0)
	findInTree(root, p, q, &ansestors)
	if len(ansestors) > 0 {
		return ansestors[0]
	} else {
		return root
	}

}

func findInTree(root, p, q *TreeNode, ancestors *([]*TreeNode)) int {
	if root == nil {
		return 0
	}
	this := 0
	if root == p || root == q {
		this = 1
	}
	left := findInTree(root.Left, p, q, ancestors)
	right := findInTree(root.Right, p, q, ancestors)
	if left+right+this == 2 {
		newAncestors := append(*ancestors, root)
		*ancestors = newAncestors
	}
	return left + right + this
}

func main() {
	q := &TreeNode{
		Val: 4,
	}
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: q,
		},
	}

	l1 := &TreeNode{
		Val:  3,
		Left: p,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	ancestor := lowestCommonAncestor(l1, p, q)
	fmt.Println(ancestor.Val)
}
