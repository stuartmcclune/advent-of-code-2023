package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type symbol string

const (
	Operational symbol = "."
	Damaged     symbol = "#"
	Unknown     symbol = "?"
)

type row struct {
	symbols []symbol
	groups  []int
}

func (r row) arrangements() int {
	memo := make(map[string]entry)
	return arrangements(r.symbols, r.groups, false, false, &memo)
}

func (r *row) unfold() *row {
	var us []symbol
	us = append(us, r.symbols...)
	us = append(us, Unknown)
	us = append(us, r.symbols...)
	us = append(us, Unknown)
	us = append(us, r.symbols...)
	us = append(us, Unknown)
	us = append(us, r.symbols...)
	us = append(us, Unknown)
	us = append(us, r.symbols...)

	var ug []int
	ug = append(ug, r.groups...)
	ug = append(ug, r.groups...)
	ug = append(ug, r.groups...)
	ug = append(ug, r.groups...)
	ug = append(ug, r.groups...)

	r.symbols = us
	r.groups = ug
	return r
}

func hash(symbols []symbol, groups []int, took bool, wasLast bool) string {
	s := ""
	for _, sym := range symbols {
		s += string(sym)
	}
	for _, g := range groups {
		s += strconv.Itoa(g)
	}
	if took {
		s += "t"
	} else {
		s += "f"
	}
	if wasLast {
		s += "t"
	} else {
		s += "f"
	}
	return s
}

type entry struct {
	hit   bool
	count int
}

func arrangements(symbols []symbol, groups []int, took bool, wasLast bool, memo *map[string]entry) int {
	hash := hash(symbols, groups, took, wasLast)
	e := (*memo)[hash]
	if e.hit {
		return e.count
	}

	symbol := symbols[0]

	if wasLast && symbol == Damaged {
		return 0
	}
	if took && !wasLast && symbol == Operational {
		return 0
	}

	if len(symbols) == 1 {
		if len(groups) == 0 {
			if symbol == Damaged {
				(*memo)[hash] = entry{true, 0}
				return 0
			}
			(*memo)[hash] = entry{true, 1}
			return 1
		} else {
			if wasLast {
				(*memo)[hash] = entry{true, 0}
				return 0
			}
			if len(groups) > 1 {
				(*memo)[hash] = entry{true, 0}
				return 0
			}
			if symbol == Operational {
				(*memo)[hash] = entry{true, 0}
				return 0
			}
			if groups[0] != 1 {
				(*memo)[hash] = entry{true, 0}
				return 0
			}
			(*memo)[hash] = entry{true, 1}
			return 1
		}
	}

	couldBeO := !(took && !wasLast)
	couldBeD := !wasLast && len(groups) != 0

	if symbol == Operational {
		if !couldBeO {
			(*memo)[hash] = entry{true, 0}
			return 0
		}
		ret := arrangements(symbols[1:], groups, false, false, memo)
		(*memo)[hash] = entry{true, ret}
		return ret
	}

	if symbol == Damaged {
		if !couldBeD {
			(*memo)[hash] = entry{true, 0}
			return 0
		}

		gs := make([]int, len(groups))
		copy(gs, groups)
		gs[0]--
		wasLast := false
		if gs[0] == 0 {
			wasLast = true
			gs = gs[1:]
		}

		ret := arrangements(symbols[1:], gs, true, wasLast, memo)
		(*memo)[hash] = entry{true, ret}
		return ret
	}

	// Unknown.

	count := 0
	if couldBeO {
		count += arrangements(symbols[1:], groups, false, false, memo)
	}

	if couldBeD {
		gs := make([]int, len(groups))
		copy(gs, groups)
		gs[0]--
		wasLast := false
		if gs[0] == 0 {
			wasLast = true
			gs = gs[1:]
		}

		count += arrangements(symbols[1:], gs, true, wasLast, memo)
	}

	(*memo)[hash] = entry{true, count}
	return count
}

func parseRow(line string) row {
	fields := strings.Fields(line)
	ss := fields[0]
	gs := strings.Split(fields[1], ",")

	symbols := make([]symbol, len(ss))

	for i, r := range ss {
		symbols[i] = symbol(string(r))
	}

	groups := make([]int, len(gs))

	for i, g := range gs {
		x, err := strconv.Atoi(g)
		if err != nil {
			log.Fatal(err)
		}
		groups[i] = x
	}

	return row{symbols, groups}
}

func main() {
	fmt.Println(twelveOne("input.txt"))
	fmt.Println(twelveTwo("input.txt"))
}

func twelveOne(fname string) int {
	sum := 0
	for _, l := range readlines(fname) {
		sum += parseRow(l).arrangements()
	}
	return sum
}

func twelveTwo(fname string) int {
	sum := 0
	for i, l := range readlines(fname) {
		row := parseRow(l)
		row = *row.unfold()
		sum += row.arrangements()
		log.Printf("Done %v", i)
	}
	return sum
}

func readlines(fname string) []string {
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
