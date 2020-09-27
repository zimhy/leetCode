package main

import "fmt"

//func Sort(input []*SortNode) []*SortNode {
////	sorted := make([]*SortNode, len(input))
////	copy(sorted, input)
////	subSort(sorted, 0, len(input)-1)
////	return sorted
////}

/**
You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].



Example 1:

Input: nums = [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.


Constraints:

0 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
Accepted
139,411
Submissions
334,431
*/
//归并排序思路
//func countSmaller(nums []int) []int {
//	toSort := make([]*SortNode, len(nums))
//	for i := 0; i < len(nums); i++ {
//		toSort[i] = NewSortNode(nums[i], i)
//	}
//	Sort(toSort)
//	for _, v := range toSort {
//		nums[v.ori] = v.rightCount
//	}
//	return nums
//}
//
//// this function is going to call itself recursively with the left
//// and the right halves of the input slice (the divide)
//// then call the "merge" function (the conquest)
//func subSort(sorted []*SortNode, leftStart int, rightEnd int) {
//	if leftStart >= rightEnd {
//		return
//	}
//	middle := (leftStart + rightEnd) / 2
//	// calling itself recursively with both halves
//	subSort(sorted, leftStart, middle)
//	subSort(sorted, middle+1, rightEnd)
//	// merging the two sorted halves
//	merge(sorted, leftStart, rightEnd)
//}
//
//func merge(sorted []*SortNode, leftStart int, rightEnd int) {
//
//	temp := make([]*SortNode, len(sorted))
//	copy(temp, sorted)
//	leftEnd := (leftStart + rightEnd) / 2
//	rightStart := leftEnd + 1
//
//	left := leftStart
//	right := rightStart
//	index := leftStart
//
//	for left <= leftEnd && right <= rightEnd {
//		if sorted[left].val <= sorted[right].val {
//			temp[index] = sorted[left]
//			left++
//		} else {
//			temp[index] = sorted[right]
//			right++
//		}
//		temp[index].curr = index
//		index++
//	}
//	copy(temp[index:rightEnd+1], sorted[right:rightEnd+1])
//	copy(temp[index:rightEnd+1], sorted[left:leftEnd+1])
//	for i, v := range temp[leftStart : rightEnd+1] {
//		v.curr = i
//	}
//
//	for i, v := range sorted[leftStart : leftEnd+1] {
//		v.rightCount += (v.curr - i)
//	}
//
//	copy(sorted, temp)
//}
//
//type SortNode struct {
//	val        int
//	curr       int
//	ori        int
//	rightCount int
//}
//
//func NewSortNode(val int, index int) *SortNode {
//	node := &SortNode{val: val, ori: index, curr: index, rightCount: 0}
//	return node
//}
//

func main() {
	nums := []int{5, 2, 6, 1}
	countSmaller := countSmaller(nums)
	fmt.Println(countSmaller)
}

//排序树思路
func countSmaller(nums []int) []int {
	if len(nums) < 1 {
		return []int{}
	}
	var root *BinaryTreeNode
	for i, v := range nums {
		node := NewBinaryTreeNode(v, i)
		root = insertSortTree(root, node)
	}
	smallerCount := make([]int, len(nums))
	getSmallerCount(root, &smallerCount)
	return smallerCount
}

func getSmallerCount(root *BinaryTreeNode, smallerCount *[]int) {
	if root != nil {
		(*smallerCount)[root.origin] = root.smallCount
		getSmallerCount(root.Left, smallerCount)
		getSmallerCount(root.Right, smallerCount)
	}
}

func smallerCountPlus(root *BinaryTreeNode) {
	if root != nil {
		root.smallCount++
		smallerCountPlus(root.Left)
		smallerCountPlus(root.Right)
	}
}

type BinaryTreeNode struct {
	Left       *BinaryTreeNode
	Right      *BinaryTreeNode
	origin     int
	val        int
	smallCount int
}

func NewBinaryTreeNode(val, origin int) *BinaryTreeNode {
	return &BinaryTreeNode{origin: origin, val: val, smallCount: 0}
}

func insertSortTree(root *BinaryTreeNode, node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return root
	}
	if root == nil {
		root = node
		return root
	}
	if node.val < root.val {
		smallerCountPlus(root.Right)
		root.smallCount++
		root.Left = insertSortTree(root.Left, node)
	} else {
		root.Right = insertSortTree(root.Right, node)
	}
	return root
}
