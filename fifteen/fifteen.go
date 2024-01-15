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

func main() {
	fmt.Println(fifteenOne("input.txt"))
	fmt.Println(fifteenTwo("input.txt"))
}

func fifteenOne(fname string) int {
	sum := 0
	l := readLines(fname)[0]
	steps := strings.Split(l, ",")
	for _, step := range steps {
		sum += hash(step)
	}
	return sum
}

type lens struct {
	label       string
	focalLength int
}

type instr struct {
	label       string
	hash        int
	focalLength int
}

func (i instr) isRemove() bool {
	return i.focalLength == -1
}

func parseInstr(s string) instr {
	if strings.Contains(s, "-") {
		label := s[0 : len(s)-1]
		return instr{label, hash(label), -1}
	} else {
		label := s[0 : len(s)-2]
		return instr{label, hash(label), rtoi(s[len(s)-1:])}
	}

}

func rtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func fifteenTwo(fname string) int {
	boxes := make(map[int][]lens)
	l := readLines(fname)[0]
	steps := strings.Split(l, ",")
	for _, step := range steps {
		instr := parseInstr(step)
		box := boxes[instr.hash]
		if instr.isRemove() {
			box = slices.DeleteFunc(box, func(l lens) bool { return l.label == instr.label })
		} else {
			matchIdx := slices.IndexFunc(box, func(l lens) bool { return l.label == instr.label })
			if matchIdx != -1 {
				box[matchIdx].focalLength = instr.focalLength
			} else {
				box = append(box, lens{instr.label, instr.focalLength})
			}
		}
		boxes[instr.hash] = box
	}

	fp := 0
	for boxNum, box := range boxes {
		for i, lens := range box {
			fp += (1 + boxNum) * (1 + i) * lens.focalLength
		}
	}
	return fp
}

func hash(s string) int {
	val := 0

	for _, r := range s {
		ascii := int(r)
		val += ascii
		val *= 17
		val %= 256
	}

	return val
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
