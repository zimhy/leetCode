package main

import (
	"fmt"
)

/**
tips : 思路二分不一定是要把已经存在的所有数组二分，也可以是把预期二分
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//TODO
	//[1,2,3,4,5,6,7,8,9,10]
	//[2,7,8,9]

	if len(nums1) == 0 {
		return findMid(nums2)
	}
	if len(nums2) == 0 {
		return findMid(nums1)
	}
	if (len(nums1)+len(nums2))%2 == 0 {
		total := findTotalOrded(nums1, nums2, (len(nums1)+len(nums2))/2, 0, 0) +
			findTotalOrded(nums1, nums2, (len(nums1)+len(nums2))/2-1, 0, 0)
		return float64(total) / 2
	} else {
		mid := findTotalOrded(nums1, nums2, (len(nums1)+len(nums2))/2, 0, 0)
		return float64(mid)
	}
	return 0
}
func findMid(nums []int) float64 {
	if len(nums)%2 == 0 {
		return (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2-1])) / 2
	} else {
		return float64(nums[len(nums)/2])
	}
}
func findTotalOrded(nums1 []int, nums2 []int, order int, start1 int, start2 int) int {

	if start1 >= len(nums1) {
		return nums2[start2+order-1]
	}
	if start2 >= len(nums2) {
		return nums1[start1+order-1]
	}
	if order == 0 {
		if nums2[start2] > nums1[start1] {
			return nums1[start1]
		} else {
			return nums2[start2]
		}
	}
	if order == 1 {
		num1 := nums1[start1]
		num2 := nums2[start2]
		if num1 < num2 {
			if len(nums1)-1 > start1 {
				if nums1[start1+1] > num2 {
					return num2
				} else {
					return nums1[start1+1]
				}
			} else {
				return num2
			}

		} else {
			if len(nums2)-1 > start2 {
				if nums2[start2+1] > num1 {
					return num1
				} else {
					return nums2[start2+1]
				}
			} else {
				return num1
			}

		}
	}
	num1, num2 := 0, 0

	if len(nums1)-1 < order/2+start1 {
		num1 = int(^uint(0) >> 1)
	} else {
		num1 = nums1[order/2+start1]
	}
	if len(nums2)-1 < order/2+start2 {
		num2 = int(^uint(0) >> 1)
	} else {
		num2 = nums2[order/2+start2]
	}
	if num1 >= num2 {
		return findTotalOrded(nums1, nums2, order-order/2, start1, start2+order/2)
	} else {
		return findTotalOrded(nums1, nums2, order-order/2, start1+order/2, start2)
	}
}

func main() {
	result := findMedianSortedArrays([]int{1}, []int{2, 3, 4})
	fmt.Println(result)
}
