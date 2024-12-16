package day16

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"math"
)

func Solve() {
	res1 := solvePart1("inputs/day16-test.txt")
	fmt.Println("Part 1:", res1)

	res11 := solvePart1("inputs/day16.txt")
	fmt.Println("Part 11:", res11)

	//res2 := solvePart2("inputs/day01.txt")
	//fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	matrix := helper.ToMatrix(path)
	startPos, endPos := findStartEnd(matrix)
	memo := make(map[Pos]int)
	visited := make(map[Pos]bool)
	res := travel(matrix, memo, visited, startPos, endPos)
	return res
}

func solvePart2(path string) int {
	return 0
}

type Direction int

const (
	NORTH Direction = 0
	EAST  Direction = 1
	SOUTH Direction = 2
	WEST  Direction = 3
)

type Pos struct {
	Row int
	Col int
	Dir Direction
}

func findStartEnd(matrix [][]byte) (start Pos, end Pos) {
	for row := range matrix {
		for col, val := range matrix[row] {
			if val == 'S' {
				start = Pos{Row: row, Col: col, Dir: EAST}
			} else if val == 'E' {
				end = Pos{Row: row, Col: col}
			}
		}
	}
	return
}

func travel(matrix [][]byte, memo map[Pos]int, visited map[Pos]bool, pos Pos, end Pos) int {
	if pos.Row == end.Row && pos.Col == end.Col {
		return 0
	}

	if val, ok := memo[pos]; ok {
		return val
	}

	visited[pos] = true
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}

	score := math.MaxInt32
	for idx := range drow {
		if helper.AbsDiff(int(pos.Dir), idx) >= 2 {
			//continue
		}

		nextRow := pos.Row + drow[idx]
		nextCol := pos.Col + dcol[idx]
		nextDir := Direction(idx)
		modScore := 1
		if nextDir != pos.Dir {
			modScore += 1000
		}

		nextPos := Pos{Row: nextRow, Col: nextCol, Dir: nextDir}

		if matrix[nextRow][nextCol] == '#' || visited[nextPos] {
			continue
		}

		tmpScore := travel(matrix, memo, visited, nextPos, end)
		if tmpScore == math.MaxInt32 {
			continue
		}
		score = min(score, modScore+tmpScore)
	}

	visited[pos] = false
	memo[pos] = score
	return score
}
