package main

import "strings"

func findRepeatedDnaSequences(s string) []string {
	totalLong := len(s)
	if totalLong <= 10 {
		return []string{}
	}
	counts := make(map[string]int)
	chars := strings.Split(s, "")
	for i := 0; i <= totalLong-10; i++ {
		sub := strings.Join(chars[i:i+10], "")
		if counts[sub] >= 0 {
			counts[sub]++
		} else {
			counts[sub] = 1
		}
	}
	res := []string{}
	for v, count := range counts {
		if count > 1 {
			res = append(res, v)
		}
	}
	return res
}
