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

	grid, src, dest := parseInput("inputs.txt")

	res1 := solvePart1(grid, src, dest)
	fmt.Println("Part 1:", res1)

	res2 := solvePart2(grid, src, dest)
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type Item struct {
	Row   int
	Col   int
	Steps int64
}

func parseInput(path string) ([][]byte, []int, []int) {
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

	src := []int{
		0,
		strings.Index(string(grid[0]), "."),
	}

	lastRowIdx := len(grid) - 1
	dest := []int{
		lastRowIdx,
		strings.Index(string(grid[lastRowIdx]), "."),
	}

	return grid, src, dest
}

func solvePart1(grid [][]byte, src []int, dest []int) int64 {
	return longestPathWithSlopes(grid, src, dest)
}

func solvePart2(grid [][]byte, src []int, dest []int) int64 {
	//optimization so we can remove logic for bounds checking
	copy := make([][]byte, len(grid))
	for row := range grid {
		copy[row] = append(copy[row], grid[row]...)
	}

	copy[src[0]][src[1]] = '#'
	copy[dest[0]][dest[1]] = '#'
	src[0]++
	dest[0]--

	return longestPathNoSlopesNoBoundsChecking(copy, src, dest) + 2
}

func longestPathWithSlopes(grid [][]byte, src []int, dest []int) int64 {
	numRows, numCols := len(grid), len(grid[0])

	var dfs func(Item, []int, [][]bool) int64
	dfs = func(item Item, dest []int, visited [][]bool) int64 {
		if visited[item.Row][item.Col] {
			return -1
		}
		visited[item.Row][item.Col] = true
		defer func() {
			visited[item.Row][item.Col] = false
		}()

		if item.Row == dest[0] && item.Col == dest[1] {
			return item.Steps
		}

		var drow, dcol []int
		switch grid[item.Row][item.Col] {
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
		default:
			panic("unknown cell")
		}

		result := int64(-1)
		nextSteps := item.Steps + 1
		for i := range drow {
			nextRow := item.Row + drow[i]
			nextCol := item.Col + dcol[i]
			if nextCol < 0 || nextCol >= numCols || nextRow < 0 || nextRow >= numRows || grid[nextRow][nextCol] == '#' {
				continue
			}

			result = max(result, dfs(Item{nextRow, nextCol, nextSteps}, dest, visited))

		}

		return result
	}

	visited := make([][]bool, numRows)
	for row := range visited {
		visited[row] = make([]bool, numCols)
	}

	return dfs(Item{src[0], src[1], 0}, dest, visited)
}

func longestPathNoSlopesNoBoundsChecking(grid [][]byte, src []int, dest []int) int64 {
	numRows, numCols := len(grid), len(grid[0])
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}

	var dfs func(Item, []int, [][]bool) int64
	dfs = func(item Item, dest []int, visited [][]bool) int64 {
		if visited[item.Row][item.Col] {
			return -1
		}
		visited[item.Row][item.Col] = true
		defer func() {
			visited[item.Row][item.Col] = false
		}()

		if item.Row == dest[0] && item.Col == dest[1] {
			return item.Steps
		}

		result := int64(-1)
		nextSteps := item.Steps + 1
		for i := range drow {
			nextRow := item.Row + drow[i]
			nextCol := item.Col + dcol[i]
			if grid[nextRow][nextCol] == '#' {
				continue
			}

			result = max(result, dfs(Item{nextRow, nextCol, nextSteps}, dest, visited))

		}

		return result
	}

	visited := make([][]bool, numRows)
	for row := range visited {
		visited[row] = make([]bool, numCols)
	}

	return dfs(Item{src[0], src[1], 0}, dest, visited)
}
