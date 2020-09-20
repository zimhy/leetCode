package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := getLength(headA)
	lenB := getLength(headB)
	if lenB > lenA {
		headC := headB
		headB = headA
		headA = headC
		lenC := lenB
		lenB = lenA
		lenA = lenC
	}
	for i := 0; i < (lenA - lenB); i++ {
		headA = headA.Next
	}

	for i := 0; i < lenB; i++ {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return headA

}

func main() {
	l1 := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: l1,
		},
	}

	l3 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  1,
				Next: l1,
			},
		},
	}

	l := getIntersectionNode(l2, l3)
	fmt.Println(l)
}
