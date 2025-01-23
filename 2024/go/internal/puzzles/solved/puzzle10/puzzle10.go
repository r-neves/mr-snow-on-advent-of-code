package puzzle10

import (
	"aoc2024/internal/io"
	"fmt"
	"strconv"
)

func RunPart1() int {
	lines := io.ReadFile("inputs/10-real.txt")
	topography := parseIntMatrix(lines)

	sum := 0
	for i, row := range topography {
		for j := range row {
			visitedTops := make(map[string]bool)
			getHowManyTopsReached(topography, i, j, 0, visitedTops)
			sum += len(visitedTops)
		}
	}

	return sum
}

func getHowManyTopsReached(topography [][]int, i, j, nextLevel int, visitedTops map[string]bool) {
	if isOutOfBounds(topography, i, j) {
		return
	}

	if topography[i][j] != nextLevel {
		return
	}

	if topography[i][j] == 9 {
		visitedTops[fmt.Sprintf("%d-%d", i, j)] = true
		return
	}

	getHowManyTopsReached(topography, i+1, j, nextLevel+1, visitedTops)
	getHowManyTopsReached(topography, i, j+1, nextLevel+1, visitedTops)
	getHowManyTopsReached(topography, i-1, j, nextLevel+1, visitedTops)
	getHowManyTopsReached(topography, i, j-1, nextLevel+1, visitedTops)
}

func isOutOfBounds(topography [][]int, i, j int) bool {
	return i < 0 || j < 0 || i >= len(topography) || j >= len(topography[0])
}

func RunPart2() int {
	lines := io.ReadFile("inputs/10-real.txt")
	topography := parseIntMatrix(lines)

	sum := 0
	for i, row := range topography {
		for j := range row {
			sum += getAllUniqueTrails(topography, i, j, 0)
		}
	}

	return sum
}

func getAllUniqueTrails(topography [][]int, i, j, nextLevel int) int {
	if isOutOfBounds(topography, i, j) {
		return 0
	}

	if topography[i][j] != nextLevel {
		return 0
	}

	if topography[i][j] == 9 {
		return 1
	}

	return getAllUniqueTrails(topography, i+1, j, nextLevel+1) +
		getAllUniqueTrails(topography, i, j+1, nextLevel+1) +
		getAllUniqueTrails(topography, i-1, j, nextLevel+1) +
		getAllUniqueTrails(topography, i, j-1, nextLevel+1)
}

func parseIntMatrix(lines []string) [][]int {
	matrix := make([][]int, 0, len(lines))

	for _, line := range lines {
		matrixRow := make([]int, 0, len(line))
		for _, char := range line {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}

			matrixRow = append(matrixRow, n)
		}

		matrix = append(matrix, matrixRow)
	}

	return matrix
}
