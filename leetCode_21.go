package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	if l1.Val > l2.Val {
//		l3 := l1
//		l1 = l2
//		l2 = l3
//	}
//	l := l1
//	for {
//		if l1.Next == nil {
//			l1.Next = l2
//			break
//		}
//		if l2 == nil {
//			//l2.Next = l1
//			break
//		}
//		if l2.Val == l1.Val || (l2.Val > l1.Val && l2.Val < l1.Next.Val) {
//			l1Next := l1.Next
//			l2Next := l2.Next
//			l1.Next = l2
//			l2.Next = l1Next
//			l2 = l2Next
//			l1 = l1.Next
//		} else if l2.Val < l1.Val {
//			l3 := l1
//			l1 = l2
//			l2 = l3
//		} else {
//			l1 = l1.Next
//		}
//
//	}
//
//	return l
//}

//func main() {
//	l1 := &ListNode{
//		Val: 1,
//		Next: &ListNode{
//			Val: 2,
//			Next: &ListNode{
//				Val: 4,
//			},
//		},
//	}
//
//	l2 := &ListNode{
//		Val: 1,
//		Next: &ListNode{
//			Val: 3,
//			Next: &ListNode{
//				Val: 4,
//			},
//		},
//	}
//	l := mergeTwoLists(l1, l2)
//	fmt.Println(l)
//}
