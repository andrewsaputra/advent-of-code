package day11

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day11.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day11.txt")
	fmt.Println("Part 2:", res2)
}

type Node struct {
	ID   string
	Next []string
}

func solvePart1(path string) int64 {
	nodes := make(map[string]Node)
	for _, line := range helper.ReadLines(path) {
		tmp := strings.Split(line, ": ")
		nodeID := tmp[0]
		next := strings.Fields(tmp[1])
		nodes[nodeID] = Node{ID: nodeID, Next: next}
	}

	dp := make(map[string]int64)
	return dfs(nodes, dp, nil, "you", "out")
}

func solvePart2(path string) int64 {
	nodes := make(map[string]Node)
	for _, line := range helper.ReadLines(path) {
		tmp := strings.Split(line, ": ")
		nodeID := tmp[0]
		next := strings.Fields(tmp[1])
		nodes[nodeID] = Node{ID: nodeID, Next: next}
	}

	dp := make(map[string]int64)
	return dfs(nodes, dp, []string{"dac", "fft"}, "svr", "out")
}

func dfs(nodes map[string]Node, dp map[string]int64, required []string, currID string, targetID string) int64 {
	key := fmt.Sprintf("%s-%s", currID, required)
	if res, ok := dp[key]; ok {
		return res
	}

	if currID == targetID {
		if len(required) == 0 {
			return 1
		}

		return 0
	}

	for idx, id := range required {
		if currID == id {
			tmp := append([]string{}, required[:idx]...)
			tmp = append(tmp, required[idx+1:]...)
			required = tmp
			defer func() {
				required = append(required, id)
			}()
			break
		}
	}

	var result int64
	for _, nextID := range nodes[currID].Next {
		result += dfs(nodes, dp, required, nextID, targetID)
	}

	dp[key] = result
	return result
}
