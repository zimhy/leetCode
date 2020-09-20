package main

import (
	"fmt"
	"math"
)

/**
Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie.
Each child i has a greed factor gi,
which is the minimum size of a cookie that the child will be content with;
and each cookie j has a size sj. If sj >= gi, we can assign the cookie j to the child i,
and the child i will be content.
Your goal is to maximize the number of your content children and output the maximum number.

Note:
You may assume the greed factor is always positive.
You cannot assign more than one cookie to one child.

Example 1:
Input: [1,2,3], [1,1]

Output: 1

Explanation: You have 3 children and 2 cookies. The greed factors of 3 children are 1, 2, 3.
And even though you have 2 cookies, since their size is both 1, you could only make the child whose greed factor is 1 content.
You need to output 1.
Example 2:
Input: [1,2], [1,2,3]

Output: 2

Explanation: You have 2 children and 3 cookies. The greed factors of 2 children are 1, 2.
You have 3 cookies and their sizes are big enough to gratify all of the children,
You need to output 2.
*/
func findContentChildren(g []int, s []int) int {
	count := 0

	for i := 0; i < len(s); i++ {
		success := false
		success, g = dispatchOne(g, s[i])
		if success {
			count++
		}
	}
	return count
}

func dispatchOne(g []int, s int) (bool, []int) {
	childIndex := -1
	contentOverValue := math.MaxInt32
	for i := 0; i < len(g); i++ {
		if s-(g)[i] >= 0 && s-g[i] < contentOverValue {
			childIndex = i
			contentOverValue = s - g[i]
		}
	}
	if childIndex >= 0 {
		g = remove(g, childIndex)
		return true, g
	}
	return false, g

}

func remove(slices []int, index int) []int {
	return append(slices[:index], slices[index+1:]...)
}

func main() {

	var result = findContentChildren([]int{1, 2}, []int{1, 2, 3})
	fmt.Println("%v", result)
}
