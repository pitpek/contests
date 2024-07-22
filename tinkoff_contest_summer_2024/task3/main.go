package main

import (
	"fmt"
	"math"
)

var n, k, a, m int64

func lcg(e int64) int64 {
	return (a*e + 11) % m
}

func abs(x int64) int64 {
	return int64(math.Abs(float64(x)))
}

func main() {
	fmt.Scan(&n, &k, &a, &m)

	if n == 0 {
		fmt.Println(0)
		return
	}

	seed := int64(0)
	ans := 0
	paid := int64(0)

	for n > 0 {
		seed = lcg(seed)
		coin := (abs(seed%3-1)*5 + abs(seed%3)*2) % 8
		ans++
		paid += coin
		if paid >= 3*k {
			n -= paid / 3
			paid %= 3
		}
	}
	fmt.Println(ans)

}
