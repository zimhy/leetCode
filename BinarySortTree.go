package main

//排序二叉树

func insert(root *TreeNode, node *TreeNode) {
	if node == nil {
		return
	}
	if root == nil {
		*root = *node
	}
	if node.Val < root.Val {
		insert(root.Left, node)
	} else {
		insert(root.Right, node)
	}
}
