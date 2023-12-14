package server

import "theTechTrailBlazer/aoc/pkg/day"

const (
    PART1 = iota + 1
    PART2
)

type AOC struct{
    day int
    part int
    answer int
    input string
}

func NewAOC(day, part int, input string) AOC {
    return AOC{
        day: day,
        part: part,
    }
}

func (aoc* AOC) GetAnswer(){

    code := string(rune(aoc.day)) + string(rune(aoc.part))
    var ans int

    switch code {
    case "11":
        ans = day.Day1(aoc.input)
    case "12":
        ans = day.Day1(aoc.input)
    case "21":
        ans = day.Day2(aoc.input)
    case "22":
        ans = day.Day2(aoc.input)
    case "31":
        ans = day.Day3(aoc.input)
    case "32":
        ans = day.Day3(aoc.input)
    case "41":
        ans = day.Day4(aoc.input)
    case "42":
        ans = day.Day4(aoc.input)
    case "51":
        ans = day.Day5(aoc.input)
    case "52":
        ans = day.Day5(aoc.input)
    case "61":
        ans = day.Day6(aoc.input)
    case "62":
        ans = day.Day6(aoc.input)
    case "71":
        ans = day.Day7(aoc.input)
    case "72":
        ans = day.Day7(aoc.input)
    }

    aoc.answer = ans
}
