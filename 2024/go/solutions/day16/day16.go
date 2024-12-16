package day16

import (
	"andrewsaputra/adventofcode2024/helper"
	"container/heap"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day16.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day16.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	matrix := helper.ToMatrix(path)
	startRow, startCol := findStartPos(matrix)
	minScore := make(map[string]int)
	return travelDjikstra(matrix, minScore, startRow, startCol)
}

func solvePart2(path string) int {
	matrix := helper.ToMatrix(path)
	pathMatrix := helper.ToMatrix(path)
	startRow, startCol := findStartPos(matrix)
	minScore := make(map[string]int)
	travelDjikstra(matrix, minScore, startRow, startCol)
	travelDFS(matrix, pathMatrix, minScore, &Item{Row: startRow, Col: startCol, Score: 0, Dir: EAST})

	var result int
	for row := range pathMatrix {
		for _, val := range pathMatrix[row] {
			if val == 'O' {
				result++
			}
		}
	}

	return result
}

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type Item struct {
	Row   int
	Col   int
	Score int
	Dir   Direction
	Prev  *Item
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Score < pq[j].Score }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(v any) { *pq = append(*pq, v.(*Item)) }
func (pq *PriorityQueue) Pop() any {
	idx := len(*pq) - 1
	item := (*pq)[idx]
	(*pq)[idx] = nil
	*pq = (*pq)[:idx]

	return item
}

func findStartPos(matrix [][]byte) (int, int) {
	for row := range matrix {
		for col, val := range matrix[row] {
			if val == 'S' {
				return row, col
			}
		}
	}
	return -1, -1
}

func cacheKey(item Item) string {
	return fmt.Sprintf("%d-%d-%d", item.Row, item.Col, item.Dir)
}

func travelDjikstra(matrix [][]byte, minScore map[string]int, startRow int, startCol int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Row: startRow, Col: startCol, Dir: EAST})

	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)

		row, col := item.Row, item.Col
		if matrix[row][col] == 'E' {
			return item.Score
		}

		for i := range drow {
			nextRow := row + drow[i]
			nextCol := col + dcol[i]
			if matrix[nextRow][nextCol] == '#' {
				continue
			}

			nextDir := Direction(i)
			nextScore := item.Score + 1

			if nextDir != item.Dir {
				nextScore += 1000
			}

			nextItem := &Item{Row: nextRow, Col: nextCol, Score: nextScore, Dir: nextDir, Prev: item}
			key := cacheKey(*nextItem)
			if val, ok := minScore[key]; !ok || nextScore < val {
				minScore[key] = nextScore
				heap.Push(pq, nextItem)
			}
		}
	}

	return -1
}

func travelDFS(matrix [][]byte, pathMatrix [][]byte, minScore map[string]int, item *Item) {
	if matrix[item.Row][item.Col] == 'E' {
		markPath(item, pathMatrix)
		return
	}

	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}

	for i := range drow {
		nextRow := item.Row + drow[i]
		nextCol := item.Col + dcol[i]
		if matrix[nextRow][nextCol] == '#' {
			continue
		}

		nextDir := Direction(i)
		nextScore := item.Score + 1
		if nextDir != item.Dir {
			nextScore += 1000
		}

		nextItem := &Item{Row: nextRow, Col: nextCol, Score: nextScore, Dir: nextDir, Prev: item}
		key := cacheKey(*nextItem)
		if val, ok := minScore[key]; ok && nextScore <= val {
			travelDFS(matrix, pathMatrix, minScore, nextItem)
		}
	}
}

func markPath(item *Item, pathMatrix [][]byte) {
	curr := item
	for curr != nil {
		pathMatrix[curr.Row][curr.Col] = 'O'
		curr = curr.Prev
	}
}
