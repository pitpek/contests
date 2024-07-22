package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Scan(&n, &m)

	s := make([]int, m)
	t := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&s[i], &t[i])
	}

	dp0 := make([]int, n+1)
	dp1 := make([]int, n+1)

	for i := s[0] - 1; i >= 1; i-- {
		dp0[i] = s[0] - i
	}
	for i := t[0] + 1; i <= n; i++ {
		dp0[i] = i - t[0]
	}

	for i := 1; i < m; i++ {
		mnVal := math.MaxInt32
		monotonicStack := make([]int, 0)

		for j := s[i]; j <= t[i]; j++ {
			mnVal = min(mnVal, dp0[j])
			for len(monotonicStack) > 0 && dp0[monotonicStack[len(monotonicStack)-1]] > dp0[j] {
				monotonicStack = monotonicStack[:len(monotonicStack)-1]
			}
			monotonicStack = append(monotonicStack, j)
		}

		lenSegment := t[i] - s[i] + 1

		for j := t[i] + 1; j <= n; j++ {
			for len(monotonicStack) > 0 && monotonicStack[0] <= j-lenSegment {
				monotonicStack = monotonicStack[1:]
			}
			for len(monotonicStack) > 0 && dp0[monotonicStack[len(monotonicStack)-1]] > dp0[j] {
				monotonicStack = monotonicStack[:len(monotonicStack)-1]
			}
			monotonicStack = append(monotonicStack, j)
			dp1[j] = j - t[i] + dp0[monotonicStack[0]]
		}

		monotonicStack = monotonicStack[:0]

		for j := t[i]; j >= s[i]; j-- {
			for len(monotonicStack) > 0 && dp0[monotonicStack[len(monotonicStack)-1]] > dp0[j] {
				monotonicStack = monotonicStack[:len(monotonicStack)-1]
			}
			monotonicStack = append(monotonicStack, j)
		}

		for j := s[i] - 1; j >= 1; j-- {
			for len(monotonicStack) > 0 && monotonicStack[0] >= j+lenSegment {
				monotonicStack = monotonicStack[1:]
			}
			for len(monotonicStack) > 0 && dp0[monotonicStack[len(monotonicStack)-1]] > dp0[j] {
				monotonicStack = monotonicStack[:len(monotonicStack)-1]
			}
			monotonicStack = append(monotonicStack, j)
			dp1[j] = s[i] - j + dp0[monotonicStack[0]]
		}

		for j := s[i]; j <= t[i]; j++ {
			dp1[j] = mnVal
		}

		copy(dp0, dp1)
		for j := range dp1 {
			dp1[j] = 0
		}
	}

	ans := math.MaxInt32
	for i := 1; i <= n; i++ {
		ans = min(ans, dp0[i])
	}
	fmt.Println(ans)
}
