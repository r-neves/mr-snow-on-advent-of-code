package puzzle18

import (
	ds "aoc2024/internal/datastructures"
	"aoc2024/internal/io"
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type nextMove struct {
	Cost        int
	NewPosition ds.Point2D
}

type nextMoveHeap []nextMove

func (h nextMoveHeap) Len() int { return len(h) }
func (h nextMoveHeap) Less(i, j int) bool {
	return h[i].Cost < h[j].Cost
}
func (h nextMoveHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *nextMoveHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(nextMove))
}

func (h *nextMoveHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var (
	minKnownCostSoFar = -1
)

func RunPart1() int {
	lines := io.ReadFile("inputs/18-real.txt")
	matrixSize := 71
	fallenBytes := 1024

	matrix := make([][]int32, matrixSize)
	for i := 0; i < matrixSize; i++ {
		matrix[i] = make([]int32, matrixSize)
		for j := 0; j < matrixSize; j++ {
			matrix[i][j] = '.'
		}
	}

	for i, line := range lines {
		if i == fallenBytes {
			break
		}
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}
		matrix[y][x] = '#'
	}

	printMatrix(matrix)

	start := ds.Point2D{I: 0, J: 0}
	end := ds.Point2D{I: matrixSize - 1, J: matrixSize - 1}

	nextMoves := nextMoveHeap{}
	heap.Init(&nextMoves)

	heap.Push(&nextMoves, nextMove{
		Cost:        0,
		NewPosition: start,
	})

	minCost := -1
	visited := make(map[ds.Point2D]bool)

	for len(nextMoves) > 0 {
		move := heap.Pop(&nextMoves).(nextMove)
		result := dfs(
			matrix,
			move.NewPosition,
			end,
			move.Cost,
			visited,
			&nextMoves,
		)

		if result != -1 && (minCost == -1 || result < minCost) {
			minCost = result
			fmt.Println("New min cost:", minCost)
		}
	}

	return minCost
}

func dfs(
	maze [][]int32,
	currPoint ds.Point2D,
	goal ds.Point2D,
	cost int,
	visited map[ds.Point2D]bool,
	nextMoves *nextMoveHeap,
) int {
	if currPoint == goal {
		if minKnownCostSoFar == -1 || cost < minKnownCostSoFar {
			minKnownCostSoFar = cost
		}
		//printMazePath(maze, visited, cost)
		return cost
	}

	if minKnownCostSoFar != -1 && cost >= minKnownCostSoFar {
		return -1
	}

	if maze[currPoint.I][currPoint.J] == '#' || visited[currPoint] {
		return -1
	}

	visited[currPoint] = true

	upPoint := ds.Point2D{I: currPoint.I - 1, J: currPoint.J}
	validUp := !isOutOfBounds(len(maze), upPoint) && maze[upPoint.I][upPoint.J] != '#' && !visited[upPoint]
	if validUp {
		heap.Push(nextMoves, nextMove{
			Cost:        cost + 1,
			NewPosition: upPoint,
		})
	}

	downPoint := ds.Point2D{I: currPoint.I + 1, J: currPoint.J}
	validDown := !isOutOfBounds(len(maze), downPoint) && maze[downPoint.I][downPoint.J] != '#' && !visited[downPoint]
	if validDown {
		heap.Push(nextMoves, nextMove{
			Cost:        cost + 1,
			NewPosition: downPoint,
		})
	}

	leftPoint := ds.Point2D{I: currPoint.I, J: currPoint.J - 1}
	validLeft := !isOutOfBounds(len(maze), leftPoint) && maze[leftPoint.I][leftPoint.J] != '#' && !visited[leftPoint]
	if validLeft {
		heap.Push(nextMoves, nextMove{
			Cost:        cost + 1,
			NewPosition: leftPoint,
		})
	}

	rightPoint := ds.Point2D{I: currPoint.I, J: currPoint.J + 1}
	validRight := !isOutOfBounds(len(maze), rightPoint) && maze[rightPoint.I][rightPoint.J] != '#' && !visited[rightPoint]
	if validRight {
		heap.Push(nextMoves, nextMove{
			Cost:        cost + 1,
			NewPosition: rightPoint,
		})
	}

	return -1
}

func isOutOfBounds(size int, point ds.Point2D) bool {
	return point.I < 0 || point.I >= size || point.J < 0 || point.J >= size
}

func printMatrix(matrix [][]int32) {
	for _, row := range matrix {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}

	fmt.Println()
}

func RunPart2() string {
	lines := io.ReadFile("inputs/18-real.txt")
	matrixSize := 71

	matrixByTimestamp := make(map[int]map[ds.Point2D]bool)

	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}

		point := ds.Point2D{I: y, J: x}
		for timestamp := i; timestamp < len(lines); timestamp++ {
			if _, ok := matrixByTimestamp[timestamp]; !ok {
				matrixByTimestamp[timestamp] = make(map[ds.Point2D]bool)
			}
			matrixByTimestamp[timestamp][point] = true
		}
	}

	for i := 0; i < len(lines); i++ {
		cost := findCostToEscape(matrixByTimestamp[i], matrixSize)
		if cost == -1 {
			fmt.Printf("%dth byte blocks path: %s\n", i, lines[i])
			return lines[i]
		}
	}

	return "-1"
}

func findCostToEscape(
	corrupted map[ds.Point2D]bool,
	matrixSize int,
) int {
	start := ds.Point2D{I: 0, J: 0}
	end := ds.Point2D{I: matrixSize - 1, J: matrixSize - 1}

	nextMoves := nextMoveHeap{}
	heap.Init(&nextMoves)

	heap.Push(&nextMoves, nextMove{
		Cost:        0,
		NewPosition: start,
	})

	visited := make(map[ds.Point2D]bool)

	for len(nextMoves) > 0 {
		move := heap.Pop(&nextMoves).(nextMove)
		result := dfs2(
			corrupted,
			move.NewPosition,
			end,
			move.Cost,
			visited,
			&nextMoves,
			matrixSize,
		)

		if result != -1 {
			return result
		}
	}

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			if corrupted[ds.Point2D{I: i, J: j}] {
				fmt.Print("#")
			} else if visited[ds.Point2D{I: i, J: j}] {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return -1
}

func dfs2(
	corrupted map[ds.Point2D]bool,
	currPoint ds.Point2D,
	goal ds.Point2D,
	cost int,
	visited map[ds.Point2D]bool,
	nextMoves *nextMoveHeap,
	size int,
) int {
	if currPoint == goal {
		//printMazePath(maze, visited, cost)
		return cost
	}

	if corrupted[currPoint] || visited[currPoint] {
		return -1
	}

	visited[currPoint] = true

	upPoint := ds.Point2D{I: currPoint.I - 1, J: currPoint.J}
	validUp := !isOutOfBounds(size, upPoint) && !corrupted[upPoint] && !visited[upPoint]
	if validUp {
		heap.Push(nextMoves, nextMove{
			Cost:        cost + 1,
			NewPosition: upPoint,
		})
	}

	downPoint := ds.Point2D{I: currPoint.I + 1, J: currPoint.J}
	validDown := !isOutOfBounds(size, downPoint) && !corrupted[downPoint] && !visited[downPoint]
	if validDown {
		heap.Push(nextMoves, nextMove{
			Cost:        cost + 1,
			NewPosition: downPoint,
		})
	}

	leftPoint := ds.Point2D{I: currPoint.I, J: currPoint.J - 1}
	validLeft := !isOutOfBounds(size, leftPoint) && !corrupted[leftPoint] && !visited[leftPoint]
	if validLeft {
		heap.Push(nextMoves, nextMove{
			Cost:        cost + 1,
			NewPosition: leftPoint,
		})
	}

	rightPoint := ds.Point2D{I: currPoint.I, J: currPoint.J + 1}
	validRight := !isOutOfBounds(size, rightPoint) && !corrupted[rightPoint] && !visited[rightPoint]
	if validRight {
		heap.Push(nextMoves, nextMove{
			Cost:        cost + 1,
			NewPosition: rightPoint,
		})
	}

	return -1
}
