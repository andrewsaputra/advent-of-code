package day25

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day25.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day25-test.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	locks, keys := parseInputs(path)
	var result int
	for _, lock := range locks {
		for _, key := range keys {
			isValid := true
			for i := 0; i < numCols; i++ {
				if lock[i]+key[i] > numRows-2 {
					isValid = false
					break
				}
			}

			if isValid {
				result++
			}
		}
	}

	return result
}

func solvePart2(path string) int {
	return 0
}

const (
	numRows = 7
	numCols = 5
)

func parseInputs(filepath string) (locks [][]int, keys [][]int) {
	matcher := "#####"
	var row int
	var res []int
	var isLock bool
	for _, line := range helper.ReadLines(filepath) {
		if row == 0 {
			res = make([]int, numCols)
			isLock = (line == matcher)
			row++
			continue
		}

		if row < numRows-1 {
			for idx, val := range line {
				if val == '#' {
					res[idx]++
				}
			}
		}

		row++
		if row == numRows {
			if isLock {
				locks = append(locks, res)
			} else {
				keys = append(keys, res)
			}
			row = 0
		}
	}

	return
}
