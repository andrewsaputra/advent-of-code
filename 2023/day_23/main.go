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

	grid := parseInput("inputs.txt")

	res1 := solvePart1(grid)
	fmt.Println("Part 1:", res1)

	//res2 := solvePart2(grid, start)
	//fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type Item struct {
	Row   int
	Col   int
	Steps int
}

func parseInput(path string) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	return grid
}

func solvePart1(grid [][]byte) int {
	numRows, numCols := len(grid), len(grid[0])
	start := [2]int{0, strings.Index(string(grid[0]), ".")}
	end := [2]int{numRows - 1, strings.Index(string(grid[numRows-1]), ".")}

	visited := make([][]bool, numRows)
	maxSteps := make([][]int, numRows)
	for row := range grid {
		visited[row] = make([]bool, numCols)
		maxSteps[row] = make([]int, numCols)
	}

	return longestPath(grid, visited, maxSteps, start, end, 0)
}

func longestPath(grid [][]byte, visited [][]bool, maxSteps [][]int, start [2]int, end [2]int, steps int) int {
	numRows, numCols := len(grid), len(grid[0])

	row, col := start[0], start[1]
	var drow, dcol []int
	if row == end[0] && col == end[1] {
		return steps
	}

	if visited[row][col] {
		return -1
	}

	visited[row][col] = true
	nextSteps := steps + 1

	switch grid[row][col] {
	case '.':
		drow = []int{-1, 0, 1, 0}
		dcol = []int{0, 1, 0, -1}
	case '^':
		drow = []int{-1}
		dcol = []int{0}
	case '>':
		drow = []int{0}
		dcol = []int{1}
	case 'v':
		drow = []int{1}
		dcol = []int{0}
	case '<':
		drow = []int{0}
		dcol = []int{-1}
	}

	result := -1
	for i := range drow {
		nextRow, nextCol := row+drow[i], col+dcol[i]
		if nextCol < 0 || nextCol >= numCols || nextRow < 0 || nextRow >= numRows || grid[nextRow][nextCol] == '#' {
			continue
		}

		if nextSteps > maxSteps[nextRow][nextCol] {
			newVisited := make([][]bool, numRows)
			for i := range visited {
				newVisited[i] = append(newVisited[i], visited[i]...)
			}

			maxSteps[nextRow][nextCol] = nextSteps
			result = max(result, longestPath(grid, newVisited, maxSteps, [2]int{nextRow, nextCol}, end, nextSteps))
		}
	}

	return result
}
