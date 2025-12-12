package day12

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day12-test.txt")
	fmt.Println("Part 1:", res1)

	//res2 := solvePart2("inputs/day12-test.txt")
	//fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {

	for _, line := range helper.ReadLines(path) {
		fmt.Println(line)
	}

	return -1
}
