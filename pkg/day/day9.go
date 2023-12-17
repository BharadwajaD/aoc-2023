package day

func get_extra(diffs []int, start, end int) int{

    all_same := true
    for i := start; i <= end; i++ {
        if diffs[i] != diffs[start] {
            all_same = false;
            break;
        }
    }

    if all_same {
        return diffs[start]
    }

    /* part1
    end_val := diffs[end]
    for i := end ; i >= start+1; i-- {
        diffs[i] = diffs[i] - diffs[i-1]
    }
    */

    start_val := diffs[start]
    for i := end ; i >= start+1; i-- {
        diffs[i] = diffs[i] - diffs[i-1]
    }

    //return end_val + get_extra(diffs, start+1, end)
    return start_val - get_extra(diffs, start+1, end)
}

func Day9(input []string) int {

    sum := 0

    for _, line := range input[:len(input)-1] {
        nums := intsFromString(line)
        sum += get_extra(nums, 0, len(nums)-1)
    }

    return sum
}
