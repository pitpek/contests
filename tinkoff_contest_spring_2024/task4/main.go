package main

import "fmt"

func rotateMatrix(matrix [][]int, direction string) {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return
	}

	n := len(matrix)

	if direction == "L" {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				new_i := n - j - 1
				new_j := i
				if matrix[new_i][new_j] != matrix[i][j] {
					fmt.Printf("%d %d %d %d\n", i, j, new_i, new_j)
				}
			}
		}
	} else {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				new_i := j
				new_j := n - i - 1
				if matrix[new_i][new_j] != matrix[i][j] {
					fmt.Printf("%d %d %d %d\n", i, j, new_i, new_j)

				}
			}
		}
	}
}

func main() {
	var n int
	var side string
	fmt.Scan(&n, &side)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	rotateMatrix(matrix, side)
}
