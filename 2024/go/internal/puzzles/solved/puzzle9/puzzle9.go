package puzzle9

import (
	"aoc2024/internal/io"
	"fmt"
)

func RunPart1() int {
	lines := io.ReadFile("inputs/9-real.txt")
	diskBlocks := lines[0]

	var disk []int
	isFreeSpace := false

	index := 0
	for _, c := range diskBlocks {
		if isFreeSpace {
			for i := 0; i < int(c)-48; i++ {
				disk = append(disk, -1)
			}
		} else {
			for i := 0; i < int(c)-48; i++ {
				disk = append(disk, index)
			}
			index++
		}
		isFreeSpace = !isFreeSpace
	}

	freeSpotIdx := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			freeSpotIdx = i
			break
		}
	}

	for i := len(disk) - 1; i > freeSpotIdx; i-- {
		if disk[i] != -1 {
			disk[freeSpotIdx], disk[i] = disk[i], -1
		}

		for disk[freeSpotIdx] != -1 || freeSpotIdx == len(disk) {
			freeSpotIdx++
		}
	}

	sum := 0

	freeSpotIdx = 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			freeSpotIdx = i
			break
		}
	}

	for i := 0; i < freeSpotIdx; i++ {
		sum += disk[i] * i
	}

	return sum
}

type diskBlock struct {
	value     int
	size      int
	diskIndex int
}

func RunPart2() int {
	lines := io.ReadFile("inputs/9-real.txt")

	var disk []int
	isFreeSpace := false
	index := 0
	var filledDiskBlocks []diskBlock

	for _, c := range lines[0] {
		if isFreeSpace {
			for i := 0; i < int(c)-48; i++ {
				disk = append(disk, -1)
			}
		} else {
			block := diskBlock{
				value:     index,
				size:      int(c) - 48,
				diskIndex: len(disk),
			}
			filledDiskBlocks = append(filledDiskBlocks, block)

			for i := 0; i < int(c)-48; i++ {
				disk = append(disk, index)
			}
			index++
		}
		isFreeSpace = !isFreeSpace
	}

	for i := len(filledDiskBlocks) - 1; i > 0; i-- {
		freeBlockIdx := findFreeBlock(disk, filledDiskBlocks[i].size)
		if freeBlockIdx != -1 && freeBlockIdx < filledDiskBlocks[i].diskIndex {
			for j := freeBlockIdx; j < freeBlockIdx+filledDiskBlocks[i].size; j++ {
				disk[j] = filledDiskBlocks[i].value
			}

			for j := filledDiskBlocks[i].diskIndex; j < filledDiskBlocks[i].diskIndex+filledDiskBlocks[i].size; j++ {
				disk[j] = -1
			}
		}
	}

	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}

		sum += disk[i] * i
	}

	return sum
}

func findFreeBlock(disk []int, size int) int {
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			potentialFreeBlock := i
			for j := potentialFreeBlock; j < potentialFreeBlock+size; j++ {
				if len(disk) <= j || disk[j] != -1 {
					potentialFreeBlock = -1
					break
				}
			}

			if potentialFreeBlock != -1 {
				return potentialFreeBlock
			}
		}
	}

	return -1
}

func printDisk(disk []int) {
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			fmt.Print(".")
			continue
		}

		fmt.Print(disk[i])
	}
	fmt.Println()
}
