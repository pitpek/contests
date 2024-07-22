package main

import "fmt"

func rotateImage(originalMatrix [][]int64) [][]int64 {
	rows := len(originalMatrix)
	cols := len(originalMatrix[0])

	rotatedMatrix := make([][]int64, cols)
	for i := range rotatedMatrix {
		rotatedMatrix[i] = make([]int64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotatedMatrix[j][rows-1-i] = originalMatrix[i][j]
		}
	}

	return rotatedMatrix
}

func solve() {
	var numRows, numCols int
	fmt.Scan(&numRows, &numCols)

	inputMatrix := make([][]int64, numRows)
	for i := range inputMatrix {
		inputMatrix[i] = make([]int64, numCols)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			fmt.Scan(&inputMatrix[i][j])
		}
	}

	rotatedMatrix := rotateImage(inputMatrix)

	for j := 0; j < numCols; j++ {
		for i := 0; i < numRows; i++ {
			fmt.Print(rotatedMatrix[j][i], " ")
		}
		fmt.Println()
	}
}

func main() {
	solve()
}
