package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Чтение количества элементов
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	total := 0

	// Чтение элементов и суммирование
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		total += x
	}

	if total%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
