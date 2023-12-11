package day

import (
	"strconv"
	"strings"
)

type Round struct {
	rcount int
	bcount int
	gcount int
}

func RoundFromString(str string) *Round {

	round := new(Round)
	splits := strings.Split(str, ",")

	for _, color := range splits {

        color = strings.TrimSpace(color)
		count_color := strings.Split(color, " ")

		switch count_color[1] {
		case "blue":
			round.bcount, _ = strconv.Atoi(count_color[0])
		case "red":
			round.rcount, _ = strconv.Atoi(count_color[0])
		case "green":
			round.gcount, _ = strconv.Atoi(count_color[0])
		}
	}
	return round
}

type Game struct {
	game_id int
	rounds  []*Round
}

func GameFromString(str string) *Game {

	splits := strings.Split(str, ":")
	gid, _ := strconv.Atoi(strings.Split(splits[0], " ")[1])

	game := new(Game)
	game.game_id = gid

	for _, set := range strings.Split(splits[1], ";") {
		game.rounds = append(game.rounds, RoundFromString(set))
	}

	return game
}

func Day2(input []string) int {

	var sum int
	for _, line := range input[:len(input)-1] {

		game := GameFromString(line)
		var max_r, max_b, max_g int

		for _, rd := range game.rounds {
			max_r = max(max_r, rd.rcount)
			max_g = max(max_g, rd.gcount)
			max_b = max(max_b, rd.bcount)
		}

		sum += max_r * max_b * max_g
	}

	return sum
}
