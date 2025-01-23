package puzzle5

import (
	"aoc2024/internal/io"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type pair struct {
	Prev int
	Next int
}

func RunPart1() {
	lines := io.ReadFile("inputs/5-real.txt")

	var rules []pair
	pagesIndex := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			pagesIndex = i + 1
			break
		}
		rule := strings.Split(lines[i], "|")
		prevNumber, err := strconv.Atoi(rule[0])
		if err != nil {
			panic(err)
		}
		nextNumber, err := strconv.Atoi(rule[1])
		if err != nil {
			panic(err)
		}
		rules = append(rules, pair{prevNumber, nextNumber})
	}

	sum := 0

	for i := pagesIndex; i < len(lines); i++ {
		pagesStr := strings.Split(lines[i], ",")
		pageIndexMap := make(map[int]int, len(pagesStr))
		for j, pageStr := range pagesStr {
			page, err := strconv.Atoi(pageStr)
			if err != nil {
				panic(err)
			}
			pageIndexMap[page] = j
		}

		valid := true
		for _, rulePair := range rules {
			if _, ok := pageIndexMap[rulePair.Prev]; ok {
				if _, ok := pageIndexMap[rulePair.Next]; ok {
					if pageIndexMap[rulePair.Prev] >= pageIndexMap[rulePair.Next] {
						valid = false
						break
					}
				}
			}
		}

		if valid {
			middleStr := pagesStr[len(pagesStr)/2]
			middle, err := strconv.Atoi(middleStr)
			if err != nil {
				panic(err)
			}
			sum += middle
		}
	}

	fmt.Println(sum)
}

func RunPart2() {
	lines := io.ReadFile("inputs/5-real.txt")

	var rules []pair
	pagesIndex := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			pagesIndex = i + 1
			break
		}
		rule := strings.Split(lines[i], "|")
		prevNumber, err := strconv.Atoi(rule[0])
		if err != nil {
			panic(err)
		}
		nextNumber, err := strconv.Atoi(rule[1])
		if err != nil {
			panic(err)
		}
		rules = append(rules, pair{prevNumber, nextNumber})
	}

	sum := 0

	for i := pagesIndex; i < len(lines); i++ {
		pagesStr := strings.Split(lines[i], ",")
		pages := make([]int, 0, len(pagesStr))
		pageIndexMap := make(map[int]int, len(pagesStr))
		for j, pageStr := range pagesStr {
			page, err := strconv.Atoi(pageStr)
			if err != nil {
				panic(err)
			}
			pageIndexMap[page] = j
			pages = append(pages, page)
		}

		valid := isValid(pageIndexMap, rules)
		if valid {
			continue
		}

		slices.SortFunc(pages, func(i, j int) int {
			for _, rulePair := range rules {
				if rulePair.Next == i && rulePair.Prev == j {
					return 1
				}
			}

			return -1
		})

		middle := pages[len(pagesStr)/2]
		sum += middle
	}

	fmt.Println(sum)
}

func isValid(pageIndexMap map[int]int, rules []pair) bool {
	for _, rulePair := range rules {
		if _, ok := pageIndexMap[rulePair.Prev]; ok {
			if _, ok := pageIndexMap[rulePair.Next]; ok {
				if pageIndexMap[rulePair.Prev] >= pageIndexMap[rulePair.Next] {
					return false
				}
			}
		}
	}

	return true
}
