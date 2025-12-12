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

func solvePart1(path string) int {
	nodes := make(map[string]Node)
	for _, line := range helper.ReadLines(path) {
		tmp := strings.Split(line, ": ")
		nodeID := tmp[0]
		next := strings.Fields(tmp[1])
		nodes[nodeID] = Node{ID: nodeID, Next: next}
	}

	var result int64
	visited := make(map[string]bool)
	dfs(&result, nodes, visited, "you", "out", nil)
	return int(result)
}

func solvePart2(path string) int64 {
	return -1

	nodes := make(map[string]Node)
	for _, line := range helper.ReadLines(path) {
		tmp := strings.Split(line, ": ")
		nodeID := tmp[0]
		next := strings.Fields(tmp[1])
		nodes[nodeID] = Node{ID: nodeID, Next: next}
	}

	var _, SvrToFft, _, DacToOut, FftToDac, _ int64

	//dfs(&SvrToDac, nodes, make(map[string]bool), "svr", "dac", nil)
	dfs(&SvrToFft, nodes, make(map[string]bool), "svr", "fft", nil)
	fmt.Println("SvrToFft", SvrToFft)
	//dfs(&DacToFft, nodes, make(map[string]bool), "dac", "fft", nil) // 0
	dfs(&DacToOut, nodes, make(map[string]bool), "dac", "out", nil) // 21653
	fmt.Println("DacToOut", DacToOut)
	dfs(&FftToDac, nodes, make(map[string]bool), "fft", "dac", nil)
	fmt.Println("FftToDac", FftToDac)
	//dfs(&FftToOut, nodes, make(map[string]bool), "fft", "out", nil)

	return SvrToFft * DacToOut * FftToDac
}

func dfs(result *int64, nodes map[string]Node, visited map[string]bool, currID string, targetID string, requiredIds []string) {
	if currID == targetID {
		for _, id := range requiredIds {
			if !visited[id] {
				return
			}
		}

		*result++
		//fmt.Println("found", *result)
		return
	}

	for _, nextID := range nodes[currID].Next {
		if visited[nextID] {
			continue
		}

		visited[nextID] = true
		dfs(result, nodes, visited, nextID, targetID, requiredIds)
		visited[nextID] = false
	}
}
