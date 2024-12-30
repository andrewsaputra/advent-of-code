package day20

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solve("inputs/day20.txt", 2)
	fmt.Println("Part 1:", res1)

	res2 := solve("inputs/day20.txt", 20)
	fmt.Println("Part 2:", res2)
}

func solve(path string, cheatsThreshold int) int {
	matrix, startPos, endPos := toInputs(path)
	minElapsed := make(map[Pos]int)
	defaultTime, item := honestShortestPath(matrix, startPos, endPos, minElapsed)
	routes := constructRoutes(item)

	var result int
	for pos, elapsed := range minElapsed {
		if elapsed > defaultTime-100 {
			continue
		}

		for _, route := range routes {
			distance := manhattanDistance(pos, route.Pos)
			if distance > cheatsThreshold ||
				elapsed+distance+route.DistanceToEnd > defaultTime-100 {
				continue
			}

			result++
		}
	}

	return result
}

type Pos struct {
	Row int
	Col int
}

type Item struct {
	Pos
	Prev *Item
}

type Route struct {
	Pos
	DistanceToEnd int
}

func honestShortestPath(matrix [][]byte, start Pos, endPos Pos, minElapsed map[Pos]int) (int, *Item) {
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}
	queue := []Item{{Pos: start, Prev: nil}}
	var elapsedTime int
	for len(queue) > 0 {
		var newQueue []Item
		for _, item := range queue {
			if item.Row == endPos.Row && item.Col == endPos.Col {
				return elapsedTime, &item
			}

			for i := range drow {
				nextRow := item.Row + drow[i]
				nextCol := item.Col + dcol[i]
				if matrix[nextRow][nextCol] == '#' {
					continue
				}

				nextPos := Pos{Row: nextRow, Col: nextCol}
				nextItem := Item{Pos: nextPos, Prev: &item}
				if val, ok := minElapsed[nextPos]; !ok || elapsedTime < val {
					minElapsed[nextPos] = elapsedTime
					newQueue = append(newQueue, nextItem)
				}
			}
		}
		queue = newQueue
		elapsedTime++
	}

	return -1, nil
}

func constructRoutes(item *Item) []Route {
	var distance int
	var result []Route
	curr := item
	for curr != nil {
		result = append(result, Route{Pos: curr.Pos, DistanceToEnd: distance})
		distance++
		curr = curr.Prev
	}
	return result
}

func manhattanDistance(a, b Pos) int {
	return helper.AbsDiff(a.Row, b.Row) + helper.AbsDiff(a.Col, b.Col)
}

func toInputs(filepath string) (matrix [][]byte, start Pos, end Pos) {
	matrix = helper.ToMatrix(filepath)
	for row := range matrix {
		for col, val := range matrix[row] {
			if val == 'S' {
				start = Pos{Row: row, Col: col}
				matrix[row][col] = '.'
			} else if val == 'E' {
				end = Pos{Row: row, Col: col}
				matrix[row][col] = '.'
			}
		}
	}
	return
}
