package day14

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day14.txt")
	fmt.Println("Part 1:", res1)

	//res2 := solvePart2("inputs/day01.txt")
	//fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	numRows, numCols := 103, 101
	midRow, midCol := numRows/2, numCols/2
	seconds := 100

	var q1, q2, q3, q4 int
	for _, line := range helper.ReadLines(path) {
		var pRow, pCol, vRow, vCol int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pCol, &pRow, &vCol, &vRow)

		row := pRow + seconds*vRow
		if row < 0 {
			div := -1 * row / numRows
			row += (div + 1) * numRows
		}
		row %= numRows

		col := pCol + seconds*vCol
		if col < 0 {
			div := -1 * col / numCols
			col += (div + 1) * numCols
		}
		col %= numCols

		if row == midRow || col == midCol {
			continue
		}

		if row < midRow {
			if col < midCol {
				q1++
			} else {
				q2++
			}
		} else {
			if col < midCol {
				q3++
			} else {
				q4++
			}
		}
	}

	return q1 * q2 * q3 * q4
}

func solvePart2(path string) int {
	return 0
}
