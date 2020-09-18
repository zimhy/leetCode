package main

import "sort"

type Points [][]int

func findMinArrowShots(points [][]int) int {
	var rightEdge []int
	sort.Sort(Points(points))
	shootCount := 0
	//overLay := 0
	for i := 0; i < len(points); i++ {
		if len(rightEdge) > 0 {
			sort.Ints(rightEdge)
			if points[i][0] > rightEdge[0] {
				rightEdge = []int{}
				shootCount++
			}
		}
		rightEdge = append(rightEdge, points[i][1])
	}
	if len(rightEdge) > 0 {
		shootCount++
	}
	return shootCount

}

func (s Points) Len() int           { return len(s) }
func (s Points) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Points) Less(i, j int) bool { return s[i][0] < s[j][0] }

func main() {
	test := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	findMinArrowShots(test)
}
