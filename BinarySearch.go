package main

func binarySearch(nums []int, start int, end int, target int) int {
	if start == end {
		if nums[start] == target {
			return start
		} else {
			return -1
		}
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binarySearch(nums, start, mid, target)
	} else {
		return binarySearch(nums, mid+1, end, target)
	}
}
