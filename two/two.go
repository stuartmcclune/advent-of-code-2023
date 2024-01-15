package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id      int
	samples []counts
}

type counts struct {
	red   int
	green int
	blue  int
}

func (c *counts) power() int {
	return c.red * c.green * c.blue
}

func main() {
	fmt.Println(twoTwo("input1.txt"))
}

func twoOne(fname string, constraint counts) int {
	lines := readLines(fname)
	games := parseGames(lines)
	possibleGames := assessPossibilities(games, constraint)
	return sumIds(possibleGames)
}

func readLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Could not open file: %v, error: %v", fname, err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseGames(lines []string) []game {
	var games []game
	for _, line := range lines {
		games = append(games, parseGame(line))
	}
	return games
}

func parseGame(line string) game {
	split := strings.Split(line, ": ")
	gameId := split[0]
	rest := split[1]

	idStr := strings.TrimPrefix(gameId, "Game ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("Could not parse game id '%v', error: %v", idStr, err)
	}
	sampleStrs := strings.Split(rest, "; ")

	samples := parseSamples(sampleStrs)

	return game{id: id, samples: samples}
}

func parseSamples(ss []string) []counts {
	samples := make([]counts, len(ss))
	for i, s := range ss {
		samples[i] = parseSample(s)
	}
	return samples
}

func parseSample(s string) counts {
	parts := strings.Split(s, ", ")
	var red int
	var green int
	var blue int
	for _, part := range parts {
		s := strings.Split(part, " ")
		countStr := s[0]
		count, err := strconv.Atoi(countStr)
		if err != nil {
			log.Fatal(err)
		}

		colour := s[1]
		switch colour {
		case "red":
			red = count
		case "green":
			green = count
		case "blue":
			blue = count

		}
	}
	return counts{red: red, green: green, blue: blue}
}

func assessPossibilities(games []game, constraints counts) []game {
	var possibleGames []game
	for _, game := range games {
		if isPossible(game, constraints) {
			possibleGames = append(possibleGames, game)
		}
	}
	return possibleGames
}

func isPossible(game game, constraints counts) bool {
	samples := game.samples
	for _, sample := range samples {
		if sample.red > constraints.red || sample.green > constraints.green || sample.blue > constraints.blue {
			return false
		}
	}
	return true
}

func sumIds(games []game) int {
	sum := 0
	for _, game := range games {
		sum += game.id
	}
	return sum
}

func twoTwo(fname string) int {
	lines := readLines(fname)
	games := parseGames(lines)
	mins := make([]counts, len(games))
	for i, game := range games {
		mins[i] = minRequired(game)
	}
	return sumPowers(mins)
}

func sumPowers(counts []counts) int {
	sum := 0
	for _, count := range counts {
		sum += count.power()
	}
	return sum
}

func minRequired(game game) counts {
	minRequired := counts{}
	for _, sample := range game.samples {
		minRequired.red = max(sample.red, minRequired.red)
		minRequired.green = max(sample.green, minRequired.green)
		minRequired.blue = max(sample.blue, minRequired.blue)
	}
	return minRequired
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
