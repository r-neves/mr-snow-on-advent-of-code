package puzzle12

import (
	"aoc2024/internal/datastructures"
	"aoc2024/internal/io"
	"fmt"
)

func RunPart1() int {
	lines := io.ReadFile("inputs/12-real.txt")

	garden := make([][]int32, 0, len(lines))
	for i, line := range lines {
		garden = append(garden, make([]int32, 0, len(line)))
		for _, c := range line {
			garden[i] = append(garden[i], c)
		}
	}

	alreadyMapped := make(map[string]bool)
	sections := make([][]datastructures.Point2D, 0)

	for i, row := range garden {
		for j, plant := range row {
			key := fmt.Sprintf("%d-%d", i, j)
			_, found := alreadyMapped[key]
			if found {
				continue
			}

			var section []datastructures.Point2D
			createGardenSection(garden, &section, alreadyMapped, i, j, plant)
			sections = append(sections, section)
		}
	}

	sum := 0

	for _, section := range sections {
		area := len(section)
		perimeter := calculatePerimeter(garden, section)
		sum += area * perimeter
	}

	return sum
}

func createGardenSection(
	garden [][]int32,
	section *[]datastructures.Point2D,
	alreadyMapped map[string]bool,
	i, j int,
	plant int32,
) {
	if isOutOfBounds(garden, i, j) || isDifferentPlant(garden, i, j, plant) {
		return
	}

	key := fmt.Sprintf("%d-%d", i, j)
	if _, found := alreadyMapped[key]; found {
		return
	}

	alreadyMapped[key] = true
	*section = append(*section, datastructures.Point2D{I: i, J: j})

	createGardenSection(garden, section, alreadyMapped, i-1, j, plant)
	createGardenSection(garden, section, alreadyMapped, i+1, j, plant)
	createGardenSection(garden, section, alreadyMapped, i, j-1, plant)
	createGardenSection(garden, section, alreadyMapped, i, j+1, plant)
}

type fenceSide int

const (
	left fenceSide = iota
	right
	top
	bottom
)

type fence struct {
	cell      datastructures.Point2D
	whichSide fenceSide
}

func (f fence) GetMapKey() string {
	return fmt.Sprintf("%d-%d-%d", f.cell.I, f.cell.J, f.whichSide)
}

func RunPart2() int {
	lines := io.ReadFile("inputs/12-real.txt")

	garden := make([][]int32, 0, len(lines))
	for i, line := range lines {
		garden = append(garden, make([]int32, 0, len(line)))
		for _, c := range line {
			garden[i] = append(garden[i], c)
		}
	}

	alreadyMapped := make(map[string]bool)
	sections := make([][]datastructures.Point2D, 0)

	for i, row := range garden {
		for j, plant := range row {
			key := fmt.Sprintf("%d-%d", i, j)
			_, found := alreadyMapped[key]
			if found {
				continue
			}

			var section []datastructures.Point2D
			createGardenSection(garden, &section, alreadyMapped, i, j, plant)
			sections = append(sections, section)
		}
	}

	sum := 0

	for _, section := range sections {
		area := len(section)
		fences := calculatePlacedFences(garden, section)
		sides := calculateSides(fences)
		sum += area * sides
	}

	return sum
}

func isOutOfBounds(garden [][]int32, i, j int) bool {
	return i < 0 || i >= len(garden) || j < 0 || j >= len(garden[i])
}

func isDifferentPlant(garden [][]int32, i, j int, plant int32) bool {
	return garden[i][j] != plant
}

func calculatePerimeter(garden [][]int32, section []datastructures.Point2D) int {
	perimeter := 0
	plant := garden[section[0].I][section[0].J]
	for _, point := range section {
		i, j := point.I, point.J
		if isOutOfBounds(garden, i+1, j) || isDifferentPlant(garden, i+1, j, plant) {
			perimeter++
		}
		if isOutOfBounds(garden, i-1, j) || isDifferentPlant(garden, i-1, j, plant) {
			perimeter++
		}
		if isOutOfBounds(garden, i, j+1) || isDifferentPlant(garden, i, j+1, plant) {
			perimeter++
		}
		if isOutOfBounds(garden, i, j-1) || isDifferentPlant(garden, i, j-1, plant) {
			perimeter++
		}
	}

	return perimeter
}

func calculatePlacedFences(garden [][]int32, section []datastructures.Point2D) map[string]fence {
	fences := make(map[string]fence)

	plant := garden[section[0].I][section[0].J]
	for _, point := range section {
		i, j := point.I, point.J
		if isOutOfBounds(garden, i+1, j) || isDifferentPlant(garden, i+1, j, plant) {
			f := fence{cell: point, whichSide: bottom}
			fences[f.GetMapKey()] = f
		}
		if isOutOfBounds(garden, i-1, j) || isDifferentPlant(garden, i-1, j, plant) {
			f := fence{cell: point, whichSide: top}
			fences[f.GetMapKey()] = f
		}
		if isOutOfBounds(garden, i, j+1) || isDifferentPlant(garden, i, j+1, plant) {
			f := fence{cell: point, whichSide: right}
			fences[f.GetMapKey()] = f
		}
		if isOutOfBounds(garden, i, j-1) || isDifferentPlant(garden, i, j-1, plant) {
			f := fence{cell: point, whichSide: left}
			fences[f.GetMapKey()] = f
		}
	}

	return fences
}

func calculateSides(fences map[string]fence) int {
	sides := 0
	alreadyMapped := make(map[string]bool)
	for _, f := range fences {
		if _, found := alreadyMapped[f.GetMapKey()]; found {
			continue
		}

		sides++

		if f.whichSide == left || f.whichSide == right {
			fenceOnTop := fence{
				cell:      datastructures.Point2D{I: f.cell.I - 1, J: f.cell.J},
				whichSide: f.whichSide,
			}
			for {
				key := fenceOnTop.GetMapKey()
				if _, found := alreadyMapped[key]; found {
					panic("fence already mapped, something went wrong")
				}

				_, found := fences[key]
				if !found {
					break
				}

				alreadyMapped[key] = true
				fenceOnTop.cell.I--
			}

			fenceOnBottom := fence{
				cell:      datastructures.Point2D{I: f.cell.I + 1, J: f.cell.J},
				whichSide: f.whichSide,
			}
			for {
				key := fenceOnBottom.GetMapKey()
				if _, found := alreadyMapped[key]; found {
					panic("fence already mapped, something went wrong")
				}

				_, found := fences[key]
				if !found {
					break
				}

				alreadyMapped[key] = true
				fenceOnBottom.cell.I++
			}
		} else { // side is top or bottom
			fenceOnLeft := fence{
				cell:      datastructures.Point2D{I: f.cell.I, J: f.cell.J - 1},
				whichSide: f.whichSide,
			}
			for {
				key := fenceOnLeft.GetMapKey()
				if _, found := alreadyMapped[key]; found {
					panic("fence already mapped, something went wrong")
				}

				_, found := fences[key]
				if !found {
					break
				}

				alreadyMapped[key] = true
				fenceOnLeft.cell.J--
			}

			fenceOnRight := fence{
				cell:      datastructures.Point2D{I: f.cell.I, J: f.cell.J + 1},
				whichSide: f.whichSide,
			}
			for {
				key := fenceOnRight.GetMapKey()
				if _, found := alreadyMapped[key]; found {
					panic("fence already mapped, something went wrong")
				}

				_, found := fences[key]
				if !found {
					break
				}

				alreadyMapped[key] = true
				fenceOnRight.cell.J++
			}
		}
	}

	return sides
}
