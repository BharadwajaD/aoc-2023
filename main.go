package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"theTechTrailBlazer/aoc/day"
)

func read_input(file string) []string {
	input, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	input_str := string(input)
	return strings.Split(input_str, "\n")
}


func main() {

	input := read_input("./input")
    fmt.Println(day.Day2(input))

}
