package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	result :=  removeKdigitsZeroAware(num, k, false)
	nums := []rune(result)
	zeroHead := 0
	for i:=0;i<len(nums)-1 ; i++ {
		if nums[i] > '0' {
			break
		}
		zeroHead++
	}

	result = string(nums[zeroHead:])
	return result
}

func removeKdigitsZeroAware(num string, k int, firstNotZero bool) string {
	if (k == 0) {
		return num
	}
	if len(num) <= k{
		return "0"
	}


	nums := []rune(num)



	minIndex := findTheMin(nums[0:k+1])
	if  num[minIndex] == '0' && firstNotZero {
		minIndex = findTheMin(nums[0:minIndex])

	}
	k=k-minIndex ;
	if k == len(nums) - (minIndex+1){
		return string(num[minIndex:minIndex+1])
	}else {
		return string(num[minIndex:minIndex+1]) + removeKdigitsZeroAware(string(num[minIndex+1:]), k, false)
	}


}

func findTheMin(num []rune) int {

	min := '9' + 1
	minIndex := -1
	for i := 0; i < len(num); i++ {
		if num[i] < min {
			min = num[i];
			minIndex = i

		}

	}

	return minIndex
}

func main() {

	var result = removeKdigits("1432219", 3)
	fmt.Println("%v", result)
}
