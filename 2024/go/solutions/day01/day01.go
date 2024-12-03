package day01

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day01.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day01.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	nums1, nums2 := []int{}, []int{}
	for _, line := range helper.ReadLines(path) {
		numbers := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	var result int
	for idx, num1 := range nums1 {
		num2 := nums2[idx]
		result += helper.AbsDiff(num1, num2)
	}
	return result
}

func solvePart2(path string) int {
	nums := []int{}
	freqs := make(map[int]int)
	for _, line := range helper.ReadLines(path) {
		numbers := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		nums = append(nums, num1)
		freqs[num2]++
	}

	var result int
	for _, num1 := range nums {
		result += num1 * freqs[num1]
	}
	return result
}
