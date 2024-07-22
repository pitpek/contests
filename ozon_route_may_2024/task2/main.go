package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, t int
	fmt.Fscanln(in, &n, &t)

	var chars []string
	for i := 0; i < n; i++ {
		var char string
		fmt.Fscan(in, &char)
		chars = append(chars, char)
	}

	totalSum := TotalSUM(chars)

	var pins []string
	for i := 0; i < t; i++ {
		var pin string
		fmt.Fscan(in, &pin)
		pins = append(pins, pin)
	}

	for _, pin := range pins {
		var sum rune
		for _, char := range pin {
			sum += char
		}
		if sum == totalSum {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func TotalSUM(chars []string) rune {
	var sum rune
	for _, str := range chars {
		for _, char := range str {
			sum += char
		}
	}
	return sum
}
