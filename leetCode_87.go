package main

import (
	"fmt"
	"strings"
)

func isScramble(s1 string, s2 string) bool {
	if strings.Compare(s1, s2) == 0 {
		return true
	}
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	length := len(runes1)
	start1 := make(map[rune]int)
	start2 := make(map[rune]int)
	end2 := make(map[rune]int)
	for i, v := range runes1 {
		if i >= length-1 {
			break
		}
		countRune(&start1, v)
		countRune(&start2, runes2[i])
		countRune(&end2, runes2[length-1-i])
		if mapEquals(&start1, &start2) {
			if isScramble(s1[:i+1], s2[:i+1]) && isScramble(s1[i+1:], s2[i+1:]) {
				return true
			}
		}
		if mapEquals(&start1, &end2) {
			if isScramble(s1[0:i+1], s2[length-i-1:length]) && isScramble(s1[i+1:length], s2[0:length-i-1]) {
				return true
			}
		}

	}
	return false

}

func countRune(countMap *map[rune]int, char rune) {
	if (*countMap)[char] >= 0 {
		(*countMap)[char] = (*countMap)[char] + 1
	} else {
		(*countMap)[char] = 1
	}
}

func mapEquals(countMap1 *map[rune]int, countMap2 *map[rune]int) bool {
	for k, v := range *countMap2 {
		if (*countMap1)[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print(isScramble("abcde", "caebd"))
}
