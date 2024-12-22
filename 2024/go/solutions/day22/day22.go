package day22

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"strconv"
	"sync"
)

func Solve() {
	res1 := solvePart1("inputs/day22.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day22.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int64 {
	var result int64
	for _, str := range helper.ReadLines(path) {
		num, _ := strconv.Atoi(str)
		for i := 0; i < 2000; i++ {
			num = nextSecretNumber(num)
		}

		result += int64(num)
	}

	return result
}

func solvePart2(path string) int {
	prices := [][]int{}
	for _, str := range helper.ReadLines(path) {
		num, _ := strconv.Atoi(str)
		price := []int{num % 10}
		for i := 0; i < 2000; i++ {
			num = nextSecretNumber(num)
			price = append(price, num%10)
		}
		prices = append(prices, price)
	}

	patterns := make(map[string]bool)
	resultMap := []map[string]int{}
	for _, price := range prices {
		diff := []int{}
		resMap := make(map[string]int)
		for i := 1; i < len(price); i++ {
			diff = append(diff, price[i]-price[i-1])
			for len(diff) > 4 {
				diff = diff[1:]
			}

			if i >= 4 {
				key := cacheKey(diff)
				if _, ok := resMap[key]; ok {
					continue
				}

				patterns[key] = true
				resMap[key] = price[i]
			}
		}
		resultMap = append(resultMap, resMap)
	}

	return getMaxValue(patterns, resultMap)
}

const mod = 16777216

func nextSecretNumber(num int) int {
	num ^= (num * 64)
	num %= mod

	num ^= (num / 32)
	num %= mod

	num ^= (num * 2048)
	num %= mod

	return num
}

func cacheKey(nums []int) string {
	return fmt.Sprintf("%d,%d,%d,%d", nums[0], nums[1], nums[2], nums[3])
}

func getMaxValue(patterns map[string]bool, resultMap []map[string]int) int {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var maxResult int

	process := func(key string) {
		defer wg.Done()

		var tmpRes int
		for _, data := range resultMap {
			if val, ok := data[key]; ok {
				tmpRes += val
			}
		}

		mutex.Lock()
		maxResult = max(maxResult, tmpRes)
		mutex.Unlock()
	}

	for key := range patterns {
		wg.Add(1)
		go process(key)
	}

	return maxResult
}
