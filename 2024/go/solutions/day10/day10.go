package day10

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day10.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day10.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	matrix := toMatrix(path)

	var result int
	for row := range matrix {
		for col, val := range matrix[row] {
			if val != '0' {
				continue
			}

			score := make(map[Pos]bool)
			trailScoreDFS(matrix, score, row, col)
			result += len(score)

		}
	}

	return result
}

func solvePart2(path string) int {
	matrix := toMatrix(path)

	var result int
	for row := range matrix {
		for col, val := range matrix[row] {
			if val != '0' {
				continue
			}

			var rating int
			trailRatingDFS(matrix, &rating, row, col)
			result += rating

		}
	}

	return result
}

type Pos struct {
	Row int
	Col int
}

func toMatrix(path string) [][]byte {
	var result [][]byte
	for _, line := range helper.ReadLines(path) {
		result = append(result, []byte(line))
	}
	return result
}

func trailScoreDFS(matrix [][]byte, result map[Pos]bool, row int, col int) {
	if matrix[row][col] == '9' {
		result[Pos{Row: row, Col: col}] = true
		return
	}

	numRows, numCols := len(matrix), len(matrix[0])
	nextVal := matrix[row][col] + 1
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}
	for i := range drow {
		nextRow := row + drow[i]
		nextCol := col + dcol[i]
		if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols || matrix[nextRow][nextCol] != nextVal {
			continue
		}

		trailScoreDFS(matrix, result, nextRow, nextCol)
	}
}

func trailRatingDFS(matrix [][]byte, result *int, row int, col int) {
	if matrix[row][col] == '9' {
		*result++
		return
	}

	numRows, numCols := len(matrix), len(matrix[0])
	nextVal := matrix[row][col] + 1
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}
	for i := range drow {
		nextRow := row + drow[i]
		nextCol := col + dcol[i]
		if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols || matrix[nextRow][nextCol] != nextVal {
			continue
		}

		trailRatingDFS(matrix, result, nextRow, nextCol)
	}
}
