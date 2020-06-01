package main

import "fmt"


//func canJump(nums []int) bool {
//	
//	if nums == nil || len(nums) <= 1 {
//		return true;
//	}
//
//	return findRout(nums)
//
//}


func canJump(nums []int) bool {
	if nums == nil || len(nums) <= 1 {
		return true;
	}
	maxStep := make([]int, len(nums))

	for i:= 0;i<len(nums);i++{
		if i== 0 {
			maxStep[i] = nums[i]
			continue
		}
		if nums[i] == 0 {
			maxStep[i] = maxStep[i-1]
			continue
		}
		if maxStep[i-1]<i {
			return  false
		}
		if(maxStep[i-1]>i+nums[i]){
			maxStep[i] = maxStep[i-1]
		}else {
			maxStep[i] = i+nums[i]
		}
	}
	return maxStep[len(nums) -1] +1>=len(nums)
	

}
//
//func findRout(nums []int) bool {
//	if nums == nil || len(nums) <= 1 {
//		return true;
//	}
//	if len(nums) <= nums[0] {
//		return true
//	}
//	if(nums[0] == 0){
//		return false
//	}
//	for i := nums[0]; i >=1; i-- {
//		if findRout(nums[i:]) {
//			return true
//		}
//
//	}
//	return false
//}

func main() {
	b := []int{2,0}
	fmt.Println("%v", canJump(b))

}
