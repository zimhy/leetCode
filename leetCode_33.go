package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	rotateIndex := getRotateIndex(nums, 0, len(nums)-1)
	if rotateIndex == -1 || rotateIndex == 0 || rotateIndex == -2 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}
	if nums[len(nums)-1] >= target {
		return binarySearch(nums, rotateIndex, len(nums)-1, target)
	} else {
		return binarySearch(nums, 0, rotateIndex-1, target)
	}

}

func getRotateIndex(nums []int, start int, end int) int {
	if (end - start) < 2 {
		if nums[start] <= nums[end] {
			return start
		} else {
			return end
		}
	}
	mid := (start + end) / 2
	if nums[mid] == nums[end] {
		if nums[start] > nums[mid] {
			getRotateIndex(nums, start, mid)
		} else if nums[mid] == nums[start] {
			return -2
		} else if nums[mid] > nums[start] {
			getRotateIndex(nums, mid, end)
		}
	}
	if nums[mid] < nums[end] {
		return getRotateIndex(nums, start, mid)
	} else {
		return getRotateIndex(nums, mid, end)
	}

}

func main() {
	fmt.Println(search([]int{1, 3}, 2))
}
