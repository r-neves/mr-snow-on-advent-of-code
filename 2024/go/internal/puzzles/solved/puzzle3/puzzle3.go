package puzzle3

import (
	"aoc2024/internal/io"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunPart1() {
	lines := io.ReadFile("inputs/3-real.txt")
	input := strings.Join(lines, "")

	sum := 0

	idx := strings.Index(input, "mul(")
	for idx != -1 {
		idx = strings.Index(input, "mul(")
		if idx == -1 {
			break
		}

		idx += 4

		n1 := ""
		for unicode.IsDigit(rune(input[idx])) {
			n1 += string(input[idx])
			idx++
		}

		if len(n1) < 1 || len(n1) > 3 {
			input = input[idx:]
			continue
		}

		if input[idx] != ',' {
			input = input[idx:]
			continue
		}

		idx++
		n2 := ""
		for unicode.IsDigit(rune(input[idx])) {
			n2 += string(input[idx])
			idx++
		}

		if len(n2) < 1 || len(n2) > 3 {
			input = input[idx:]
			continue
		}

		if input[idx] != ')' {
			input = input[idx:]
			continue
		}

		num1, err := strconv.Atoi(n1)
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(n2)
		if err != nil {
			panic(err)
		}

		sum += num1 * num2
		input = input[idx+1:]
	}

	fmt.Println(sum)
}

func RunPart2() {
	const (
		doKeyword   = "do()"
		dontKeyword = "don't()"
		mulKeyword  = "mul("
	)

	lines := io.ReadFile("inputs/3-real.txt")
	input := strings.Join(lines, "")

	sum := 0
	enabled := true

	mulIdx := strings.Index(input, mulKeyword)
	for mulIdx != -1 {
		mulIdx = strings.Index(input, mulKeyword)
		doIdx := strings.Index(input, doKeyword)
		dontIdx := strings.Index(input, dontKeyword)

		if doIdx != -1 && (dontIdx == -1 || doIdx < dontIdx) && doIdx < mulIdx {
			enabled = true
			input = input[doIdx+len(doKeyword):]
			continue
		}

		if dontIdx != -1 && (doIdx == -1 || dontIdx < doIdx) && dontIdx < mulIdx {
			enabled = false
			input = input[dontIdx+len(dontKeyword):]
			continue
		}

		if !enabled {
			input = input[mulIdx+len(mulKeyword):]
			continue
		}

		if mulIdx == -1 {
			break
		}

		mulIdx += 4

		n1 := ""
		for unicode.IsDigit(rune(input[mulIdx])) {
			n1 += string(input[mulIdx])
			mulIdx++
		}

		if len(n1) < 1 || len(n1) > 3 {
			input = input[mulIdx:]
			continue
		}

		if input[mulIdx] != ',' {
			input = input[mulIdx:]
			continue
		}

		mulIdx++
		n2 := ""
		for unicode.IsDigit(rune(input[mulIdx])) {
			n2 += string(input[mulIdx])
			mulIdx++
		}

		if len(n2) < 1 || len(n2) > 3 {
			input = input[mulIdx:]
			continue
		}

		if input[mulIdx] != ')' {
			input = input[mulIdx:]
			continue
		}

		num1, err := strconv.Atoi(n1)
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(n2)
		if err != nil {
			panic(err)
		}

		sum += num1 * num2
		input = input[mulIdx+1:]
	}

	fmt.Println(sum)
}
