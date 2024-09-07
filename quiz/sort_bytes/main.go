package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	max_size_num := 255
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please provide a number as parameter")
		os.Exit(-1)
	}

	s_size := args[0]
	size, err := strconv.Atoi(s_size)

	// now cluster those into groups of 8 bytes
	step := 8
	size_groups := size / step

	// init groups variable
	var groups [][]int

	fmt.Printf("Groups: %d \n", size_groups)

	// check if size is a power of 2
	if size&(size-1) != 0 {
		fmt.Println("Please provide a number that is a power of 2")
		os.Exit(-1)
	}

	if err != nil {
		fmt.Println("Please provide a valid number")
		os.Exit(-1)
	}

	numbers := make([]int, size)

	// fulfill numbers with random numbers
	for i := range numbers {
		numbers[i] = rand.Intn(max_size_num)
	}

	for i := 0; i < len(numbers); i += step {
		groups = append(groups, numbers[i:i+step])
	}

	fmt.Printf("Bytes:%d \n", groups)
}
