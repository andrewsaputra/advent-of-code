package day11

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solve("inputs/day11.txt", 25)
	fmt.Println("Part 1:", res1)

	res2 := solve("inputs/day11.txt", 75)
	fmt.Println("Part 2:", res2)
}

func solve(path string, maxBlink int) int64 {
	nums := toInput(path)
	freqsMap := make(map[int64]int64)
	for _, num := range nums {
		freqsMap[num]++
	}

	for count := 0; count < maxBlink; count++ {
		newFreqsMap := make(map[int64]int64)
		for num, freq := range freqsMap {
			for _, val := range blink(num) {
				newFreqsMap[val] += freq
			}
		}
		freqsMap = newFreqsMap
	}

	var result int64
	for _, val := range freqsMap {
		result += val
	}
	return result
}

func toInput(path string) []int64 {
	var result []int64
	for _, str := range strings.Fields(helper.ReadLines(path)[0]) {
		num, _ := strconv.ParseInt(str, 10, 64)
		result = append(result, num)
	}
	return result
}

func blink(num int64) []int64 {
	if num == 0 {
		return []int64{1}
	}

	str := strconv.FormatInt(num, 10)
	if len(str)%2 == 0 {
		mid := len(str) / 2
		numLeft, _ := strconv.ParseInt(str[:mid], 10, 64)
		numRight, _ := strconv.ParseInt(str[mid:], 10, 64)
		return []int64{numLeft, numRight}
	}

	return []int64{num * 2024}
}
