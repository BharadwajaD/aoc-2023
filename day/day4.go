package day

import (
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	id           int
	winning_nums []int
	nums         []int
}

func (c *Card) get_score() int {

	//O(nlogn)

	sort.Slice(c.winning_nums, func(i, j int) bool {
		return c.winning_nums[i] < c.winning_nums[j]
	})

	sort.Slice(c.nums, func(i, j int) bool {
		return c.nums[i] < c.nums[j]
	})

	//score := -1
	score := 0
	var ptw, ptn int

	for ptw < len(c.winning_nums) && ptn < len(c.nums) {
		if c.nums[ptn] == c.winning_nums[ptw] {
			/*
				if score == -1 {
					score = 1
				} else {
					score *= 2
				}
			*/

			score++
			ptn++

		} else if c.nums[ptn] > c.winning_nums[ptw] {
			ptw++
		} else {
			ptn++
		}
	}

	/*
		if score == -1 {
			return 0
		}
	*/

	return score
}

func IntFromString(str string) []int {

	var ints []int
	strs := strings.SplitN(str, " ", -1)

	for i := 0; i < len(strs); i++ {
		s := strs[i]
		if len(s) > 0 {
			n, _ := strconv.Atoi(s)
			ints = append(ints, n)
		}
	}

	return ints
}

func CardFromString(str string) Card {
	splits := strings.Split(str, ":")

    var id int
	for _, split := range strings.Split(splits[0], " ") {
		if len(split) > 0 {
			id, _ = strconv.Atoi(split)
		}
	}

	num_splits := strings.Split(splits[1], "|")

	return Card{
		id:           id,
		winning_nums: IntFromString(num_splits[0]),
		nums:         IntFromString(num_splits[1]),
	}
}


func Day4(input []string) int {
	var sum int
	scores_mp := make(map[int]int)

	for _, line := range input[:len(input)-1] {
		card := CardFromString(line)
		score := card.get_score()
		for cnt := 0; cnt <= scores_mp[card.id]; cnt++ {
			for i := 1; i <= score; i++ {
				scores_mp[card.id+i]++
			}
		}
		sum += scores_mp[card.id] + 1
		//for part1: sum += score
	}
    return sum
}
