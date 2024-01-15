package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type galaxy struct {
	id int
	p  pos
}

type pos struct {
	x, y int
}

type universe struct {
	xmin, xmax, ymin, ymax int
	galaxies               []*galaxy
}

func (u *universe) expand(ef int) {
	for c := u.xmin; c <= u.xmax; c++ {
		containsGalaxy := false
		for _, g := range u.galaxies {
			if g.p.x == c {
				containsGalaxy = true
				break
			}
		}
		if containsGalaxy {
			continue
		}

		u.xmin -= ef
		for _, g := range u.galaxies {
			if g.p.x < c {
				g.p.x -= ef
			}
		}
	}

	for r := u.ymin; r <= u.ymax; r++ {
		containsGalaxy := false
		for _, g := range u.galaxies {
			if g.p.y == r {
				containsGalaxy = true
				break
			}
		}
		if containsGalaxy {
			continue
		}

		u.ymin -= ef
		for _, g := range u.galaxies {
			if g.p.y < r {
				g.p.y -= ef
			}
		}
	}

}

func (p *pos) dist(o pos) int {
	return (max(p.x, o.x) - min(p.x, o.x)) + (max(p.y, o.y) - min(p.y, o.y))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func parseUniverse(lines []string) *universe {
	ymax := len(lines) - 1
	xmax := len(lines[0]) - 1

	var galaxies []*galaxy
	id := 1

	for i, l := range lines {
		for j, r := range l {
			if string(r) == "#" {
				galaxies = append(galaxies, &galaxy{id, pos{j, i}})
				id++
			}
		}
	}

	return &universe{0, xmax, 0, ymax, galaxies}
}

func main() {
	fmt.Println(eleven("input.txt", 1))
	fmt.Println(eleven("input.txt", 999999))
}

func eleven(fname string, ef int) int {
	u := parseUniverse(readLines(fname))
	// log.Println(u)
	u.expand(ef)
	// log.Println(u)
	sum := 0
	for _, g := range u.galaxies {
		for _, og := range u.galaxies {
			if og.id > g.id {
				sum += g.p.dist(og.p)
			}
		}
	}
	return sum
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
