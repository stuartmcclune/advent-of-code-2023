package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type rock string

const (
	Empty rock = "."
	Round rock = "O"
	Cube  rock = "#"
)

type platform struct {
	rocks         [][]rock
	width, height int
}

func parsePlatform(lines []string) platform {
	width := len(lines[0])
	height := len(lines)

	rocks := make([][]rock, width)
	for i := 0; i < width; i++ {
		rocks[i] = make([]rock, height)
	}

	for y, l := range lines {
		for x, r := range l {
			rocks[x][y] = rock(r)
		}
	}
	return platform{rocks, width, height}
}

func (p *platform) tiltN() *platform {

	for x := 0; x < p.width; x++ {
		for y := 0; y < p.height; y++ {
			rock := p.rocks[x][y]
			if rock != Round {
				continue
			}
			for dy := -1; y+dy >= 0; dy-- {
				if p.rocks[x][y+dy] == Empty {
					p.rocks[x][y+dy] = Round
					p.rocks[x][y+dy+1] = Empty
				} else {
					break
				}
			}
		}
	}

	return p
}

func (p *platform) tiltW() *platform {

	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			rock := p.rocks[x][y]
			if rock != Round {
				continue
			}
			for dx := -1; x+dx >= 0; dx-- {
				if p.rocks[x+dx][y] == Empty {
					p.rocks[x+dx][y] = Round
					p.rocks[x+dx+1][y] = Empty
				} else {
					break
				}
			}
		}
	}

	return p
}

func (p *platform) tiltS() *platform {

	for x := 0; x < p.width; x++ {
		for y := p.height - 1; y >= 0; y-- {
			rock := p.rocks[x][y]
			if rock != Round {
				continue
			}
			for dy := 1; y+dy < p.height; dy++ {
				if p.rocks[x][y+dy] == Empty {
					p.rocks[x][y+dy] = Round
					p.rocks[x][y+dy-1] = Empty
				} else {
					break
				}
			}
		}
	}

	return p
}

func (p *platform) tiltE() *platform {

	for y := 0; y < p.height; y++ {
		for x := p.width - 1; x >= 0; x-- {
			rock := p.rocks[x][y]
			if rock != Round {
				continue
			}
			for dx := 1; x+dx < p.width; dx++ {
				if p.rocks[x+dx][y] == Empty {
					p.rocks[x+dx][y] = Round
					p.rocks[x+dx-1][y] = Empty
				} else {
					break
				}
			}
		}
	}

	return p
}

func (p *platform) hash() string {
	s := ""
	for i := 0; i < p.width; i++ {
		for j := 0; j < p.height; j++ {
			s += string(p.rocks[i][j])
		}
	}
	return s
}

func (p *platform) spinAndGetLoadN(cycles int) int {
	seen := make(map[string]int)
	loads := make(map[int]int)
	cycleStart := 0
	cycleLen := 0
	for i := 1; i <= cycles; i++ {
		hash := p.hash()
		loads[i] = p.loadN()
		if seen[hash] != 0 {
			cycleStart = seen[hash]
			cycleLen = i - cycleStart
			log.Printf("Seen %v before, at iteration %v!", i, seen[hash])
			break
		}
		seen[hash] = i
		p.tiltN().tiltW().tiltS().tiltE()
	}

	// after n cycles, I'm at iteration n * cycleLen + cycleStart
	// c = n*cl+cs + idx
	// idx = c - cs - n*cl

	n := (cycles - (cycleStart - 1)) / cycleLen
	idx := cycles - (cycleStart - 1) - n*cycleLen
	log.Println(cycleStart + idx)

	return loads[cycleStart+idx]
}

func (p *platform) loadN() int {
	load := 0
	for x := 0; x < p.width; x++ {
		for y := 0; y < p.height; y++ {
			if p.rocks[x][y] != Round {
				continue
			}
			load += p.height - y
		}
	}
	return load
}

func main() {
	fmt.Println(fourteenOne("input.txt"))
	fmt.Println(fourteenTwo("input.txt"))
}

func fourteenOne(fname string) int {
	platform := parsePlatform(readLines(fname))
	return platform.tiltN().loadN()
}

func fourteenTwo(fname string) int {
	platform := parsePlatform(readLines(fname))
	return platform.spinAndGetLoadN(1_000_000_000)
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
