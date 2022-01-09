package main

import "fmt"

func main() {

	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	result := countNegatives(grid)
	fmt.Println(result)
}

func countNegatives(grid [][]int) int {

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] < 0 {
				count++
			}
		}
	}

	return count
}
