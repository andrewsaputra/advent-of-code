package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res1 := solvePart1("inputs.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs.txt")
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solvePart1(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	return countEnergized(grid, 0, 0, East)
}

func solvePart2(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	maxEnergized := 0
	numRows, numCols := len(grid), len(grid[0])
	for row := 0; row < numRows; row++ {
		maxEnergized = max(maxEnergized, countEnergized(grid, row, 0, East))
		maxEnergized = max(maxEnergized, countEnergized(grid, row, numCols-1, West))
	}

	for col := 0; col < numCols; col++ {
		maxEnergized = max(maxEnergized, countEnergized(grid, 0, col, South))
		maxEnergized = max(maxEnergized, countEnergized(grid, numRows-1, col, North))
	}

	return maxEnergized
}

type Direction int

const (
	North Direction = 0
	South Direction = 1
	East  Direction = 2
	West  Direction = 3
)

func countEnergized(grid [][]byte, row int, col int, direction Direction) int {
	cache := [][][4]bool{}
	for _, data := range grid {
		cache = append(cache, make([][4]bool, len(data)))
	}

	traverse(grid, cache, row, col, direction)
	total := 0
	for _, data := range cache {
		for _, data2 := range data {
			for _, val := range data2 {
				if val {
					total++
					break
				}
			}
		}
	}

	return total
}

func traverse(grid [][]byte, cache [][][4]bool, row int, col int, direction Direction) {
	numRows, numCols := len(grid), len(grid[0])
	if row < 0 || row >= numRows || col < 0 || col >= numCols || cache[row][col][direction] {
		return
	}

	cache[row][col][direction] = true

	switch grid[row][col] {
	case '.':
		switch direction {
		case North:
			traverse(grid, cache, row-1, col, direction)
		case South:
			traverse(grid, cache, row+1, col, direction)
		case East:
			traverse(grid, cache, row, col+1, direction)
		case West:
			traverse(grid, cache, row, col-1, direction)
		}
	case '/':
		switch direction {
		case North:
			traverse(grid, cache, row, col+1, East)
		case South:
			traverse(grid, cache, row, col-1, West)
		case East:
			traverse(grid, cache, row-1, col, North)
		case West:
			traverse(grid, cache, row+1, col, South)
		}
	case '\\':
		switch direction {
		case North:
			traverse(grid, cache, row, col-1, West)
		case South:
			traverse(grid, cache, row, col+1, East)
		case East:
			traverse(grid, cache, row+1, col, South)
		case West:
			traverse(grid, cache, row-1, col, North)
		}
	case '|':
		switch direction {
		case North:
			traverse(grid, cache, row-1, col, direction)
		case South:
			traverse(grid, cache, row+1, col, direction)
		case East, West:
			traverse(grid, cache, row-1, col, North)
			traverse(grid, cache, row+1, col, South)
		}
	case '-':
		switch direction {
		case North, South:
			traverse(grid, cache, row, col-1, West)
			traverse(grid, cache, row, col+1, East)
		case East:
			traverse(grid, cache, row, col+1, direction)
		case West:
			traverse(grid, cache, row, col-1, direction)
		}
	}
}
