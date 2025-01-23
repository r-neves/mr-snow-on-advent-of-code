package datastructures

import "fmt"

type Point2D struct {
	I int
	J int
}

func (p Point2D) String() string {
	return fmt.Sprint(p.I, ",", p.J)
}

func ManhattanDistance(p1, p2 Point2D) int {
	iDistance := p1.I - p2.I
	if iDistance < 0 {
		iDistance = -iDistance
	}

	jDistance := p1.J - p2.J
	if jDistance < 0 {
		jDistance = -jDistance
	}

	return iDistance + jDistance
}
