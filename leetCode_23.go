package main

import (
	"sort"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.



Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []


Constraints:

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] is sorted in ascending order.
The sum of lists[i].length won't exceed 10^4.
*/

type NodeSlice []*ListNodeWithLen

func (s NodeSlice) Len() int           { return len(s) }
func (s NodeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s NodeSlice) Less(i, j int) bool { return s[i].len > s[j].len }

type ListNodeWithLen struct {
	listNode *ListNode
	len      int
}

func NewListNodeWithLen(listNode *ListNode) *ListNodeWithLen {
	return &ListNodeWithLen{listNode: listNode}
}

func calculateLength(l *ListNodeWithLen) int {
	l.len = GetLength(l.listNode)
	return l.len
}

func mergeKLists(lists []*ListNode) *ListNode {
	var linkedLists []*ListNodeWithLen = make([]*ListNodeWithLen, len(lists))
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	for i := 0; i < len(lists); i++ {
		node := lists[i]
		var tempNode *ListNodeWithLen = NewListNodeWithLen(node)
		linkedLists[i] = tempNode
	}
	var nodeSlices NodeSlice = linkedLists
	for nodeSlices.Len() > 1 {
		sort.Sort(nodeSlices)
		twoLists := mergeTwoLists(nodeSlices[1].listNode, nodeSlices[0].listNode)
		var tempNode *ListNodeWithLen = NewListNodeWithLen(twoLists)
		nodeSlices = nodeSlices[2:]
		nodeSlices = append(nodeSlices, tempNode)
	}
	return nodeSlices[0].listNode

}

func GetLength(l *ListNode) int {
	if l == nil {
		return 0
	}
	if l.Next == nil {
		return 1
	} else {
		return GetLength(l.Next) + 1
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l3 := l1
		l1 = l2
		l2 = l3
	}
	l := l1
	for {
		if l1.Next == nil {
			l1.Next = l2
			break
		}
		if l2 == nil {
			//l2.Next = l1
			break
		}
		if l2.Val == l1.Val || (l2.Val > l1.Val && l2.Val < l1.Next.Val) {
			l1Next := l1.Next
			l2Next := l2.Next
			l1.Next = l2
			l2.Next = l1Next
			l2 = l2Next
			l1 = l1.Next
		} else if l2.Val < l1.Val {
			l3 := l1
			l1 = l2
			l2 = l3
		} else {
			l1 = l1.Next
		}

	}

	return l
}
