package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res := solve("inputs.txt")
	fmt.Println(res)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solve(path string) int {
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
		maxEnergized = max(maxEnergized, countEnergized(grid, row, 0, "east"))
		maxEnergized = max(maxEnergized, countEnergized(grid, row, numCols-1, "west"))
	}

	for col := 0; col < numCols; col++ {
		maxEnergized = max(maxEnergized, countEnergized(grid, 0, col, "south"))
		maxEnergized = max(maxEnergized, countEnergized(grid, numRows-1, col, "north"))
	}

	return maxEnergized
}

func countEnergized(grid [][]byte, row int, col int, direction string) int {
	cache := [][][]string{}
	for _, data := range grid {
		cache = append(cache, make([][]string, len(data)))
	}

	traverse(grid, cache, row, col, direction)
	total := 0
	for _, data := range cache {
		for _, val := range data {
			if len(val) > 0 {
				total++
			}
		}
	}

	return total
}

func traverse(grid [][]byte, cache [][][]string, row int, col int, direction string) {
	numRows, numCols := len(grid), len(grid[0])
	if row < 0 || row >= numRows || col < 0 || col >= numCols {
		return
	}

	for _, dir := range cache[row][col] {
		if dir == direction {
			return
		}
	}

	cache[row][col] = append(cache[row][col], direction)

	switch grid[row][col] {
	case '.':
		switch direction {
		case "north":
			traverse(grid, cache, row-1, col, direction)
		case "south":
			traverse(grid, cache, row+1, col, direction)
		case "east":
			traverse(grid, cache, row, col+1, direction)
		case "west":
			traverse(grid, cache, row, col-1, direction)
		}
	case '/':
		switch direction {
		case "north":
			traverse(grid, cache, row, col+1, "east")
		case "south":
			traverse(grid, cache, row, col-1, "west")
		case "east":
			traverse(grid, cache, row-1, col, "north")
		case "west":
			traverse(grid, cache, row+1, col, "south")
		}
	case '\\':
		switch direction {
		case "north":
			traverse(grid, cache, row, col-1, "west")
		case "south":
			traverse(grid, cache, row, col+1, "east")
		case "east":
			traverse(grid, cache, row+1, col, "south")
		case "west":
			traverse(grid, cache, row-1, col, "north")
		}
	case '|':
		switch direction {
		case "north":
			traverse(grid, cache, row-1, col, direction)
		case "south":
			traverse(grid, cache, row+1, col, direction)
		case "east", "west":
			traverse(grid, cache, row-1, col, "north")
			traverse(grid, cache, row+1, col, "south")
		}
	case '-':
		switch direction {
		case "north", "south":
			traverse(grid, cache, row, col-1, "west")
			traverse(grid, cache, row, col+1, "east")
		case "east":
			traverse(grid, cache, row, col+1, direction)
		case "west":
			traverse(grid, cache, row, col-1, direction)
		}
	}
}
