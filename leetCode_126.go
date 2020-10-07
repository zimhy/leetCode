package main

import (
	"fmt"
	"strings"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	fromToGroup := make(map[string]*[]string)
	visitmap := make(map[string]bool)
	for _, from := range append(wordList, beginWord) {
		toGroup := make([]string, 0)
		for _, to := range wordList {
			if isTransferEnable(from, to) {
				toGroup = append(toGroup, to)
			}
		}
		fromToGroup[from] = &toGroup
		visitmap[from] = false
	}

	queue := Queue{}
	queue.Push(NewNode(beginWord, nil))
	queue.Push(NewNode("", nil))
	matched := false
	var res []*Node = make([]*Node, 0)
	for queue.Any() {
		from := queue.Pop().(*Node)
		if strings.Compare("", from.value) == 0 {
			if matched {
				break
			}
			if queue.Any() {
				queue.Push(NewNode("", from))
			}
		} else {
			if strings.Compare(endWord, from.value) == 0 {
				res = append(res, from)
				matched = true
			}
			if fromToGroup[from.value] != nil && len(*fromToGroup[from.value]) > 0 {
				for _, v := range *fromToGroup[from.value] {
					if !visitmap[v] {
						visitmap[v] = true
					} else {

					}
					queue.Push(NewNode(v, from))
				}
			}
		}
	}
	if !matched {
		return [][]string{}
	} else {
		result := [][]string{}
		for _, node := range res {

			result = append(result, getTrace(node))
		}
		return result
	}

}

func getTrace(node *Node) []string {
	if node.pre == nil {
		return []string{node.value}
	} else {
		return append(getTrace(node.pre), node.value)
	}
}

type Node struct {
	value string
	pre   *Node
}

func NewNode(value string, pre *Node) *Node {
	return &Node{value: value, pre: pre}
}

func isTransferEnable(from string, to string) bool {
	count := 0
	froms := []byte(from)
	tos := []byte(to)
	for i, _ := range froms {
		if froms[i] != tos[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	if count == 1 {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(findLadders("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}))
}
