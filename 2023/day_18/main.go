package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

// conventional solution by rendering the grid and using flood fill to exclude the outer area
func solvePart1(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	digPlans := []Plan{}
	for scanner.Scan() {
		line := scanner.Text()
		digPlans = append(digPlans, lineToPlan(line))
	}

	grid := planToGrid(digPlans)
	fillOuter(grid, '.', '*')

	return countTiles(grid, '*')
}

type Plan struct {
	Direction string
	Steps     int64
}

func lineToPlan(line string) Plan {
	data := strings.Split(line, " ")

	direction := data[0]
	steps, _ := strconv.Atoi(data[1])

	return Plan{
		Direction: direction,
		Steps:     int64(steps),
	}
}

func planToGrid(digPlans []Plan) [][]byte {
	row, col := int64(0), int64(0)
	minRow, maxRow, minCol, maxCol := int64(0), int64(0), int64(0), int64(0)

	for _, plan := range digPlans {
		switch plan.Direction {
		case "U":
			row -= plan.Steps
			if row < minRow {
				minRow = row
			}
		case "D":
			row += plan.Steps
			if row > maxRow {
				maxRow = row
			}
		case "R":
			col += plan.Steps
			if col > maxCol {
				maxCol = col
			}
		case "L":
			col -= plan.Steps
			if col < minCol {
				minCol = col
			}
		}
	}

	numRows := maxRow - minRow + 1
	numCols := maxCol - minCol + 1

	grid := make([][]byte, numRows)
	for i := range grid {
		grid[i] = make([]byte, numCols)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	row, col = -minRow, -minCol
	grid[row][col] = '#'
	for _, plan := range digPlans {
		for step := int64(0); step < plan.Steps; step++ {
			switch plan.Direction {
			case "U":
				row--
			case "D":
				row++
			case "R":
				col++
			case "L":
				col--
			}

			grid[row][col] = '#'
		}
	}

	return grid
}

func fillOuter(grid [][]byte, target byte, replace byte) {
	numRows, numCols := len(grid), len(grid[0])

	var floodFill func(int, int)
	floodFill = func(row int, col int) {
		if row < 0 || col < 0 || row >= numRows || col >= numCols || grid[row][col] != target {
			return
		}

		grid[row][col] = replace

		floodFill(row-1, col)
		floodFill(row+1, col)
		floodFill(row, col-1)
		floodFill(row, col+1)
	}

	for row := 0; row < numRows; row++ {
		floodFill(row, 0)
		floodFill(row, numCols-1)
	}

	for col := 0; col < numCols; col++ {
		floodFill(0, col)
		floodFill(numRows-1, col)
	}
}

func countTiles(grid [][]byte, exclude byte) int64 {
	total := int64(0)
	for _, data := range grid {
		for _, val := range data {
			if val != exclude {
				total++
			}
		}
	}

	return total
}

// using modified shoelace formula since the grid is too large for rendering
// same approach can also be used for part 1
func solvePart2(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	polygons := [][]int64{}
	borderLength := int64(0)
	row, col := int64(0), int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		var coord []int64
		var steps int64
		coord, steps, row, col = lineToCoordinate(line, row, col)
		polygons = append(polygons, coord)
		borderLength += steps
	}

	area := calculateAreaShoelace(polygons)
	area += borderLength/2 + 1 //adjustments for 2D grid

	return area
}

func lineToCoordinate(line string, row int64, col int64) ([]int64, int64, int64, int64) {
	data := strings.Split(line, " ")
	color, n := data[2], len(data[2])

	steps, _ := strconv.ParseInt(color[2:n-2], 16, 64)
	switch color[n-2] {
	case '0': //R
		col += steps
	case '1': //D
		row += steps
	case '2': //L
		col -= steps
	case '3': //U
		row -= steps
	}

	coordinate := []int64{row, col}
	return coordinate, steps, row, col
}

// https://www.themathdoctors.org/polygon-coordinates-and-areas/
func calculateAreaShoelace(coordinates [][]int64) int64 {
	n := len(coordinates)
	area := int64(0)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += coordinates[i][1] * coordinates[j][0]
		area -= coordinates[j][1] * coordinates[i][0]
	}

	if area < 0 {
		area = -area
	}
	return area / 2
}
