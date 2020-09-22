package main

import "fmt"

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
func countSmaller(nums []int) []int {
	toSort := make([]*SortNode, len(nums))
	for i := 0; i < len(nums); i++ {
		toSort[i] = NewSortNode(nums[i], i)
	}
	Sort(toSort)
	for _, v := range toSort {
		nums[v.ori] = v.rightCount
	}
	return nums
}

func Sort(input []*SortNode) []*SortNode {
	sorted := make([]*SortNode, len(input))
	copy(sorted, input)
	subSort(sorted, 0, len(input)-1)
	return sorted
}

// this function is going to call itself recursively with the left
// and the right halves of the input slice (the divide)
// then call the "merge" function (the conquest)
func subSort(sorted []*SortNode, leftStart int, rightEnd int) {
	if leftStart >= rightEnd {
		return
	}
	middle := (leftStart + rightEnd) / 2
	// calling itself recursively with both halves
	subSort(sorted, leftStart, middle)
	subSort(sorted, middle+1, rightEnd)
	// merging the two sorted halves
	merge(sorted, leftStart, rightEnd)
}

func merge(sorted []*SortNode, leftStart int, rightEnd int) {

	temp := make([]*SortNode, len(sorted))
	copy(temp, sorted)
	leftEnd := (leftStart + rightEnd) / 2
	rightStart := leftEnd + 1

	left := leftStart
	right := rightStart
	index := leftStart

	for left <= leftEnd && right <= rightEnd {
		if sorted[left].val <= sorted[right].val {
			temp[index] = sorted[left]
			left++
		} else {
			temp[index] = sorted[right]
			right++
		}
		temp[index].curr = index
		index++
	}
	copy(temp[index:rightEnd+1], sorted[right:rightEnd+1])
	copy(temp[index:rightEnd+1], sorted[left:leftEnd+1])
	for i, v := range temp[leftStart : rightEnd+1] {
		v.curr = i
	}

	for i, v := range sorted[leftStart : leftEnd+1] {
		v.rightCount += (v.curr - i)
	}

	copy(sorted, temp)
}

type SortNode struct {
	val        int
	curr       int
	ori        int
	rightCount int
}

func NewSortNode(val int, index int) *SortNode {
	node := &SortNode{val: val, ori: index, curr: index, rightCount: 0}
	return node
}

func main() {
	nums := []int{5, 2, 6, 1}
	countSmaller(nums)
	fmt.Println(nums)
}
