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

type Sign struct {
	Sign rune
	Idx  int
}

func solvePart2(path string) int64 {
	var nums [][]string
	var signs []Sign

	inputLines := helper.ReadLines(path)
	for idx, val := range inputLines[len(inputLines)-1] {
		if val == ' ' {
			continue
		}

		signs = append(signs, Sign{Sign: val, Idx: idx})
	}

	for i := 0; i < len(inputLines)-1; i++ {
		line := inputLines[i]

		var tmp []string
		for idx, sign := range signs {
			if idx == len(signs)-1 {
				tmp = append(tmp, line[sign.Idx:])
				continue
			}

			tmp = append(tmp, line[sign.Idx:signs[idx+1].Idx-1])
		}

		nums = append(nums, tmp)
	}

	numRows, numCols := len(nums), len(nums[0])

	var result int64
	for col := range numCols {
		sign := signs[col].Sign
		lenStr := len(nums[0][col])
		var tmpResult int64
		for i := lenStr - 1; i >= 0; i-- {
			var tmpVal int64
			for row := range numRows {
				val := nums[row][col][i]
				if val == ' ' {
					continue
				}

				tmpVal = tmpVal*10 + int64(val-'0')
			}

			if tmpResult == 0 {
				tmpResult = tmpVal
			} else {
				switch sign {
				case '+':
					tmpResult += tmpVal
				case '*':
					tmpResult *= tmpVal
				}
			}
		}

		result += tmpResult
	}

	return result
}
