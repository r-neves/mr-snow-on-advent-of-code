package puzzle8

import (
	"aoc2024/internal/datastructures"
	"aoc2024/internal/io"
	"fmt"
)

func RunPart1() int {
	lines := io.ReadFile("inputs/8-real.txt")
	antennasBySignal := make(map[int32][]datastructures.Point2D)
	uniqueAntinodes := make(map[string]bool)

	for i, line := range lines {
		for j, antenna := range line {
			if antenna == '.' {
				continue
			}
			antennasBySignal[antenna] = append(antennasBySignal[antenna], datastructures.Point2D{I: i, J: j})
		}
	}

	iSize := len(lines)
	jSize := len(lines[0])

	for _, antennas := range antennasBySignal {
		for i := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				distanceI, distanceJ := distanceBetween2Points(antennas[i], antennas[j])
				antiNode1 := datastructures.Point2D{I: antennas[i].I + distanceI, J: antennas[i].J + distanceJ}
				antiNode2 := datastructures.Point2D{I: antennas[j].I - distanceI, J: antennas[j].J - distanceJ}
				if isWithinBounds(antiNode1, iSize, jSize) {
					uniqueAntinodes[fmt.Sprintf("%d-%d", antiNode1.I, antiNode1.J)] = true
				}

				if isWithinBounds(antiNode2, iSize, jSize) {
					uniqueAntinodes[fmt.Sprintf("%d-%d", antiNode2.I, antiNode2.J)] = true
				}
			}
		}
	}

	return len(uniqueAntinodes)
}

func isWithinBounds(p datastructures.Point2D, iSize, jSize int) bool {
	return p.I >= 0 && p.I < iSize && p.J >= 0 && p.J < jSize
}

func distanceBetween2Points(p1, p2 datastructures.Point2D) (int, int) {
	return p1.I - p2.I, p1.J - p2.J
}

func RunPart2() int {
	lines := io.ReadFile("inputs/8-real.txt")
	antennasBySignal := make(map[int32][]datastructures.Point2D)
	uniqueAntinodes := make(map[string]bool)

	for i, line := range lines {
		for j, antenna := range line {
			if antenna == '.' {
				continue
			}
			antennasBySignal[antenna] = append(antennasBySignal[antenna], datastructures.Point2D{I: i, J: j})
		}
	}

	iSize := len(lines)
	jSize := len(lines[0])

	for _, antennas := range antennasBySignal {
		for i := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				uniqueAntinodes[fmt.Sprintf("%d-%d", antennas[i].I, antennas[i].J)] = true
				uniqueAntinodes[fmt.Sprintf("%d-%d", antennas[j].I, antennas[j].J)] = true
				distanceI, distanceJ := distanceBetween2Points(antennas[i], antennas[j])

				newPoint1 := datastructures.Point2D{I: antennas[i].I + distanceI, J: antennas[i].J + distanceJ}
				for isWithinBounds(newPoint1, iSize, jSize) {
					uniqueAntinodes[fmt.Sprintf("%d-%d", newPoint1.I, newPoint1.J)] = true
					newPoint1.I += distanceI
					newPoint1.J += distanceJ
				}

				newPoint2 := datastructures.Point2D{I: antennas[j].I - distanceI, J: antennas[j].J - distanceJ}
				for isWithinBounds(newPoint2, iSize, jSize) {
					uniqueAntinodes[fmt.Sprintf("%d-%d", newPoint2.I, newPoint2.J)] = true
					newPoint2.I -= distanceI
					newPoint2.J -= distanceJ
				}
			}
		}
	}

	return len(uniqueAntinodes)
}
