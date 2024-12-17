package day17

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day17.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day17-test.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) string {
	registers, program := toInputs(path)

	var outputs []string
	for i := 0; i < len(program)-1; i += 2 {
		opcode, operand := program[i], program[i+1]
		switch opcode {
		case 0:
			opCode0(registers, operand)
		case 1:
			opCode1(registers, operand)
		case 2:
			opCode2(registers, operand)
		case 3:
			jumpTarget := opCode3(registers, operand)
			if jumpTarget != -1 {
				i = jumpTarget - 2
			}
		case 4:
			opCode4(registers, operand)
		case 5:
			outputs = append(outputs, strconv.Itoa(opCode5(registers, operand)))
		case 6:
			opCode6(registers, operand)
		case 7:
			opCode7(registers, operand)
		}
	}

	return strings.Join(outputs, ",")
}

func solvePart2(path string) int {
	return 0
}

func toInputs(path string) (registers []int, program []int) {
	lines := helper.ReadLines(path)

	registers = make([]int, 3)

	fmt.Sscanf(lines[0], "Register A: %d", &registers[0])
	fmt.Sscanf(lines[1], "Register B: %d", &registers[1])
	fmt.Sscanf(lines[2], "Register C: %d", &registers[2])

	var strProgram string
	fmt.Sscanf(lines[3], "Program: %s", &strProgram)

	for _, str := range strings.Split(strProgram, ",") {
		num, _ := strconv.Atoi(str)
		program = append(program, num)
	}
	return
}

func opCode0(registers []int, operand int) {
	operandValue := comboOperandValue(registers, operand)
	denominator := int(math.Exp2(float64(operandValue)))
	registers[0] /= denominator
}

func opCode1(registers []int, operand int) {
	registers[1] ^= operand
}

func opCode2(registers []int, operand int) {
	operandValue := comboOperandValue(registers, operand)
	registers[1] = operandValue % 8
}

func opCode3(registers []int, operand int) (jumpTarget int) {
	if registers[0] == 0 {
		jumpTarget = -1
	} else {
		jumpTarget = operand
	}
	return
}

func opCode4(registers []int, operand int) {
	registers[1] ^= registers[2]
	return
}

func opCode5(registers []int, operand int) (output int) {
	operandValue := comboOperandValue(registers, operand)
	output = operandValue % 8
	return
}

func opCode6(registers []int, operand int) {
	operandValue := comboOperandValue(registers, operand)
	denominator := int(math.Exp2(float64(operandValue)))
	registers[1] = registers[0] / denominator
}

func opCode7(registers []int, operand int) {
	operandValue := comboOperandValue(registers, operand)
	denominator := int(math.Exp2(float64(operandValue)))
	registers[2] = registers[0] / denominator
}

func comboOperandValue(registers []int, operand int) int {
	switch operand {
	case 4:
		return registers[0]
	case 5:
		return registers[1]
	case 6:
		return registers[2]
	default:
		return operand
	}
}
