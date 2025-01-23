package puzzle7

import (
	"aoc2024/internal/io"
	"strconv"
	"strings"
)

type equation struct {
	result   int
	operands []int
}

func RunPart1() int {
	lines := io.ReadFile("inputs/7-real.txt")

	sum := 0
	for _, line := range lines {
		eq := parseEquation(line)
		if isPossibleEquation1(eq) {
			sum += eq.result
		}
	}

	return sum
}

func parseEquation(line string) equation {
	parts := strings.Split(line, ":")
	result, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	var operands []int
	parts[1] = strings.TrimSpace(parts[1])
	for _, operand := range strings.Split(parts[1], " ") {
		value, err := strconv.Atoi(operand)
		if err != nil {
			panic(err)
		}

		operands = append(operands, value)
	}

	return equation{
		result:   result,
		operands: operands,
	}
}

func isPossibleEquation1(equation equation) bool {
	acc := equation.operands[0]
	return isCorrectResult1(acc, equation.operands[1:], equation.result)
}

func isCorrectResult1(acc int, operands []int, result int) bool {
	if len(operands) == 0 {
		return acc == result
	}

	nextOperand := operands[0]

	if isCorrectResult1(acc+nextOperand, operands[1:], result) {
		return true
	}

	if isCorrectResult1(acc*nextOperand, operands[1:], result) {
		return true
	}

	return false
}
func RunPart2() int {
	lines := io.ReadFile("inputs/7-real.txt")

	sum := 0
	for _, line := range lines {
		eq := parseEquation(line)
		if isPossibleEquation2(eq) {
			sum += eq.result
		}
	}

	return sum
}

func isPossibleEquation2(equation equation) bool {
	acc := equation.operands[0]
	return isCorrectResult2(acc, equation.operands[1:], equation.result)
}

func isCorrectResult2(acc int, operands []int, result int) bool {
	if len(operands) == 0 {
		return acc == result
	}

	nextOperand := operands[0]

	if isCorrectResult2(acc+nextOperand, operands[1:], result) {
		return true
	}

	if isCorrectResult2(acc*nextOperand, operands[1:], result) {
		return true
	}

	joined := concatInts(acc, nextOperand)
	if isCorrectResult2(joined, operands[1:], result) {
		return true
	}

	return false
}

func concatInts(a, b int) int {
	concatenated := strconv.Itoa(a) + strconv.Itoa(b)
	result, err := strconv.Atoi(concatenated)
	if err != nil {
		panic(err)
	}
	return result
}
