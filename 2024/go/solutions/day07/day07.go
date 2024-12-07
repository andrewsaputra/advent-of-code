package day07

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day07.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day07.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int64 {
	var result int64
	for _, line := range helper.ReadLines(path) {
		target, nums := toInputs(line)
		if isPossible(target, nums, 0, 0) {
			result += target
		}
	}

	return result
}

func solvePart2(path string) int64 {
	var result int64
	for _, line := range helper.ReadLines(path) {
		target, nums := toInputs(line)
		if isPossible2(target, nums, 0, 0) {
			result += target
		}
	}

	return result
}

func toInputs(line string) (target int64, nums []int64) {
	tmp := strings.Split(line, ": ")
	target, _ = strconv.ParseInt(tmp[0], 10, 64)

	nums = []int64{}
	for _, str := range strings.Fields(tmp[1]) {
		num, _ := strconv.ParseInt(str, 10, 64)
		nums = append(nums, num)
	}
	return
}

func isPossible(target int64, nums []int64, curr int64, idx int) bool {
	if idx == len(nums) || curr > target {
		return curr == target
	}

	return isPossible(target, nums, curr*nums[idx], idx+1) || isPossible(target, nums, curr+nums[idx], idx+1)
}

func isPossible2(target int64, nums []int64, curr int64, idx int) bool {
	if idx == len(nums) || curr > target {
		return curr == target
	}

	return isPossible2(target, nums, combineNums(curr, nums[idx]), idx+1) ||
		isPossible2(target, nums, curr*nums[idx], idx+1) ||
		isPossible2(target, nums, curr+nums[idx], idx+1)
}

func combineNums(a, b int64) int64 {
	digits := len(strconv.FormatInt(b, 10))
	return a*int64(math.Pow10(digits)) + b
}
