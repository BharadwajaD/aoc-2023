package day

import "strings"

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

func Day3(inp string) int {

    input := strings.Split(inp, "\n")

	var sum int
	var num int

	mp := make(map[Loc][]int)
	special_chars := make(map[Loc]rune)
	start := -1
	end := -1

	for ri, line := range input {
		for ci, c := range line {
			if isSpecialChar(c) {
				//insert locations of all special chars
                mp[Loc{x: ri, y: ci}] = make([]int, 0)
                special_chars[Loc{x: ri, y: ci}] = c
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
							if mp[adj] != nil && special_chars[adj] == '*'{
                                v := mp[adj]
                                v = append(v, num)
                                mp[adj] = v

                                //for part1: sum += num
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
    //checking for gears
    for k,v := range mp {
        if special_chars[k] == '*'  && len(v) > 1{
            sum += v[0]*v[1]
        }
    }

    return sum
}
