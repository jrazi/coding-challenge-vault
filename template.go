package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		lines = append(lines, words)
	}

	fmt.Printf("Test: %s\n", lines[0][0])
}
