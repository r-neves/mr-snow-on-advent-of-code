package puzzle11

import (
	"aoc2024/internal/io"
	"math"
	"strconv"
	"strings"
)

func RunPart1() int {
	const blinkTimes = 25
	lines := io.ReadFile("inputs/11-real.txt")
	stonesStr := strings.Split(lines[0], " ")

	stones := make(map[int]int)
	for _, number := range stonesStr {
		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		stones[n]++
	}

	for i := 0; i < blinkTimes; i++ {
		//fmt.Println(stones)
		stones = blink(stones)
	}

	//fmt.Println(stones)

	sum := 0
	for _, count := range stones {
		sum += count
	}

	return sum
}

func RunPart2() int {
	const blinkTimes = 75
	lines := io.ReadFile("inputs/11-real.txt")
	stonesStr := strings.Split(lines[0], " ")

	stones := make(map[int]int)
	for _, number := range stonesStr {
		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		stones[n]++
	}

	for i := 0; i < blinkTimes; i++ {
		stones = blink(stones)
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}

	return sum
}

func blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)

	for stone, count := range stones {
		if stone == 0 {
			newStones[1] += count
			continue
		}

		if getNumOfDigitsByDivisions(stone)%2 == 0 {
			stoneA, stoneB := splitInTwo(stone)
			newStones[stoneA] += count
			newStones[stoneB] += count
			continue
		}

		newStone := stone * 2024
		newStones[newStone] += count
	}

	return newStones
}

func splitInTwo(stone int) (int, int) {
	numOfDigits := getNumOfDigitsByDivisions(stone)

	a := 0
	b := 0

	for i := 0; i < numOfDigits/2; i++ {
		b += stone % 10 * int(math.Pow10(i))
		stone /= 10
	}

	for i := 0; i < numOfDigits/2; i++ {
		a += stone % 10 * int(math.Pow10(i))
		stone /= 10
	}

	return a, b
}

func getNumOfDigitsByDivisions(number int) int {
	numOfDigits := 0
	for number > 0 {
		number /= 10
		numOfDigits++
	}

	return numOfDigits
}
