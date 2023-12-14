package day

import (
	"sort"
	"strconv"
	"strings"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KLIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type Bid struct {
	hand  []rune
	bid   int
	ctype int
}

func bidFromString(str string) Bid {
	splits := strings.Split(str, " ")
	amt, _ := strconv.Atoi(strings.TrimSpace(splits[1]))
	hand := []rune(strings.TrimSpace(splits[0]))

	cnt_mp := make(map[rune]int)
	for i, rn := range hand {
		switch rn {
		case rune('A'):
			hand[i] = rune('9' + 5)
		case rune('K'):
			hand[i] = rune('9' + 4)
		case rune('Q'):
			hand[i] = rune('9' + 3)
		/* part 1
		        case rune('J'):
					hand[i] = rune('9'+2)
		*/
		case rune('T'):
			hand[i] = rune('9' + 1)
		case rune('J'): //for part 2
			hand[i] = rune('2' - 1)
		}
		cnt_mp[hand[i]]++
	}

	cnts := make([]int, 0, 5)
	var ctype int

	var njokers int
	for k, v := range cnt_mp {
		if k == rune('2'-1) {
			njokers = v
			continue
		}
		cnts = append(cnts, v)
	}

	sort.Slice(cnts, func(i, j int) bool {
		return cnts[i] > cnts[j]
	})

    //for part 2
    //TODO: Understand this
    if len(cnts) == 0 {
        cnts = append(cnts, 5)
    } else {
        diff := max(0, 3-cnts[0])
        cnts[0] += njokers
        njokers -= diff
        if len(cnts) > 1 && njokers > 0 {
            cnts[1] += njokers

        }
    }

	ctype = HIGH_CARD
	if cnts[0] >= 5 {
		ctype = FIVE_OF_A_KIND
	} else if cnts[0] >= 4 {
		ctype = FOUR_OF_A_KIND
	} else if cnts[0] >= 3 && cnts[1] >= 2 {
		ctype = FULL_HOUSE
	} else if cnts[0] >= 3 {
		ctype = THREE_OF_A_KLIND
	} else if cnts[0] >= 2 && cnts[1] >= 2 {
		ctype = TWO_PAIR
	} else if cnts[0] >= 2 {
		ctype = ONE_PAIR
	}

	return Bid{
		hand:  hand,
		bid:   amt,
		ctype: ctype,
	}

}

func compare_bid(hd1, hd2 Bid) bool {

	if hd1.ctype == hd2.ctype {
		for i := 0; i < len(hd1.hand); i++ {
			if hd1.hand[i] == hd2.hand[i] {
				continue
			}

			return hd1.hand[i] < hd2.hand[i]
		}
	}

	return hd1.ctype < hd2.ctype
}

func Day7(inp string) int {

    input := strings.Split(inp, "\n")

	var bids []Bid
	for _, line := range input[:len(input)-1] {
		bids = append(bids, bidFromString(line))
	}

	//ascending order
	sort.Slice(bids, func(i, j int) bool {
		return compare_bid(bids[i], bids[j])
	})

	var ans int
	for i, bid := range bids {
		ans += (i + 1) * bid.bid
	}

    return ans
}
