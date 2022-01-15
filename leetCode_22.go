package main

import "fmt"

/**
示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]


提示：

1 <= n <= 8
用栈来解是不是更好
https://leetcode-cn.com/problems/generate-parentheses/solution/zui-jian-dan-yi-dong-de-dong-tai-gui-hua-bu-lun-da/
动归的最好思路

状态划分的时候思路: 分割or 不分割
而不是显而易见的左中右
 * }
*/
func generateParenthesis(n int) []string {
	var totalResult = generateParenthesisMap(n)
	var mapResult = totalResult[n]
	result := []string{}
	for one, _ := range mapResult {
		result = append(result, one)
	}
	return result
}

func generateParenthesisMap(n int) map[int]map[string]bool {
	if n == 1 {
		var result1 = map[string]bool{}
		var result0 = map[string]bool{}
		result1["()"] = true
		result0[""] = true
		var result = map[int]map[string]bool{}
		result[1] = result1
		result[0] = result0
		return result
	}
	var preData = generateParenthesisMap(n - 1)
	var resultn = map[string]bool{}
	for i := 0; i < n && i >= 0; i++ {
		lefts := preData[i]
		rights := preData[n-i-1]
		for left, _ := range lefts {
			for right, _ := range rights {
				resultn["("+left+")"+right] = true
				resultn["("+left+right+")"] = true
			}
		}
	}
	preData[n] = resultn
	return preData
}

func main() {

	fmt.Println(generateParenthesis(3))
}
