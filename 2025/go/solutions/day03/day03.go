package day03

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
	"math"
)

func Solve() {
	res1 := solvePart1("inputs/day03.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day03.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int64 {
	var result int64
	for _, line := range helper.ReadLines(path) {
		dp := make(map[string]int64)
		joltage := getJoltage(dp, line, 0, 2)
		result += joltage
	}

	return result
}

func solvePart2(path string) int64 {
	var result int64
	for _, line := range helper.ReadLines(path) {
		dp := make(map[string]int64)
		joltage := getJoltage(dp, line, 0, 12)
		result += joltage
	}

	return result
}

func getJoltage(dp map[string]int64, line string, startIdx int, kLeft int) int64 {
	if kLeft == 0 || startIdx >= len(line)-kLeft+1 {
		return 0
	}

	key := fmt.Sprintf("%d-%d", startIdx, kLeft)
	if res, ok := dp[key]; ok {
		return res
	}

	var prevDigit int64
	var res int64
	for i := startIdx; i < len(line)-kLeft+1; i++ {
		digit := int64(line[i] - '0')
		if digit <= prevDigit {
			continue
		}

		prevDigit = digit
		val := int64(float64(digit) * math.Pow(10, float64(kLeft-1)))
		for j := i + 1; j < len(line)-kLeft+2; j++ {
			res = max(res, val+getJoltage(dp, line, j, kLeft-1))
		}
	}

	dp[key] = res
	return res
}
