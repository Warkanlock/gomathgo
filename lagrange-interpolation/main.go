package main

import (
	"fmt"
	"gomathgo/common"
	"os"
	"os/signal"
	"syscall"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func showCoordinatesTable(points []common.Point) {
	fmt.Println("\nCoordinates Table:")
	// list all points provided as input
	fmt.Printf("| x | y |\n")
	for _, point := range points {
		fmt.Printf("| %f | %f |\n", point.X, point.Y)
	}
}

func getPolynomial(points []common.Point) {
	var coefficients []common.Coefficient

	for _, point := range points {
		// create a list of coefficients based on the x values
		coefficients = append(coefficients, common.Coefficient(point.X))
	}

	// create and show the polynomial
	polynomial := common.Polynomial{}.New(coefficients)

	fmt.Println(polynomial)
}

func waitForInterrupt(stop chan bool) {
	// make a channel to listen to inputs
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// wait for an interrupt signal
	<-sigs
	fmt.Println("\nStopping...")

	// send a true to the channel as response
	stop <- true
}

func main() {
	var x, y int
	var points []common.Point

	fmt.Printf("Insert (x,y) coordinates from points \n")

	stop := make(chan bool, 1)

	// wait for the signal to stop the loop
	go waitForInterrupt(stop)

	go func() {
		for {
			fmt.Printf("coordinates > ")
			_, err := fmt.Scanf("%d,%d\n", &x, &y)

			if err != nil {
				clearScreen()
				fmt.Println("An error occur while parsing:", err)
				continue
			}

			// add x,y to points array
			points = append(points, common.Point{X: float64(x), Y: float64(y)})
		}
	}()

	// wait for this channel to be ready to start consuming from points object
	<-stop
	showCoordinatesTable(points)
	getPolynomial(points)
}
