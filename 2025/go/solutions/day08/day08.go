package day08

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
	"math"
	"sort"
)

func Solve() {
	res1 := solvePart1("inputs/day08.txt", 1000)
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day08.txt", 1000)
	fmt.Println("Part 2:", res2)
}

type Box struct {
	ID int
	X  int
	Y  int
	Z  int
}

type Distance struct {
	Distance int
	ID1      int
	ID2      int
}

func solvePart1(path string, targetShortest int) int {
	var boxes []Box
	var id int
	for _, line := range helper.ReadLines(path) {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		boxes = append(boxes, Box{ID: id, X: x, Y: y, Z: z})
		id++
	}

	var distances []Distance
	for i := 0; i < len(boxes)-1; i++ {
		box1 := boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			box2 := boxes[j]
			dist := calcDistance(box1, box2)
			distances = append(distances, Distance{Distance: dist, ID1: box1.ID, ID2: box2.ID})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	id = 0
	circuitMap := make(map[int]map[int]bool)
	for i := range targetShortest {
		distance := distances[i]
		id1, id2 := distance.ID1, distance.ID2

		var existsInCircuitMap bool
		for circuitID, circuitData := range circuitMap {
			if circuitData[id1] || circuitData[id2] {
				circuitMap[circuitID][id1] = true
				circuitMap[circuitID][id2] = true
				existsInCircuitMap = true
			}
		}

		if !existsInCircuitMap {
			id++
			circuitMap[id] = map[int]bool{id1: true, id2: true}
		}
	}

	mergeCircuitMap(&circuitMap)

	var circuitSizes []int
	for _, circuitData := range circuitMap {
		circuitSizes = append(circuitSizes, len(circuitData))
	}

	sort.Slice(circuitSizes, func(i, j int) bool {
		return circuitSizes[i] >= circuitSizes[j]
	})

	result := 1
	for i := range 3 {
		result *= circuitSizes[i]
	}

	return result
}

func calcDistance(box1 Box, box2 Box) int {
	return int(math.Pow(float64(helper.AbsDiff(box1.X, box2.X)), 2) +
		math.Pow(float64(helper.AbsDiff(box1.Y, box2.Y)), 2) +
		math.Pow(float64(helper.AbsDiff(box1.Z, box2.Z)), 2))
}

func solvePart2(path string, targetShortest int) int {
	var boxes []Box
	var id int
	for _, line := range helper.ReadLines(path) {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		boxes = append(boxes, Box{ID: id, X: x, Y: y, Z: z})
		id++
	}

	var distances []Distance
	for i := 0; i < len(boxes)-1; i++ {
		box1 := boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			box2 := boxes[j]
			dist := calcDistance(box1, box2)
			distances = append(distances, Distance{Distance: dist, ID1: box1.ID, ID2: box2.ID})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	id = 0
	circuitMap := make(map[int]map[int]bool)
	for i := range distances {
		distance := distances[i]
		id1, id2 := distance.ID1, distance.ID2

		var existsInCircuitMap bool
		for circuitID, circuitData := range circuitMap {
			if circuitData[id1] || circuitData[id2] {
				circuitMap[circuitID][id1] = true
				circuitMap[circuitID][id2] = true
				existsInCircuitMap = true
			}
		}

		if !existsInCircuitMap {
			id++
			circuitMap[id] = map[int]bool{id1: true, id2: true}
		}

		if i > targetShortest {
			mergeCircuitMap(&circuitMap)
			if isAllBoxesConnected(circuitMap, boxes) {
				return boxes[id1].X * boxes[id2].X
			}
		}
	}

	return -1
}

func mergeCircuitMap(circuitMap *map[int]map[int]bool) {
	for {
		var hasMerge bool

		for circuitID1, circuitData1 := range *circuitMap {
			for circuitID2, circuitData2 := range *circuitMap {
				if circuitID1 == circuitID2 {
					continue
				}

				for id1 := range circuitData1 {
					if circuitData2[id1] {
						for id2 := range circuitData2 {
							circuitData1[id2] = true
						}

						delete(*circuitMap, circuitID2)

						hasMerge = true
						break
					}
				}

				if hasMerge {
					break
				}
			}

			if hasMerge {
				break
			}
		}

		if !hasMerge {
			break
		}
	}
}

func isAllBoxesConnected(circuitMap map[int]map[int]bool, boxes []Box) bool {
	if len(circuitMap) != 1 {
		return false
	}

	for _, circuitData := range circuitMap {
		for _, box := range boxes {
			if !circuitData[box.ID] {
				return false
			}
		}
	}

	return true
}
