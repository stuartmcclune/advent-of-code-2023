package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type card int

const (
	Joker card = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type hand struct {
	cards []card
	bid   int
}

type category int

const (
	HighCard category = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

func createCard(label string, jIsJoker bool) card {
	switch label {
	case "2":
		return Two
	case "3":
		return Three
	case "4":
		return Four
	case "5":
		return Five
	case "6":
		return Six
	case "7":
		return Seven
	case "8":
		return Eight
	case "9":
		return Nine
	case "T":
		return Ten
	case "J":
		if jIsJoker {
			return Joker
		} else {
			return Jack
		}
	case "Q":
		return Queen
	case "K":
		return King
	case "A":
		return Ace
	}
	return Ace
}

func parseHand(line string, jIsJoker bool) hand {
	fields := strings.Fields(line)
	cards := parseCards(fields[0], jIsJoker)
	bid := atoi(fields[1])
	return hand{cards, bid}
}

func parseCards(s string, jIsJoker bool) []card {
	cards := make([]card, len(s))
	for i, c := range s {
		cards[i] = createCard(string(c), jIsJoker)
	}
	return cards
}

func (hand *hand) category() category {
	counts := make(map[card]int)
	highestCount := 0
	var commonCard card
	for _, c := range hand.cards {
		counts[c]++
		if counts[c] > highestCount && c != Joker {
			highestCount = counts[c]
			commonCard = c
		}
	}

	jokers := counts[Joker]
	if jokers == 5 {
		return FiveKind
	}

	counts[commonCard] += jokers

	pairs := 0
	threes := 0

	for k, v := range counts {
		if k == Joker {
			continue
		}
		if v == 5 {
			return FiveKind
		}
		if v == 4 {
			return FourKind
		}
		if v == 3 {
			threes++
		}
		if v == 2 {
			pairs++
		}
	}

	if threes == 1 {
		if pairs == 1 {
			return FullHouse
		}
		return ThreeKind
	}

	switch pairs {
	case 2:
		return TwoPair
	case 1:
		return OnePair
	}
	return HighCard

}

func (hand *hand) points(rank int) int {
	return hand.bid * rank
}

func parseHands(lines []string, jIsJoker bool) []hand {
	hands := make([]hand, len(lines))
	for i, l := range lines {
		hands[i] = parseHand(l, jIsJoker)
	}
	return hands
}

func compareHands(a hand, b hand) int {
	comp := cmp.Compare(a.category(), b.category())
	if comp != 0 {
		return comp
	}

	for i := 0; i < 5; i++ {
		comp := cmp.Compare(a.cards[i], b.cards[i])
		if comp != 0 {
			return comp
		}
	}
	return 0
}

func main() {
	fmt.Println(sevenOne("input.txt"))
	fmt.Println(sevenTwo("input.txt"))
}

func sevenOne(fname string) int {
	hands := parseHands(readLines(fname), false)
	slices.SortFunc(hands, compareHands)
	points := 0
	for i, h := range hands {
		points += h.points(i + 1)
	}
	return points
}

func sevenTwo(fname string) int {
	hands := parseHands(readLines(fname), true)
	slices.SortFunc(hands, compareHands)
	points := 0
	for i, h := range hands {
		points += h.points(i + 1)
	}
	return points
}

func readLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
