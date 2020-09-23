package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	flatternOne(root)
}

func flatternOne(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	var leftEnd *TreeNode
	rightStart := root.Right
	leftStart := root.Left
	if root.Left != nil {
		leftEnd = flatternOne(root.Left)
		root.Right = root.Left
		if leftEnd != nil {
			leftEnd.Right = rightStart
		}
		root.Left = nil
	}
	if rightStart != nil {
		return flatternOne(rightStart)
	} else if leftStart != nil {
		return leftEnd
	} else {
		return root
	}

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
	flatten(l1)
}
