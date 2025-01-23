package puzzle2

import (
	"aoc2024/internal/io"
	"fmt"
	"strconv"
	"strings"
)

const (
	maxDiff = 3
	minDiff = 1

	ascIndex  = 0
	descIndex = 1
)

func RunPart1() {
	lines := io.ReadFile("inputs/2-real.txt")

	unsafeSum := 0

	for _, line := range lines {
		split := strings.Split(line, " ")
		numbers := toIntSlice(split)
		ascending := numbers[0] < numbers[1]

		for i := 0; i < len(numbers)-1; i++ {
			diff := numbers[i+1] - numbers[i]
			if diff < 0 {
				diff = -diff
			}

			if diff < minDiff || diff > maxDiff {
				unsafeSum++
				break
			}

			if ascending && numbers[i] > numbers[i+1] {
				unsafeSum++
				break
			} else if !ascending && numbers[i] < numbers[i+1] {
				unsafeSum++
				break
			}
		}
	}

	fmt.Println(len(lines) - unsafeSum)
}

func toIntSlice(strs []string) []int {
	var err error
	result := make([]int, len(strs))

	for i, str := range strs {
		result[i], err = strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
	}

	return result
}

func RunPart2() {
	lines := io.ReadFile("inputs/2-real.txt")

	validSum := 0

	for _, line := range lines {
		split := strings.Split(line, " ")
		numbers := toIntSlice(split)

		if isValidReportWithError(numbers) {
			validSum++
		}
	}

	fmt.Println(validSum)
}

func isValidReportWithError(numbers []int) bool {
	if isValidReport(numbers) {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		var subNumbers []int
		subNumbers = append(subNumbers, numbers[:i]...)
		subNumbers = append(subNumbers, numbers[i+1:]...)

		if isValidReport(subNumbers) {
			return true
		}
	}

	return false
}

func isValidReport(numbers []int) bool {
	ascending := numbers[0] < numbers[1]
	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		if diff < 0 {
			diff = -diff
		}

		if diff < minDiff || diff > maxDiff {
			return false
		}

		if ascending && numbers[i] > numbers[i+1] {
			return false
		} else if !ascending && numbers[i] < numbers[i+1] {
			return false
		}
	}

	return true
}
