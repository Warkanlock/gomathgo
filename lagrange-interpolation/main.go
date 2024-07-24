package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Coordinate struct {
	x int
	y int
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func showCoordinatesTable(points []Coordinate) {
	fmt.Println("\n Coordinates Table:")
	// list all points provided as input
	fmt.Printf(" | x | y |\n")
	for _, point := range points {
		fmt.Printf(" | %d | %d |\n", point.x, point.y)
	}
}

func getPolynomial(points []Coordinate) {
	panic("Not implemented yet")
}

func waitForInterrupt(stop chan bool) {
	// make a channel to listen to inputs
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// wait for an interrupt signal
	<-sigs
	fmt.Println("\n Stopping...")

	// send a true to the channel as response
	stop <- true
}

func main() {
	var x, y int
	var points []Coordinate

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
			points = append(points, Coordinate{x, y})
		}
	}()

	// wait for this channel to be ready to start consuming from points object
	<-stop
	showCoordinatesTable(points)
	getPolynomial(points)
}
