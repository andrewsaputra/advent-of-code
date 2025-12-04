package day04

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day04.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day04.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	drow := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	dcol := []int{-1, 0, 1, 1, 1, 0, -1, -1}

	grid := helper.ToMatrix(path)
	numRows, numCols := len(grid), len(grid[0])

	var result int
	for row := range numRows {
		for col := range numCols {
			if grid[row][col] != '@' {
				continue
			}

			var count int
			for i := range drow {
				nextRow := row + drow[i]
				nextCol := col + dcol[i]
				if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols {
					continue
				}

				if grid[nextRow][nextCol] == '@' {
					count++
				}
			}

			if count < 4 {
				result++
			}
		}
	}

	return result
}

type Item struct {
	Row int
	Col int
}

func solvePart2(path string) int {
	drow := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	dcol := []int{-1, 0, 1, 1, 1, 0, -1, -1}

	grid := helper.ToMatrix(path)
	numRows, numCols := len(grid), len(grid[0])

	var result int
	for {
		var queue []Item

		for row := range numRows {
			for col := range numCols {
				if grid[row][col] != '@' {
					continue
				}

				var count int
				for i := range drow {
					nextRow := row + drow[i]
					nextCol := col + dcol[i]
					if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols {
						continue
					}

					if grid[nextRow][nextCol] == '@' {
						count++
					}
				}

				if count < 4 {
					queue = append(queue, Item{Row: row, Col: col})
					result++
				}
			}
		}

		if len(queue) == 0 {
			break
		}

		for _, item := range queue {
			grid[item.Row][item.Col] = '.'
		}
	}

	return result
}
