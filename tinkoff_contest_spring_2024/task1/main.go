package main

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scan(&n)

	examMarks := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&examMarks[i])
	}

	maxConsecutiveFives := -1
	for i := 6; i < n; i++ {
		hasLowMarks := false
		countFives := 0
		for j := 0; j < 7; j++ {
			if examMarks[i-j] <= 3 {
				hasLowMarks = true
				break
			}
			if examMarks[i-j] == 5 {
				countFives++
			}
		}
		if !hasLowMarks {
			if countFives > maxConsecutiveFives {
				maxConsecutiveFives = countFives
			}
		}
	}
	fmt.Println(maxConsecutiveFives)
}

func main() {
	solve()
}
