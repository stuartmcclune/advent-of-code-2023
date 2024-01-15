package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type almanac struct {
	seeds []int
	maps  map[string]almanacMap
}

type almanacr struct {
	seeds []seedRange
	maps  map[string]almanacMap
}

type seedRange struct {
	start  int
	length int
}

type almanacMap struct {
	sCat     string
	dCat     string
	mappings []mapping
}

type mapping struct {
	dRangeStart int
	sRangeStart int
	rLen        int
}

func main() {
	fmt.Println(closestLocation("almanac.txt"))
	fmt.Println(closestLocationRanges("almanac.txt"))
}

func closestLocation(fname string) int {
	almanac := parseAlmanac(fname)
	seedToLocation := make(map[int]int)
	for _, seed := range almanac.seeds {
		seedToLocation[seed] = mapSeedToLocation(seed, almanac)
	}

	minLoc := math.MaxInt
	for _, location := range seedToLocation {
		if location < minLoc {
			minLoc = location
		}
	}
	return minLoc
}

func closestLocationRanges(fname string) int {
	almanacr := parseAlmanacRanges(fname)
	seedRangeToMinLocation := make(map[seedRange]int)
	for _, seedRange := range almanacr.seeds {
		seedRangeToMinLocation[seedRange] = mapSeedRangeToMinLocation(seedRange, almanacr)
	}

	minLoc := math.MaxInt
	for _, location := range seedRangeToMinLocation {
		if location < minLoc {
			minLoc = location
		}
	}
	return minLoc
}

func (mapping *mapping) inRange(x int) bool {
	return x >= mapping.sRangeStart && x < mapping.sRangeStart+mapping.rLen
}

func (mapping *mapping) conv(x int) int {
	if !mapping.inRange(x) {
		return x
	}
	dx := x - mapping.sRangeStart
	res := mapping.dRangeStart + dx
	return res
}

func (am *almanacMap) conv(x int) int {
	for _, m := range am.mappings {
		if m.inRange(x) {
			return m.conv(x)
		}
	}
	return x
}

func mapSeedToLocation(seed int, a almanac) int {
	sts := a.maps["seed"]
	soil := sts.conv(seed)
	stf := a.maps["soil"]
	fertilizer := stf.conv(soil)
	ftw := a.maps["fertilizer"]
	water := ftw.conv(fertilizer)
	wtl := a.maps["water"]
	light := wtl.conv(water)
	ltt := a.maps["light"]
	temperature := ltt.conv(light)
	tth := a.maps["temperature"]
	humidity := tth.conv(temperature)
	htl := a.maps["humidity"]
	location := htl.conv(humidity)
	return location
}

func mapSeedToLocationR(seed int, a almanacr) int {
	sts := a.maps["seed"]
	soil := sts.conv(seed)
	stf := a.maps["soil"]
	fertilizer := stf.conv(soil)
	ftw := a.maps["fertilizer"]
	water := ftw.conv(fertilizer)
	wtl := a.maps["water"]
	light := wtl.conv(water)
	ltt := a.maps["light"]
	temperature := ltt.conv(light)
	tth := a.maps["temperature"]
	humidity := tth.conv(temperature)
	htl := a.maps["humidity"]
	location := htl.conv(humidity)
	return location
}

func mapSeedRangeToMinLocation(sr seedRange, a almanacr) int {
	minLoc := math.MaxInt
	for i := 0; i < sr.length; i++ {
		loc := mapSeedToLocationR(sr.start+i, a)
		if loc < minLoc {
			minLoc = loc
		}
	}
	return minLoc
}

func parseAlmanac(fname string) almanac {
	lines := slices.DeleteFunc(readLines(fname), isEmpty)
	seedsLine := lines[0]
	seeds := parseSeeds(seedsLine)
	maps := parseMaps(lines[1:])
	return almanac{seeds, maps}
}

func parseAlmanacRanges(fname string) almanacr {
	lines := slices.DeleteFunc(readLines(fname), isEmpty)
	seedsLine := lines[0]
	seeds := parseSeedsRanges(seedsLine)
	maps := parseMaps(lines[1:])
	return almanacr{seeds, maps}
}

func parseSeeds(line string) []int {
	seedStrs := slices.DeleteFunc(strings.Split(strings.TrimPrefix(line, "seeds: "), " "), isEmpty)
	seeds := make([]int, len(seedStrs))
	for i, str := range seedStrs {
		seed, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		seeds[i] = seed
	}
	return seeds
}

func parseSeedsRanges(line string) []seedRange {
	seedStrs := slices.DeleteFunc(strings.Split(strings.TrimPrefix(line, "seeds: "), " "), isEmpty)
	var seedRanges []seedRange
	for i := 0; i < len(seedStrs); i += 2 {
		start, err := strconv.Atoi(seedStrs[i])
		if err != nil {
			log.Fatal(err)
		}
		length, err := strconv.Atoi(seedStrs[i+1])
		if err != nil {
			log.Fatal(err)
		}
		seedRanges = append(seedRanges, seedRange{start, length})
		// for j := 0; j < length; j++ {
		// 	seeds = append(seeds, start+j)
		// }
	}
	return seedRanges
}

func parseMaps(lines []string) map[string]almanacMap {
	maps := make(map[string]almanacMap)

	aMap := almanacMap{}
	for _, line := range lines {
		if strings.Contains(line, "map") {
			if aMap.sCat != "" {
				maps[aMap.sCat] = aMap
			}
			aMap = almanacMap{}
			s, d := parseSourceAndDestination(line)
			aMap.sCat = s
			aMap.dCat = d
		} else {
			aMap.mappings = append(aMap.mappings, parseMapping(line))
		}
	}
	maps[aMap.sCat] = aMap

	return maps
}

func parseSourceAndDestination(line string) (string, string) {
	str := strings.TrimSuffix(strings.TrimSpace(line), " map:")
	sd := strings.Split(str, "-to-")
	return sd[0], sd[1]
}

func parseMapping(line string) mapping {
	strs := strings.Split(strings.TrimSpace(line), " ")
	dRangeStart, err := strconv.Atoi(strs[0])
	if err != nil {
		log.Fatal(err)
	}
	sRangeStart, err := strconv.Atoi(strs[1])
	if err != nil {
		log.Fatal(err)
	}
	rLen, err := strconv.Atoi(strs[2])
	if err != nil {
		log.Fatal(err)
	}
	return mapping{dRangeStart, sRangeStart, rLen}
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func readLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
