package puzzle4

import (
	"aoc2024/internal/io"
	"fmt"
)

const xmas = "XMAS"

func RunPart1() {
	lines := io.ReadFile("inputs/4-real.txt")

	matrix := make([][]uint8, 0, len(lines))
	for _, line := range lines {
		matrix = append(matrix, []byte(line))
	}

	sum := 0

	for i, row := range matrix {
		for j := range row {
			sum += findXmasWordForDirection(matrix, i, j, 1, 0, 0)
			sum += findXmasWordForDirection(matrix, i, j, 1, 1, 0)
			sum += findXmasWordForDirection(matrix, i, j, 1, -1, 0)
			sum += findXmasWordForDirection(matrix, i, j, 0, 1, 0)
			sum += findXmasWordForDirection(matrix, i, j, 0, -1, 0)
			sum += findXmasWordForDirection(matrix, i, j, -1, 0, 0)
			sum += findXmasWordForDirection(matrix, i, j, -1, 1, 0)
			sum += findXmasWordForDirection(matrix, i, j, -1, -1, 0)
		}
	}

	fmt.Println(sum)
}

func findXmasWordForDirection(matrix [][]uint8, i, j, iDirection, jDirection, wordIndex int) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
		return 0
	}

	if matrix[i][j] != xmas[wordIndex] {
		return 0
	}

	if wordIndex == len(xmas)-1 {
		return 1
	}

	wordIndex++
	i += iDirection
	j += jDirection

	return findXmasWordForDirection(matrix, i, j, iDirection, jDirection, wordIndex)
}

func RunPart2() {
	lines := io.ReadFile("inputs/4-real.txt")

	matrix := make([][]uint8, 0, len(lines))
	for _, line := range lines {
		matrix = append(matrix, []byte(line))
	}

	sum := 0

	for i, row := range matrix {
		for j := range row {
			sum += findMASCross(matrix, i, j)
		}
	}

	fmt.Println(sum)
}

func findMASCross(matrix [][]uint8, i, j int) int {
	if !isLetter(matrix, i, j, 'A') {
		return 0
	}

	masCount := 0

	for iDirection := -1; iDirection <= 1; iDirection += 2 {
		for jDirection := -1; jDirection <= 1; jDirection += 2 {
			if isLetter(matrix, i+iDirection, j+jDirection, 'M') &&
				isLetter(matrix, i-iDirection, j-jDirection, 'S') {
				masCount++
			}
		}
	}

	if masCount == 2 {
		return 1
	}

	return 0
}

func isLetter(matrix [][]uint8, i, j int, letter uint8) bool {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
		return false
	}

	return matrix[i][j] == letter
}
