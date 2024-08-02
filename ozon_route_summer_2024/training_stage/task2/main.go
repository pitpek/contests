package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// t - количество наборов данных
	// n - количество проданных товаров
	// p - процент комиссии
	// a - стоимость проданных товаров
	var t int
	fmt.Fscan(in, &t)

	for ti := 0; ti < t; ti++ {
		var n, p int
		fmt.Fscan(in, &n, &p)

		var lostAmount float64

		for i := 0; i < n; i++ {
			var ai int
			fmt.Fscan(in, &ai)

			// Правильная комиссия
			correctCommission := math.Floor(float64(ai)*float64(p)) / 100

			// Ошибочная комиссия
			incorrectCommission := float64(ai * p / 100)

			// Потерянная сумма
			lostAmount += correctCommission - incorrectCommission
		}

		// Выводим потерянную сумму
		fmt.Fprintf(out, "%.2f\n", lostAmount)
	}
}
