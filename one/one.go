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

func main() {
	lines := readLines("input1.txt")

	fmt.Println(calibrate2(lines))

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

func calibrate(document []string) uint64 {
	sum := uint64(0)

	for _, line := range document {
		first, last := firstAndLast(line)
		toAdd := combine(first, last)
		sum += toAdd
	}

	return sum
}

func firstAndLast(line string) (uint64, uint64) {
	runes := []rune(line)
	var digits []uint64

	for _, r := range runes {
		digit, err := runeToNum(r)
		if err == nil {
			digits = append(digits, digit)
		}
	}

	return digits[0], digits[len(digits)-1]
}

func firstAndLast2(line string) (uint64, uint64) {
	spellings := make([]string, len(numberWords))
	i := 0
	for spelling := range numberWords {
		spellings[i] = spelling
		i++
	}

	var digits []uint64
	letters := strings.Split(line, "")

	for i, c := range letters {
		digit, err := letterToNum(c)
		if err == nil {
			digits = append(digits, digit)
			continue
		}

		remaining := strings.Join(letters[i:], "")

		for _, spelling := range spellings {
			if strings.HasPrefix(remaining, spelling) {
				digits = append(digits, numberWords[spelling])
				break
			}
		}

	}

	return digits[0], digits[len(digits)-1]
}

func combine(first uint64, last uint64) uint64 {
	firstS := strconv.FormatUint(first, 10)
	lastS := strconv.FormatUint(last, 10)
	combined := firstS + lastS
	combinedUint, _ := strconv.ParseUint(combined, 10, 64)
	return combinedUint

}

func runeToNum(r rune) (uint64, error) {
	if '0' <= r && r <= '9' {
		return uint64(r) - '0', nil
	}
	return 0, errors.New("NaN")
}

func letterToNum(c string) (uint64, error) {
	return strconv.ParseUint(c, 10, 64)
}

func calibrate2(document []string) uint64 {
	sum := uint64(0)

	for _, line := range document {
		first, last := firstAndLast2(line)
		toAdd := combine(first, last)
		sum += toAdd
	}

	return sum
}

var numberWords = map[string]uint64{
	// "zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
