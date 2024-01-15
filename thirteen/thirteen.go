package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type symbol string

const (
	Ash   symbol = "."
	Rocks symbol = "#"
)

var InvalidMirror = mirror{false, -1}

type pattern struct {
	width, height int
	symbols       [][]symbol
}

type mirror struct {
	horizontal bool
	after      int
}

func parsePattern(lines []string) pattern {
	height := len(lines)
	width := len(lines[0])

	symbols := make([][]symbol, width)

	for x := 0; x < width; x++ {
		symbols[x] = make([]symbol, height)
		for y := 0; y < height; y++ {
			symbols[x][y] = symbol(string(lines[y][x]))
		}
	}
	return pattern{width, height, symbols}
}

func (p pattern) fixSmudge() mirror {
	m := p.findMirror()

	for x := 0; x < p.width; x++ {
		for y := 0; y < p.height; y++ {
			dup := make([][]symbol, len(p.symbols))
			for i := range p.symbols {
				dup[i] = make([]symbol, len(p.symbols[i]))
				copy(dup[i], p.symbols[i])
			}
			if dup[x][y] == Rocks {
				dup[x][y] = Ash
			} else {
				dup[x][y] = Rocks
			}
			np := pattern{p.width, p.height, dup}
			nm := np.findNewMirror(m)
			if nm != InvalidMirror {
				return nm
			}
		}
	}
	return InvalidMirror
}

func (p pattern) findNewMirror(old mirror) mirror {
	for xm := 0; xm < p.width-1; xm++ {
		found := true
		maxDx := min(xm+1, p.width-(xm+1))
		for dx := 1; dx <= maxDx; dx++ {
			colL := p.col(xm + 1 - dx)
			colR := p.col(xm + dx)
			if !equalSlices(colL, colR) {
				found = false
			}
		}
		if found {
			m := mirror{false, xm + 1}
			if m != old {
				return m
			}
		}
	}
	for ym := 0; ym < p.height-1; ym++ {
		found := true
		maxDy := min(ym+1, p.height-(ym+1))
		for dy := 1; dy <= maxDy; dy++ {
			rowL := p.row(ym + 1 - dy)
			rowR := p.row(ym + dy)
			if !equalSlices(rowL, rowR) {
				found = false
			}
		}
		if found {
			m := mirror{true, ym + 1}
			if m != old {
				return m
			}
		}
	}
	return InvalidMirror
}

func (p pattern) findMirror() mirror {
	return p.findNewMirror(InvalidMirror)
}

func (p pattern) col(i int) []symbol {
	return p.symbols[i]
}

func (p pattern) row(i int) []symbol {
	row := make([]symbol, p.width)
	for x := 0; x < p.width; x++ {
		row[x] = p.symbols[x][i]
	}
	return row
}

func equalSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if t != b[i] {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m mirror) score() int {
	mult := 1
	if m.horizontal {
		mult = 100
	}
	return m.after * mult
}

func main() {
	fmt.Println(thirteenOne("input.txt"))
	fmt.Println(thirteenTwo("input.txt"))
}

func thirteenOne(fname string) int {
	lines := readLines(fname)
	sum := 0
	var pls []string
	for _, l := range lines {
		if l == "" {
			sum += parsePattern(pls).findMirror().score()

			pls = nil
		} else {
			pls = append(pls, l)
		}
	}
	sum += parsePattern(pls).findMirror().score()
	return sum
}

func thirteenTwo(fname string) int {
	lines := readLines(fname)
	sum := 0
	var pls []string
	for _, l := range lines {
		if l == "" {
			sum += parsePattern(pls).fixSmudge().score()

			pls = nil
		} else {
			pls = append(pls, l)
		}
	}
	sum += parsePattern(pls).fixSmudge().score()
	return sum
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
