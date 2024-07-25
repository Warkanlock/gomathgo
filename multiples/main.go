package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/**
  If we list all the natural numbers below that are multiples of N, we get and M. The sum of these multiples is Z.

  The following is an algorithm to provide F(N) = Z_1 + F(M) = Z_2
*/

const MIN_ARG int = 2

func calculateMultiples(number int, size int) float32 {
	a := number // first term
	d := number // common difference
	l := int(size/number) * number

	// if we exceed from size, we replace for (l - number)
	if l >= size {
		l = l - number
	}

	// total terms / number of multiples
	n := l / number

	// f(n) = (2a + (n - 1)d)
	start := 2 * a
	end := (n - 1) * d

	// S_n = (n/2) * (a + l)
	return float32(float32(float32(n)/2) * (float32(start) + float32(end)))
}

func parseNumber(input string) int {
	numberCast, err := strconv.ParseFloat(input, 32)

	if err != nil {
		fmt.Println("There was an error parsing this number:", err)
		os.Exit(-1)
	}

	return int(float32(numberCast))
}

func main() {
	if len(os.Args) < MIN_ARG {
		fmt.Println("usage: <TOTAL_N_FOR_SUM> <N_1> <N_2> ... <N_n>")
		os.Exit(-1)
	}

	var totalNumbers int = int(parseNumber(os.Args[1]))
	var numbers []string = os.Args[2:len(os.Args)]

	fmt.Println(numbers)

	var totalMultiples float32
	for _, number := range numbers {
		number_parsed := parseNumber(number)

		// parse and compute the number all multiples for this
		sumOfMultiples := calculateMultiples(int(math.Abs(float64(number_parsed))), totalNumbers)

		// get sign of number and substract or add depending on the case
		isNegative := math.Signbit(float64(number_parsed))

		if isNegative {
			sumOfMultiples = sumOfMultiples * (-1)
		}

		fmt.Printf("Multiples of: %s [=%f]\n", number, sumOfMultiples)
		// we add the result to the previous sum calculated
		totalMultiples += sumOfMultiples

	}

	fmt.Println("Total multiples sum is:", totalMultiples)
}
