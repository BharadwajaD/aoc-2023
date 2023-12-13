package day

import (
	"fmt"
	"math"
	"strings"
)

type Race struct {
	max_time        int
	record_distance int
}

func racesFromString(time, distance string) []Race {
	max_times := intsFromString(time)
	record_distances := intsFromString(distance)

	var races []Race
	for i := 0; i < len(max_times); i++ {
		races = append(races, Race{max_time: max_times[i], record_distance: record_distances[i]})
	}

	return races
}

func digits_count(num int) int {
	cnt := 0

	for num > 0 {
		cnt++
		num /= 10
	}

	return cnt
}

func singleraceFromString(time, distance string) Race {
	max_times := intsFromString(time)
	record_distances := intsFromString(distance)

	var rdist, mtime, ndigits int

	for _, mt := range max_times {
		ndigits = digits_count(mt)
		mtime = mtime*int(math.Pow10(ndigits)) + mt
		fmt.Println(mtime)
	}

	ndigits = 0
	for _, rd := range record_distances {
        ndigits = digits_count(rd)
		rdist = rdist*int(math.Pow10(ndigits)) + rd
	}

	return Race{max_time: mtime, record_distance: rdist}
}

// no of ways to win the given race
func (r *Race) winning_ways_count() int {

	max_time := r.max_time
	start := 0
	end := 0

	for mt := 0; mt <= max_time/2; mt++ {
		if mt*(max_time-mt) > r.record_distance {
			start = mt
			break
		}
	}

	for mt := max_time; mt >= max_time/2; mt-- {
		if mt*(max_time-mt) > r.record_distance {
			end = mt
			break
		}
	}

	return end - start + 1
}

func Day6(input []string) int{
	times := strings.Split(input[0], ":")[1]
	distances := strings.Split(input[1], ":")[1]

	race := singleraceFromString(times, distances)
	fmt.Println(race)

	/*
	   //part 1
	   races := racesFromString(times, distances)
	   ans := 1
	   for _, race := range races {
	       ans *= race.winning_ways_count()
	   }
	*/

    return race.winning_ways_count()

}
