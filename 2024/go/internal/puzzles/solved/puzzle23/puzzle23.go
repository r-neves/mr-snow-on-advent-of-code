package puzzle23

import (
	"aoc2024/internal/io"
	"slices"
	"strings"
)

type node struct {
	Name  string
	Links map[string]*node
}

func RunPart1() int {
	lines := io.ReadFile("inputs/23-real.txt")

	nodes := make(map[string]*node)
	for _, line := range lines {
		machines := strings.Split(line, "-")

		machine1, ok := nodes[machines[0]]
		if !ok {
			machine1 = &node{Name: machines[0], Links: make(map[string]*node)}
			nodes[machines[0]] = machine1
		}

		machine2, ok := nodes[machines[1]]
		if !ok {
			machine2 = &node{Name: machines[1], Links: make(map[string]*node)}
			nodes[machines[1]] = machine2
		}

		machine1.Links[machine2.Name] = machine2
		machine2.Links[machine1.Name] = machine1
	}

	trios := make(map[string][]*node)
	for _, n1 := range nodes {
		for _, n2 := range n1.Links {
			for _, n3 := range n2.Links {
				if _, found := n1.Links[n3.Name]; found {
					names := []string{n1.Name, n2.Name, n3.Name}
					slices.Sort(names)
					key := strings.Join(names, ",")
					trios[key] = []*node{n1, n2, n3}
				}
			}
		}
	}

	var sortedKeys []string
	for key := range trios {
		sortedKeys = append(sortedKeys, key)
	}

	slices.Sort(sortedKeys)

	sum := 0

	for _, key := range sortedKeys {
		//fmt.Println(key)
		if trios[key][0].Name[0] == 't' || trios[key][1].Name[0] == 't' || trios[key][2].Name[0] == 't' {
			sum++
		}
	}

	return sum
}

func RunPart2() string {
	lines := io.ReadFile("inputs/23-real.txt")

	nodes := make(map[string]*node)
	for _, line := range lines {
		machines := strings.Split(line, "-")

		machine1, ok := nodes[machines[0]]
		if !ok {
			machine1 = &node{Name: machines[0], Links: make(map[string]*node)}
			nodes[machines[0]] = machine1
		}

		machine2, ok := nodes[machines[1]]
		if !ok {
			machine2 = &node{Name: machines[1], Links: make(map[string]*node)}
			nodes[machines[1]] = machine2
		}

		machine1.Links[machine2.Name] = machine2
		machine2.Links[machine1.Name] = machine1
	}

	// For each node:
	// - Check if every link is connected to every other link

	var interconnectedSets []map[string]*node
	for _, n1 := range nodes {
		for _, n2 := range n1.Links {
			set := make(map[string]*node)
			set[n1.Name] = n1
			set[n2.Name] = n2
			for _, n3 := range n1.Links {
				if n2 == n3 {
					continue
				}

				if _, found := n2.Links[n3.Name]; found {
					belongs := true
					for _, n := range set {
						if _, found := n3.Links[n.Name]; !found {
							belongs = false
							break
						}
					}

					if belongs {
						set[n3.Name] = n3
					}
				}
			}

			interconnectedSets = append(interconnectedSets, set)
		}
	}

	maxIndex := 0
	maxNodes := 0
	for i, set := range interconnectedSets {
		if len(set) > maxNodes {
			maxNodes = len(set)
			maxIndex = i
		}
	}

	var largestSet []string
	for name := range interconnectedSets[maxIndex] {
		largestSet = append(largestSet, name)
	}

	slices.Sort(largestSet)

	return strings.Join(largestSet, ",")
}
