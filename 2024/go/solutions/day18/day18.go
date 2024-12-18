package day18

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day18.txt", 71, 1024)
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day18.txt", 71, 1024)
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string, gridSize int, numFallen int) int {
	matrix := make([][]byte, gridSize)
	for row := range matrix {
		matrix[row] = make([]byte, gridSize)
		for col := range matrix[row] {
			matrix[row][col] = '.'
		}
	}

	var count int
	for _, line := range helper.ReadLines(path) {
		if count == numFallen {
			break
		}

		var row, col int
		fmt.Sscanf(line, "%d,%d", &col, &row)
		matrix[row][col] = '#'
		count++
	}

	return findMinSteps(matrix)
}

func solvePart2(path string, gridSize int, startingNumFallen int) string {
	matrix := make([][]byte, gridSize)
	for row := range matrix {
		matrix[row] = make([]byte, gridSize)
		for col := range matrix[row] {
			matrix[row][col] = '.'
		}
	}

	inputs := helper.ReadLines(path)
	var numFallen int
	for _, line := range inputs {
		if numFallen == startingNumFallen {
			break
		}

		var row, col int
		fmt.Sscanf(line, "%d,%d", &col, &row)
		matrix[row][col] = '#'
		numFallen++
	}

	for findMinSteps(matrix) != -1 {
		var row, col int
		fmt.Sscanf(inputs[numFallen], "%d,%d", &col, &row)
		matrix[row][col] = '#'
		numFallen++
	}

	return inputs[numFallen-1]
}

type Pos struct {
	Row int
	Col int
}

type Item struct {
	Pos
	Steps int
}

func parseMap(filepath string, gridSize int, numFallen int) [][]byte {
	matrix := make([][]byte, gridSize)
	for row := range matrix {
		matrix[row] = make([]byte, gridSize)
		for col := range matrix[row] {
			matrix[row][col] = '.'
		}
	}

	var count int
	for _, line := range helper.ReadLines(filepath) {
		if count == numFallen {
			break
		}

		var row, col int
		fmt.Sscanf(line, "%d,%d", &col, &row)
		matrix[row][col] = '#'
		count++
	}

	return matrix
}

func findMinSteps(matrix [][]byte) int {
	numRows, numCols := len(matrix), len(matrix[0])
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}
	queue := []Item{{Pos: Pos{Row: 0, Col: 0}, Steps: 0}}
	minSteps := make(map[Pos]int)

	for len(queue) > 0 {
		var newQueue []Item
		for _, item := range queue {
			if item.Row == numRows-1 && item.Col == numCols-1 {
				return item.Steps
			}

			for i := range drow {
				nextRow := item.Row + drow[i]
				nextCol := item.Col + dcol[i]
				if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols || matrix[nextRow][nextCol] == '#' {
					continue
				}

				nextSteps := item.Steps + 1
				nextPos := Pos{Row: nextRow, Col: nextCol}
				if val, ok := minSteps[nextPos]; !ok || nextSteps < val {
					minSteps[nextPos] = nextSteps
					newQueue = append(newQueue, Item{Pos: nextPos, Steps: nextSteps})
				}
			}
		}

		queue = newQueue
	}

	return -1
}
