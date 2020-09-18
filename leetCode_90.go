package main

import (
	"fmt"
	"sort"
)

/**
Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/
//func subsetsWithDup(nums []int) [][]int {
//	 numsSorted  := sortWithDup(nums)
//	 return getSubsets(numsSorted)
//
//}
//
//func getSubsets(numsSorted [][]int  )  [][]int{
//	if len(numsSorted) == 1 {
//		return getSamnumSubSet(numsSorted[0])
//	} else {
//		nextSubset := getSubsets(numsSorted[1:])
//
//		var containedSubSet [][]int
//		for i := 0; i < len(nextSubset); i++ {
//			subsetCopy := copyMatrix(nextSubset)
//			samenumSets := getSamnumSubSet(numsSorted[0])
//			for j := 0; j < len(samenumSets); j++ {
//				sameNums := samenumSets[j]
//				contained := append(subsetCopy[i], sameNums...)
//				containedSubSet = append(containedSubSet, contained)
//			}
//
//		}
//		return containedSubSet
//	}
//}
//
//func copyMatrix(nextSubset [][]int) [][]int {
//	var subsetCopy [][]int
//	for i := 0; i < len(nextSubset); i++ {
//		cpy := make([]int, len(nextSubset[i]))
//		copy(cpy, nextSubset[i])
//		subsetCopy = append(subsetCopy, cpy)
//	}
//	return subsetCopy
//}
//
//func getSamnumSubSet(sameNums []int  )  [][]int{
//	var numsSubset [][]int
//	for i:= 0 ;i<=len(sameNums);i++{
//		set := []int{}
//		num := sameNums[0]
//		for j:= 0 ;j<i;j++{
//			set = append(set,num)
//		}
//		numsSubset = append(numsSubset,set)
//	}
//	return numsSubset
//}
//
//
//func sortWithDup(nums []int) [][]int {
//	var numsSorted [][]int
//	sort.Ints(nums)
//	var sameNums =  &[]int{}
//	for i:=0;i<len(nums);i++{
//		*sameNums = append(*sameNums,nums[i])
//		if (i<len(nums)-1&&nums[i] !=nums[i+1])|| i == len(nums) -1 {
//			numsSorted = append(numsSorted,*sameNums)
//			sameNums = &[]int{}
//		}
//	}
//	return numsSorted
//
//}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	solve(nums, &res, nil)
	return res
}

func solve(nums []int, res *[][]int, cur []int) {
	*res = append(*res, append([]int{}, cur...))
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		solve(nums[i+1:], res, append(cur, nums[i]))
	}
}
