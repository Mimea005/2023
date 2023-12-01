package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("[Day 01]: Hello World!")
	var res int = 0

	f, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	lines := bufio.NewScanner(f)

	for lines.Scan() {
		parsed := extractNumber(lines.Text())
		res += parsed
	}

	fmt.Println(res)
}

func extractNumber(line string) int {
	var result, last int
	first := true

	fmt.Println(line)

	for _, letter := range line {
		if letter <= '9' && letter >= '0' {
			last = int(letter - '0')
			if first {
				result = last
				first = false
			}
		}
	}
	fmt.Printf("%d%d\n", result, last)
	result *= 10
	result += last

	return result
}
