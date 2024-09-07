package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// technically a race it's just a sort from lowest to highest time
func race(horses []int, total_races *int) []int {
	sort.Slice(horses, func(i, j int) bool {
		return horses[i] < horses[j]
	})

	*total_races++

	return horses
}

func assign_horses(horses []int, size uint) {
	for i := 1; i <= int(size); i++ {
		// to avoid starting it in zero
		horses[i-1] = i
	}
}

func get_index(horse int, groups [][]int) int {
	for i, group := range groups {
		for _, h := range group {
			if h == horse {
				return i
			}
		}
	}

	return -1
}

func main() {
	horses_size := uint(25)
	total_races := 0

	// slices are already by reference type
	horses := make([]int, horses_size)

	// assign to each horse an specific id
	assign_horses(horses, horses_size)

	// shuffle array
	rand.Shuffle(len(horses), func(i, j int) { horses[i], horses[j] = horses[j], horses[i] })

	// make groups to divide horses in
	groups := [][]int{}
	for i := 0; i < int(horses_size); i += 5 {
		groups = append(groups, race(horses[i:i+5], &total_races))
	}

	// winners of each group
	winners := []int{groups[0][0], groups[1][0], groups[2][0], groups[3][0], groups[4][0]}
	finalist := race(winners, &total_races)

	// the first horse is the fastest of all
	fastest := finalist[0]

	// compete between the rest of horses
	// candidates will be: a2,a3,b1,b2,c1
	group_1 := groups[get_index(fastest, groups)]
	group_2 := groups[get_index(finalist[1], groups)]
	group_3 := groups[get_index(finalist[2], groups)]

	// create an slice of candidates
	candidates := race([]int{group_1[1], group_1[2], group_2[0], group_2[1], group_3[0]}, &total_races)

	second_fastest := candidates[0]
	third_fastest := candidates[1]

	fmt.Printf("All horses: %d\n", groups)
	fmt.Printf("Total races: %d\n", total_races)
	fmt.Printf("---\n")
	fmt.Printf(">\tFastest: %d\n", fastest)
	fmt.Printf(">\t2nd Fastest: %d\n", second_fastest)
	fmt.Printf(">\t3rd Fastest: %d\n", third_fastest)
}
