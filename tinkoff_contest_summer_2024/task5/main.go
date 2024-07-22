package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	edges   [][]int
	color   []int
	topSort []int
	has     map[pair]bool
)

type pair struct {
	u, v int
}

func dfs(v int) bool {
	color[v] = 1
	for _, to := range edges[v] {
		if color[to] == 1 {
			return true
		}
		if color[to] == 0 {
			if dfs(to) {
				return true
			}
		}
	}
	topSort = append(topSort, v)
	color[v] = 2
	return false
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	edges = make([][]int, n+1)
	color = make([]int, n+1)
	has = make(map[pair]bool)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < m; i++ {
		var u, v int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &u)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &v)
		if has[pair{u, v}] {
			continue
		}
		has[pair{u, v}] = true
		edges[u] = append(edges[u], v)
	}

	for v := 1; v <= n; v++ {
		if color[v] == 0 {
			if dfs(v) {
				fmt.Println("No")
				return
			}
		}
	}
	sort.Ints(topSort)

	fmt.Println("Yes")
	for _, v := range topSort {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
