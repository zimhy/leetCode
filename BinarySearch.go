package main

func binarySearch(nums []int, start int, end int, target int) (bool, int) {
	if start == end {
		if nums[start] == target {
			return true, start
		} else {
			return false, start
		}
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return true, mid
	} else if nums[mid] > target {
		return binarySearch(nums, start, mid, target)
	} else {
		return binarySearch(nums, mid+1, end, target)
	}
}
