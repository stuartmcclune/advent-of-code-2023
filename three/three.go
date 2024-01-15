package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type schematic struct {
	raw    [][]string
	width  int
	height int
}

type index struct {
	x int
	y int
}

type digit struct {
	digit string
	index index
}

type gear struct {
	pn1 int
	pn2 int
}

func (g *gear) ratio() int {
	return g.pn1 * g.pn2
}

func (schematic *schematic) at(index index) string {
	return schematic.raw[index.x][index.y]
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println(sumPartNumbers("input1.txt"))
	fmt.Println(sumGearRatios("input1.txt"))
}

func sumPartNumbers(fname string) int {
	schematic := createSchematic(readLines(fname))
	symbolIndices := findSymbols(schematic)
	var partNumberDigits []digit
	for _, idx := range symbolIndices {
		partNumberDigits = append(partNumberDigits, findAdjacentDigits(idx, schematic)...)
	}
	partNumbers := constructPartNumbers(partNumberDigits, schematic)

	return sum(partNumbers)
}

func sumGearRatios(fname string) int {
	schematic := createSchematic(readLines(fname))
	gears := findGears(schematic)
	gearRatios := make([]int, len(gears))
	for i, g := range gears {
		gearRatios[i] = g.ratio()
	}

	return sum(gearRatios)
}

func findGears(s schematic) []gear {
	stars := findStars(s)

	var gears []gear

	for _, star := range stars {
		partNumberDigits := findAdjacentDigits(star, s)
		partNumbers := constructPartNumbers(partNumberDigits, s)
		if len(partNumbers) == 2 {
			gears = append(gears, gear{partNumbers[0], partNumbers[1]})
		}
	}

	return gears
}

func findStars(schematic schematic) []index {
	var indices []index
	for x := 0; x < schematic.width; x++ {
		for y := 0; y < schematic.height; y++ {
			index := index{x, y}
			s := schematic.at(index)
			if isStar(s) {
				indices = append(indices, index)
			}
		}
	}
	return indices
}

func findSymbols(schematic schematic) []index {
	var indices []index
	for x := 0; x < schematic.width; x++ {
		for y := 0; y < schematic.height; y++ {
			index := index{x, y}
			s := schematic.at(index)
			if isSymbol(s) {
				indices = append(indices, index)
			}
		}
	}
	return indices
}

func contains(s []index, e index) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func constructPartNumbers(digits []digit, schematic schematic) []int {
	var seen []index
	var partNumbers []int
	for _, digit := range digits {
		if contains(seen, digit.index) {
			continue
		}
		partNumber, sn := constructPartNumber(digit, schematic)
		seen = append(seen, sn...)
		partNumbers = append(partNumbers, partNumber)
	}
	return partNumbers
}

func (d *digit) left(s schematic) *digit {
	return d.dx(-1, s)
}

func (d *digit) right(s schematic) *digit {
	return d.dx(1, s)
}

func (d *digit) dx(dx int, s schematic) *digit {
	x := d.index.x + dx
	if x < 0 || x >= s.width {
		return nil
	}
	idx := index{x, d.index.y}
	item := s.at(idx)
	if !isDigit(item) {
		return nil
	}
	return &digit{item, idx}
}

func constructPartNumber(d digit, s schematic) (int, []index) {
	var seen []index
	var pnStrs []string

	// log.Println(d)

	var leftmost digit
	for x := &d; x != nil; x = x.left(s) {
		leftmost = *x
	}
	// log.Println(leftmost)

	for x := &leftmost; x != nil; x = x.right(s) {
		seen = append(seen, x.index)
		pnStrs = append(pnStrs, x.digit)
	}

	pnStr := strings.Join(pnStrs, "")
	pn, err := strconv.Atoi(pnStr)
	if err != nil {
		log.Fatal(err)
	}
	return pn, seen
}

func findAdjacentDigits(index index, schematic schematic) []digit {
	var digits []digit
	for _, adjacent := range findAdjacentIndices(index, schematic.width, schematic.height) {
		item := schematic.at(adjacent)
		if isDigit(item) {
			digits = append(digits, digit{item, adjacent})
		}
	}
	return digits
}

func findAdjacentIndices(idx index, width int, height int) []index {
	x := idx.x
	y := idx.y

	var adjacent []index

	for i := max(0, x-1); i <= min(width-1, x+1); i++ {
		for j := max(0, y-1); j <= min(height-1, y+1); j++ {
			adjacent = append(adjacent, index{i, j})
		}
	}
	return adjacent
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isSymbol(s string) bool {
	return !(isDigit(s) || s == ".")
}

func isStar(s string) bool {
	return s == "*"
}

func readLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func createSchematic(lines []string) schematic {
	height := len(lines)
	width := len(lines[0])

	raw := make([][]string, width)
	for i := range raw {
		raw[i] = make([]string, height)
	}

	for y, line := range lines {
		split := strings.Split(line, "")
		for x, s := range split {
			raw[x][y] = s
		}
	}

	return schematic{raw: raw, width: width, height: height}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func sum(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}
