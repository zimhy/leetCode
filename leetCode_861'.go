package main

import "fmt"

/**
输入：[[0,0,1,1],[1,0,1,0],[1,1,0,0]]
输出：39
解释：
转换为 [[1,1,1,1],[1,0,0,1],[1,1,1,1]]
0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/score-after-flipping-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func matrixScore(A [][]int) int {
	m := len(A)
	n := len(A[0])
	for i := 0; i < n; i++ {
		flags := getFlaLines(A, i)
		if i == 0 {
			for _, v := range flags {
				revertOneLine(A, v)
			}
		} else {
			if len(flags) > (m / 2) {
				reversOneRow(A, i)
			}
		}
	}
	result := 0
	for _, v := range A {
		result += calculateOneLine(v, n)
	}
	return result

}

func calculateOneLine(line []int, n int) int {
	value := 0
	for i, v := range line {
		shift := (uint)(n - 1 - i)
		value += v << shift
	}
	return value
}
func processOneRow(A [][]int, rowNum int) {
	flagLines := getFlaLines(A, rowNum)
	if len(flagLines) == len(A) {
		return
	}
	if len(flagLines) == 0 {
		reversOneRow(A, rowNum)
		return
	}
	for _, v := range flagLines {
		revertOneLine(A, v)
	}
	reversOneRow(A, rowNum)
	for _, v := range flagLines {
		revertOneLine(A, v)
	}
	return
}

func getFlaLines(A [][]int, rowNum int) []int {
	flagLines := []int{}
	for i := 0; i < len(A); i++ {
		if A[i][rowNum] == 0 {
			flagLines = append(flagLines, i)
		}
	}
	return flagLines
}

func revertOneLine(A [][]int, lineNum int) {
	for k, v := range A[lineNum] {
		A[lineNum][k] = v ^ 1
		//if(v == 1){
		//	A[lineNum][k] = 0
		//}else {
		//	A[lineNum][k] = 1
		//}
	}
}

func reversOneRow(A [][]int, rowNum int) {
	for i := 0; i < len(A); i++ {
		A[i][rowNum] = A[i][rowNum] ^ 1
		//if A[i][rowNum] == 1 {
		//	A[i][rowNum] = 0
		//}else {
		//	A[i][rowNum] = 1
		//}
	}
}

func main() {
	fmt.Println(matrixScore([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}))
}
