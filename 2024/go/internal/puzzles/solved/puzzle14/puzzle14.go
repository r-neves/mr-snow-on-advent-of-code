package puzzle14

import (
	"aoc2024/internal/datastructures"
	"aoc2024/internal/io"
	"fmt"
	"strconv"
	"strings"
)

type robot struct {
	position datastructures.Point2D
	velocity datastructures.Point2D
}

const (
	//width  = 11
	//height = 7
	width  = 101
	height = 103

	seconds = 100
)

func RunPart1() int {
	lines := io.ReadFile("inputs/14-real.txt")

	var robots []robot

	for _, line := range lines {
		robots = append(robots, parseRobotInput(line))
	}

	for i := range robots {
		robots[i].position.I = (robots[i].position.I + robots[i].velocity.I*seconds) % height
		if robots[i].position.I < 0 {
			robots[i].position.I += height
		}

		robots[i].position.J = (robots[i].position.J + robots[i].velocity.J*seconds) % width
		if robots[i].position.J < 0 {
			robots[i].position.J += width
		}
	}

	//printRobotMap(robots)

	finalPositions := make(map[string]int, len(robots))
	for _, robot := range robots {
		finalPositions[robot.position.String()]++
	}

	q1 := 0
	for i := 0; i < height/2; i++ {
		for j := 0; j < width/2; j++ {
			count, found := finalPositions[datastructures.Point2D{i, j}.String()]
			if found {
				q1 += count
			}
		}
	}

	q2 := 0
	for i := height/2 + 1; i < height; i++ {
		for j := width/2 + 1; j < width; j++ {
			count, found := finalPositions[datastructures.Point2D{i, j}.String()]
			if found {
				q2 += count
			}
		}
	}

	q3 := 0
	for i := height/2 + 1; i < height; i++ {
		for j := 0; j < width/2; j++ {
			count, found := finalPositions[datastructures.Point2D{i, j}.String()]
			if found {
				q3 += count
			}
		}
	}

	q4 := 0
	for i := 0; i < height/2; i++ {
		for j := width/2 + 1; j < width; j++ {
			count, found := finalPositions[datastructures.Point2D{i, j}.String()]
			if found {
				q4 += count
			}
		}
	}

	return q1 * q2 * q3 * q4
}

func RunPart2() int {
	lines := io.ReadFile("inputs/14-real.txt")

	var robots []robot

	for _, line := range lines {
		robots = append(robots, parseRobotInput(line))
	}

	for s := 1; s < 100000000000; s++ {
		for i := range robots {
			robots[i].position.I = (robots[i].position.I + robots[i].velocity.I) % height
			if robots[i].position.I < 0 {
				robots[i].position.I += height
			}

			robots[i].position.J = (robots[i].position.J + robots[i].velocity.J) % width
			if robots[i].position.J < 0 {
				robots[i].position.J += width
			}
		}

		found := checkForXmasTree(robots)
		if found {
			printRobotMap(robots, false)
			return s
		}
	}

	//printRobotMap(robots)

	finalPositions := make(map[string]int, len(robots))
	for _, robot := range robots {
		finalPositions[robot.position.String()]++
	}

	return 0
}

func parseRobotInput(line string) robot {
	start := strings.Index(line, "=") + 1
	end := strings.Index(line, " ")

	coordinates := strings.Split(line[start:end], ",")
	j, err := strconv.Atoi(coordinates[0])
	if err != nil {
		panic(err)
	}

	i, err := strconv.Atoi(coordinates[1])
	if err != nil {
		panic(err)
	}

	position := datastructures.Point2D{i, j}

	line = line[end+1:]
	start = strings.Index(line, "=") + 1

	vectors := strings.Split(line[start:], ",")
	j, err = strconv.Atoi(vectors[0])
	if err != nil {
		panic(err)
	}

	i, err = strconv.Atoi(vectors[1])
	if err != nil {
		panic(err)
	}

	velocity := datastructures.Point2D{i, j}

	return robot{position, velocity}
}

func printRobotMap(robots []robot, excludeMiddle bool) {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]string, width)
		for j := 0; j < width; j++ {
			grid[i][j] = "."
		}
	}

	finalPositions := make(map[string]int, len(robots))
	for _, robot := range robots {
		finalPositions[robot.position.String()]++
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if excludeMiddle && (i == height/2 || j == width/2) {
				fmt.Print(" ")
				continue
			}

			count, found := finalPositions[datastructures.Point2D{i, j}.String()]
			if found {
				grid[i][j] = strconv.Itoa(count)
			}

			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

}

func checkForXmasTree(robots []robot) bool {
	positions := make(map[string]int, len(robots))
	for _, r := range robots {
		positions[r.position.String()]++
	}

	for _, r := range robots {
		top := r.position
		found := true
		for stack := 1; stack <= 3; stack++ {
			for j := top.J - stack; j <= top.J+1; j++ {
				if j < 0 || j >= width || top.I+stack >= height {
					found = false
					break
				}

				_, hasRobot := positions[datastructures.Point2D{top.I + stack, j}.String()]

				if !hasRobot {
					found = false
					break
				}
			}
		}

		if found {
			return true
		}
	}

	return false
}
