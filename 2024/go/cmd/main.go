package main

import (
	"aoc2024/internal/puzzles/puzzle25"
	"fmt"
	"os"
	"runtime/pprof"
)

const doProfiling = false

func main() {
	if doProfiling {
		f, err := os.Create("cpu.pprof")
		if err != nil {
			panic(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			panic(err)
		}

		defer pprof.StopCPUProfile()
	}

	result := puzzle25.RunPart1()

	fmt.Println(result)
}
