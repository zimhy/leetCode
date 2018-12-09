package main

import "fmt"

/**
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func lengthOfLongestSubstring(s string) int {
	//if len(s) == 0 {
	//	return 0
	//}

	chaMap := make(map[string]int)
	max := len(chaMap)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			pos, exist := chaMap[string(s[j])]
			if exist {
				if len(chaMap) > max {
					max = len(chaMap)
				}
				i = pos
				chaMap = make(map[string]int)
				break
			} else {
				chaMap[string(s[j])] = j
			}
		}
		if len(chaMap) > max {
			max = len(chaMap)
		}
	}
	return max
}
func main() {
	s := " "
	maxstr := lengthOfLongestSubstring(s)
	fmt.Println(maxstr)
}
