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
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	resolve(root, sum, make([]*TreeNode, 0), &res)
	return res
}

func resolve(root *TreeNode, sum int, curr []*TreeNode, res *([][]int)) {
	if root == nil {
		return
	}
	if sum == root.Val && root.Right == nil && root.Left == nil {
		curr = append(curr, root)
		line := make([]int, len(curr))
		for i, v := range curr {
			line[i] = v.Val
		}
		*res = append(*res, line)
	}
	if root == nil {
		return
	}
	if root.Left != nil {
		resolve(root.Left, sum-root.Val, append(curr, root), res)
	}
	if root.Right != nil {
		resolve(root.Right, sum-root.Val, append(curr, root), res)
	}

}

func main() {
	l1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(pathSum(l1, 22))
}
