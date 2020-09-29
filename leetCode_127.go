package main

import (
	"fmt"
	"strings"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	fromToGroup := make(map[string]*[]string)
	visitmap := make(map[string]bool)
	for _, from := range append(wordList, beginWord) {
		toGroup := make([]string, 0)
		for _, to := range wordList {
			if isTransfer(from, to) {
				toGroup = append(toGroup, to)
			}
		}
		fromToGroup[from] = &toGroup
		visitmap[from] = false
	}
	count := 0
	queue := Queue{}
	queue.Push(beginWord)
	queue.Push("")
	for queue.Any() {
		from := queue.Pop().(string)
		if strings.Compare("", from) == 0 {
			if queue.Any() {
				count++
				queue.Push("")
			}
		} else {
			if strings.Compare(endWord, from) == 0 {
				return (count + 1)
			}
			if fromToGroup[from] != nil && len(*fromToGroup[from]) > 0 {
				for _, v := range *fromToGroup[from] {
					if !visitmap[v] {
						queue.Push(v)
						visitmap[v] = true
					}
				}
			}
		}
	}
	return 0
}

func isTransfer(from string, to string) bool {
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
	return true
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
