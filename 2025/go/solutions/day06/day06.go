package day06

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day06.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day06.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int64 {
	var nums [][]int64
	var signs []string
	for _, line := range helper.ReadLines(path) {
		if line[0] == '+' || line[0] == '*' {
			signs = strings.Fields(line)
			continue
		}

		var tmp []int64
		for _, str := range strings.Fields(line) {
			num, _ := strconv.ParseInt(str, 10, 64)
			tmp = append(tmp, num)
		}

		nums = append(nums, tmp)
	}

	numRows, numCols := len(nums), len(nums[0])

	var result int64
	for col := range numCols {
		sign := signs[col]
		tmp := nums[0][col]
		for row := range numRows {
			val := nums[row][col]
			if row == 0 {
				tmp = val
				continue
			}

			switch sign {
			case "+":
				tmp += val
			case "*":
				tmp *= val
			}
		}

		result += tmp
	}

	return result
}

func solvePart2(path string) int64 {
	var nums [][]byte
	var signs []string
	for _, line := range helper.ReadLines(path) {
		if line[0] == '+' || line[0] == '*' {
			signs = strings.Fields(line)
			break
		}

		nums = append(nums, []byte(line))
	}

	numRows, numCols := len(nums), len(nums[0])
	signIdx := len(signs) - 1
	var result int64
	var tmpResult int64
	for col := numCols - 1; col >= 0; col-- {
		var tmpVal int64
		for row := range numRows {
			val := nums[row][col]
			if val == ' ' {
				continue
			}

			tmpVal = tmpVal*10 + int64(val-'0')
		}

		if tmpVal == 0 {
			signIdx--
			result += tmpResult
			tmpResult = 0
			continue
		}

		if tmpResult == 0 {
			tmpResult = tmpVal
		} else {
			switch signs[signIdx] {
			case "+":
				tmpResult += tmpVal
			case "*":
				tmpResult *= tmpVal
			}
		}
	}

	result += tmpResult
	return result
}
