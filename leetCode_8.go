package main

import (
	"fmt"
	"math"
	"strings"
)

/***
tips :

*/
func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	splited := []rune(str)
	isPositive := true
	nums := make([]int, 0)
	result := int64(0)
	startNum := -1
	for i := 0; i < len(splited); i++ {
		if i == 0 {
			if !(splited[i] >= rune('0') && splited[i] <= rune('9')) && splited[i] != rune('-') && splited[i] != rune('+') {
				return 0
			} else if splited[i] == rune('-') {
				isPositive = false
			}
		}
		if splited[i] >= rune('0') && splited[i] <= rune('9') {
			nums = append(nums, int(splited[i]-rune('0')))
			if startNum == -1 {
				startNum = i
			}
		} else {
			if startNum > -1 {
				break
			}
		}

	}
	if len(nums) < 1 {
		return 0
	}
	if startNum > 1 {
		return 0
	}
	if startNum == 1 && splited[0] != rune('-') && splited[0] != rune('+') {
		return 0
	}
	if len(strings.TrimLeft(string([]rune(splited[startNum:len(nums)])), "0")) > 11 {
		if isPositive {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}

	}
	//fmt.Println(splited)
	for i := 0; i < len(nums); i++ {
		si := nums[i]
		posi := math.Pow10(len(nums) - i - 1)
		tmpnum := int64(si) * int64(posi)
		result += tmpnum
	}
	if !isPositive {
		result = -result
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	return int(result)
}

func main() {
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("20000000000000000000"))
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("123-"))
}
