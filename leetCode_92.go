package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
 */

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if(m==n){
		return head
	}
	if m == 1 {
		reversedHead :=reverseFirstN(head,n-m)
		return  reversedHead;
	}else {
		reverseStart := head.Next
		leftList :=head
		for i := 0; i < m-2; i++{
			reverseStart = reverseStart.Next
			leftList = leftList.Next
		}
		reversedHead :=reverseFirstN(reverseStart,n-m)
		leftList.Next = reversedHead
		return head
	}
}

func reverseFirstN(head *ListNode, length int) *ListNode {
	end := head
	left := head
	right := head.Next
	left.Next = nil
	for i := 0; i < length; i++ {
		next :=right.Next
		right.Next = left
		left = right
		right = next
		end.Next = next
	}
	return left
}

func main() {
	l1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 5,

		},
	}


	l := reverseBetween(l1, 1,2)
	fmt.Println(l)
}
