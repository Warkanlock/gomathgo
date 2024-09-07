package main

import (
	"fmt"
	"sort"
)

// technically a race it's just a sort from lowest to highest time
func race(horses []int) []int {
	sort.Slice(horses, func(i, j int) bool {
		return horses[i] > horses[j]
	})

	return horses
}

func assign_horses(horses []int, size uint) {
	for i := 1; i <= int(size); i++ {
		// to avoid starting it in zero
		horses[i-1] = i
	}
}

func main() {
	horses_size := uint(25)

	// slices are already by reference type
	horses := make([]int, horses_size)

	// assign to each horse an specific id
	assign_horses(horses, horses_size)

	// we need to:
	// 1. group into groups
	// 2. race the first of each group
	// 3. race the candidates
	// 4. get the first, second and third position

	fmt.Printf("All horses: %d\n", horses)
}
