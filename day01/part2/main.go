package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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
	var numbers []int

	for i, letter := range line {
		if isDigit(letter) {
			numbers = append(numbers, int(letter-'0'))
			fmt.Print("i")
		} else {
			num, err := wordToNumber(line[i:])
			if err == nil {
				numbers = append(numbers, num)
				fmt.Print("n")
			} else {
				fmt.Print("!")
			}
		}
	}
	result := 0
	if len(numbers) > 0 {
		result = numbers[0] * 10
		result += numbers[len(numbers)-1]
	}

	fmt.Printf(" %s -> %d\n", line, result)
	return result
}

func isDigit(symbol rune) bool {
	return symbol >= '0' && symbol <= '9'
}

func wordToNumber(slice string) (int, error) {
	switch {
	case strings.HasPrefix(slice, "one"):
		return 1, nil
	case strings.HasPrefix(slice, "two"):
		return 2, nil
	case strings.HasPrefix(slice, "three"):
		return 3, nil
	case strings.HasPrefix(slice, "four"):
		return 4, nil
	case strings.HasPrefix(slice, "five"):
		return 5, nil
	case strings.HasPrefix(slice, "six"):
		return 6, nil
	case strings.HasPrefix(slice, "seven"):
		return 7, nil
	case strings.HasPrefix(slice, "eight"):
		return 8, nil
	case strings.HasPrefix(slice, "nine"):
		return 9, nil
	case strings.HasPrefix(slice, "zero"):
		return 0, nil
	}
	return 0, errors.New("not digit")
}
