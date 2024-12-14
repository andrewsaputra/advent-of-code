package day13

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day13.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day13-test.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int64 {
	var result int64
	var ax, ay, bx, by, px, py int64
	for idx, line := range helper.ReadLines(path) {
		switch idx % 3 {
		case 0:
			ax, ay = parsePos(line[10:], "+")
		case 1:
			bx, by = parsePos(line[10:], "+")
		case 2:
			px, py = parsePos(line[7:], "=")

			res := int64(math.MaxInt32)
			var pressA int64
			for pressA < 100 {
				diffX := px - (pressA * ax)
				diffY := py - (pressA * ay)

				divX, modX := diffX/bx, diffX%bx
				divY, modY := diffY/by, diffY%by

				if divX == divY && modX == 0 && modY == 0 {
					res = min(res, pressA*3+divX)
				}

				pressA++
			}
			if res != math.MaxInt32 {
				result += res
			}
		}
	}
	return result
}

func solvePart2(path string) int64 {
	return 0

	var result int64
	var ax, ay, bx, by, px, py int64
	for idx, line := range helper.ReadLines(path) {
		switch idx % 3 {
		case 0:
			ax, ay = parsePos(line[10:], "+")
			//fmt.Println(ax, ay)
		case 1:
			bx, by = parsePos(line[10:], "+")
			//fmt.Println(bx, by)
		case 2:
			px, py = parsePos(line[7:], "=")
			px += 10000000000000
			py += 10000000000000
			//fmt.Println(px, py)

			//lcmX, lcmY := LCM(ax, bx), LCM(ay, by)
			//fmt.Println(lcmX, lcmY)

			res := int64(math.MaxInt64)
			var pressA int64
			for {
				diffX := px - (pressA * ax)
				diffY := py - (pressA * ay)

				if diffX < 0 || diffY < 0 {
					break
				}

				divX, modX := diffX/bx, diffX%bx
				divY, modY := diffY/by, diffY%by

				if divX == divY && modX == 0 && modY == 0 {
					res = min(res, pressA*3+divX)
				}

				pressA++
			}
			if res != math.MaxInt64 {
				result += res
			}
		}
	}
	return result
}

func parsePos(str string, separator string) (x, y int64) {
	for i, val := range strings.Split(str, ", ") {
		tmp := strings.Split(val, separator)
		if i == 0 {
			x, _ = strconv.ParseInt(tmp[1], 10, 64)
		} else {
			y, _ = strconv.ParseInt(tmp[1], 10, 64)
		}
	}
	return
}

func LCM(a int64, b int64) int64 { //least common multiple
	gcd := func(a, b int64) int64 { //general common divisor
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	return a * b / gcd(a, b)
}
