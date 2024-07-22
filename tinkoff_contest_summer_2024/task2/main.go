package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	DAYS_IN_WEEK  = 7
	DAYS_IN_MONTH = 28
)

var dayToIndex = map[string]int{
	"MON": 0, "TUE": 1, "WED": 2,
	"THU": 3, "FRI": 4, "SAT": 5, "SUN": 6,
}

func getDayIndex(day string) int {
	return dayToIndex[day]
}

func main() {
	occupied := make([]bool, DAYS_IN_MONTH)
	scanner := bufio.NewScanner(os.Stdin)
	dayNumber := 1

	for week := 0; week < 4; week++ {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line != "" {
			parts := strings.Split(line, " ")
			for _, day := range parts {
				index := getDayIndex(day)
				occupied[dayNumber+index-1] = true
			}
		}
		dayNumber += DAYS_IN_WEEK
	}

	maxStart, maxEnd, maxLength := 0, 0, 0
	currentStart, currentLength := 0, 0

	for i := 0; i < DAYS_IN_MONTH; i++ {
		if !occupied[i] {
			if currentLength == 0 {
				currentStart = i + 1
			}
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
				maxStart = currentStart
				maxEnd = i + 1
			}
		} else {
			currentLength = 0
		}
	}

	if maxLength == 0 {
		fmt.Println("0 0")
	} else {
		fmt.Println(maxStart, maxEnd)
	}
}
