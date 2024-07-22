package main

import (
	"container/heap"
	"fmt"
)

const MAX_N = 100
const INF = 1e9

var boardSize int
var distanceMatrix [MAX_N][MAX_N][2]int
var board [MAX_N]string

type Item struct {
	distance int
	typeIdx  int
	row      int
	col      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func inBoard(x, y int) bool {
	return x >= 0 && y >= 0 && x < boardSize && y < boardSize
}

func bfs(startX, startY int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			distanceMatrix[i][j][0] = INF
			distanceMatrix[i][j][1] = INF
		}
	}
	distanceMatrix[startX][startY][1] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{distance: 0, typeIdx: 1, row: startX, col: startY})

	dxG := []int{0, 0, 1, 1, 1, -1, -1, -1}
	dyG := []int{1, -1, 0, 1, -1, 0, 1, -1}
	dxK := []int{1, 1, 2, 2, -1, -1, -2, -2}
	dyK := []int{2, -2, 1, -1, 2, -2, 1, -1}

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(*Item)
		typeIdx := curr.typeIdx
		posX := curr.row
		posY := curr.col

		if typeIdx == 0 {
			for i := 0; i < len(dxG); i++ {
				newX := posX + dxG[i]
				newY := posY + dyG[i]

				if !inBoard(newX, newY) {
					continue
				}
				newType := 0
				if board[newX][newY] == 'K' {
					newType = 1
				}
				if distanceMatrix[newX][newY][newType] > distanceMatrix[posX][posY][typeIdx]+1 {
					distanceMatrix[newX][newY][newType] = distanceMatrix[posX][posY][typeIdx] + 1
					heap.Push(&pq, &Item{distance: distanceMatrix[newX][newY][newType], typeIdx: newType, row: newX, col: newY})
				}
			}
		} else {
			for i := 0; i < len(dxK); i++ {
				newX := posX + dxK[i]
				newY := posY + dyK[i]

				if !inBoard(newX, newY) {
					continue
				}
				newType := 1
				if board[newX][newY] == 'G' {
					newType = 0
				}

				if distanceMatrix[newX][newY][newType] > distanceMatrix[posX][posY][typeIdx]+1 {
					distanceMatrix[newX][newY][newType] = distanceMatrix[posX][posY][typeIdx] + 1
					heap.Push(&pq, &Item{distance: distanceMatrix[newX][newY][newType], typeIdx: newType, row: newX, col: newY})
				}
			}
		}
	}
}

func solve() {
	fmt.Scan(&boardSize)
	var startX, startY, endX, endY int
	for i := 0; i < boardSize; i++ {
		fmt.Scan(&board[i])
		for j := 0; j < boardSize; j++ {
			if board[i][j] == 'S' {
				startX = i
				startY = j
			}
			if board[i][j] == 'F' {
				endX = i
				endY = j
			}
		}
	}

	bfs(startX, startY)
	minDistance := min(distanceMatrix[endX][endY][0], distanceMatrix[endX][endY][1])
	if minDistance == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(minDistance)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	solve()
}
