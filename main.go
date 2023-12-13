package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"theTechTrailBlazer/aoc/day"
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


func main() {

	tstart := time.Now()
	input := read_input("./input")
	tend := time.Now()

	fmt.Println(day.Day7(input), tend.Sub(tstart))
}
