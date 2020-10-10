package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	includeds := make([]int, len(nums))
	includeds[0] = 1

	for i, _ := range nums {
		if (i) == 0 {
			includeds[i] = 1
		} else {
			includeds[i] = 1
			for j := 0; j < i; j++ {
				if nums[j] < nums[i] && includeds[j]+1 > includeds[i] {
					includeds[i] = includeds[j] + 1
				}
			}
		}
	}
	max := includeds[0]
	for _, v := range includeds {
		if (v) > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLIS([]int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}))
}
