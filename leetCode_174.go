package main

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	length := len(dungeon)
	wide := len(dungeon[0])
	if length == 1 && wide == 1 {
		if dungeon[0][0] > 0 {
			return 1
		} else {
			return 1 - dungeon[0][0]
		}
		return dungeon[0][0]
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, wide)
	}
	for i := length - 1; i >= 0; i-- {
		if i == length-1 {
			if dungeon[i][wide-1] >= 0 {
				dp[i][wide-1] = 1
			} else {
				dp[i][wide-1] = 1 - dungeon[i][wide-1]
			}
		} else {
			if dp[i+1][wide-1]-dungeon[i][wide-1] < 1 {
				dp[i][wide-1] = 1
			} else {
				dp[i][wide-1] = dp[i+1][wide-1] - dungeon[i][wide-1]
			}
		}
	}
	for i := wide - 1; i >= 0; i-- {
		if i == wide-1 {
			if dungeon[length-1][i] >= 0 {
				dp[length-1][i] = 1
			} else {
				dp[length-1][i] = 1 - dungeon[length-1][i]
			}
		} else {
			if dp[length-1][i+1]-dungeon[length-1][i] < 1 {
				dp[length-1][i] = 1
			} else {
				dp[length-1][i] = dp[length-1][i+1] - dungeon[length-1][i]
			}
		}
	}
	if length > 1 && wide > 1 {
		for i := length - 2; i >= 0; i-- {
			for j := wide - 2; j >= 0; j-- {
				step := 1
				if dp[i+1][j] > dp[i][j+1] {
					step = dp[i][j+1] - dungeon[i][j]
				} else {
					step = dp[i+1][j] - dungeon[i][j]
				}
				if step > 1 {
					dp[i][j] = step
				} else {
					dp[i][j] = 1
				}
			}
		}
	}
	return dp[0][0]
}

func main() {
	fmt.Println(calculateMinimumHP([][]int{{1, -2, 3}, {2, -2, -2}}))
	//fmt.Println(calculateMinimumHP([][]int{{0,-3}}))
	//fmt.Println(calculateMinimumHP([][]int{{1},{2}}))

}
