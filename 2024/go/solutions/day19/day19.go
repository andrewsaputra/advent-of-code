package day19

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day19.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day19.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	inputs := helper.ReadLines(path)
	towels := strings.Split(inputs[0], ", ")

	var result int
	for i := 1; i < len(inputs); i++ {
		dp := make(map[int]int)
		if possiblePatterns(towels, dp, 0, inputs[i]) > 0 {
			result++
		}
	}

	return result
}

func solvePart2(path string) int64 {
	inputs := helper.ReadLines(path)
	towels := strings.Split(inputs[0], ", ")

	var result int64
	for i := 1; i < len(inputs); i++ {
		dp := make(map[int]int)
		result += int64(possiblePatterns(towels, dp, 0, inputs[i]))
	}

	return result
}

func possiblePatterns(towels []string, dp map[int]int, idx int, pattern string) int {
	if idx >= len(pattern) {
		return 1
	}

	if val, ok := dp[idx]; ok {
		return val
	}

	var result int
	for _, twl := range towels {
		if !strings.HasPrefix(pattern[idx:], twl) {
			continue
		}

		result += possiblePatterns(towels, dp, idx+len(twl), pattern)
	}

	dp[idx] = result
	return result
}
