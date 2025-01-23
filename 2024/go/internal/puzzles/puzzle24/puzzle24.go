package puzzle24

import (
	"aoc2024/internal/io"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type operation int

type gate struct {
	wire1      string
	wire2      string
	operation  operation
	outputWire string
}

const (
	AND operation = iota
	OR
	XOR
)

func RunPart1() int {
	lines := io.ReadFile("inputs/24-real.txt")

	wireOutputs, gates := parseInput(lines)

	for len(gates) > 0 {
		var newGates []gate
		for _, g := range gates {
			if _, ok := wireOutputs[g.wire1]; !ok {
				newGates = append(newGates, g)
				continue
			}

			if _, ok := wireOutputs[g.wire2]; !ok {
				newGates = append(newGates, g)
				continue
			}

			switch g.operation {
			case AND:
				wireOutputs[g.outputWire] = wireOutputs[g.wire1] & wireOutputs[g.wire2]
			case OR:
				wireOutputs[g.outputWire] = wireOutputs[g.wire1] | wireOutputs[g.wire2]
			case XOR:
				wireOutputs[g.outputWire] = wireOutputs[g.wire1] ^ wireOutputs[g.wire2]
			}
		}

		gates = newGates
	}

	var keys []string
	for k := range wireOutputs {
		if k[0] == 'z' {
			keys = append(keys, k)
		}
	}

	slices.Sort(keys)
	slices.Reverse(keys)

	result := ""
	for _, k := range keys {
		result += strconv.Itoa(wireOutputs[k])
	}

	fmt.Println(result)
	number, err := strconv.ParseInt(result, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(number)
}

func parseInput(lines []string) (map[string]int, []gate) {
	wireOuputs := make(map[string]int)
	var gates []gate

	index := 0
	for {
		if lines[index] == "" {
			break
		}

		split := strings.Split(lines[index], ": ")
		value, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		wireOuputs[split[0]] = value
		index++
	}

	for i := index + 1; i < len(lines); i++ {
		split := strings.Split(lines[i], " ")

		var op operation
		switch split[1] {
		case "AND":
			op = AND
		case "OR":
			op = OR
		case "XOR":
			op = XOR
		default:
			panic("Invalid operation")
		}

		g := gate{
			wire1:      split[0],
			wire2:      split[2],
			operation:  op,
			outputWire: split[4],
		}

		gates = append(gates, g)
	}

	return wireOuputs, gates
}

func RunPart2() string {
	lines := io.ReadFile("inputs/24-real.txt")

	wireOutputs, gates := parseInput(lines)
	gatesByOutput := make(map[string]gate)
	for _, g := range gates {
		gatesByOutput[g.outputWire] = g
	}

	xValue := getWireFullValue(wireOutputs, "x")
	yValue := getWireFullValue(wireOutputs, "y")
	expectedValue := xValue + yValue

	for len(gates) > 0 {
		var newGates []gate
		for _, g := range gates {
			if _, ok := wireOutputs[g.wire1]; !ok {
				newGates = append(newGates, g)
				continue
			}

			if _, ok := wireOutputs[g.wire2]; !ok {
				newGates = append(newGates, g)
				continue
			}

			switch g.operation {
			case AND:
				wireOutputs[g.outputWire] = wireOutputs[g.wire1] & wireOutputs[g.wire2]
			case OR:
				wireOutputs[g.outputWire] = wireOutputs[g.wire1] | wireOutputs[g.wire2]
			case XOR:
				wireOutputs[g.outputWire] = wireOutputs[g.wire1] ^ wireOutputs[g.wire2]
			}
		}

		gates = newGates
	}

	number := getWireFullValue(wireOutputs, "z")

	fmt.Println("X value:", xValue)
	fmt.Println("Y value:", yValue)
	fmt.Println("Expected value:", expectedValue)
	fmt.Println("Actual value:", number)

	intDiff := expectedValue ^ number
	fmt.Println("Diff:", intDiff)

	strDiff := strconv.FormatInt(int64(intDiff), 2)
	bitsDiff := make(map[string]bool)
	for i := len(strDiff) - 1; i >= 0; i-- {
		if strDiff[i] == '1' {
			strIndex := strconv.Itoa(len(strDiff) - 1 - i)
			if len(strIndex) == 1 {
				strIndex = "0" + strIndex
			}
			bitsDiff["z"+strIndex] = true
		}
	}

	currLen := len(bitsDiff)
	for {
		for wire := range bitsDiff {
			g, found := gatesByOutput[wire]
			if found {
				_, alreadyIn := bitsDiff[g.wire1]
				if !alreadyIn {
					bitsDiff[g.wire1] = true
				}

				_, alreadyIn2 := bitsDiff[g.wire2]
				if !alreadyIn2 {
					bitsDiff[g.wire2] = true
				}
			}
		}

		if currLen == len(bitsDiff) {
			break
		}

		currLen = len(bitsDiff)
	}

	for wire := range bitsDiff {
		_, isOutput := gatesByOutput[wire]
		if !isOutput {
			delete(bitsDiff, wire)
		}
	}

	var keys []string
	for k := range bitsDiff {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	return strings.Join(keys, ",")
}

func getWireFullValue(wireOutputs map[string]int, prefix string) int {
	var relevantWires []string

	for wire := range wireOutputs {
		if strings.HasPrefix(wire, prefix) {
			relevantWires = append(relevantWires, wire)
		}
	}

	slices.Sort(relevantWires)
	slices.Reverse(relevantWires)

	result := ""
	for _, wire := range relevantWires {
		result += strconv.Itoa(wireOutputs[wire])
	}

	number, err := strconv.ParseInt(result, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(number)
}
