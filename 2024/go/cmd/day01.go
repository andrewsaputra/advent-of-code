package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res1 := solvePart1("inputs/day01.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day01.txt")
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solvePart1(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	nums1, nums2 := []int{}, []int{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

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
		result += abs(num1, num2)
	}
	return result
}

func solvePart2(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	nums := []int{}
	freqs := make(map[int]int)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

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

func abs(v1, v2 int) int {
	if v1 > v2 {
		return v1 - v2
	}
	return v2 - v1
}
