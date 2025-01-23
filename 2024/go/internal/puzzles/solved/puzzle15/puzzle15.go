package puzzle15

import (
	ds "aoc2024/internal/datastructures"
	"aoc2024/internal/io"
	"os"
)

const (
	wallChar         = '#'
	boxChar          = 'O'
	robotChar        = '@'
	emptyChar        = '.'
	boxLeftSideChar  = '['
	boxRightSideChar = ']'

	moveUp    = '^'
	moveDown  = 'v'
	moveLeft  = '<'
	moveRight = '>'
)

func RunPart1() int {
	lines := io.ReadFile("inputs/15-real.txt")

	matrix, moves, robot := parseInput1(lines)
	for _, move := range moves {
		//printMatrix(matrix)
		direction := ds.Point2D{}
		switch move {
		case moveUp:
			direction = ds.Point2D{I: -1, J: 0}
		case moveDown:
			direction = ds.Point2D{I: 1, J: 0}
		case moveLeft:
			direction = ds.Point2D{I: 0, J: -1}
		case moveRight:
			direction = ds.Point2D{I: 0, J: 1}
		}

		moveToDirection1(&matrix, &robot, direction)
	}

	sum := 0
	for i, row := range matrix {
		for j, c := range row {
			if c == boxChar {
				sum += i*100 + j
			}
		}
	}

	return sum
}

func moveToDirection1(matrix *[][]int32, robot *ds.Point2D, direction ds.Point2D) {
	robotNewPosition := ds.Point2D{I: robot.I + direction.I, J: robot.J + direction.J}
	if (*matrix)[robotNewPosition.I][robotNewPosition.J] == wallChar {
		return
	}

	if (*matrix)[robotNewPosition.I][robotNewPosition.J] == emptyChar {
		(*matrix)[robot.I][robot.J] = emptyChar
		(*matrix)[robotNewPosition.I][robotNewPosition.J] = robotChar
		*robot = robotNewPosition
		return
	}

	boxCandidatePosition := robotNewPosition
	for {
		if (*matrix)[boxCandidatePosition.I][boxCandidatePosition.J] == wallChar {
			return
		}

		if (*matrix)[boxCandidatePosition.I][boxCandidatePosition.J] == emptyChar {
			(*matrix)[robot.I][robot.J] = emptyChar
			(*matrix)[boxCandidatePosition.I][boxCandidatePosition.J] = boxChar
			(*matrix)[robotNewPosition.I][robotNewPosition.J] = robotChar
			*robot = robotNewPosition
			return
		}

		boxCandidatePosition = ds.Point2D{
			I: boxCandidatePosition.I + direction.I,
			J: boxCandidatePosition.J + direction.J,
		}
	}
}

type box struct {
	Left  ds.Point2D
	Right ds.Point2D
}

func RunPart2() int {
	lines := io.ReadFile("inputs/15-real.txt")
	//output, err := os.Create("output.txt")
	//if err != nil {
	//	panic(err)
	//}

	matrix, moves, robot := parseInput2(lines)
	for _, move := range moves {
		//printMatrix(matrix, output)
		//output.WriteString("Moving: " + string(move) + "\n")

		moveToDirection2(&matrix, &robot, move)
	}

	//printMatrix(matrix, output)

	sum := 0
	for i, row := range matrix {
		for j, c := range row {
			if c == boxLeftSideChar {
				sum += i*100 + j
			}
		}
	}

	return sum
}

func moveToDirection2(matrixPtr *[][]int32, robot *ds.Point2D, move int32) {
	matrix := *matrixPtr
	direction := ds.Point2D{}
	switch move {
	case moveUp:
		direction = ds.Point2D{I: -1, J: 0}
	case moveDown:
		direction = ds.Point2D{I: 1, J: 0}
	case moveLeft:
		direction = ds.Point2D{I: 0, J: -1}
	case moveRight:
		direction = ds.Point2D{I: 0, J: 1}
	}

	robotNewPosition := ds.Point2D{I: robot.I + direction.I, J: robot.J + direction.J}
	if matrix[robotNewPosition.I][robotNewPosition.J] == wallChar {
		return
	}

	if matrix[robotNewPosition.I][robotNewPosition.J] == emptyChar {
		matrix[robot.I][robot.J] = emptyChar
		matrix[robotNewPosition.I][robotNewPosition.J] = robotChar
		*robot = robotNewPosition
		return
	}

	boxCandidatePosition := robotNewPosition
	var boxLinesToUpdate [][]box
	for {
		switch move {
		case moveLeft:
			if matrix[boxCandidatePosition.I][boxCandidatePosition.J] == wallChar {
				return
			}

			if matrix[boxCandidatePosition.I][boxCandidatePosition.J] == emptyChar {
				updatePositions(matrixPtr, direction, boxLinesToUpdate, robot, robotNewPosition)
				return
			}

			boxLinesToUpdate = append(
				boxLinesToUpdate,
				[]box{
					{
						Left:  ds.Point2D{I: boxCandidatePosition.I, J: boxCandidatePosition.J - 1},
						Right: ds.Point2D{I: boxCandidatePosition.I, J: boxCandidatePosition.J},
					},
				},
			)

			boxCandidatePosition.J += direction.J * 2
		case moveRight:
			if matrix[boxCandidatePosition.I][boxCandidatePosition.J] == wallChar {
				return
			}

			if matrix[boxCandidatePosition.I][boxCandidatePosition.J] == emptyChar {
				updatePositions(matrixPtr, direction, boxLinesToUpdate, robot, robotNewPosition)
				return
			}

			boxLinesToUpdate = append(
				boxLinesToUpdate,
				[]box{
					{
						Left:  ds.Point2D{I: boxCandidatePosition.I, J: boxCandidatePosition.J},
						Right: ds.Point2D{I: boxCandidatePosition.I, J: boxCandidatePosition.J + 1},
					},
				},
			)

			boxCandidatePosition.J += direction.J * 2
		case moveUp, moveDown:
			var boxes []box
			if len(boxLinesToUpdate) == 0 {
				var b box
				if matrix[boxCandidatePosition.I][boxCandidatePosition.J] == boxLeftSideChar {
					b = box{
						Left:  ds.Point2D{I: boxCandidatePosition.I, J: boxCandidatePosition.J},
						Right: ds.Point2D{I: boxCandidatePosition.I, J: boxCandidatePosition.J + 1},
					}
				} else {
					b = box{
						Left:  ds.Point2D{I: boxCandidatePosition.I, J: boxCandidatePosition.J - 1},
						Right: ds.Point2D{I: boxCandidatePosition.I, J: boxCandidatePosition.J},
					}
				}
				boxes = append(boxes, b)
			} else {
				lastLine := boxLinesToUpdate[len(boxLinesToUpdate)-1]
				for i := 0; i < len(lastLine); i++ {
					currBox := lastLine[i]
					if matrix[currBox.Left.I+direction.I][currBox.Left.J] == wallChar ||
						matrix[currBox.Right.I+direction.I][currBox.Right.J] == wallChar {
						return
					}

					if matrix[currBox.Left.I+direction.I][currBox.Left.J] == boxLeftSideChar {
						boxes = append(boxes, box{
							Left:  ds.Point2D{I: currBox.Left.I + direction.I, J: currBox.Left.J},
							Right: ds.Point2D{I: currBox.Right.I + direction.I, J: currBox.Right.J},
						})
					} else if matrix[currBox.Left.I+direction.I][currBox.Left.J] == boxRightSideChar {
						boxes = append(boxes, box{
							Left:  ds.Point2D{I: currBox.Left.I + direction.I, J: currBox.Left.J - 1},
							Right: ds.Point2D{I: currBox.Right.I + direction.I, J: currBox.Left.J},
						})
					}
					// ignore right side of the top box since it was added by the left side of the box before
					if matrix[currBox.Right.I+direction.I][currBox.Right.J] == boxLeftSideChar {
						boxes = append(boxes, box{
							Left:  ds.Point2D{I: currBox.Right.I + direction.I, J: currBox.Right.J},
							Right: ds.Point2D{I: currBox.Right.I + direction.I, J: currBox.Right.J + 1},
						})
					}
				}
			}

			if len(boxes) == 0 {
				if matrix[boxCandidatePosition.I][boxCandidatePosition.J] == wallChar {
					updatePositions(matrixPtr, direction, boxLinesToUpdate, robot, robotNewPosition)
					return
				}

				if matrix[boxCandidatePosition.I][boxCandidatePosition.J] == emptyChar {
					updatePositions(matrixPtr, direction, boxLinesToUpdate, robot, robotNewPosition)
					return
				}
			}

			boxLinesToUpdate = append(boxLinesToUpdate, boxes)
			boxCandidatePosition.I += direction.I
		}

	}
}

func updatePositions(
	matrix *[][]int32,
	direction ds.Point2D,
	boxLinesToUpdate [][]box,
	robot *ds.Point2D,
	robotNewPosition ds.Point2D,
) {
	for i := len(boxLinesToUpdate) - 1; i >= 0; i-- {
		for j := 0; j < len(boxLinesToUpdate[i]); j++ {
			boxToUpdate := boxLinesToUpdate[i][j]
			(*matrix)[boxToUpdate.Left.I][boxToUpdate.Left.J] = emptyChar
			(*matrix)[boxToUpdate.Right.I][boxToUpdate.Right.J] = emptyChar
			(*matrix)[boxToUpdate.Left.I+direction.I][boxToUpdate.Left.J+direction.J] =
				boxLeftSideChar
			(*matrix)[boxToUpdate.Right.I+direction.I][boxToUpdate.Right.J+direction.J] =
				boxRightSideChar
		}
	}
	(*matrix)[robot.I][robot.J] = emptyChar
	(*matrix)[robotNewPosition.I][robotNewPosition.J] = robotChar
	*robot = robotNewPosition
}

func parseInput1(lines []string) ([][]int32, string, ds.Point2D) {
	var matrix [][]int32
	robot := ds.Point2D{}

	movesStartLine := -1
	for i, line := range lines {
		if line == "" {
			movesStartLine = i + 1
			break
		}

		matrixLine := make([]int32, len(line))
		for j, c := range line {
			if c == robotChar {
				robot = ds.Point2D{I: j, J: i}
			}
			matrixLine[j] = c
		}

		matrix = append(matrix, matrixLine)
	}

	moves := ""
	for i := movesStartLine; i < len(lines); i++ {
		moves += lines[i]
	}

	return matrix, moves, robot
}

func parseInput2(lines []string) ([][]int32, string, ds.Point2D) {
	var matrix [][]int32
	robot := ds.Point2D{}

	movesStartLine := -1
	for i, line := range lines {
		if line == "" {
			movesStartLine = i + 1
			break
		}

		var matrixLine []int32
		for _, c := range line {
			switch c {
			case boxChar:
				matrixLine = append(matrixLine, boxLeftSideChar, boxRightSideChar)
			case wallChar:
				matrixLine = append(matrixLine, wallChar, wallChar)
			case robotChar:
				matrixLine = append(matrixLine, robotChar, emptyChar)
				robot = ds.Point2D{I: i, J: len(matrixLine) - 2}
			case emptyChar:
				matrixLine = append(matrixLine, emptyChar, emptyChar)
			}
		}

		matrix = append(matrix, matrixLine)
	}

	moves := ""
	for i := movesStartLine; i < len(lines); i++ {
		moves += lines[i]
	}

	return matrix, moves, robot
}

func printMatrix(matrix [][]int32, file *os.File) {
	for _, row := range matrix {
		for _, c := range row {
			file.WriteString(string(c))
		}
		file.WriteString("\n")
	}
}
