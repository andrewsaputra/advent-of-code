package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day05.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day05.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	freshRanges, ids := parseInput(path)

	var result int
	for _, id := range ids {
		for _, data := range freshRanges {
			if id >= data[0] && id <= data[1] {
				result++
				break
			}
		}
	}

	return result
}

func solvePart2(path string) int64 {
	freshRanges, _ := parseInput(path)

	sort.Slice(freshRanges, func(i, j int) bool {
		if freshRanges[i][0] == freshRanges[j][0] {
			return freshRanges[i][1] < freshRanges[j][1]
		}

		return freshRanges[i][0] < freshRanges[j][0]
	})

	for i := 0; i < len(freshRanges)-1; i++ {
		curr, next := freshRanges[i], freshRanges[i+1]
		if next[0] <= curr[1] && next[1] >= curr[0] {
			next[1] = max(curr[1], next[1])
			next[0] = curr[0]
			freshRanges[i] = nil
		}
	}

	var result int64
	for _, data := range freshRanges {
		if data == nil {
			continue
		}

		result += data[1] - data[0] + 1
	}

	return result
}

func parseInput(path string) ([][]int64, []int64) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var freshRanges [][]int64
	var ids []int64
	var checkIds bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			checkIds = true
			continue
		}

		if checkIds {
			id, _ := strconv.ParseInt(line, 10, 64)
			ids = append(ids, id)
		} else {
			tmp := strings.Split(line, "-")
			start, _ := strconv.ParseInt(tmp[0], 10, 64)
			end, _ := strconv.ParseInt(tmp[1], 10, 64)
			freshRanges = append(freshRanges, []int64{start, end})
		}
	}

	return freshRanges, ids
}
