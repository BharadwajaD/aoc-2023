package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"theTechTrailBlazer/aoc/pkg/day"
	"time"
)

func read_input(file string) []string {
	input, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	input_str := string(input)
	return strings.Split(input_str, "\n")
}

func intsFromString(str string) []int {
	var nums []int
	splits := strings.Split(str, " ")
	for _, split := range splits {
		if len(split) > 0 {
			n, _ := strconv.Atoi(split)
			nums = append(nums, n)
		}
	}
	return nums
}

type Location struct {
	row, col int
}

func add_pipe(adjList map[Location][]Location, loc Location, c rune) {

	row := loc.row
	col := loc.col
	var locs []Location

	switch c {
	case '|':
		locs = []Location{{row: row - 1, col: col}, {row: row + 1, col: col}}
	case '-':
		locs = []Location{{row: row , col: col-1}, {row: row, col: col+1}}
	case 'L':
		locs = []Location{{row: row-1 , col: col}, {row: row, col: col+1}}
	case 'J':
		locs = []Location{{row: row-1 , col: col}, {row: row, col: col-1}}
	case '7':
		locs = []Location{{row: row+1 , col: col}, {row: row, col: col-1}}
	case 'F':
		locs = []Location{{row: row+1 , col: col}, {row: row, col: col+1}}
	case 'S':
		locs = []Location{{row: row-1 , col: col-1},{row: row+1 , col: col},{row: row+1 , col: col}, {row: row, col: col+1}}
	}

	adjList[loc] = append(adjList[loc], locs...)
}

func main() {

	tstart := time.Now()
	input := read_input("./input")
	tend := time.Now()

	fmt.Println(day.Day9(input), tend.Sub(tstart))

}
