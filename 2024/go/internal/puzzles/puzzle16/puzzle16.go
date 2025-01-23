package puzzle16

import (
	ds "aoc2024/internal/datastructures"
	"aoc2024/internal/io"
	"container/heap"
	"fmt"
)

type nextMove struct {
	Cost        int
	Direction   int
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

const (
	rotateCost = 1000

	leftDir  = 0
	rightDir = 1
	upDir    = 2
	downDir  = 3
)

var (
	minKnownCostSoFar = -1
)

func RunPart1() int {
	lines := io.ReadFile("inputs/16-real.txt")
	maze, start, end := parseInput(lines)
	maze = reduceSearchSpace(maze)

	nextMoves := nextMoveHeap{}
	heap.Init(&nextMoves)

	heap.Push(&nextMoves, nextMove{
		Cost:        0,
		Direction:   rightDir,
		NewPosition: start,
	})

	minCost := -1
	visited := make(map[ds.Point2D]bool)

	for len(nextMoves) > 0 {
		move := heap.Pop(&nextMoves).(nextMove)
		result := dfs(
			maze,
			move.NewPosition,
			end,
			move.Cost,
			move.Direction,
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

func reduceSearchSpace(maze [][]int32) [][]int32 {
	deadEnds := getDeadEnds(maze)
	for len(deadEnds) > 0 {
		for _, deadEnd := range deadEnds {
			maze[deadEnd.I][deadEnd.J] = '#'
		}

		deadEnds = getDeadEnds(maze)
	}

	fmt.Println("Finished reducing search space")
	return maze
}

func getDeadEnds(maze [][]int32) []ds.Point2D {
	var deadEnds []ds.Point2D
	for i, row := range maze {
		for j, cell := range row {
			if cell != '.' {
				continue
			}

			walls := 0
			if maze[i-1][j] == '#' {
				walls++
			}

			if maze[i+1][j] == '#' {
				walls++
			}

			if maze[i][j-1] == '#' {
				walls++
			}

			if maze[i][j+1] == '#' {
				walls++
			}

			if walls == 3 {
				deadEnds = append(deadEnds, ds.Point2D{I: i, J: j})
			}
		}
	}

	return deadEnds
}

func dfs2(
	maze [][]int32,
	currPoint ds.Point2D,
	goal ds.Point2D,
	cost int,
	facing int,
	visited map[ds.Point2D]bool,
	nextMoves *nextMoveHeap,
	graph map[ds.Point2D][]ds.Point2D,
) int {
	if currPoint == goal {
		if minKnownCostSoFar == -1 || cost < minKnownCostSoFar {
			minKnownCostSoFar = cost
		}
		//printMazePath(maze, visited, cost)
		return cost
	}

	if minKnownCostSoFar != -1 && cost > minKnownCostSoFar {
		return -1
	}

	if maze[currPoint.I][currPoint.J] == '#' || visited[currPoint] {
		return -1
	}

	visited[currPoint] = true

	possiblePaths := 0
	upPoint := ds.Point2D{I: currPoint.I - 1, J: currPoint.J}
	validUp := facing != downDir && maze[upPoint.I][upPoint.J] != '#' && !visited[upPoint]
	if validUp {
		possiblePaths++
		heap.Push(nextMoves, nextMove{
			Cost:        cost + calculateDirectionCost(facing, upDir),
			Direction:   upDir,
			NewPosition: upPoint,
		})
		graph[upPoint] = append(graph[upPoint], currPoint)
	}

	downPoint := ds.Point2D{I: currPoint.I + 1, J: currPoint.J}
	validDown := facing != upDir && maze[downPoint.I][downPoint.J] != '#' && !visited[downPoint]
	if validDown {
		possiblePaths++
		heap.Push(nextMoves, nextMove{
			Cost:        cost + calculateDirectionCost(facing, downDir),
			Direction:   downDir,
			NewPosition: downPoint,
		})
		graph[downPoint] = append(graph[downPoint], currPoint)
	}

	leftPoint := ds.Point2D{I: currPoint.I, J: currPoint.J - 1}
	validLeft := facing != rightDir && maze[leftPoint.I][leftPoint.J] != '#' && !visited[leftPoint]
	if validLeft {
		possiblePaths++
		heap.Push(nextMoves, nextMove{
			Cost:        cost + calculateDirectionCost(facing, leftDir),
			Direction:   leftDir,
			NewPosition: leftPoint,
		})
		graph[leftPoint] = append(graph[leftPoint], currPoint)
	}

	rightPoint := ds.Point2D{I: currPoint.I, J: currPoint.J + 1}
	validRight := facing != leftDir && maze[rightPoint.I][rightPoint.J] != '#' && !visited[rightPoint]
	if validRight {
		possiblePaths++
		heap.Push(nextMoves, nextMove{
			Cost:        cost + calculateDirectionCost(facing, rightDir),
			Direction:   rightDir,
			NewPosition: rightPoint,
		})
		graph[rightPoint] = append(graph[rightPoint], currPoint)
	}

	return -1
}

func dfs(
	maze [][]int32,
	currPoint ds.Point2D,
	goal ds.Point2D,
	cost int,
	facing int,
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

	possiblePaths := 0
	upPoint := ds.Point2D{I: currPoint.I - 1, J: currPoint.J}
	validUp := facing != downDir && maze[upPoint.I][upPoint.J] != '#' && !visited[upPoint]
	if validUp {
		possiblePaths++
		heap.Push(nextMoves, nextMove{
			Cost:        cost + calculateDirectionCost(facing, upDir),
			Direction:   upDir,
			NewPosition: upPoint,
		})
	}

	downPoint := ds.Point2D{I: currPoint.I + 1, J: currPoint.J}
	validDown := facing != upDir && maze[downPoint.I][downPoint.J] != '#' && !visited[downPoint]
	if validDown {
		possiblePaths++
		heap.Push(nextMoves, nextMove{
			Cost:        cost + calculateDirectionCost(facing, downDir),
			Direction:   downDir,
			NewPosition: downPoint,
		})
	}

	leftPoint := ds.Point2D{I: currPoint.I, J: currPoint.J - 1}
	validLeft := facing != rightDir && maze[leftPoint.I][leftPoint.J] != '#' && !visited[leftPoint]
	if validLeft {
		possiblePaths++
		heap.Push(nextMoves, nextMove{
			Cost:        cost + calculateDirectionCost(facing, leftDir),
			Direction:   leftDir,
			NewPosition: leftPoint,
		})
	}

	rightPoint := ds.Point2D{I: currPoint.I, J: currPoint.J + 1}
	validRight := facing != leftDir && maze[rightPoint.I][rightPoint.J] != '#' && !visited[rightPoint]
	if validRight {
		possiblePaths++
		heap.Push(nextMoves, nextMove{
			Cost:        cost + calculateDirectionCost(facing, rightDir),
			Direction:   rightDir,
			NewPosition: rightPoint,
		})
	}

	return -1
}

func RunPart2() int {
	lines := io.ReadFile("inputs/16-sample1.txt")
	maze, start, end := parseInput(lines)
	maze = reduceSearchSpace(maze)

	nextMoves := nextMoveHeap{}
	heap.Init(&nextMoves)

	heap.Push(&nextMoves, nextMove{
		Cost:        0,
		Direction:   rightDir,
		NewPosition: start,
	})

	minCost := -1
	visited := make(map[ds.Point2D]bool)
	graph := make(map[ds.Point2D][]ds.Point2D)

	for len(nextMoves) > 0 {
		move := heap.Pop(&nextMoves).(nextMove)
		result := dfs2(
			maze,
			move.NewPosition,
			end,
			move.Cost,
			move.Direction,
			visited,
			&nextMoves,
			graph,
		)

		if result != -1 && (minCost == -1 || result < minCost) {
			minCost = result
			fmt.Println("New min cost:", minCost)
		}
	}

	path := make(map[ds.Point2D]bool)
	count := 1
	queue := []ds.Point2D{end}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		queue = append(queue, graph[curr]...)
		count++
		path[curr] = true
	}

	printMazePath(maze, path, minCost)

	return count
}

func parseInput(lines []string) ([][]int32, ds.Point2D, ds.Point2D) {
	maze := make([][]int32, 0, len(lines))
	start := ds.Point2D{}
	end := ds.Point2D{}

	for i, line := range lines {
		maze = append(maze, make([]int32, 0, len(line)))
		for j, char := range line {
			if char == 'S' {
				start = ds.Point2D{I: i, J: j}
			}

			if char == 'E' {
				end = ds.Point2D{I: i, J: j}
			}

			maze[i] = append(maze[i], char)
		}
	}

	return maze, start, end
}

func calculateDirectionCost(currFacing, nextFacing int) int {
	if currFacing == nextFacing {
		return 1
	}

	if currFacing == leftDir && nextFacing == rightDir || currFacing == rightDir && nextFacing == leftDir ||
		currFacing == upDir && nextFacing == downDir || currFacing == downDir && nextFacing == upDir {
		return rotateCost*2 + 1
	}

	return rotateCost + 1
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

func printMazePath(matrix [][]int32, visited map[ds.Point2D]bool, cost int) {
	fmt.Printf("Path with %d steps, cost: %d\n", len(visited), cost)
	for i, row := range matrix {
		for j, c := range row {
			if visited[ds.Point2D{I: i, J: j}] {
				fmt.Print("o")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}

	fmt.Println()
}
