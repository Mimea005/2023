package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	var file, err = os.Open("input")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	fileBuffer := bufio.NewScanner(file)

	var possibleSum int

	for fileBuffer.Scan() {
		game, r, g, b, err := parseGame(fileBuffer.Text())
		if err != nil {
			log.Printf("Failed to parse: %s\n", fileBuffer.Text())
			log.Print(err)
			continue
		}
		if r <= maxRed && g <= maxGreen && b <= maxBlue {
			possibleSum += game
		}
	}

	fmt.Println(possibleSum)
}

func parseGame(line string) (game, r, g, b int, err error) {
	if !strings.HasPrefix(line, "Game ") {
		err = errors.New("wrong formatting")
		return
	}
	meta := strings.Split(line, ":")
	num, err := strconv.ParseInt(meta[0][len("Game "):], 10, 0)
	if err != nil {
		return
	}
	game = int(num)

	rounds := strings.Split(meta[1], ";")
	for _, round := range rounds {
		var rr, rg, rb int
		rr, rg, rb, err = parseRound(round)
		if err != nil {
			return
		}
		if rr > r {
			r = rr
		}
		if rg > g {
			g = rg
		}
		if rb > b {
			b = rb
		}

	}
	return
}

func parseRound(round string) (r, g, b int, err error) {
	groups := strings.Split(round, ",")
	if len(groups) == 0 {
		return
	}
	for _, group := range groups {
		group = strings.TrimPrefix(group, " ")
		pair := strings.Split(group, " ")
		var count int64
		count, err = strconv.ParseInt(pair[0], 10, 0)
		if err != nil {
			return
		}
		switch pair[1] {
		case "red":
			r = int(count)
		case "green":
			g = int(count)
		case "blue":
			b = int(count)
		}
	}

	return
}
