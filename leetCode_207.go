package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {

	grapMap := make([][]int, numCourses)
	for i, _ := range grapMap {
		grapMap[i] = []int{}
	}
	for _, v := range prerequisites {
		grapMap[v[0]] = append(grapMap[v[0]], v[1])
	}
	canfinishSingle := make([]bool, numCourses)

	for i, v := range grapMap {
		if len(v) == 0 {
			canfinishSingle[i] = true
		}
	}
	for j := 0; j < len(canfinishSingle)-1; j++ {
		for i := 0; i < len(canfinishSingle); i++ {
			if len(grapMap[i]) > 0 && !canfinishSingle[i] {
				calResult := true
				for _, dependency := range grapMap[i] {
					if !canfinishSingle[dependency] {
						calResult = false
					}
				}
				canfinishSingle[i] = calResult
			}
		}
	}
	for _, v := range canfinishSingle {
		if !v {
			return false
		}
	}
	return true

}

func main() {
	result := canFinish(2, [][]int{{0, 1}})
	fmt.Println(result)
}
