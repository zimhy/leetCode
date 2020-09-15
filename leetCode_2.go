package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lenl1 := getLength(l1)
	lenl2 := getLength(l2)
	l3 := l1
	if lenl2 > lenl1 {

		l1 = l2
		l2 = l3
	}
	l3 = l1
	tmpNode := ListNode{
		0,
		nil,
	}
	for {
		if l1 == nil {
			break
		}
		if l2 == nil {
			l2 = &tmpNode
		}
		l1.Val = l1.Val + l2.Val
		if l1.Val >= 10 {
			l1.Val = l1.Val % 10
			if nil == l1.Next {
				l1.Next = &ListNode{
					0,
					nil,
				}
			}
			l1.Next.Val = l1.Next.Val + 1
		}

		l1 = l1.Next
		l2 = l2.Next
	}
	return l3
}

func getLength(l *ListNode) int {
	if l == nil {
		return 0
	}
	if l.Next == nil {
		return 1
	} else {
		return getLength(l.Next) + 1
	}
}

//
//func main() {
//	l1 := &ListNode{
//		1,
//		nil,
//		//&ListNode{
//		//	4,
//		//	&ListNode{3, nil},
//		//},
//	}
//	l2 := &ListNode{
//		9,
//		&ListNode{
//			9,
//			nil,
//			//&ListNode{4, nil},
//		},
//	}
//
//	ret := addTwoNumbers(l1, l2)
//	fmt.Println("%v", ret)
//}
