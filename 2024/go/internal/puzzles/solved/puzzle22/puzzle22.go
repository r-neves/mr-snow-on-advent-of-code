package puzzle22

import (
	"aoc2024/internal/io"
	"strconv"
)

func RunPart1() int {
	const iterations = 2000

	lines := io.ReadFile("inputs/22-real.txt")

	sum := 0
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		for i := 0; i < iterations; i++ {
			newNumber := calculateNextSecret(number)
			number = newNumber
		}
		sum += number
	}

	return sum
}

func calculateNextSecret(number int) int {
	const divisor = 16777216

	mul64 := ((number << 6) ^ number) % divisor
	div32 := ((mul64 >> 5) ^ mul64) % divisor
	mul2048 := ((div32 << 11) ^ div32) % divisor
	return mul2048
}

type marketSequence struct {
	Change1 int
	Change2 int
	Change3 int
	Change4 int
}

func RunPart2() int {
	const iterations = 2000

	lines := io.ReadFile("inputs/22-real.txt")

	pricesPerBuyer := make([]map[marketSequence]int, len(lines))
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		price := number % 10
		buyerPrices := make(map[marketSequence]int, iterations)
		marketWindow := make([]int, 0, iterations)
		for i := 0; i < iterations; i++ {
			newNumber := calculateNextSecret(number)
			newPrice := newNumber % 10

			marketWindow = append(marketWindow, newPrice-price)
			if i >= 3 {
				marketSeq := marketSequence{
					Change1: marketWindow[i-3],
					Change2: marketWindow[i-2],
					Change3: marketWindow[i-1],
					Change4: marketWindow[i],
				}

				_, exists := buyerPrices[marketSeq]
				if !exists {
					buyerPrices[marketSeq] = newPrice
				}
			}

			number = newNumber
			price = newPrice
		}

		pricesPerBuyer = append(pricesPerBuyer, buyerPrices)
	}

	bestSellingPrice := 0
	globalPrices := make(map[marketSequence]int)
	for _, buyerPrices := range pricesPerBuyer {
		for marketSeq, price := range buyerPrices {
			globalPrices[marketSeq] += price
			if globalPrices[marketSeq] > bestSellingPrice {
				bestSellingPrice = globalPrices[marketSeq]
			}
		}
	}

	return bestSellingPrice
}
