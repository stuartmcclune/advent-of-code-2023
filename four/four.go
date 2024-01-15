package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	id      int
	winning []int
	numbers []int
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println(calculatePoints("input.txt"))
	fmt.Println(calclateNumberOfCards("input.txt"))
}

func calculatePoints(fname string) int {
	lines := readLines(fname)
	points := 0
	for _, l := range lines {
		card := parseCard(l)
		points += card.value()
	}
	return points
}

func (card *card) value() int {
	value := 0

	for _, n := range card.numbers {
		if slices.Contains(card.winning, n) {
			if value == 0 {
				value = 1
			} else {
				value *= 2
			}
		}
	}

	return value
}

func (card *card) matches() int {
	count := 0

	for _, n := range card.numbers {
		if slices.Contains(card.winning, n) {
			count += 1
		}
	}

	return count
}

func (card *card) winnings() []int {
	matches := card.matches()
	winnings := make([]int, matches)
	for i := range winnings {
		winnings[i] = card.id + i + 1
	}
	return winnings
}

func parseCard(s string) card {
	s1 := strings.Split(s, ":")

	id, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(s1[0], "Card ")))

	if err != nil {
		log.Fatal(err)
	}

	wsns := strings.Split(s1[1], "|")
	ws := parseNumberStrip(wsns[0])
	ns := parseNumberStrip(wsns[1])
	return card{id, ws, ns}
}

func parseNumberStrip(s string) []int {
	nstrs := slices.DeleteFunc(strings.Split(s, " "), isSpaceOrEmpty)
	ns := make([]int, len(nstrs))
	for i, nstr := range nstrs {
		n, err := strconv.Atoi(nstr)
		if err != nil {
			log.Fatal(err)
		}
		ns[i] = n
	}
	return ns
}

func isSpaceOrEmpty(s string) bool {
	return s == " " || s == ""
}

func readLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func calclateNumberOfCards(fname string) int {
	lines := readLines(fname)
	cards := make([]card, len(lines))
	for i, l := range lines {
		cards[i] = parseCard(l)
	}
	slices.Reverse(cards)

	winnings := make(map[int]int, len(cards))
	for _, card := range cards {
		sum := 1
		for _, id := range card.winnings() {
			sum += winnings[id]
		}
		winnings[card.id] = sum
	}

	won := 0
	for _, w := range winnings {
		won += w
	}
	return won
}
