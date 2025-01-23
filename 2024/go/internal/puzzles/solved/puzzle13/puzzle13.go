package puzzle13

import (
	"aoc2024/internal/io"
	"strconv"
	"strings"
)

type machineInput struct {
	ButtonAX int
	ButtonAY int
	ButtonBX int
	ButtonBY int
	PrizeX   int
	PrizeY   int
}

const (
	buttonACost   = 3
	buttonBCost   = 1
	maxIterations = 100
)

func RunPart1() int {
	lines := io.ReadFile("inputs/13-real.txt")
	var machineInputs []machineInput
	for i := 0; i < len(lines); i += 4 {
		parsedMachineInput := parseMachineInput(lines[i], lines[i+1], lines[i+2], false)
		machineInputs = append(machineInputs, parsedMachineInput)
	}

	sum := 0
	for _, machine := range machineInputs {
		result := getMinimumCostToWin1(machine)
		if result == -1 {
			continue
		}

		sum += result
	}

	return sum
}

func RunPart2() int {
	lines := io.ReadFile("inputs/13-real.txt")
	var machineInputs []machineInput
	for i := 0; i < len(lines); i += 4 {
		parsedMachineInput := parseMachineInput(lines[i], lines[i+1], lines[i+2], true)
		machineInputs = append(machineInputs, parsedMachineInput)
	}

	sum := 0
	for _, machine := range machineInputs {
		result := getMinimumCostToWin2(machine)
		if result == -1 {
			continue
		}

		sum += result
	}

	return sum
}

func parseMachineInput(buttonALine, buttonBLine, prizeLine string, isSecondPart bool) machineInput {
	start := strings.Index(buttonALine, "+")
	end := strings.Index(buttonALine, ",")
	buttonAX, err := strconv.Atoi(buttonALine[start+1 : end])
	if err != nil {
		panic(err)
	}

	buttonALine = buttonALine[end+1:]
	start = strings.Index(buttonALine, "+")
	buttonAY, err := strconv.Atoi(buttonALine[start+1:])
	if err != nil {
		panic(err)
	}

	start = strings.Index(buttonBLine, "+")
	end = strings.Index(buttonBLine, ",")
	buttonBX, err := strconv.Atoi(buttonBLine[start+1 : end])
	if err != nil {
		panic(err)
	}

	buttonBLine = buttonBLine[end+1:]
	start = strings.Index(buttonBLine, "+")
	buttonBY, err := strconv.Atoi(buttonBLine[start+1:])
	if err != nil {
		panic(err)
	}

	start = strings.Index(prizeLine, "=")
	end = strings.Index(prizeLine, ",")
	prizeX, err := strconv.Atoi(prizeLine[start+1 : end])
	if err != nil {
		panic(err)
	}

	prizeLine = prizeLine[end+1:]
	start = strings.Index(prizeLine, "=")
	prizeY, err := strconv.Atoi(prizeLine[start+1:])
	if err != nil {
		panic(err)
	}

	if isSecondPart {
		prizeX += 10000000000000
		prizeY += 10000000000000
	}

	return machineInput{
		ButtonAX: buttonAX,
		ButtonAY: buttonAY,
		ButtonBX: buttonBX,
		ButtonBY: buttonBY,
		PrizeX:   prizeX,
		PrizeY:   prizeY,
	}
}

func getMinimumCostToWin1(machine machineInput) int {
	minCost := -1

	for i := 0; i <= maxIterations; i++ {
		for j := 0; j <= maxIterations; j++ {
			if i*machine.ButtonAX+j*machine.ButtonBX > machine.PrizeX ||
				i*machine.ButtonAY+j*machine.ButtonBY > machine.PrizeY {
				break
			}

			if i*machine.ButtonAX+j*machine.ButtonBX == machine.PrizeX &&
				i*machine.ButtonAY+j*machine.ButtonBY == machine.PrizeY {
				cost := i*buttonACost + j*buttonBCost
				if minCost == -1 || cost < minCost {
					minCost = cost
				}
			}
		}
	}

	return minCost
}

func getMinimumCostToWin2(machine machineInput) int {
	determinant := machine.ButtonAX*machine.ButtonBY - machine.ButtonAY*machine.ButtonBX
	if determinant == 0 {
		return -1
	}

	// a1x (button A X) + b1y (button B X) = c1 (machine X)
	// a2x (button A Y) + b2y (button B Y) = c2 (machine Y)
	// determinant = a1 b2 - a2 b1
	// x = (c1b2 - c2b1)/determinant
	// y = (a1c2 - a2c1)/determinant

	first := machine.PrizeX*machine.ButtonBY - machine.PrizeY*machine.ButtonBX
	second := machine.ButtonAX*machine.PrizeY - machine.ButtonAY*machine.PrizeX

	if first%determinant != 0 {
		return -1
	}

	if second%determinant != 0 {
		return -1
	}

	buttonAClicks := first / determinant
	buttonBClicks := second / determinant

	return buttonACost*buttonAClicks + buttonBCost*buttonBClicks
}
