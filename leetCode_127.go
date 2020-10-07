package main

import (
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
	count, matched := bfsGetlength(beginWord, endWord, fromToGroup, visitmap)
	if matched {
		return count
	} else {
		return 0
	}
}

func bfsGetlength(beginWord string, endWord string, fromToGroup map[string]*[]string, visitmap map[string]bool) (int, bool) {
	count := 0
	queue := Queue{}
	queue.Push(beginWord)
	queue.Push("")
	matched := false
	for queue.Any() {
		from := queue.Pop().(string)
		if strings.Compare("", from) == 0 {
			if queue.Any() {
				count++
				queue.Push("")
			}
		} else {
			if strings.Compare(endWord, from) == 0 {
				matched = true
				count++
				break
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
	return count, matched
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
