package puzzle1

import (
	"aoc2024/internal/datastructures"
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunPart1() {
	list1, list2 := readInputToHeaps("inputs/1-real.txt")

	sum := 0
	for list1.Len() > 0 {
		sum += heap.Pop(list1).(int) - heap.Pop(list2).(int)
	}

	fmt.Println(sum)
}

func RunPart2() {
	map1, map2 := readInputToMaps("inputs/1-real.txt")

	sum := 0
	for key, _ := range map1 {
		sum += key * map1[key] * map2[key]
	}

	fmt.Println(sum)
}

func readInputToHeaps(filename string) (*datastructures.IntHeap, *datastructures.IntHeap) {
	list1 := &datastructures.IntHeap{}
	list2 := &datastructures.IntHeap{}
	heap.Init(list1)
	heap.Init(list2)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "   ")
		first, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}

		heap.Push(list1, first)
		heap.Push(list2, second)
	}

	return list1, list2
}

func readInputToMaps(filename string) (map[int]int, map[int]int) {
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "   ")
		first, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}

		map1[first]++
		map2[second]++
	}

	return map1, map2
}
