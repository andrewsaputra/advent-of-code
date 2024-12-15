package day15

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day15.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day15.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	matrix, movements := toInputs(path, false)
	guardRow, guardCol := findStartPos(matrix)
	for _, move := range movements {
		doMove(matrix, &guardRow, &guardCol, move)
	}

	var result int
	for row := range matrix {
		for col, val := range matrix[row] {
			if val != 'O' {
				continue
			}

			result += 100*row + col
		}
	}

	return result
}

func solvePart2(path string) int {
	matrix, movements := toInputs(path, true)
	guardRow, guardCol := findStartPos(matrix)
	for _, move := range movements {
		doWideMove(matrix, &guardRow, &guardCol, move)
	}

	var result int
	for row := range matrix {
		for col, val := range matrix[row] {
			if val != '[' {
				continue
			}

			result += 100*row + col
		}
	}

	return result
}

func toInputs(filepath string, wideMode bool) (matrix [][]byte, movements []byte) {
	matrix = make([][]byte, 0)
	movements = make([]byte, 0)

	var isMovements bool
	for _, line := range helper.ReadLines(filepath) {
		if !isMovements && !strings.HasPrefix(line, "#") {
			isMovements = true
		}

		if isMovements {
			movements = append(movements, []byte(line)...)
		} else {
			if wideMode {
				var sb strings.Builder
				for _, v := range line {
					switch v {
					case '#':
						sb.WriteString("##")
					case 'O':
						sb.WriteString("[]")
					case '.':
						sb.WriteString("..")
					case '@':
						sb.WriteString("@.")
					}
				}
				matrix = append(matrix, []byte(sb.String()))
			} else {
				matrix = append(matrix, []byte(line))
			}
		}
	}
	return
}

func findStartPos(matrix [][]byte) (int, int) {
	for row := range matrix {
		for col, val := range matrix[row] {
			if val == '@' {
				return row, col
			}
		}
	}
	return -1, -1
}

func printMap(matrix [][]byte) {
	for row := range matrix {
		fmt.Println(string(matrix[row]))
	}
	fmt.Println()
}

func doMove(matrix [][]byte, guardRow *int, guardCol *int, move byte) {
	nextRow := *guardRow
	nextCol := *guardCol
	switch move {
	case '<':
		nextCol--
	case '>':
		nextCol++
	case '^':
		nextRow--
	case 'v':
		nextRow++
	}

	switch matrix[nextRow][nextCol] {
	case '#':
		return
	case '.':
		matrix[*guardRow][*guardCol] = '.'
		*guardRow = nextRow
		*guardCol = nextCol
		matrix[*guardRow][*guardCol] = '@'
	case 'O':
		switch move {
		case '<':
			for col := nextCol; col >= 0; col-- {
				if matrix[*guardRow][col] == '.' {
					matrix[*guardRow][*guardCol] = '.'
					for col2 := col; col2 < nextCol; col2++ {
						matrix[*guardRow][col2] = matrix[*guardRow][col2+1]
					}

					*guardCol = nextCol
					matrix[*guardRow][*guardCol] = '@'
					return
				} else if matrix[*guardRow][col] == '#' {
					return
				}
			}
		case '>':
			for col := nextCol; col < len(matrix[0]); col++ {
				if matrix[*guardRow][col] == '.' {
					matrix[*guardRow][*guardCol] = '.'
					for col2 := col; col2 > nextCol; col2-- {
						matrix[*guardRow][col2] = matrix[*guardRow][col2-1]
					}

					*guardCol = nextCol
					matrix[*guardRow][*guardCol] = '@'
					return
				} else if matrix[*guardRow][col] == '#' {
					return
				}
			}
		case '^':
			for row := nextRow; row >= 0; row-- {
				if matrix[row][*guardCol] == '.' {
					matrix[*guardRow][*guardCol] = '.'
					for row2 := row; row2 < nextRow; row2++ {
						matrix[row2][*guardCol] = matrix[row2+1][*guardCol]
					}

					*guardRow = nextRow
					matrix[*guardRow][*guardCol] = '@'
					return
				} else if matrix[row][*guardCol] == '#' {
					return
				}
			}
		case 'v':
			for row := nextRow; row < len(matrix); row++ {
				if matrix[row][*guardCol] == '.' {
					matrix[*guardRow][*guardCol] = '.'
					for row2 := row; row2 > nextRow; row2-- {
						matrix[row2][*guardCol] = matrix[row2-1][*guardCol]
					}

					*guardRow = nextRow
					matrix[*guardRow][*guardCol] = '@'
					return
				} else if matrix[row][*guardCol] == '#' {
					return
				}
			}
		}
	}
}

func doWideMove(matrix [][]byte, guardRow *int, guardCol *int, move byte) {
	nextRow := *guardRow
	nextCol := *guardCol
	switch move {
	case '<':
		nextCol--
	case '>':
		nextCol++
	case '^':
		nextRow--
	case 'v':
		nextRow++
	}

	switch matrix[nextRow][nextCol] {
	case '#':
		return
	case '.':
		matrix[*guardRow][*guardCol] = '.'
		*guardRow = nextRow
		*guardCol = nextCol
		matrix[*guardRow][*guardCol] = '@'
		return
	}

	switch move {
	case '<':
		for col := nextCol; col >= 0; col-- {
			if matrix[*guardRow][col] == '.' {
				matrix[*guardRow][*guardCol] = '.'
				for col2 := col; col2 < nextCol; col2++ {
					matrix[*guardRow][col2] = matrix[*guardRow][col2+1]
				}

				*guardCol = nextCol
				matrix[*guardRow][*guardCol] = '@'
				return
			} else if matrix[*guardRow][col] == '#' {
				return
			}
		}
	case '>':
		for col := nextCol; col < len(matrix[0]); col++ {
			if matrix[*guardRow][col] == '.' {
				matrix[*guardRow][*guardCol] = '.'
				for col2 := col; col2 > nextCol; col2-- {
					matrix[*guardRow][col2] = matrix[*guardRow][col2-1]
				}

				*guardCol = nextCol
				matrix[*guardRow][*guardCol] = '@'
				return
			} else if matrix[*guardRow][col] == '#' {
				return
			}
		}
	case '^', 'v':
		boxRow, boxCol := nextRow, nextCol
		if matrix[nextRow][nextCol] == ']' {
			boxCol--
		}

		validMove, boxes := canMoveVertical(matrix, boxRow, boxCol, move)
		if validMove {
			matrix[*guardRow][*guardCol] = '.'
			for idx := len(boxes) - 1; idx >= 0; idx-- {
				boxRow, boxCol := boxes[idx][0], boxes[idx][1]
				nextBoxRow := boxRow
				if move == '^' {
					nextBoxRow--
				} else {
					nextBoxRow++
				}
				matrix[nextBoxRow][boxCol] = '['
				matrix[nextBoxRow][boxCol+1] = ']'

				matrix[boxRow][boxCol] = '.'
				matrix[boxRow][boxCol+1] = '.'
			}

			*guardRow = nextRow
			matrix[*guardRow][*guardCol] = '@'
		}
	}
}

func canMoveVertical(matrix [][]byte, boxRow int, boxCol int, move byte) (bool, [][]int) {
	boxes := [][]int{}
	queue := [][]int{{boxRow, boxCol}}

	for len(queue) > 0 {
		var newQueue [][]int
		for _, item := range queue {
			boxes = append(boxes, []int{item[0], item[1]})
			nextRow := item[0]
			nextCol := item[1]
			if move == '^' {
				nextRow--
			} else {
				nextRow++
			}

			if matrix[nextRow][nextCol] == '#' || matrix[nextRow][nextCol+1] == '#' {
				return false, nil
			}

			if matrix[nextRow][nextCol] == '.' && matrix[nextRow][nextCol+1] == '.' {
				continue
			}

			if matrix[nextRow][nextCol] == '[' {
				newQueue = append(newQueue, []int{nextRow, nextCol})
			}

			if matrix[nextRow][nextCol] == ']' {
				newQueue = append(newQueue, []int{nextRow, nextCol - 1})
			}

			if matrix[nextRow][nextCol+1] == '[' {
				newQueue = append(newQueue, []int{nextRow, nextCol + 1})
			}
		}

		queue = newQueue
	}

	return true, boxes
}
