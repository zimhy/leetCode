package main

func searchInsert(nums []int, target int) int {
	success, index := binarySearch(nums, 0, len(nums)-1, target)
	if success {
		return index
	} else {
		if nums[index] > target {
			return index
		} else {
			return index + 1
		}
	}
}
