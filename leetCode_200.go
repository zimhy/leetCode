package main

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != byte('0') {
				count++
				searchIland(&grid, i, j)
			}
		}
	}
	return count

}

func searchIland(grid *[][]byte, i int, j int) {
	if i < 0 || j < 0 {
		return
	}
	if len(*grid) <= i || len((*grid)[i]) <= j {
		return
	}
	fmt.Println((*grid)[i][j])
	if (*grid)[i][j] == byte('0') {
		return
	}
	(*grid)[i][j] = byte('0')
	searchIland(grid, i+1, j)
	searchIland(grid, i-1, j)
	searchIland(grid, i, j+1)
	searchIland(grid, i, j-1)
}
func main() {
	numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}})
}
