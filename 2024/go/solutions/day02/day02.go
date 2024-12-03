package day02

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day02.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day02.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	var result int
	for _, line := range helper.ReadLines(path) {
		if isSafe(helper.StringToInts(line)) {
			result++
		}
	}
	return result
}

func solvePart2(path string) int {
	var result int
	for _, line := range helper.ReadLines(path) {
		nums := helper.StringToInts(line)

		if isSafe(nums) {
			result++
		} else {
			for i := range nums {
				nums2 := append([]int{}, nums[:i]...)
				nums2 = append(nums2, nums[i+1:]...)
				if isSafe(nums2) {
					result++
					break
				}
			}
		}
	}

	return result
}

func isValid(prev int, curr int, increasing bool) bool {
	if (increasing && curr <= prev) || (!increasing && curr >= prev) {
		return false
	}

	var diff int
	if increasing {
		diff = curr - prev
	} else {
		diff = prev - curr
	}

	return diff >= 1 && diff <= 3
}

func isSafe(nums []int) bool {
	increasing := nums[1] > nums[0]
	for i := 1; i < len(nums); i++ {
		prev, curr := nums[i-1], nums[i]
		if !isValid(prev, curr, increasing) {
			return false
		}
	}

	return true
}
