package puzzle19

import (
	ds "aoc2024/internal/datastructures"
	"aoc2024/internal/io"
	"fmt"
	"strings"
)

var iterations = 0

func RunPart1() int {
	lines := io.ReadFile("inputs/19-real.txt")

	availableTowels, designs := parseInput(lines)
	availableTowels = reduceTowelSet(availableTowels)

	trie := ds.NewCharTrie()
	for towel := range availableTowels {
		trie.Insert(towel)
	}

	count := 0
	for _, design := range designs {

		possible := findTowelPattern(design, trie, 0)
		//fmt.Println("Evaluated design", i+1, "/", len(designs), ":", design, "=>", possible)
		if possible {
			count++
		}

	}

	return count
}

func findTowelPattern(design string, trie *ds.CharTrie, currIndex int) bool {
	if currIndex == len(design) {
		return true
	}

	node := trie.FindLastPossibleChar(design[currIndex:])
	for node != nil {
		if !node.IsWordEnd {
			node = node.Parent
			continue
		}

		if findTowelPattern(design, trie, currIndex+node.Depth) {
			return true
		}

		node = node.Parent
	}

	return false
}

func RunPart2() int {
	lines := io.ReadFile("inputs/19-real.txt")

	availableTowels, designs := parseInput(lines)
	reducedTowels := reduceTowelSet(availableTowels)

	reducedTrie := ds.NewCharTrie()
	for towel := range reducedTowels {
		reducedTrie.Insert(towel)
	}

	count := 0

	for _, design := range designs {
		possible := findTowelPattern(design, reducedTrie, 0)
		//fmt.Println("Evaluated design", i+1, "/", len(designs), ":", design, "=>", possible)
		if !possible {
			continue
		}

		usableTowels := make(map[string]bool)
		for towel := range availableTowels {
			if strings.Contains(design, towel) {
				usableTowels[towel] = true
			}
		}

		usableTrie := ds.NewCharTrie()
		for towel := range usableTowels {
			usableTrie.Insert(towel)
		}
		fmt.Println("Built usable trie for design", design, "with", len(usableTowels), "towels")

		possibleCombs := findAllPossibleTowelPatterns(design, usableTrie, 0)
		fmt.Println("Found", possibleCombs, "possible combinations for design", design)
		count += possibleCombs
	}

	return count
}

func findAllPossibleTowelPatterns(design string, trie *ds.CharTrie, currIndex int) int {
	if currIndex == len(design) {
		return 1
	}

	count := 0
	node := trie.FindLastPossibleChar(design[currIndex:])
	for node != nil {
		if !node.IsWordEnd {
			node = node.Parent
			continue
		}

		count += findAllPossibleTowelPatterns(design, trie, currIndex+node.Depth)

		node = node.Parent
	}

	return count
}

func parseInput(lines []string) (map[string]bool, []string) {
	var designs []string
	availableTowels := make(map[string]bool)

	availableTowelsSlice := strings.Split(lines[0], ", ")
	for _, towel := range availableTowelsSlice {
		availableTowels[towel] = true
	}

	for _, line := range lines[2:] {
		designs = append(designs, line)
	}

	return availableTowels, designs
}

func reduceTowelSet(towels map[string]bool) map[string]bool {
	newTowels := make(map[string]bool)

	for towel := range towels {
		trie := ds.NewCharTrie()
		for t := range towels {
			if t == towel {
				continue
			}
			trie.Insert(t)
		}

		if !findTowelPattern(towel, trie, 0) {
			newTowels[towel] = true
		}
	}

	fmt.Println("Reduced towel set from", len(towels), "to", len(newTowels))

	return newTowels
}
