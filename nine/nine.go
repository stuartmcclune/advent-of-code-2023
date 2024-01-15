package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sequence []int

func (seq *sequence) diffs() sequence {
	res := make(sequence, len(*seq)-1)
	for i := 0; i < len(*seq)-1; i++ {
		res[i] = (*seq)[i+1] - (*seq)[i]
	}
	return res
}

func (seq *sequence) isZero() bool {
	for _, i := range *seq {
		if i != 0 {
			return false
		}
	}
	return true
}

func (seq *sequence) extrapolate() int {
	if seq.isZero() {
		return 0
	}

	diffs := seq.diffs()
	e := diffs.extrapolate()
	return (*seq)[len(*seq)-1] + e
}

func (seq *sequence) extrapolateBack() int {
	if seq.isZero() {
		return 0
	}

	diffs := seq.diffs()
	e := diffs.extrapolateBack()
	return (*seq)[0] - e
}

func parseSequence(line string) sequence {
	var seq sequence
	for _, s := range strings.Fields(line) {
		seq = append(seq, atoi(s))
	}
	return seq
}

func main() {
	fmt.Println(nineOne("input.txt"))
	fmt.Println(nineTwo("input.txt"))
}

func nineOne(fname string) int {
	total := 0
	for _, l := range readLines(fname) {
		seq := parseSequence(l)
		total += seq.extrapolate()
	}
	return total
}

func nineTwo(fname string) int {
	total := 0
	for _, l := range readLines(fname) {
		seq := parseSequence(l)
		total += seq.extrapolateBack()
	}
	return total
}

func readLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
