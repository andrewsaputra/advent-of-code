package day07

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day07.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day07.txt")
	fmt.Println("Part 2:", res2)
}

type Pos struct {
	Row int
	Col int
}

func solvePart1(path string) int {
	grid := helper.ToMatrix(path)
	numRows, numCols := len(grid), len(grid[0])
	beamColumns := make(map[int]bool)
	for col := range numCols {
		if grid[0][col] == 'S' {
			beamColumns[col] = true
			break
		}
	}

	var result int
	for row := 2; row < numRows; row += 2 {
		newBeamColumns := make(map[int]bool)
		for col := range beamColumns {
			if grid[row][col] == '^' {
				result++

				if col-1 >= 0 {
					newBeamColumns[col-1] = true
				}

				if col+1 < numCols {
					newBeamColumns[col+1] = true
				}
			} else {
				newBeamColumns[col] = true
			}
		}

		beamColumns = newBeamColumns
	}

	return result
}

func solvePart2(path string) int64 {
	grid := helper.ToMatrix(path)
	numCols := len(grid[0])
	var startColumn int
	for col := range numCols {
		if grid[0][col] == 'S' {
			startColumn = col
			break
		}
	}

	dp := make(map[string]int64)
	return traverse(grid, dp, 2, startColumn)
}

func traverse(grid [][]byte, dp map[string]int64, row int, beamColumn int) int64 {
	numRows, numCols := len(grid), len(grid[0])
	if row >= numRows {
		return 1
	}

	key := fmt.Sprintf("%d-%d", row, beamColumn)
	if res, ok := dp[key]; ok {
		return res
	}

	var res int64
	if grid[row][beamColumn] == '^' {
		if beamColumn-1 >= 0 {
			res += traverse(grid, dp, row+2, beamColumn-1)
		}

		if beamColumn+1 < numCols {
			res += traverse(grid, dp, row+2, beamColumn+1)
		}
	} else {
		res = traverse(grid, dp, row+2, beamColumn)
	}

	dp[key] = res
	return res
}
