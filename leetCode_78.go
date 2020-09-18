package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{}, {nums[0]}}
	} else {
		nextSubset := subsets(nums[:len(nums)-1])
		var subsetCopy [][]int
		for i := 0; i < len(nextSubset); i++ {
			cpy := make([]int, len(nextSubset[i]))
			copy(cpy, nextSubset[i])
			subsetCopy = append(subsetCopy, cpy)
		}
		//for i:= 0 ;i<len(nextSubset);i++ {
		//	tmp := []int(nextSubset[i])
		//	containedSubSet = append(containedSubSet, nextSubset...)
		//}
		var containedSubSet [][]int
		for i := 0; i < len(subsetCopy); i++ {
			contained := append(subsetCopy[i], nums[len(nums)-1])
			containedSubSet = append(containedSubSet, contained)
		}
		containedSubSet = append(nextSubset, containedSubSet...)
		//for i:= 0 ;i<len(containedSubSet);i++ {
		//	sort.Ints(containedSubSet[i])
		//}
		return containedSubSet
	}
}

func main() {
	fmt.Println(subsets([]int{
		9, 0, 3, 5, 7}))
}
