package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res := solve("inputs.txt")
	fmt.Println(res)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solve(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	calculate := func(strs []string) int {
		point := findSmudgedHorizontalMirrorPoint(strs)
		if point != -1 {
			return point * 100
		}

		return findSmudgedVerticalMirrorPoint(strs)
	}

	total := 0
	strs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			strs = append(strs, line)
			continue
		}

		if len(strs) > 0 {
			total += calculate(strs)
			strs = []string{}
		}
	}

	if len(strs) > 0 {
		total += calculate(strs)
	}

	return total
}

func findSmudgedHorizontalMirrorPoint(strs []string) int {
	countDiff := func(str1, str2 string) int {
		res := 0
		for i, v1 := range str1 {
			v2 := rune(str2[i])

			if v1 != v2 {
				res++
			}
		}

		return res
	}

	m := len(strs)
	for i := 1; i < m; i++ {
		numDiff := countDiff(strs[i], strs[i-1])
		if numDiff >= 2 {
			continue
		}

		mirror := true
		prevIdx, forwardIdx := i-2, i+1
		for prevIdx >= 0 && forwardIdx < m {
			numDiff += countDiff(strs[prevIdx], strs[forwardIdx])
			if numDiff >= 2 {
				mirror = false
				break
			}

			prevIdx--
			forwardIdx++
		}

		if mirror && numDiff == 1 {
			return i
		}
	}

	return -1
}

func findSmudgedVerticalMirrorPoint(strs []string) int {
	m, n := len(strs), len(strs[0])

	strs2 := []string{}
	for col := 0; col < n; col++ {
		tmp := []byte{}

		for row := 0; row < m; row++ {
			tmp = append(tmp, strs[row][col])
		}
		strs2 = append(strs2, string(tmp))
	}

	return findSmudgedHorizontalMirrorPoint(strs2)
}
