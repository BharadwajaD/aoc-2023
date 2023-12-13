package day

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

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

func mapped_to(num int, trans []string) int {

	for _, t := range trans {

		if len(t) == 0 {
			continue
		}

		dst_src_rng := intsFromString(t)
		if num >= dst_src_rng[1] && num < dst_src_rng[1]+dst_src_rng[2] {
			return dst_src_rng[0] + num - dst_src_rng[1]
		}
	}

	return num
}

func reverse_mapped_to(num int, trans []string) int {

	for _, t := range trans {

		if len(t) == 0 {
			continue
		}

		src_dst_rng := intsFromString(t)
		src := num + src_dst_rng[1] - src_dst_rng[0]
		if num >= src_dst_rng[1] && num < src_dst_rng[1]+src_dst_rng[2] {
			return src
		}
	}

	return num
}

func Day5(){
	input, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	input_str := string(input)
	trans_blocks := strings.Split(input_str, "\n\n")

	seeds := intsFromString(strings.Split(trans_blocks[0], ":")[1])

	converts := make([][]string, 0, 8)
	for _, block := range trans_blocks[1:] {
		converts = append(converts, strings.Split(block, "\n")[1:])
	}

	/* part 1
		ans := math.MaxInt
	    for _, seed := range seeds {
	        search_for := seed
	        for _, converter := range converts {
	            search_for = mapped_to(search_for, converter)
	        }

	        ans = min(ans, search_for)
	    }
	*/

	var s int
	var l int

	slices.Reverse(converts)

    //optimisation:
    //1. sorting seeds
    //2. using go routines

    go func(){
    }()

	for {
		s = l
		for _, t := range converts {
			s = reverse_mapped_to(s, t)
		}

		// check if s is valid seed number
        fmt.Println(l)
		for i := 0; i < len(seeds); i += 2 {
			if s > seeds[i] && s < seeds[i]+seeds[i+1] {
				fmt.Println(l)
				return
			}
		}

		l++
	}
	/*
		    for loc := 0; loc < math.MaxInt; loc ++ {
		        search_for := loc
		        for i := 6; i >= 0; i-- {
		            search_for = reverse_mapped_to(search_for, converts[i])
		        }

		        //if this loc has a valid seed then return it
		        for i := 0; i < len(seeds); i += 2 {
		            if search_for >= seeds[i] && search_for < seeds[i] + seeds[i+1] {
			            fmt.Println(loc)
		                return
		            }
		        }
		    }
	*/
}
