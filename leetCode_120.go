package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, len(triangle), len(triangle))
	dp[len(triangle)-1] = make([]int, len(triangle[len(triangle)-1]))
	for i, v := range triangle[len(triangle)-1] {
		dp[len(triangle)-1][i] = v
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		dp[i] = make([]int, len(triangle[i]))
		for j, v := range triangle[i] {
			if dp[i+1][j] < dp[i+1][j+1] {
				dp[i][j] = dp[i+1][j] + v
			} else {
				dp[i][j] = dp[i+1][j+1] + v
			}
		}
	}
	return dp[0][0]
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
