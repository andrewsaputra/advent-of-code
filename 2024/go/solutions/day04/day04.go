package day04

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day04.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day04.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	matrix := toMatrix(helper.ReadLines(path))

	var result int
	for row := range matrix {
		for col := range matrix[row] {
			result += countXMAS(matrix, row, col)
		}
	}

	return result
}

func solvePart2(path string) int {
	matrix := toMatrix(helper.ReadLines(path))

	var result int
	for row := range matrix {
		for col := range matrix[row] {
			result += countXMAS2(matrix, row, col)
		}
	}

	return result
}

func toMatrix(lines []string) [][]byte {
	matrix := make([][]byte, len(lines))
	for i := range matrix {
		matrix[i] = []byte(lines[i])
	}
	return matrix
}

func countXMAS(matrix [][]byte, startRow int, startCol int) int {
	if matrix[startRow][startCol] != 'X' {
		return 0
	}

	var result int
	numRows, numCols := len(matrix), len(matrix[0])

	//left to right
	if startCol <= numCols-4 {
		if string(matrix[startRow][startCol:startCol+4]) == "XMAS" {
			result++
		}
	}

	//right to left
	if startCol >= 3 {
		if string(matrix[startRow][startCol-3:startCol+1]) == "SAMX" {
			result++
		}
	}

	//top to bottom
	if startRow <= numRows-4 {
		if matrix[startRow+1][startCol] == 'M' && matrix[startRow+2][startCol] == 'A' && matrix[startRow+3][startCol] == 'S' {
			result++
		}
	}

	//bottom to top
	if startRow >= 3 {
		if matrix[startRow-1][startCol] == 'M' && matrix[startRow-2][startCol] == 'A' && matrix[startRow-3][startCol] == 'S' {
			result++
		}
	}

	//top-left to bottom-right
	if startRow <= numRows-4 && startCol <= numCols-4 {
		if matrix[startRow+1][startCol+1] == 'M' && matrix[startRow+2][startCol+2] == 'A' && matrix[startRow+3][startCol+3] == 'S' {
			result++
		}
	}

	//bottom-left to top-right
	if startRow >= 3 && startCol <= numCols-4 {
		if matrix[startRow-1][startCol+1] == 'M' && matrix[startRow-2][startCol+2] == 'A' && matrix[startRow-3][startCol+3] == 'S' {
			result++
		}
	}

	//top-right to bottom-left
	if startRow <= numRows-4 && startCol >= 3 {
		if matrix[startRow+1][startCol-1] == 'M' && matrix[startRow+2][startCol-2] == 'A' && matrix[startRow+3][startCol-3] == 'S' {
			result++
		}
	}

	//bottom-right to top-left
	if startRow >= 3 && startCol >= 3 {
		if matrix[startRow-1][startCol-1] == 'M' && matrix[startRow-2][startCol-2] == 'A' && matrix[startRow-3][startCol-3] == 'S' {
			result++
		}
	}

	return result
}

func countXMAS2(matrix [][]byte, startRow int, startCol int) int {
	if matrix[startRow][startCol] != 'A' {
		return 0
	}

	numRows, numCols := len(matrix), len(matrix[0])
	if startRow < 1 || startRow > numRows-2 || startCol < 1 || startCol > numCols-2 {
		return 0
	}

	var count int
	//top-left to bottom-right || bottom-right to top-left
	if (matrix[startRow-1][startCol-1] == 'M' && matrix[startRow+1][startCol+1] == 'S') || (matrix[startRow+1][startCol+1] == 'M' && matrix[startRow-1][startCol-1] == 'S') {
		count++
	}

	//bottom-left to top-right || top-right to bottom-left
	if (matrix[startRow+1][startCol-1] == 'M' && matrix[startRow-1][startCol+1] == 'S') || (matrix[startRow-1][startCol+1] == 'M' && matrix[startRow+1][startCol-1] == 'S') {
		count++
	}

	if count == 2 {
		return 1
	}
	return 0
}
