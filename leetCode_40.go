package main

import (
	"fmt"
	"sort"
	"strings"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	resolve(candidates, target, nil, &res)
	return unique(res)

}

func resolve(candidates []int, target int, cur []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, cur)
	}

	if len(candidates) == 0 {
		return
	}
	if candidates[0] > target {
		return
	}
	resolve(candidates[1:], target-candidates[0], append(cur, candidates[0]), res)
	resolve(candidates[1:], target, cur, res)

}

func unique(intMatrix [][]int) [][]int {
	keys := make(map[string]bool)
	list := [][]int{}
	for _, line := range intMatrix {
		if _, value := keys[stringValOfArray(line)]; !value {
			keys[stringValOfArray(line)] = true
			list = append(list, line)
		}
	}
	return list
}

func stringValOfArray(intSlice []int) string {
	sort.Ints(intSlice)
	return strings.Join(strings.Fields(fmt.Sprint(intSlice)), ",")

}
