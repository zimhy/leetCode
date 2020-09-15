package main

import "fmt"

func isMatch(s string, p string) bool {
	S := []rune(s)
	P := []rune(p)
	return arrayRuneIsMatch(S, P)
	//return true
}

/**
S: "a"
P: "a*"
*/
func arrayRuneIsMatch(S []rune, P []rune) bool {
	var dp [][]bool
	for i := 0; i <= len(S); i++ {
		line := make([]bool, 0)
		for j := 0; j <= len(P); j++ {
			line = append(line, false)
		}
		dp = append(dp, line)
	}
	dp[0][0] = true
	for i := 1; i <= len(P); i++ {
		if P[i-1] == '*' && dp[0][i-2] {
			dp[0][i] = true
		}
	}
	for i := 1; i <= len(S); i++ {
		for j := 1; j <= len(P); j++ {
			if P[j-1] == '.' || P[j-1] == S[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if P[j-1] == '*' {
				if P[j-2] != S[i-1] && P[j-2] != '.' {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j]
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[len(S)][len(P)]
}

func main() {
	fmt.Println(isMatch("aab", "c*a*b*"))
}
