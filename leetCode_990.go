package main

import (
	"fmt"
	"strings"
)

func equationsPossible(equations []string) bool {

	setMap := make(map[string]*map[string]bool)
	for _, v := range equations {
		params := strings.Split(v, "")
		A := params[0]
		B := params[3]
		if params[1] == "=" {
			result := mergeSet(setMap[A], setMap[B])
			(*result)[A] = true
			(*result)[B] = true
			setMap[A] = result
			setMap[B] = result
		}
	}
	for _, v := range equations {
		params := strings.Split(v, "")
		A := params[0]
		B := params[3]
		if params[1] == "!" {
			if A == B {
				return false
			}
			resultA := setMap[A]
			if resultA != nil && (*resultA)[B] {
				return false
			}
			resultB := setMap[B]
			if resultB != nil && (*resultB)[A] {
				return false
			}

		}
	}
	return true

}

func mergeSet(A *map[string]bool, B *map[string]bool) *map[string]bool {
	if B != nil && A != nil {
		for k, v := range *B {
			(*A)[k] = v
		}
		return A
	} else if A != nil {
		return A
	} else if B != nil {
		return B
	} else {
		result := make(map[string]bool)
		return &result
	}

}

func main() {
	fmt.Print(equationsPossible([]string{"c==c", "b==d", "x!=z"}))
}
