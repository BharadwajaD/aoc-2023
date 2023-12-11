package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func read_input(file string) []string {
	input, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	input_str := string(input)
	return strings.Split(input_str, "\n")
}

type Loc struct {
	x int
	y int
}

func adjList(e1 *Loc) []Loc {

	return []Loc{
		{x: e1.x - 1, y: e1.y - 1},
		{x: e1.x - 1, y: e1.y},
		{x: e1.x - 1, y: e1.y + 1},
		{x: e1.x + 1, y: e1.y - 1},
		{x: e1.x + 1, y: e1.y},
		{x: e1.x + 1, y: e1.y + 1},
		{x: e1.x, y: e1.y - 1},
		{x: e1.x, y: e1.y + 1},
	}

}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isSpecialChar(c rune) bool {
	return !isDigit(c) && c != '.'
}

func main() {
	input := read_input("./input")

	var sum int
	var num int

	mp := make(map[Loc]bool)
	start := -1
	end := -1

	for ri, line := range input {
		for ci, c := range line {
			if isSpecialChar(c) {
				//insert locations of all special chars
				mp[Loc{x: ri, y: ci}] = true
			}
		}
	}

	for ri, line := range input {
		line += "." //if num is at end of line
		for ci, c := range line {
			if isDigit(c) {
				if start == -1 {
					start = ci
				}
				end = ci
				num = num*10 + int(c-'0')
			} else {
				if num > 0 {
					//check if [start,end] is adjacent to any special char
				out:
					for i := start; i <= end; i++ {
						adjs := adjList(&Loc{ri, i})
						for _, adj := range adjs {
							if mp[adj] {
								sum += num
								break out
							}
						}
					}
				}
				start = -1
				end = -1
				num = 0
			}
		}
	}

	fmt.Println(sum)
}
