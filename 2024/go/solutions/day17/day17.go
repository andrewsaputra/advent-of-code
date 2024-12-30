package day17

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day17.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day17.txt")
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
				i = int(jumpTarget - 2)
			}
		case 4:
			opCode4(registers, operand)
		case 5:
			outputs = append(outputs, strconv.FormatInt(opCode5(registers, operand), 10))
		case 6:
			opCode6(registers, operand)
		case 7:
			opCode7(registers, operand)
		}
	}

	return strings.Join(outputs, ",")
}

func solvePart2(path string) int64 {
	return 0
	//originalRegisters, program := toInputs(path)

	var result int64

	return result
}

func toInputs(path string) (registers []int64, program []int64) {
	lines := helper.ReadLines(path)

	registers = make([]int64, 3)

	fmt.Sscanf(lines[0], "Register A: %d", &registers[0])
	fmt.Sscanf(lines[1], "Register B: %d", &registers[1])
	fmt.Sscanf(lines[2], "Register C: %d", &registers[2])

	var strProgram string
	fmt.Sscanf(lines[3], "Program: %s", &strProgram)

	for _, str := range strings.Split(strProgram, ",") {
		num, _ := strconv.ParseInt(str, 10, 64)
		program = append(program, num)
	}
	return
}

func opCode0(registers []int64, operand int64) {
	operandValue := comboOperandValue(registers, operand)
	//denominator := int64(math.Exp2(float64(operandValue)))
	//registers[0] /= denominator

	registers[0] >>= operandValue
}

func opCode1(registers []int64, operand int64) {
	registers[1] ^= operand
}

func opCode2(registers []int64, operand int64) {
	operandValue := comboOperandValue(registers, operand)
	registers[1] = operandValue % 8
}

func opCode3(registers []int64, operand int64) (jumpTarget int64) {
	if registers[0] == 0 {
		jumpTarget = -1
	} else {
		jumpTarget = operand
	}
	return
}

func opCode4(registers []int64, operand int64) {
	registers[1] ^= registers[2]
}

func opCode5(registers []int64, operand int64) (output int64) {
	operandValue := comboOperandValue(registers, operand)
	output = operandValue % 8
	return
}

func opCode6(registers []int64, operand int64) {
	operandValue := comboOperandValue(registers, operand)
	//denominator := int64(math.Exp2(float64(operandValue)))
	//registers[1] = registers[0] / denominator

	registers[1] = registers[0] >> operandValue
}

func opCode7(registers []int64, operand int64) {
	operandValue := comboOperandValue(registers, operand)
	//denominator := int64(math.Exp2(float64(operandValue)))
	//registers[2] = registers[0] / denominator

	registers[2] = registers[0] >> operandValue
}

func comboOperandValue(registers []int64, operand int64) int64 {
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
