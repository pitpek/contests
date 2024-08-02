package main

import (
	"bufio"
	"fmt"
	"os"
)

// t  - количество наборов данных
// m  - длина последовательности
// a  - код дерева, ai элемент дерева
// u  - первое число номер вершины
// du - второе число количество её детей
// d  - номера сыновей в произвольном порядке

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	roots := []int{}

	for ti := 0; ti < t; ti++ {
		var n int
		fmt.Fscan(in, &n)

		// Карта для отслеживания всех вершин и их упоминаний как ребенка
		children := make(map[int]bool)
		vertices := make(map[int]bool)

		// Создаем слайс для хранения последовательности дерева
		sequence := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &sequence[i])
		}

		i := 0
		for i < n {
			node := sequence[i]
			childrenCount := sequence[i+1]
			vertices[node] = true

			for j := 0; j < childrenCount; j++ {
				child := sequence[i+2+j]
				children[child] = true
			}

			i += 2 + childrenCount
		}

		// Поиск вершины, которая не упоминается как ребенок
		for vertex := range vertices {
			if !children[vertex] {
				roots = append(roots, vertex)
				break
			}
		}
	}

	for _, root := range roots {
		fmt.Fprintf(out, "%v\n", root)
	}
}
