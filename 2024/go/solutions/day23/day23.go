package day23

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"sort"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day23.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day23-test.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	relations := toRelations(path)
	lanMap := map[string]bool{}
	for g1, item := range relations {
		if !strings.HasPrefix(g1, "t") {
			continue
		}

		for i := 0; i < len(item.ConnArr)-1; i++ {
			g2 := item.ConnArr[i]
			for j := i + 1; j < len(item.ConnArr); j++ {
				g3 := item.ConnArr[j]
				if relations[g2].ConnMap[g3] {
					key := cacheKey([]string{g1, g2, g3})
					lanMap[key] = true
				}
			}
		}
	}

	return len(lanMap)
}

func solvePart2(path string) int {
	return 0
}

type Item struct {
	ConnArr []string
	ConnMap map[string]bool
}

func toRelations(filepath string) map[string]*Item {
	relations := make(map[string]*Item)
	for _, line := range helper.ReadLines(filepath) {
		tmp := strings.Split(line, "-")
		g1, g2 := tmp[0], tmp[1]

		if item, ok := relations[g1]; ok {
			item.ConnArr = append(item.ConnArr, g2)
			item.ConnMap[g2] = true
		} else {
			relations[g1] = &Item{
				ConnArr: []string{g2},
				ConnMap: map[string]bool{g2: true},
			}
		}

		if item, ok := relations[g2]; ok {
			item.ConnArr = append(item.ConnArr, g1)
			item.ConnMap[g1] = true
		} else {
			relations[g2] = &Item{
				ConnArr: []string{g1},
				ConnMap: map[string]bool{g1: true},
			}
		}
	}

	return relations
}

func cacheKey(graphs []string) string {
	sort.Strings(graphs)
	return strings.Join(graphs, ",")
}
