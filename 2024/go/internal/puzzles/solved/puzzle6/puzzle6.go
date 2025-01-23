package puzzle6

import (
	"aoc2024/internal/io"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func RunPart1() {
	lines := io.ReadFile("inputs/6-real.txt")
	fieldMap := make([][]uint8, 0, len(lines))
	iCoord := 0
	jCoord := 0
	for i, line := range lines {
		fieldMap = append(fieldMap, []uint8(line))
		playerIndex := slices.Index(fieldMap[i], '^')
		if playerIndex != -1 {
			iCoord = i
			jCoord = playerIndex
		}
	}

	fmt.Println(len(getGuardVisitedPlaces(fieldMap, iCoord, jCoord)))
}

func isOutOfBounds(fieldMap [][]uint8, i, j int) bool {
	return i < 0 || j < 0 || i >= len(fieldMap) || j >= len(fieldMap[0])
}

func getGuardVisitedPlaces(fieldMap [][]uint8, i, j int) map[string]bool {
	visited := make(map[string]bool)
	iDir := -1
	jDir := 0
	for {
		key := fmt.Sprintf("%d,%d", i, j)
		visited[key] = true

		if isOutOfBounds(fieldMap, i+iDir, j+jDir) {
			break
		}

		nextPosition := fieldMap[i+iDir][j+jDir]
		if nextPosition == '#' {
			if iDir == -1 {
				iDir = 0
				jDir = 1
			} else if iDir == 1 {
				iDir = 0
				jDir = -1
			} else if jDir == -1 {
				iDir = -1
				jDir = 0
			} else if jDir == 1 {
				iDir = 1
				jDir = 0
			}

			continue
		}

		i += iDir
		j += jDir
	}

	return visited
}

func RunPart2() {
	lines := io.ReadFile("inputs/6-real.txt")

	fieldMap := make([][]uint8, 0, len(lines))
	iCoord := 0
	jCoord := 0
	for i, line := range lines {
		fieldMap = append(fieldMap, []uint8(line))
		playerIndex := slices.Index(fieldMap[i], '^')
		if playerIndex != -1 {
			iCoord = i
			jCoord = playerIndex
		}
	}

	visited := getGuardVisitedPlaces(fieldMap, iCoord, jCoord)
	sum := 0

	for place, _ := range visited {
		coords := strings.Split(place, ",")
		i, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}

		j, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}

		if i == iCoord && j == jCoord {
			continue
		}

		fieldMap[i][j] = '#'
		if isStuckInALoop(fieldMap, iCoord, jCoord) {
			sum++
		}
		fieldMap[i][j] = '.'
	}

	fmt.Println(sum)
}

func isStuckInALoop(fieldMap [][]uint8, i, j int) bool {
	visited := make(map[string]bool)
	iDir := -1
	jDir := 0
	for {
		key := fmt.Sprintf("%d,%d,%d,%d", i, j, iDir, jDir)
		if visited[key] {
			return true
		}

		if isOutOfBounds(fieldMap, i+iDir, j+jDir) {
			return false
		}

		visited[key] = true

		nextPosition := fieldMap[i+iDir][j+jDir]
		if nextPosition == '#' {
			if iDir == -1 {
				iDir = 0
				jDir = 1
			} else if iDir == 1 {
				iDir = 0
				jDir = -1
			} else if jDir == -1 {
				iDir = -1
				jDir = 0
			} else if jDir == 1 {
				iDir = 1
				jDir = 0
			}

			continue
		}

		i += iDir
		j += jDir
	}
}
