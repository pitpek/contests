package main

import (
	"fmt"
)

var numRows int

func solve() {
	fmt.Scan(&numRows)
	dynamicProgramming := make([][]int, numRows)
	grid := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		dynamicProgramming[i] = make([]int, 3)
		grid[i] = make([]int, 3)

		for j := 0; j < 3; j++ {
			dynamicProgramming[i][j] = -1
		}

		var row string
		fmt.Scan(&row)
		for j := 0; j < 3; j++ {
			if row[j] == 'W' {
				grid[i][j] = -1
			} else if row[j] == '.' {
				grid[i][j] = 0
			} else if row[j] == 'C' {
				grid[i][j] = 1
			}
		}
	}

	for j := 0; j < 3; j++ {
		dynamicProgramming[0][j] = grid[0][j]
	}

	for i := 1; i < numRows; i++ {
		for j := 0; j < 3; j++ {
			for dj := -1; dj <= 1; dj++ {
				if j+dj >= 0 && j+dj < 3 && dynamicProgramming[i-1][j+dj] != -1 && grid[i][j] != -1 {
					dynamicProgramming[i][j] = max(dynamicProgramming[i][j], dynamicProgramming[i-1][j+dj]+grid[i][j])
				}
			}
		}
	}

	maxEaten := 0
	for i := 0; i < numRows; i++ {
		for j := 0; j < 3; j++ {
			maxEaten = max(maxEaten, dynamicProgramming[i][j])
		}
	}
	fmt.Println(maxEaten)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	solve()
}
