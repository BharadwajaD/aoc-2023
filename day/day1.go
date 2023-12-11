package day

import (
	"strings"
)

func get_indexes(str, substr string) []int {

	indices := []int{}

	for i := 0; i < len(str); i++ {
		index := strings.Index(str[i:], substr)
		if index == -1 {
			break
		}
		indices = append(indices, i+index)
		i += index + len(substr) - 1
	}

	return indices
}

var nums map[string]int

func Day1(input []string) int {

	nums = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sum := 0

	for _, line := range input[:len(input)-1] {
		var digits []int

		for i, c := range line {
			if c >= '0' && c <= '9' {
				digits = append(digits, int(c)-'0')
			}

			for k, v := range nums {
				if strings.HasPrefix(line[i:], k) {
					digits = append(digits, v)
				}
			}

		}
		sum += 10*(digits[0]) + digits[len(digits)-1]

	}

	return sum
}
