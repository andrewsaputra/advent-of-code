package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	bricks := parseInput("inputs.txt")
	countFallenBricks(bricks, -1, true)

	var res1, res2 int
	for i := range bricks {
		numFallen := countFallenBricks(bricks, i, false)
		res2 += numFallen
		if numFallen == 0 {
			res1++
		}
	}

	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type Brick struct {
	x, y, z    int
	dx, dy, dz int
}

func parseInput(input string) []Brick {
	file, err := os.Open(input)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	bricks := []Brick{}
	for scanner.Scan() {
		line := scanner.Text()
		var brick Brick
		var x2, y2, z2 int
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &brick.x, &brick.y, &brick.z, &x2, &y2, &z2)
		brick.dx = x2 - brick.x
		brick.dy = y2 - brick.y
		brick.dz = z2 - brick.z
		bricks = append(bricks, brick)
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].z < bricks[j].z
	})

	return bricks
}

func isColliding(a, b Brick) bool {
	return a.x <= b.x+b.dx && a.x+a.dx >= b.x && a.y <= b.y+b.dy && a.y+a.dy >= b.y
}

func countFallenBricks(bricks []Brick, idxRemove int, applyRemoval bool) int {
	copy := slices.Clone(bricks)
	total := 0
	for i := idxRemove + 1; i < len(copy); i++ {
		a := copy[i]
		newZ := 1
		for j := i - 1; j >= 0; j-- {
			b := copy[j]
			if j == idxRemove || b.z+b.dz < newZ {
				continue
			}

			if isColliding(a, b) {
				newZ = max(newZ, b.z+b.dz+1)
			}
		}

		if a.z != newZ {
			total++
			copy[i].z = newZ
		}
	}

	if applyRemoval {
		sort.Slice(copy, func(i, j int) bool {
			return copy[i].z < copy[j].z
		})

		for i, brick := range copy {
			bricks[i] = brick
		}
	}

	return total
}

func isSafeToRemove(bricks []Brick, idxRemove int) bool {
	for i := idxRemove + 1; i < len(bricks); i++ {
		a := bricks[i]
		newZ := 1
		for j := i - 1; j >= 0; j-- {
			b := bricks[j]
			if j == idxRemove || b.z+b.dz < newZ {
				continue
			}

			if isColliding(a, b) {
				newZ = max(newZ, b.z+b.dz+1)
			}
		}

		if a.z != newZ {
			return false
		}
	}

	return true
}

func visualizeBricks(bricks []Brick) {
	largestX, largestY, largestZ := -1, -1, -1
	for _, brick := range bricks {
		x2 := brick.x + brick.dx
		y2 := brick.y + brick.dy
		z2 := brick.z + brick.dz

		if x2 > largestX {
			largestX = x2
		}
		if y2 > largestY {
			largestY = y2
		}
		if z2 > largestZ {
			largestZ = z2
		}
	}

	numRows := largestZ + 1
	numCols := largestX + 1
	gridXZ := make([][]byte, numRows)
	for row := range gridXZ {
		gridXZ[row] = make([]byte, numCols)
		for col := range gridXZ[row] {
			if row == 0 {
				gridXZ[row][col] = '-'
			} else {
				gridXZ[row][col] = '.'
			}
		}
	}

	for idx, brick := range bricks {
		for row := brick.z; row <= brick.z+brick.dz; row++ {
			for col := brick.x; col <= brick.x+brick.dx; col++ {
				if gridXZ[row][col] == '.' {
					gridXZ[row][col] = byte(idx + '0')
				} else {
					gridXZ[row][col] = '?'
				}

			}
		}
	}

	numCols = largestY + 1
	gridYZ := make([][]byte, numRows)
	for row := range gridYZ {
		gridYZ[row] = make([]byte, numCols)
		for col := range gridYZ[row] {
			if row == 0 {
				gridYZ[row][col] = '-'
			} else {
				gridYZ[row][col] = '.'
			}
		}
	}

	for idx, brick := range bricks {
		for row := brick.z; row <= brick.z+brick.dz; row++ {
			for col := brick.y; col <= brick.y+brick.dy; col++ {
				if gridYZ[row][col] == '.' {
					gridYZ[row][col] = byte(idx + '0')
				} else {
					gridYZ[row][col] = '?'
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("== Grid XZ ==")
	for row := len(gridXZ) - 1; row >= 0; row-- {
		fmt.Println(string(gridXZ[row]))
	}

	fmt.Println()
	fmt.Println("== Grid YZ ==")
	for row := len(gridYZ) - 1; row >= 0; row-- {
		fmt.Println(string(gridYZ[row]))
	}
}
