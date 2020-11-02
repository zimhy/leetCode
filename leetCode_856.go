package main

import (
	"fmt"
)

type ComputeNode struct {
	value int
}

func NewComputeNode() *ComputeNode {
	return &ComputeNode{value: 0}
}
func scoreOfParentheses(S string) int {

	stack := NewStack()
	runes := []rune(S)
	flag := false
	value := 0
	for i, v := range runes {
		fmt.Println(i, "-", v)
		if '(' == v {
			stack.Push(NewComputeNode())
			flag = false
		} else {
			node := stack.Pop().(*ComputeNode)
			if !flag {
				node.value = 1
			} else {
				node.value = node.value * 2
			}
			if stack.Empty() {
				value += node.value
			} else {
				higher := stack.Peak().(*ComputeNode)
				higher.value = higher.value + node.value
			}
			flag = true
		}
	}
	return value

}

func main() {
	fmt.Println(scoreOfParentheses("(()()()())(()())"))
}
