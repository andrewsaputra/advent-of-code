package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res1 := solve("inputs.txt", 1, 3)
	fmt.Println("Part 1:", res1)

	res2 := solve("inputs.txt", 4, 10)
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solve(path string, minSteps int, maxSteps int) int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		tmp := []int{}
		for _, v := range line {
			tmp = append(tmp, int(v-'0'))
		}

		grid = append(grid, tmp)
	}

	return djikstra(grid, minSteps, maxSteps)
}

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Item struct {
	Row   int
	Col   int
	Cost  int
	Steps int
	Dir   Direction
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Cost < pq[j].Cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(v any) { *pq = append(*pq, v.(*Item)) }
func (pq *PriorityQueue) Pop() any {
	idx := len(*pq) - 1
	item := (*pq)[idx]
	(*pq)[idx] = nil
	*pq = (*pq)[:idx]

	return item
}

func djikstra(grid [][]int, minSteps int, maxSteps int) int {
	numRows, numCols := len(grid), len(grid[0])
	minCosts := make(map[string]int)

	cacheKey := func(item *Item) string {
		return fmt.Sprintf("%d,%d,%d,%d", item.Row, item.Col, item.Dir, item.Steps)
	}

	directionMapping := map[Direction][]Direction{
		North: {North, East, West},
		South: {South, East, West},
		East:  {East, North, South},
		West:  {West, North, South},
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{0, minSteps, getCost(grid, 0, 0, 0, minSteps), minSteps, East})
	heap.Push(pq, &Item{minSteps, 0, getCost(grid, 0, 0, minSteps, 0), minSteps, South})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)

		row, col := item.Row, item.Col
		if row == numRows-1 && col == numCols-1 {
			return item.Cost
		}

		for _, nextDir := range directionMapping[item.Dir] {
			nextRow, nextCol := row, col
			nextSteps := minSteps
			if item.Dir == nextDir {
				switch nextDir {
				case North:
					nextRow--
				case South:
					nextRow++
				case East:
					nextCol++
				case West:
					nextCol--
				}

				nextSteps = item.Steps + 1
			} else {
				switch nextDir {
				case North:
					nextRow -= minSteps
				case South:
					nextRow += minSteps
				case East:
					nextCol += minSteps
				case West:
					nextCol -= minSteps
				}
			}

			if nextRow < 0 || nextCol < 0 || nextRow >= numRows || nextCol >= numCols || nextSteps > maxSteps {
				continue
			}

			newCost := getCost(grid, row, col, nextRow, nextCol)
			if newCost == -1 {
				continue
			}

			newCost += item.Cost
			nextItem := &Item{nextRow, nextCol, newCost, nextSteps, nextDir}
			key := cacheKey(nextItem)
			cachedCost, ok := minCosts[key]
			if !ok || newCost < cachedCost {
				minCosts[key] = newCost
				heap.Push(pq, nextItem)
			}
		}
	}

	return -1
}

func getCost(grid [][]int, row, col, nextRow, nextCol int) int {
	numRows, numCols := len(grid), len(grid[0])
	if nextCol < 0 || nextCol >= numCols || nextRow < 0 || nextRow >= numRows {
		return -1
	}

	cost := 0
	modRow, modCol := 0, 0
	var steps int
	switch {
	case row < nextRow:
		modRow++
		steps = nextRow - row
	case row > nextRow:
		modRow--
		steps = row - nextRow
	case col < nextCol:
		modCol++
		steps = nextCol - col
	case col > nextCol:
		modCol--
		steps = col - nextCol
	}

	for i := 0; i < steps; i++ {
		row += modRow
		col += modCol
		cost += grid[row][col]
	}

	return cost
}
