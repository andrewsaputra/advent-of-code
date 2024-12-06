package day06

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day06.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day06.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	matrix, startRow, startCol := toInputs(path)
	patrol(matrix, startRow, startCol, UP)

	var result int
	for row := range matrix {
		for _, val := range matrix[row] {
			if val == 'x' {
				result++
			}
		}
	}

	return result
}

func solvePart2(path string) int {
	matrix, startRow, startCol := toInputs(path)

	var result int
	for row := range matrix {
		for col, val := range matrix[row] {
			if val != '.' {
				continue
			}

			matrix[row][col] = '#'
			path := Path{Row: startRow, Col: startCol, Dir: UP}
			explored := make(map[Path]bool)
			if patrolHasLoop(matrix, explored, path) {
				result++
			}
			matrix[row][col] = '.'
		}
	}

	return result
}

func toInputs(path string) (matrix [][]byte, startRow int, startCol int) {
	matrix = [][]byte{}
	for _, line := range helper.ReadLines(path) {
		matrix = append(matrix, []byte(line))

		if col := strings.Index(line, "^"); col != -1 {
			startRow = len(matrix) - 1
			startCol = col
		}
	}
	return
}

type Direction int

const (
	UP    Direction = 0
	DOWN  Direction = 1
	LEFT  Direction = 2
	RIGHT Direction = 3
)

func patrol(matrix [][]byte, row int, col int, dir Direction) {
	numRows, numCols := len(matrix), len(matrix[0])
	if row < 0 || row >= numRows || col < 0 || col >= numCols {
		return
	}

	matrix[row][col] = 'x'

	var nextRow, nextCol int
	for {
		nextRow, nextCol = nextPos(row, col, dir)
		if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols {
			return
		}

		if matrix[nextRow][nextCol] == '#' {
			dir = turnRight(dir)
			continue
		}

		break
	}

	patrol(matrix, nextRow, nextCol, dir)
}

func nextPos(row, col int, dir Direction) (nextRow, nextCol int) {
	switch dir {
	case UP:
		nextRow = row - 1
		nextCol = col
	case DOWN:
		nextRow = row + 1
		nextCol = col
	case LEFT:
		nextRow = row
		nextCol = col - 1
	case RIGHT:
		nextRow = row
		nextCol = col + 1
	}

	return
}

func turnRight(curr Direction) Direction {
	switch curr {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	default:
		return UP
	}
}

type Path struct {
	Row int
	Col int
	Dir Direction
}

func patrolHasLoop(matrix [][]byte, explored map[Path]bool, path Path) bool {
	numRows, numCols := len(matrix), len(matrix[0])
	if path.Row < 0 || path.Row >= numRows || path.Col < 0 || path.Col >= numCols {
		return false
	}

	if explored[path] {
		return true
	}

	explored[path] = true

	var nextRow, nextCol int
	nextDir := path.Dir
	for {
		nextRow, nextCol = nextPos(path.Row, path.Col, nextDir)
		if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols {
			return false
		}

		if matrix[nextRow][nextCol] == '#' {
			nextDir = turnRight(nextDir)
			continue
		}

		break
	}

	return patrolHasLoop(matrix, explored, Path{Row: nextRow, Col: nextCol, Dir: nextDir})
}
