package main

type ListNode struct {
	Val  int
	Next *ListNode
	len  int
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
