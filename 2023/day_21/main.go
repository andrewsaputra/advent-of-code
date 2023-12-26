package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	grid, start := parseInput("inputs.txt")

	res1 := calculateReachableCells(grid, start, 64)
	fmt.Println("Part 1:", res1)

	//fmt.Println("coba", coba(grid, start, 64))

	res2 := solvePart2(grid, start, 26501365)
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type Item struct {
	Row   int
	Col   int
	Steps int
}

func calculateReachableCells(grid [][]byte, start []int, targetSteps int) int64 {
	indexKey := func(row, col, steps int) string {
		return fmt.Sprintf("%d-%d-%d", row, col, steps)
	}

	numRows, numCols := len(grid), len(grid[0])

	//used only for visualizing result
	resultGrid := make([][]byte, numRows)
	for row := range resultGrid {
		resultGrid[row] = make([]byte, numCols)
		for col := range resultGrid[row] {
			resultGrid[row][col] = grid[row][col]
		}
	}

	visited := make(map[string]bool)
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}
	queue := []Item{{start[0], start[1], 0}}

	var bfs func([]Item) int64
	bfs = func(queue []Item) int64 {
		result := int64(0)
		newQueue := []Item{}
		for _, item := range queue {
			row, col := item.Row, item.Col
			key := indexKey(row, col, item.Steps)
			if visited[key] {
				continue
			}
			visited[key] = true

			if item.Steps == targetSteps {
				resultGrid[row][col] = '*'
				result++
				continue
			}

			for i := range drow {
				nextRow := row + drow[i]
				nextCol := col + dcol[i]
				if nextRow < 0 || nextCol < 0 || nextRow >= numRows || nextCol >= numCols || grid[nextRow][nextCol] == '#' {
					continue
				}

				nextSteps := item.Steps + 1
				newQueue = append(newQueue, Item{nextRow, nextCol, nextSteps})
			}
		}

		if len(newQueue) > 0 {
			result += bfs(newQueue)
		}

		return result
	}

	result := bfs(queue)
	//writeResult("output.txt", resultGrid)
	return result
}

/*
Excellent Summary Here : https://github.com/villuna/aoc23/wiki/A-Geometric-solution-to-advent-of-code-2023,-day-21

Important Notes :
- 26501365 steps = (202300 * 131) + 65
- grid size = 131 x 131
- S = center of grid = (65, 65)
- n = number of grids
- if n % 2 == 0 :
  - we have (n+1)^2 odd grids + n^2 even grids
  - we have n+1 odd outer grid corners and n even inner grid corners

- if n % 2 == 1 :
  - we have n^2 odd grids + (n+1)^2 even grids
  - we have n+1 even outer grid corners and n odd inner grid corners

- final formula if n % 2 == 0 :
(n+1)^2 * odd grids + n^2 * even grids - (n+1) * odd outer grid corners + n * even inner grid corners
*/
func solvePart2(grid [][]byte, start []int, targetSteps int) int64 {
	numRows := len(grid)
	if numRows%2 == 0 {
		panic("solution only works for odd length grid")
	}

	if (targetSteps-start[0])%numRows != 0 {
		panic("solution requires steps which covers full grid length")
	}

	numGrids := int64((targetSteps - start[0]) / numRows)

	amountFullOddGrid := calculateReachableCells(grid, start, numRows)
	amountFullEvenGrid := calculateReachableCells(grid, start, numRows-1)
	amountCornersOddGrid := amountFullOddGrid - calculateReachableCells(grid, start, start[0])
	amountCornersEvenGrid := amountFullEvenGrid - calculateReachableCells(grid, start, start[0]-1)

	//fmt.Println("amountFullEvenGrid", "amountFullOddGrid", "amountCornersEvenGrid", "amountCornersOddGrid")
	//fmt.Println(amountFullEvenGrid, amountFullOddGrid, amountCornersEvenGrid, amountCornersOddGrid)

	var result int64
	result = int64(math.Pow(float64(numGrids+1), 2)) * amountFullOddGrid
	result += int64(math.Pow(float64(numGrids), 2)) * amountFullEvenGrid
	result -= (numGrids + 1) * amountCornersOddGrid
	result += numGrids * amountCornersEvenGrid

	return result
}

func parseInput(path string) ([][]byte, []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := [][]byte{}
	var start []int
	numRows := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))

		if colIdx := strings.Index(line, "S"); colIdx != -1 {
			start = []int{numRows, colIdx}
		}

		numRows++
	}

	return grid, start
}

func writeResult(output string, grid [][]byte) {
	file, _ := os.Create(output)
	defer file.Close()

	writer := bufio.NewWriter(file)

	for row := range grid {
		writer.Write(grid[row])
		writer.WriteByte('\n')
	}

	writer.Flush()
}
