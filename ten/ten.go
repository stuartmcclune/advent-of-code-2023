package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type pipe struct {
	label       string
	connections []connection
}

func (pipe pipe) String() string {
	return pipe.label
}

type connection struct {
	dx int
	dy int
}

var N = connection{0, -1}
var E = connection{1, 0}
var S = connection{0, 1}
var W = connection{-1, 0}

var NS = pipe{"|", []connection{N, S}}
var EW = pipe{"-", []connection{E, W}}
var NE = pipe{"L", []connection{N, E}}
var NW = pipe{"J", []connection{N, W}}
var SW = pipe{"7", []connection{S, W}}
var SE = pipe{"F", []connection{S, E}}
var Ground = pipe{".", []connection{}}
var Start = pipe{"S", []connection{N, S, E, W}}
var None pipe

type grid [][]pipe

type tinfo struct {
	outside bool
	p       pipe
	nx      int
	ny      int
	x       int
	y       int
	nCon    connection
}

func parsePipe(s string) pipe {
	switch s {
	case "|":
		return NS
	case "-":
		return EW
	case "L":
		return NE
	case "J":
		return NW
	case "7":
		return SW
	case "F":
		return SE
	case ".":
		return Ground
	case "S":
		return Start
	}
	return Ground
}

func parseGrid(lines []string) grid {
	g := make(grid, len(lines[0]))
	for i := range g {
		g[i] = make([]pipe, len(lines))
	}
	for y, l := range lines {
		for x, s := range l {
			g[x][y] = parsePipe(string(s))
		}
	}
	return g
}

func findStart(g grid) (int, int) {
	for x, col := range g {
		for y, p := range col {
			if p.label == Start.label {
				return x, y
			}
		}
	}
	return -1, -1
}

func (connection connection) opp() connection {
	switch connection {
	case N:
		return S
	case S:
		return N
	case E:
		return W
	}
	return E
}

func (pipe pipe) connects(con connection) bool {
	opp := con.opp()

	for _, c := range pipe.connections {
		if c == opp {
			return true
		}
	}
	return false
}

func main() {
	// log.SetOutput(io.Discard)
	fmt.Println(tenOne("input.txt"))
	fmt.Println(tenTwo("input.txt"))
}

func tenOne(fname string) int {
	g := parseGrid(readLines(fname))
	maxX := len(g) - 1
	maxY := len(g[0]) - 1
	sx, sy := findStart(g)
	at := g[sx][sy]
	var lastCon connection
	count := 0
	for {
		log.Printf("At %v, %v, %v", at, sx, sy)
		for _, con := range at.connections {
			log.Printf("Considering connection %v", con)
			if con.opp() == lastCon {
				// Don't go back the way we came
				log.Printf("Not going back: %v", con)
				continue
			}
			nx := sx + con.dx
			if nx < 0 || nx > maxX {
				log.Printf("%v would escape X bounds: %v", con, nx)
				continue
			}
			ny := sy + con.dy
			if ny < 0 || ny > maxY {
				log.Printf("%v would escape Y bounds: %v", con, ny)
				continue
			}
			n := g[nx][ny]
			log.Printf("I connect to %v at %v, %v", n, nx, ny)
			if !n.connects(con) {
				log.Printf("It doesn't connect to me")
				continue
			}
			log.Println("Moving to it")
			count++
			at = n
			sx = nx
			sy = ny
			lastCon = con
			// n connects & wasn't the previous
			if n.label == Start.label {
				return count / 2
			}
			break

		}
	}
}

func tenTwo(fname string) int {
	g := parseGrid(readLines(fname))
	maxX := len(g) - 1
	maxY := len(g[0]) - 1
	sx, sy := findStart(g)
	at := g[sx][sy]
	var lastCon connection

	info := make([][]tinfo, len(g))
	for i := 0; i <= maxX; i++ {
		info[i] = make([]tinfo, maxY+1)
	}

	for {
		log.Printf("At %v, %v, %v", at, sx, sy)
		for _, con := range at.connections {
			log.Printf("Considering connection %v", con)
			if con.opp() == lastCon {
				// Don't go back the way we came
				log.Printf("Not going back: %v", con)
				continue
			}
			nx := sx + con.dx
			if nx < 0 || nx > maxX {
				log.Printf("%v would escape X bounds: %v", con, nx)
				continue
			}
			ny := sy + con.dy
			if ny < 0 || ny > maxY {
				log.Printf("%v would escape Y bounds: %v", con, ny)
				continue
			}
			n := g[nx][ny]
			log.Printf("I connect to %v at %v, %v", n, nx, ny)
			if !n.connects(con) {
				log.Printf("It doesn't connect to me")
				continue
			}
			log.Println("Moving to it")

			info[sx][sy] = tinfo{false, at, nx, ny, sx, sy, con}

			at = n
			sx = nx
			sy = ny
			lastCon = con

			break

		}
		if at.label == Start.label {
			break
		}
	}

	ssx, ssy := findStart(g)
	categoriseStart(&info, ssx, ssy, lastCon)

	log.Println(info[ssx][ssy])

	var si tinfo

	for x := 0; x <= maxX; x++ {
		inf := info[x][sy]
		if inf.p.label == None.label {
			info[x][sy].outside = true
		} else {
			si = inf
			break
		}
	}
	outsideDir := Left

	ai := si
	log.Println(ai)

	for {
		// On corners, mark 2
		// On straights, mark 1
		//FIXME: newDir calc might break if start is a corner
		// Seems OK, since by chance start is L (NE) & ncon is N, but not in TEST
		// UPDATE: NOT OK - would need horizontal movement from corner for it be to OK. I can figure out what I should use as starting outsideDir using just the label though
		// Plus the newDir calc just seems inexplicably fucked anyway (see notes there)
		outsideDir = markO(&info, ai, outsideDir)

		// TODO: potentially could save myself the arse ache of the above fixme by defining that I WILL navigate clockwise this time
		ai = info[ai.nx][ai.ny]
		log.Printf("Moved to %v", ai)
		// time.Sleep(time.Second)

		if ai.x == si.x && ai.y == si.y {
			log.Println("Back to start!")
			break
		}
	}

	debugCount := 0
	for _, col := range info {
		for _, ai := range col {
			log.Print(ai.outside)
			if (!ai.outside) && (ai.p.label == None.label) {
				debugCount++
			}
		}
		log.Println()
	}
	log.Printf("dbg count %v", debugCount)

	for propogateO(&info) {
	}

	count := 0
	for _, col := range info {
		for _, ai := range col {
			log.Print(ai.outside)
			if (!ai.outside) && (ai.p.label == None.label) {
				count++
			}
		}
		log.Println()
	}

	log.Println(count)

	return count
}

func propogateO(info *[][]tinfo) bool {
	changed := false

	for x, col := range *info {
		for y, ai := range col {
			if ai.p.label != None.label {
				continue
			}
			if ai.outside {
				continue
			}

			for _, n := range neighbours(ai, info) {
				if n.outside {
					(*info)[x][y].outside = true
					changed = true
					break
				}
			}
		}
	}

	log.Printf("Changed: %v", changed)

	return changed
}

var allDirs = []dir{Left, Right, Up, Down}

func neighbours(ai tinfo, info *[][]tinfo) []tinfo {
	maxX := len(*info)
	maxY := len((*info)[0])

	var ns []tinfo

	for _, d := range allDirs {
		nx := ai.x + d.dx
		ny := ai.y + d.dy
		if nx < 0 || nx >= maxX {
			continue
		}
		if ny < 0 || ny >= maxY {
			continue
		}
		ns = append(ns, (*info)[nx][ny])
	}

	return ns
}

func categoriseStart(info *[][]tinfo, sx int, sy int, conIn connection) {
	conOut := (*info)[sx][sy].nCon
	cons := []connection{conIn.opp(), conOut}

	var p pipe

	if slices.Contains(cons, N) {
		if slices.Contains(cons, E) {
			p = NE
		} else if slices.Contains(cons, S) {
			p = NS
		} else {
			p = NW
		}
	} else if slices.Contains(cons, S) {
		if slices.Contains(cons, E) {
			p = SE
		} else {
			p = SW
		}
	} else {
		p = EW
	}

	(*info)[sx][sy].p = p
}

func markO(info *[][]tinfo, ai tinfo, odir dir) dir {

	//FIXME: why tf is F, Left giving DTM Left Down?
	log.Printf("MarkO at %v, odir: %v", ai, odir)

	p := ai.p
	newDir := odir

	dtm := []dir{odir}
	switch odir {
	case Left:
		switch p.label {
		case NE.label:
			dtm = append(dtm, Up)
			newDir = Up
			break
		case NW.label:
			dtm = append(dtm, Down)
			newDir = Down
			break
		case SE.label:
			dtm = append(dtm, Down)
			newDir = Down
			break
		case SW.label:
			dtm = append(dtm, Up)
			newDir = Up
			break
		}
		break
	case Right:
		switch p.label {
		case NE.label:
			dtm = append(dtm, Down)
			newDir = Down
			break
		case NW.label:
			dtm = append(dtm, Up)
			newDir = Up
			break
		case SE.label:
			dtm = append(dtm, Up)
			newDir = Up
			break
		case SW.label:
			dtm = append(dtm, Down)
			newDir = Down
			break
		}
	case Up:
		switch p.label {
		case NE.label:
			dtm = append(dtm, Left)
			newDir = Left
		case NW.label:
			dtm = append(dtm, Right)
			newDir = Right
		case SE.label:
			dtm = append(dtm, Right)
			newDir = Right
		case SW.label:
			dtm = append(dtm, Left)
			newDir = Left
		}
	case Down:
		switch p.label {
		case NE.label:
			dtm = append(dtm, Right)
			newDir = Right
		case NW.label:
			dtm = append(dtm, Left)
			newDir = Left
		case SE.label:
			dtm = append(dtm, Left)
			newDir = Left
		case SW.label:
			dtm = append(dtm, Right)
			newDir = Right
		}
	}

	log.Printf("DTM: %v", dtm)

	for _, d := range dtm {

		ox := ai.x + d.dx
		oy := ai.y + d.dy
		maxX := len(*info)
		maxY := len((*info)[0])
		if ox < 0 || ox >= maxX {
			continue
		}
		if oy < 0 || oy >= maxY {
			continue
		}
		oi := &(*info)[ox][oy]
		if oi.p.label != None.label {
			continue
		}
		log.Printf("Marking %v as outside", d)
		oi.outside = true
	}

	return newDir
}

type dir struct {
	dx int
	dy int
}

func (d dir) String() string {
	switch d {
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Up:
		return "Up"
	case Down:
		return "Down"
	}
	return "NA"
}

var Left = dir{-1, 0}
var Right = dir{1, 0}
var Up = dir{0, -1}
var Down = dir{0, 1}

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
