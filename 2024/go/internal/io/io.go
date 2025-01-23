package io

import (
	"bufio"
	"os"
)

func ReadFile(filepath string) []string {
	var lines []string

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
