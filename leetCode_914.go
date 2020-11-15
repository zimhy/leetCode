package main

func hasGroupsSizeX(deck []int) bool {
	countMap := make(map[int]int)
	for _, v := range deck {
		if countMap[v] <= 0 {
			countMap[v] = 1
		} else {
			countMap[v] = countMap[v] + 1
		}
	}
	minCount := len(deck)
	for _, v := range countMap {
		if v < minCount {
			minCount = v
		}
	}
	if minCount >= 2 {
		for i := 2; i <= minCount; i++ {
			allMatched := true
			for _, v := range countMap {
				if v%i != 0 {
					allMatched = false
					break
				}
			}
			if allMatched {
				return true
			}
		}
	}
	return false
}
