package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction string

const (
	Left  instruction = "L"
	Right instruction = "R"
)

type label string

func (label *label) ends(s string) bool {
	return string((*label)[len(*label)-1]) == s
}

type node struct {
	id    label
	left  label
	right label
}

func (node *node) nav(instr instruction) label {
	switch instr {
	case Left:
		return node.left
	case Right:
		return node.right
	}
	return node.id
}

type network map[label]node

func parseInstrs(line string) []instruction {
	instrs := make([]instruction, len(line))
	for i, r := range line {
		if string(r) == "L" {
			instrs[i] = Left
		} else {
			instrs[i] = Right
		}
	}
	return instrs
}

func parseNetwork(lines []string) network {
	network := make(network, len(lines))
	for _, line := range lines {
		node := parseNode(line)
		network[node.id] = node
	}
	return network
}

func navigate(n network, instrs []instruction, start label, end label) int {
	loc := start
	for i := 0; true; i++ {
		if loc == end {
			return i
		}
		instrIndex := i % len(instrs)
		instr := instrs[instrIndex]
		currentNode := n[loc]
		loc = currentNode.nav(instr)
	}
	return -1
}

func navigateGhost(n network, instrs []instruction, start string, end string) int {
	var starts []label
	for l := range n {
		if l.ends(start) {
			starts = append(starts, l)
		}
	}
	locs := starts
	zs := make([]int, len(starts))
	found := 0
	for i := 0; true; i++ {
		atEnd := true
		for g, loc := range locs {
			z := loc.ends(end)
			if z {
				log.Printf("Ghost %v: %v", g, i-zs[g])
				if zs[g] == 0 {
					found++
				}
				zs[g] = i
			}
			atEnd = atEnd && z
		}
		if found == len(starts) {
			return LCM(zs[0], zs[1], zs[2:]...)
		}
		if atEnd {
			return i
		}
		instrIndex := i % len(instrs)
		instr := instrs[instrIndex]
		for j, loc := range locs {
			currentNode := n[loc]
			locs[j] = currentNode.nav(instr)
		}
	}
	return -1
}

func parseNode(line string) node {
	idlr := strings.Split(strings.ReplaceAll(line, " ", ""), "=")
	id := label(idlr[0])
	lr := strings.Split(strings.ReplaceAll(strings.ReplaceAll(idlr[1], "(", ""), ")", ""), ",")
	return node{id, label(lr[0]), label(lr[1])}
}

func main() {
	fmt.Println(eightOne("input.txt"))
	fmt.Println(eightTwo("input.txt"))
}

func eightOne(fname string) int {
	lines := readLines(fname)
	instrLine := lines[0]
	networkLines := lines[2:]

	instrs := parseInstrs(instrLine)
	network := parseNetwork(networkLines)

	return navigate(network, instrs, label("AAA"), label("ZZZ"))
}

func eightTwo(fname string) int {
	lines := readLines(fname)
	instrLine := lines[0]
	networkLines := lines[2:]

	instrs := parseInstrs(instrLine)
	network := parseNetwork(networkLines)

	return navigateGhost(network, instrs, "A", "Z")
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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	res := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		res = LCM(res, integers[i])
	}
	return res
}
