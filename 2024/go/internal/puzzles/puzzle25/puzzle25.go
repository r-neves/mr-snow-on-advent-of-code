package puzzle25

import (
	"aoc2024/internal/io"
)

func RunPart1() int {
	lines := io.ReadFile("inputs/25-real.txt")

	locks, keys := parseInput(lines)

	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			fits := true
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					fits = false
					break
				}
			}

			if fits {
				count++
			}
		}
	}

	return count
}

func parseInput(lines []string) ([][5]int, [][5]int) {
	locks := make([][5]int, 0)
	keys := make([][5]int, 0)

	for i := 0; i < len(lines); i += 8 {
		pattern := [5]int{}
		for column := 0; column < 5; column++ {
			count := 0
			for j := 1; j < 6; j++ {
				if lines[i+j][column] == '#' {
					count++
				}
			}

			pattern[column] = count
		}

		if lines[i][0] == '#' {
			locks = append(locks, pattern)
		} else {
			keys = append(keys, pattern)
		}
	}

	return locks, keys
}
