package day03

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day03.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day03.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	var result int
	for _, line := range helper.ReadLines(path) {
		pattern := `mul\((\d+),(\d+)\)`
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			result += num1 * num2
		}
	}

	return result
}

func solvePart2(path string) int {
	var result int
	isEnabled := true
	for _, line := range helper.ReadLines(path) {
		pattern := `mul\((\d+),(\d+)\)`
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllStringSubmatch(line, -1)
		matchesIdx := regex.FindAllStringIndex(line, -1)
		if len(matches) == 0 {
			continue
		}

		matchIdx := 0
		lineIdx := 0
		for lineIdx < len(line)-len("don't()") && matchIdx < len(matches) {
			if strings.HasPrefix(line[lineIdx:], "do()") {
				isEnabled = true
			} else if strings.HasPrefix(line[lineIdx:], "don't()") {
				isEnabled = false
			} else {
				if lineIdx == matchesIdx[matchIdx][0] {
					if isEnabled {
						num1, _ := strconv.Atoi(matches[matchIdx][1])
						num2, _ := strconv.Atoi(matches[matchIdx][2])
						result += num1 * num2
					}

					matchIdx++
				}
			}

			lineIdx++
		}
	}

	return result
}
