package main

import (
	"bufio"
	"fmt"
	"os"
)

// // t - количество наборов данных
// // n - количество запросов
// // a - id клиента

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ti := 0; ti < t; ti++ {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		if n == 1 {
			fmt.Fprintln(out, 1)
			continue
		}

		maxLen := 0

		for i := 0; i < n; i++ {
			clientCounts := make(map[int]int)
			count := 0

			for j := i; j < n; j++ {
				clientCounts[a[j]]++
				fmt.Println(clientCounts)
				if len(clientCounts) > 2 {
					break
				}
				count++
			}

			if count > maxLen {
				maxLen = count
			}
		}

		fmt.Fprintln(out, maxLen)
	}
}
