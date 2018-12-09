package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//TODO
	//[1,2,3,4,5,6,7,8,9,10]
	//[2,7,8,9]
	if (len(nums1)+len(nums2))%2 == 0 {
		total := findTotalOrded(nums1, nums2, (len(nums1)+len(nums2))/2) +
			findTotalOrded(nums1, nums2, (len(nums1)+len(nums2))/2+1)
		return float64(total) / 2
	} else {
		mid := findTotalOrded(nums1, nums2, (len(nums1)+len(nums2))/2+1)
		return float64(mid)
	}
	return 0
}

//[1,2,2,3,4,5,6,7,7,8,8,9,9,10]
//order 从0 开始
func findTotalOrded(nums1 []int, nums2 []int, order int) int {
	if len(nums1) == 0 {
		return nums2[order]
	}
	if len(nums2) == 0 {
		return nums2[order]
	}

	mid1 := nums1[len(nums1)/2]
	mid2 := nums2[len(nums2)/2]
	if mid1 == mid2 {
		if order == len(nums2)/2+len(nums1)/2 {
			return mid1
		} else if order > len(nums2)/2+len(nums1)/2 {
			return findTotalOrded(nums1[len(nums1)/2:], nums2[len(nums2)/2:], (order - len(nums1)/2 - len(nums2)/2))
		} else {
			return findTotalOrded(nums1[0:len(nums1)/2], nums2[0:len(nums2)/2], (order - len(nums1)/2 - len(nums2)/2))
		}
	} else {
		if mid1 > mid2 {
			mid := mid2
			mid2 = mid1
			mid1 = mid
			nums := nums1
			nums1 = nums2
			nums2 = nums
		}
		mid1InN2Less := findIndex(nums2, mid1, 0, len(nums2)-1)
		mid2InN1Less := findIndex(nums1, mid2, 0, len(nums1)-1)

		mid1AllIndex := mid1InN2Less + len(nums1)/2
		mid2AllIndex := mid2InN1Less + len(nums2)/2

		if order > mid2AllIndex {
			return findTotalOrded(nums1[mid2InN1Less:], nums2[len(nums2)/2:], (order - mid2InN1Less - len(nums2)))
		} else if order < mid1AllIndex {
			return findTotalOrded(nums1[:len(nums1)/2], nums2[:mid1InN2Less], order)
		}
		if order > mid1AllIndex && order < mid2AllIndex {
			return findTotalOrded(nums1[len(nums1)/2:], nums2[mid1InN2Less:], order-len(nums1)/2)
		}

	}
	return mid1
}

func findIndex(nums []int, num int, start int, end int) (lessIndex int) {
	mid := (start + end) / 2
	if nums[mid] == num {
		return mid
	} else if nums[mid] > num {
		if mid == 0 {
			return 0
		} else if nums[mid-1] < num {
			return mid
		} else {
			return findIndex(nums, num, start, mid-1)
		}
	} else if nums[mid] < num {
		if mid == len(nums)-1 {
			return len(nums)
		} else if nums[mid+1] > num {
			return mid + 1
		} else {
			return findIndex(nums, num, mid+1, end)
		}
	}
	return mid

}
func main() {
	findMedianSortedArrays([]int{3, 4, 4}, []int{3, 4, 4})
}
