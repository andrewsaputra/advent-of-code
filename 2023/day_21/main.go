package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	grid, start := parseInput("inputs.txt")

	res1 := moveNormal(grid, start, 64)
	fmt.Println("Part 1:", res1)

	res2 := solvePart2(grid, start, 500) //WIP solve me
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type Item struct {
	Row   int
	Col   int
	Steps int
}

func solvePart2(grid [][]byte, start [2]int, targetSteps int) int64 {
	fullGrid := int64(0)
	for _, rows := range grid {
		for _, val := range rows {
			if val != '#' {
				fullGrid++
			}
		}
	}

	total := int64(0)
	numRows, numCols := len(grid), len(grid[0])
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}

	for i := range drow {
		var nextTargetSteps int
		nextRow := (start[0] + drow[i]*targetSteps) % numRows
		nextCol := (start[1] + dcol[i]*targetSteps) % numCols

		switch {
		case nextRow < 0:
			nextTargetSteps = -nextRow
			nextRow = numRows - 1
		case nextRow >= 0:
			nextTargetSteps = nextRow
			nextRow = 0
		case nextCol < 0:
			nextTargetSteps = -nextCol
			nextCol = numCols - 1
		case nextCol >= 0:
			nextTargetSteps = nextCol
			nextCol = 0
		}

		fmt.Println(nextTargetSteps)

	}

	fmt.Println(fullGrid)

	return total
}

func parseInput(path string) ([][]byte, [2]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := [][]byte{}
	var start [2]int
	numRows := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))

		if colIdx := strings.Index(line, "S"); colIdx != -1 {
			start = [2]int{numRows, colIdx}
		}

		numRows++
	}

	return grid, start
}

func moveNormal(grid [][]byte, start [2]int, targetSteps int) int64 {
	indexKey := func(row, col, steps int) string {
		return fmt.Sprintf("%d-%d-%d", row, col, steps)
	}

	numRows, numCols := len(grid), len(grid[0])
	visited := make(map[string]bool)
	result := int64(0)
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}
	queue := []Item{{start[0], start[1], 0}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		row, col := item.Row, item.Col
		if item.Steps == targetSteps {
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
			key := indexKey(nextRow, nextCol, nextSteps)
			if visited[key] {
				continue
			}

			visited[key] = true
			queue = append(queue, Item{nextRow, nextCol, nextSteps})
		}
	}

	return result
}
