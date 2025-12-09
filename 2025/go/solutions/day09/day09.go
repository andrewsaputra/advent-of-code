package day09

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
	"sort"
)

func Solve() {
	res1 := solvePart1("inputs/day09.txt")
	fmt.Println("Part 1:", res1)

	//res2 := solvePart2("inputs/day01.txt")
	//fmt.Println("Part 2:", res2)
}

type Point struct {
	X int64
	Y int64
}

func solvePart1(path string) int64 {
	var points []Point
	for _, line := range helper.ReadLines(path) {
		var x, y int64
		fmt.Sscanf(line, "%d,%d", &x, &y)
		points = append(points, Point{X: x, Y: y})
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})

	var maxArea int64
	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			area := calcWidth(p1.X, p2.X) * calcWidth(p1.Y, p2.Y)
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}

func calcWidth(a, b int64) int64 {
	if a > b {
		return a - b + 1
	}
	return b - a + 1
}
