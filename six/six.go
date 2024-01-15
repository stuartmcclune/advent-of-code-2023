package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sixOne("input.txt"))
	fmt.Println(sixTwo("input.txt"))
}

func sixOne(fname string) int {
	lines := readLines(fname)
	times := numbers(lines[0])
	distances := numbers(lines[1])

	res := 1
	for i := 0; i < len(times); i++ {
		res *= waysToWin(times[i], distances[i])
	}

	return res
}

func sixTwo(fname string) int {
	lines := readLines(fname)
	time := number(lines[0])
	distance := number(lines[1])
	return waysToWin(time, distance)
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

func numbers(l string) []int {
	strs := strings.Fields(strings.Split(l, ":")[1])
	nums := make([]int, len(strs))
	for i, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = n
	}
	return nums
}

func number(l string) int {
	str := strings.Split(strings.ReplaceAll(l, " ", ""), ":")[1]
	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func waysToWin(time int, distance int) int {
	count := 0
	for hold := 1; hold < time; hold++ {
		d := (time - hold) * hold
		if d > distance {
			count++
		}
	}
	return count
}
