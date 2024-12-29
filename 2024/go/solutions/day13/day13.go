package day13

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day13.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day13.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int64 {
	var result int64
	inputs := helper.ReadLines(path)
	for i := 0; i < len(inputs); i += 3 {
		ax, ay, bx, by, px, py := parseInput(inputs, i)

		result += min(
			calculate1(ax, ay, bx, by, px, py),
			calculate2(ax, ay, bx, by, px, py),
		)
	}

	return result
}

func solvePart2(path string) int64 {
	var result int64
	inputs := helper.ReadLines(path)
	for i := 0; i < len(inputs); i += 3 {
		ax, ay, bx, by, px, py := parseInput(inputs, i)
		px += 10000000000000
		py += 10000000000000

		result += min(
			calculate1(ax, ay, bx, by, px, py),
			calculate2(ax, ay, bx, by, px, py),
		)
	}

	return result
}

func parseInput(inputs []string, idx int) (ax, ay, bx, by, px, py int64) {
	fmt.Sscanf(inputs[idx], "Button A: X+%d, Y+%d", &ax, &ay)
	fmt.Sscanf(inputs[idx+1], "Button B: X+%d, Y+%d", &bx, &by)
	fmt.Sscanf(inputs[idx+2], "Prize: X=%d, Y=%d", &px, &py)
	return
}

/*
Equations :
pressA * ax + pressB * bx = px
pressA * ay + pressB * by = py

calculate1 is prioritizing pressA
pressA = (px * by - py * bx) / (ax * by - ay * bx)

calculate2 is prioritizing pressB
pressB = (px * ay - ax * py) / (ay * bx - ax * by)
*/

func calculate1(ax, ay, bx, by, px, py int64) int64 {
	numerator := px*by - py*bx
	divisor := ax*by - ay*bx
	if divisor == 0 || numerator%divisor != 0 {
		return 0
	}
	pressA := numerator / divisor
	pressB := (px - (ax * pressA)) / bx
	return 3*pressA + pressB
}

func calculate2(ax, ay, bx, by, px, py int64) int64 {
	numerator := px*ay - ax*py
	divisor := ay*bx - ax*by
	if divisor == 0 || numerator%divisor != 0 {
		return 0
	}
	pressB := numerator / divisor
	pressA := (px - (bx * pressB)) / ax
	return 3*pressA + pressB
}
