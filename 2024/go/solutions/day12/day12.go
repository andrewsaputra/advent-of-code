package day12

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day12.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day12.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	matrix := helper.ToMatrix(path)

	var result int
	visited := make(map[Pos]bool)
	for row := range matrix {
		for col := range matrix[row] {
			if visited[Pos{Row: row, Col: col}] {
				continue
			}

			result += explorePt1(matrix, visited, row, col)
		}
	}
	return result
}

func solvePart2(path string) int {
	matrix := helper.ToMatrix(path)

	var result int
	visited := make(map[Pos]bool)
	for row := range matrix {
		for col := range matrix[row] {
			if visited[Pos{Row: row, Col: col}] {
				continue
			}

			result += explorePt2(matrix, visited, row, col)
		}
	}
	return result
}

type Pos struct {
	Row int
	Col int
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Perimeter struct {
	Pos
	Dir Direction
}

func explorePt1(matrix [][]byte, visited map[Pos]bool, row int, col int) int {
	numRows, numCols := len(matrix), len(matrix[1])

	var area, perimeter int

	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}

	startPos := Pos{Row: row, Col: col}
	visited[startPos] = true
	target := matrix[row][col]
	queue := []Pos{startPos}

	for len(queue) > 0 {
		var newQueue []Pos
		for _, pos := range queue {
			area++

			for i := range drow {
				nextRow := pos.Row + drow[i]
				nextCol := pos.Col + dcol[i]

				if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols || matrix[nextRow][nextCol] != target {
					perimeter++
					continue
				}

				nextPos := Pos{Row: nextRow, Col: nextCol}
				if visited[nextPos] {
					continue
				}

				visited[nextPos] = true
				newQueue = append(newQueue, nextPos)
			}
		}

		queue = newQueue
	}

	//fmt.Printf("target:%s, area:%d, perimeter:%d\n", string(target), area, perimeter)

	return area * perimeter
}

func explorePt2(matrix [][]byte, visited map[Pos]bool, row int, col int) int {
	numRows, numCols := len(matrix), len(matrix[1])

	var area, perimeter int

	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}
	dirs := []Direction{UP, RIGHT, DOWN, LEFT}

	startPos := Pos{Row: row, Col: col}
	visited[startPos] = true
	explored := make(map[Perimeter]bool)
	target := matrix[row][col]
	queue := []Pos{startPos}

	for len(queue) > 0 {
		var newQueue []Pos
		for _, pos := range queue {
			area++

			for i := range drow {
				nextRow := pos.Row + drow[i]
				nextCol := pos.Col + dcol[i]
				dir := dirs[i]

				nextPos := Pos{Row: nextRow, Col: nextCol}
				if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols || matrix[nextRow][nextCol] != target {
					explored[Perimeter{Pos: nextPos, Dir: dir}] = true

					var alreadyCounted bool
					switch dir {
					case UP, DOWN:
						for _, neighborDst := range []Perimeter{{Pos: Pos{Row: nextRow, Col: nextCol - 1}, Dir: dir}, {Pos: Pos{Row: nextRow, Col: nextCol + 1}, Dir: dir}} {
							if _, ok := explored[neighborDst]; ok {
								alreadyCounted = true
								break
							}
						}

					case LEFT, RIGHT:
						for _, neighborDst := range []Perimeter{{Pos: Pos{Row: nextRow - 1, Col: nextCol}, Dir: dir}, {Pos: Pos{Row: nextRow + 1, Col: nextCol}, Dir: dir}} {
							if _, ok := explored[neighborDst]; ok {
								alreadyCounted = true
								break
							}
						}
					}

					if !alreadyCounted {
						perimeter++
					}

					continue
				}

				if visited[nextPos] {
					continue
				}

				visited[nextPos] = true
				newQueue = append(newQueue, nextPos)
			}
		}

		queue = newQueue
	}

	//fmt.Printf("target:%s, area:%d, perimeter:%d\n", string(target), area, perimeter)

	return area * perimeter
}
