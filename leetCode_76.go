package main

import (
	"fmt"
	"math"
	"strings"
)

func minWindow(s string, t string) string {
	windowToMove := initWindow(t)
	stringArr := strings.Split(s, "")
	for i, v := range stringArr {
		processNext(i, v, &windowToMove)
	}
	if windowToMove.matched {
		return strings.Join(stringArr[windowToMove.start:windowToMove.end+1], "")
	} else {
		return ""
	}

}

type Window struct {
	start     int
	end       int
	collected map[string]*[]int
	startChar string
	pattern   map[string]int
	minWindow int
	index     int
	matched   bool
}

func initWindow(t string) Window {
	pattern := make(map[string]int)
	arrs := strings.Split(t, "")
	for _, v := range arrs {
		if pattern[v] >= 0 {
			pattern[v]++
		} else {
			pattern[v] = 1
		}
	}
	return Window{start: -1, end: 0, collected: map[string]*[]int{}, startChar: "", pattern: pattern, minWindow: math.MaxInt32, matched: false}
}

func processNext(index int, char string, window *Window) {
	if window.pattern[char] <= 0 {
		return
	}
	if window.start == -1 {
		window.start = index
		window.startChar = char
	}
	if window.collected[char] == nil {
		window.collected[char] = &[]int{}
	}
	charCollected := append(*(window.collected[char]), index)
	window.collected[char] = &charCollected
	window.index = index
	if patternMatched(window) {
		calculateWindow(window)
	}
}

func calculateWindow(window *Window) {
	collected := window.collected
	start := math.MaxInt32
	end := math.MinInt32
	for char, v := range collected {
		if len(*v) > window.pattern[char] {
			newCollected := (*v)[len(*v)-window.pattern[char]:]
			window.collected[char] = &newCollected
		}
		for _, i := range *collected[char] {
			if (i) < start {
				start = i
				window.startChar = char
			}
			if (i) > end {
				end = i
				window.startChar = char
			}
		}
	}
	partternMatched := patternMatched(window)
	if partternMatched && end-start < window.minWindow {
		window.end = end
		window.start = start
		window.matched = partternMatched
		window.minWindow = window.end - window.start
	}

}

func patternMatched(window *Window) bool {
	partternMatched := true
	for i, v := range window.pattern {
		if window.collected[i] == nil || v > len(*window.collected[i]) {
			return false
		}
	}
	return partternMatched
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
