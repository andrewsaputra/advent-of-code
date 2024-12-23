package day23

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"slices"
	"sort"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day23.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day23.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	graph := toGraph(path)
	lanMap := map[string]bool{}
	for g1, connections := range graph {
		if !strings.HasPrefix(g1, "t") {
			continue
		}

		for i := 0; i < len(connections)-1; i++ {
			g2 := connections[i]
			for j := i + 1; j < len(connections); j++ {
				g3 := connections[j]
				if slices.Contains(graph[g2], g3) {
					key := cacheKey([]string{g1, g2, g3})
					lanMap[key] = true
				}
			}
		}
	}

	return len(lanMap)
}

func solvePart2(path string) string {
	graph := toGraph(path)

	var R, P, X []string
	for val := range graph {
		P = append(P, val)
	}

	cliques := [][]string{}
	BronKerbosch(graph, R, P, X, &cliques)

	var result string
	var maxLen int
	for _, res := range cliques {
		if len(res) <= maxLen {
			continue
		}

		maxLen = len(res)
		result = cacheKey(res)
	}

	return result
}

type Item struct {
	ConnArr []string
	ConnMap map[string]bool
}

func toGraph(filepath string) map[string][]string {
	relations := make(map[string][]string)
	for _, line := range helper.ReadLines(filepath) {
		tmp := strings.Split(line, "-")
		g1, g2 := tmp[0], tmp[1]

		relations[g1] = append(relations[g1], g2)
		relations[g2] = append(relations[g2], g1)
	}

	return relations
}

func cacheKey(graphs []string) string {
	sort.Strings(graphs)
	return strings.Join(graphs, ",")
}

func arrIntersect(strs1, strs2 []string) []string {
	m := make(map[string]bool)
	for _, str := range strs1 {
		m[str] = true
	}

	var result []string
	for _, str := range strs2 {
		if m[str] {
			result = append(result, str)
		}
	}
	return result
}

func arrRemove(strs []string, val string) []string {
	var result []string
	for _, str := range strs {
		if str != val {
			result = append(result, str)
		}
	}
	return result
}

/*
Reference : https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
- R = current clique being explored
- P = set of candidates to be added to current clique
- X = already processed
*/
func BronKerbosch(graph map[string][]string, R, P, X []string, cliques *[][]string) {
	if len(P) == 0 && len(X) == 0 {
		*cliques = append(*cliques, append([]string{}, R...))
		return
	}

	for _, v := range P {
		neighbors := graph[v]

		BronKerbosch(
			graph,
			append(R, v),
			arrIntersect(P, neighbors),
			arrIntersect(X, neighbors),
			cliques,
		)

		P = arrRemove(P, v)
		X = append(X, v)
	}

}
