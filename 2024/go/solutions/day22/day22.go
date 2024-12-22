package day22

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"strconv"
)

func Solve() {
	res1 := solvePart1("inputs/day22.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day22-test.txt")
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

func solvePart2(path string) int64 {
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

	for _, tmp := range prices {
		fmt.Println(tmp)
	}

	return 0
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
