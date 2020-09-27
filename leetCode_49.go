package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	mappingList := make(map[string]*(map[string]int))
	for _, v := range strs {
		processInsert(v, mappingList)
	}
	res := make([][]string, len(mappingList))
	i := 0
	for _, v := range mappingList {
		tempResult := make([]string, 0)
		for value, count := range *v {
			for j := 0; j < count; j++ {
				tempResult = append(tempResult, value)
			}
		}
		res[i] = tempResult
		i++
	}
	return res

}

func processInsert(s string, mappingList map[string]*map[string]int) {
	sorted := getSorted(s)
	var grouped map[string]int
	if mappingList[sorted] == nil {
		grouped = make(map[string]int)
	} else {
		grouped = *(mappingList[sorted])
	}

	if grouped == nil {

	}
	if grouped[s] >= 0 {
		grouped[s]++
	} else {
		grouped[s] = 1
	}
	mappingList[sorted] = &grouped

}

func getSorted(s string) string {
	splited := strings.Split(s, "")
	sort.Strings(splited)
	return strings.Join(splited, "")
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
