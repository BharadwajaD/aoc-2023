package day

import (
	"strings"
)


const (
	LEFT = iota
	RIGHT
)

type DirNode struct {
	direction int
	node      string
}

func add_to_map(dir_node string, mp map[string][]DirNode) {

	splits := strings.Split(dir_node, "=")
	from_node := strings.TrimSpace(splits[0])
	splits[1] = strings.TrimSpace(splits[1])
	dir_node = splits[1][1 : len(splits[1])-1]
	splits = strings.Split(dir_node, ",")

	mp[from_node] = []DirNode{
		{direction: LEFT, node: strings.TrimSpace(splits[0])},
		{direction: RIGHT, node: strings.TrimSpace(splits[1])},
	}
}

func part1(instructions string, mp map[string][]DirNode) int {

	src := "AAA"
	dst := "ZZZ"
	count := 0

	for src != dst {
		for _, ins := range instructions {
			if src == dst {
				break
			}
			if ins == rune('L') {
				src = mp[src][0].node
			} else {
				src = mp[src][1].node
			}
			count++
		}
	}

	return count
}

func prime_factors(n int) []int {
	pfs := make([]int, 0, 8)
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// if n is a prime number itself
	if n > 2 {
		pfs = append(pfs, n)
	}

	return pfs
}

func part2(instructions string, mp map[string][]DirNode) int {
	var srcs []string

	for k := range mp {
		if k[2] == 'A' {
			srcs = append(srcs, k)
		}
	}

	steps := make([]int, len(srcs))

	for i, src := range srcs {
		st := src
		for st[2] != 'Z' {
			for _, ins := range instructions {
				if st[2] == 'Z' {
					break
				}
				if ins == rune('L') {
					st = mp[st][0].node
				} else {
					st = mp[st][1].node
				}

				steps[i]++
			}
		}
	}

	lcm := 1
	temp_mp := make(map[int]bool)
	for _, step := range steps {
		pfs := prime_factors(step)
		for _, pf := range pfs {
			if !temp_mp[pf] {
				lcm *= pf
				temp_mp[pf] = true
			}
		}
	}

	return lcm
}

func Day8(input []string) int{

	mp := make(map[string][]DirNode)

	instructions := input[0]
	for _, line := range input[2 : len(input)-1] {
		add_to_map(line, mp)
	}

    return part2(instructions, mp)
}
