package day20

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day20.txt")
	fmt.Println("Part 1:", res1)

	//res2 := solvePart2("inputs/day20.txt")
	//fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	matrix := helper.ToMatrix(path)
	startPos := findStartPos(matrix)

	defaultTime := honestShortestPath(matrix, startPos)
	fmt.Println("defaultTime", defaultTime)
	numRows, numCols := len(matrix), len(matrix[0])
	var result int
	for row := range matrix {
		for col, val := range matrix[row] {
			if row == 0 || row >= numRows-1 || col == 0 || col >= numCols-1 || val != '#' {
				continue
			}

			matrix[row][col] = '.'
			if honestShortestPath(matrix, startPos) <= defaultTime-100 {
				result++
			}
			matrix[row][col] = '#'
		}
	}

	return result
}

func solvePart2(path string) int {
	return 0
	/*
		matrix := helper.ToMatrix(path)
		startPos := findStartPos(matrix)

		defaultTime := shortestPath(matrix, startPos, 0, math.MaxInt32)
		numRows, numCols := len(matrix), len(matrix[0])
		var result int
		for row := range matrix {
			for col, val := range matrix[row] {
				if row == 0 || row >= numRows-1 || col == 0 || col >= numCols-1 || val != '#' {
					continue
				}

				matrix[row][col] = '.'
				if shortestPath(matrix, startPos, 19, defaultTime-100) != -1 {
					result++
				}
				matrix[row][col] = '#'
			}
		}

		return result
	*/
}

type Pos struct {
	Row int
	Col int
}

type Item struct {
	Pos
	CheatsLeft int
	Cheat      *Cheat
}

type Cheat struct {
	Start Pos
	End   Pos
}

func findStartPos(matrix [][]byte) (pos Pos) {
	for row := range matrix {
		for col := range matrix {
			if matrix[row][col] == 'S' {
				pos = Pos{Row: row, Col: col}
				return
			}
		}
	}
	return
}

func honestShortestPath(matrix [][]byte, start Pos) int {
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}
	minTime := make(map[Pos]int)
	queue := []Pos{start}
	var elapsedTime int
	for len(queue) > 0 {
		var newQueue []Pos
		for _, item := range queue {
			if matrix[item.Row][item.Col] == 'E' {
				return elapsedTime
			}

			for i := range drow {
				nextRow := item.Row + drow[i]
				nextCol := item.Col + dcol[i]
				if matrix[nextRow][nextCol] == '#' {
					continue
				}

				nextPos := Pos{Row: nextRow, Col: nextCol}
				if val, ok := minTime[nextPos]; !ok || elapsedTime < val {
					minTime[nextPos] = elapsedTime
					newQueue = append(newQueue, nextPos)
				}
			}
		}
		queue = newQueue
		elapsedTime++
	}
	return -1
}

func solve(matrix [][]byte, start Pos, cheatsLeft int, timeLimit int) int {
	/*
		cacheKey := func(item Item) string {
			return fmt.Sprintf("%d-%d-%d", item.Pos.Row, item.Pos.Col, item.CheatsLeft)
		}
	*/

	numRows, numCols := len(matrix), len(matrix[0])
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}
	//minTime := make(map[Item]int)
	usedCheats := make(map[Cheat]bool)

	queue := []Item{{Pos: start, CheatsLeft: cheatsLeft, Cheat: nil}}
	var elapsedTime int
	for len(queue) > 0 {
		if elapsedTime > timeLimit {
			break
		}

		var newQueue []Item
		for _, item := range queue {
			if matrix[item.Row][item.Col] == '.' || matrix[item.Row][item.Col] == 'E' {
				if item.Cheat != nil && item.CheatsLeft > 0 {
					item.CheatsLeft = 0
					item.Cheat.End = item.Pos
				}
			}

			if matrix[item.Row][item.Col] == 'E' {
				if item.Cheat != nil && !usedCheats[*item.Cheat] {
					usedCheats[*item.Cheat] = true
				}
				continue
			}

			for i := range drow {
				nextRow := item.Row + drow[i]
				nextCol := item.Col + dcol[i]
				if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols {
					continue
				}

				if matrix[nextRow][nextCol] == '#' {
					// todo
				}

				/*
					nextItem := Item{Pos: Pos{Row: nextRow, Col: nextCol}, NumCheats: nextNumCheats}
					if val, ok := minTime[nextItem]; !ok || elapsedTime < val {
						minTime[nextItem] = elapsedTime
						newQueue = append(newQueue, nextItem)
					}
				*/
			}
		}
		queue = newQueue
		elapsedTime++
	}
	return len(usedCheats)
}
