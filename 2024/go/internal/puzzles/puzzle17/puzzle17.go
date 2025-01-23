package puzzle17

import (
	"aoc2024/internal/io"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type opcode int

const (
	adv opcode = 0
	blx opcode = 1
	bst opcode = 2
	jnz opcode = 3
	bxc opcode = 4
	out opcode = 5
	bdv opcode = 6
	cdv opcode = 7
)

func RunPart1() string {
	lines := io.ReadFile("inputs/17-real.txt")
	registerA, registerB, registerC, program := parseInput(lines)

	fmt.Println("Register A:", registerA)
	fmt.Println("Register B:", registerB)
	fmt.Println("Register C:", registerC)
	fmt.Println("Program:", program)

	_, _, _, output := runProgram(registerA, registerB, registerC, program)
	return output
}

func runProgram(registerA, registerB, registerC int, program []int) (int, int, int, string) {
	i := 0
	output := ""
	for i < len(program) {
		//fmt.Println(i)
		op := opcode(program[i])
		operand := program[i+1]
		switch op {
		case adv:
			comboOperand := getComboOperandValue(operand, registerA, registerB, registerC)
			result := math.Trunc(float64(registerA) / math.Pow(2, float64(comboOperand)))
			registerA = int(result)
		case blx:
			registerB = registerB ^ operand
		case bst:
			comboOperand := getComboOperandValue(operand, registerA, registerB, registerC)
			registerB = comboOperand % 8
		case jnz:
			if registerA == 0 {
				break
			}
			i = operand
		case bxc:
			registerB = registerB ^ registerC
		case out:
			comboOperand := getComboOperandValue(operand, registerA, registerB, registerC)
			output += strconv.Itoa(comboOperand%8) + ","
			//fmt.Println("Current output:", output)
		case bdv:
			comboOperand := getComboOperandValue(operand, registerA, registerB, registerC)
			result := math.Trunc(float64(registerA) / math.Pow(2, float64(comboOperand)))
			registerB = int(result)
		case cdv:
			comboOperand := getComboOperandValue(operand, registerA, registerB, registerC)
			result := math.Trunc(float64(registerA) / math.Pow(2, float64(comboOperand)))
			registerC = int(result)
		}

		if op != jnz || registerA == 0 {
			i += 2
		}
	}

	output = strings.TrimSuffix(output, ",")

	return registerA, registerB, registerC, output
}

func getComboOperandValue(operand, registerA, registerB, registerC int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registerA
	case 5:
		return registerB
	case 6:
		return registerC
	default:
		panic("Invalid operand")
	}
}

// Program 2,4,1,1,7,5,1,5,4,5,0,3,5,5,3,0
// To print the first 2, value in register B needs to be %8 == 2:
//   Sequence:
//     - 2,4: regB = regA % 8
//     - 1,1: regB = regB XOR 1
//     - 7,5: regC = regA / 2^regB
//     - 1,5: regB = regB XOR 5
//     - 4,5: regB = regB XOR regC
//     - 0,3: regA = regA / 8
//     - 5,5 (print regB % 8)
//     - 3,0 (loop)
//
// Reverse engineering the program:
// regA gets divided by 8 until it reaches 0,
// more specifically, 15 times.

var maxIndex = 0

func RunPart2() int {
	i := int64(math.Pow(8, 15))

	expected := []int{2, 4, 1, 1, 7, 5, 1, 5, 4, 5, 0, 3, 5, 5, 3, 0}

	found := false
	for !found {
		found = optimizedPart2(int(i), expected)
		i += 1
	}

	return int(i)
}

// Program 2,4,1,1,7,5,1,5,4,5,0,3,5,5,3,0
//
//	Sequence:
//	  - 2,4: regB = regA % 8
//	  - 1,1: regB = regB XOR 1
//	  - 7,5: regC = regA / 2^regB
//	  - 1,5: regB = regB XOR 5
//	  - 4,5: regB = regB XOR regC
//	  - 0,3: regA = regA / 8
//	  - 5,5 (print regB % 8)
//	  - 3,0 (loop)
func optimizedPart2(initialValue int, expected []int) bool {
	regA := initialValue
	for i := 0; i < len(expected); i++ {
		if maxIndex < i {
			maxIndex = i
			fmt.Println("Max index:", maxIndex)
			fmt.Println("Initial value:", initialValue)
			fmt.Println("Initial value binary:", strconv.FormatInt(int64(initialValue), 2))
		}

		regB := (regA % 8) ^ 1
		tmp := 1
		if regB > 0 {
			tmp = 2 << max(regB-1)
		}
		regC := int(math.Trunc(float64(regA) / float64(tmp)))
		regB = regB ^ 5 ^ regC
		regA = int(math.Trunc(float64(regA) / 8))

		res := regB % 8
		if res != expected[i] {
			return false
		}
	}

	return true
}

func parseInput(lines []string) (int, int, int, []int) {
	registerA, err := strconv.Atoi(strings.Split(lines[0], " ")[2])
	if err != nil {
		panic(err)
	}
	registerB, err := strconv.Atoi(strings.Split(lines[1], " ")[2])
	if err != nil {
		panic(err)
	}

	registerC, err := strconv.Atoi(strings.Split(lines[2], " ")[2])
	if err != nil {
		panic(err)
	}

	programStr := strings.Split(lines[4], " ")[1]
	split := strings.Split(programStr, ",")
	var program []int
	for _, num := range split {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		program = append(program, n)
	}

	return registerA, registerB, registerC, program
}
