package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(path string) []string {
	var lines []string

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
