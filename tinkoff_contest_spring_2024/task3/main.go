package main

import (
	"fmt"
	"sort"
	"strings"
)

func split(str string, delimiter rune) []string {
	var tokens []string
	var token strings.Builder
	for _, c := range str {
		if c == delimiter {
			tokens = append(tokens, token.String())
			token.Reset()
		} else {
			token.WriteRune(c)
		}
	}
	tokens = append(tokens, token.String())
	return tokens
}

func solve() {
	var numDirs int
	fmt.Scan(&numDirs)
	directoryPaths := make([][]string, numDirs)
	var rootName string
	for i := 0; i < numDirs; i++ {
		var dirPath string
		fmt.Scan(&dirPath)
		directoryPaths[i] = split(dirPath, '/')
		rootName = directoryPaths[i][0]
		directoryPaths[i][0] = "@"
	}
	sort.Slice(directoryPaths, func(i, j int) bool {
		// Сортируем директории по их глубине.
		lenI, lenJ := len(directoryPaths[i]), len(directoryPaths[j])
		minLen := lenI
		if lenJ < minLen {
			minLen = lenJ
		}
		for k := 0; k < minLen; k++ {
			if directoryPaths[i][k] != directoryPaths[j][k] {
				return directoryPaths[i][k] < directoryPaths[j][k]
			}
		}
		return lenI < lenJ

	})
	for _, dir := range directoryPaths {
		for j := 0; j < len(dir)-1; j++ {
			fmt.Print("  ")
		}
		currentDir := dir[len(dir)-1]
		if currentDir == "@" {
			fmt.Println(rootName)
		} else {
			fmt.Println(currentDir)
		}
	}
}

func main() {
	solve()
}
