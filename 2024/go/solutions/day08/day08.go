package day08

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day08.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day08.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	matrix := toMatrix(path)
	freqsMap := make(map[byte][]Pos)
	for row := range matrix {
		for col, val := range matrix[row] {
			if val == '.' {
				continue
			}

			freqsMap[val] = append(freqsMap[val], Pos{Row: row, Col: col})
		}
	}

	copyMatrix := toMatrix(path)
	numRows, numCols := len(matrix), len(matrix[0])
	for _, positions := range freqsMap {
		for i, pos1 := range positions {
			for j := i + 1; j < len(positions); j++ {
				pos2 := positions[j]
				diffRow := pos1.Row - pos2.Row
				diffCol := pos1.Col - pos2.Col

				newRow := pos1.Row + diffRow
				newCol := pos1.Col + diffCol
				if newRow >= 0 && newRow < numRows && newCol >= 0 && newCol < numCols {
					copyMatrix[newRow][newCol] = '#'
				}

				newRow = pos2.Row - diffRow
				newCol = pos2.Col - diffCol
				if newRow >= 0 && newRow < numRows && newCol >= 0 && newCol < numCols {
					copyMatrix[newRow][newCol] = '#'
				}
			}
		}
	}

	var result int
	for row := range copyMatrix {
		for _, val := range copyMatrix[row] {
			if val == '#' {
				result++
			}
		}
	}

	return result
}

func solvePart2(path string) int {
	matrix := toMatrix(path)
	freqsMap := make(map[byte][]Pos)
	for row := range matrix {
		for col, val := range matrix[row] {
			if val == '.' {
				continue
			}

			freqsMap[val] = append(freqsMap[val], Pos{Row: row, Col: col})
		}
	}

	copyMatrix := toMatrix(path)
	for _, positions := range freqsMap {
		for i, pos1 := range positions {
			for j := i + 1; j < len(positions); j++ {
				pos2 := positions[j]
				diffRow := pos1.Row - pos2.Row
				diffCol := pos1.Col - pos2.Col

				route := Pos{
					Row: diffRow,
					Col: diffCol,
				}
				harmonize(copyMatrix, pos1, route)

				route = Pos{
					Row: -diffRow,
					Col: -diffCol,
				}
				harmonize(copyMatrix, pos1, route)
			}
		}
	}

	var result int
	for row := range copyMatrix {
		for _, val := range copyMatrix[row] {
			if val == '#' {
				result++
			}
		}
	}

	return result
}

func toMatrix(path string) [][]byte {
	var result [][]byte
	for _, line := range helper.ReadLines(path) {
		result = append(result, []byte(line))
	}
	return result
}

type Pos struct {
	Row int
	Col int
}

func harmonize(resultMatrix [][]byte, pos Pos, route Pos) {
	numRows, numCols := len(resultMatrix), len(resultMatrix[0])
	if pos.Row < 0 || pos.Row >= numRows || pos.Col < 0 || pos.Col >= numCols {
		return
	}

	resultMatrix[pos.Row][pos.Col] = '#'
	nextPos := Pos{
		Row: pos.Row + route.Row,
		Col: pos.Col + route.Col,
	}
	harmonize(resultMatrix, nextPos, route)
}
