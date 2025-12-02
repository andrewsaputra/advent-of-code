package day02

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day02.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day02.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int64 {
	line := helper.ReadLines(path)[0]

	var result int64
	for _, numRange := range strings.Split(line, ",") {
		tmp := strings.Split(numRange, "-")
		start, _ := strconv.Atoi(tmp[0])
		end, _ := strconv.Atoi(tmp[1])

		for num := start; num <= end; num++ {
			strNum := strconv.Itoa(num)
			mid := len(strNum) / 2
			if strNum[:mid] == strNum[mid:] {
				result += int64(num)
			}
		}
	}

	return result
}

func solvePart2(path string) int64 {
	line := helper.ReadLines(path)[0]

	var result int64
	for _, numRange := range strings.Split(line, ",") {
		tmp := strings.Split(numRange, "-")
		start, _ := strconv.Atoi(tmp[0])
		end, _ := strconv.Atoi(tmp[1])

		for num := start; num <= end; num++ {
			strNum := strconv.Itoa(num)
			strLen := len(strNum)

			//check for all possible repeating lengths
			halfLen := strLen / 2
			for targetLen := 1; targetLen <= halfLen; targetLen++ {
				if strLen%targetLen != 0 {
					continue
				}

				invalidNumber := true
				base := strNum[:targetLen]
				for idx := 1; idx*targetLen < strLen; idx++ {
					if strNum[idx*targetLen:(idx+1)*targetLen] != base {
						invalidNumber = false
						break
					}
				}

				if invalidNumber {
					result += int64(num)
					break
				}
			}
		}
	}

	return result
}
