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
		point := findHorizontalMirrorPoint(strs)
		if point != -1 {
			return point * 100
		}

		return findVerticalMirrorPoint(strs)
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

func findHorizontalMirrorPoint(strs []string) int {
	m := len(strs)
	for i := 1; i < m; i++ {
		if strs[i] != strs[i-1] {
			continue
		}

		mirror := true

		prevIdx, forwardIdx := i-2, i+1
		for prevIdx >= 0 && forwardIdx < m {
			if strs[prevIdx] != strs[forwardIdx] {
				mirror = false
				break
			}

			prevIdx--
			forwardIdx++
		}

		if mirror {
			return i
		}

	}

	return -1
}

func findVerticalMirrorPoint(strs []string) int {
	m, n := len(strs), len(strs[0])

	strs2 := []string{}
	for col := 0; col < n; col++ {
		tmp := []byte{}

		for row := 0; row < m; row++ {
			tmp = append(tmp, strs[row][col])
		}
		strs2 = append(strs2, string(tmp))
	}

	return findHorizontalMirrorPoint(strs2)
}
