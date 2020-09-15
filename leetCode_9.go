package main

import "fmt"

func isPalindrome(x int) bool {
	splited := make([]int, 0)
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	for {
		if x > 0 {
			splited = append(splited, x%10)
			x = x / 10
		} else {
			break
		}
	}
	for i := 0; i <= len(splited)/2; i++ {
		if splited[i] == splited[len(splited)-1-i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(0))
}
