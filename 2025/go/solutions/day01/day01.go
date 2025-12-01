package day01

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
	"strconv"
)

func Solve() {
	res1 := solvePart1("inputs/day01.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day01.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	var result int
	curr := 50
	for _, line := range helper.ReadLines(path) {
		dir := line[0]
		distance, _ := strconv.Atoi(line[1:])

		distance %= 100

		switch dir {
		case 'L':
			curr -= distance
			if curr < 0 {
				curr += 100
			}
		case 'R':
			curr = (curr + distance) % 100
		}

		if curr == 0 {
			result++
		}
	}

	return result
}

func solvePart2(path string) int {
	var result int
	curr := 50
	for _, line := range helper.ReadLines(path) {
		dir := line[0]
		distance, _ := strconv.Atoi(line[1:])

		result += distance / 100
		distance %= 100

		switch dir {
		case 'L':
			if curr > 0 && curr-distance <= 0 {
				result++
			}

			curr -= distance
			if curr < 0 {
				curr += 100
			}

		case 'R':
			result += (curr + distance) / 100
			curr = (curr + distance) % 100

		}
	}

	return result
}
