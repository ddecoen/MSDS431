package main

//Import the packages for the program
import (
	"fmt"

	"github.com/montanaflynn/stats"

	"time"
)

//Create a variable to input the data
var (
	//Anscombe's Quartet Dataset I
	x1 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y1 = []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
)

//use Stats package to Coordinate data from the arrays
func makeStatsData() []stats.Coordinate {
	data1 := make([]stats.Coordinate, len(x1))

	for i := 0; i < len(x1); i++ {
		data1[i] = stats.Coordinate{X: x1[i], Y: y1[i]}
	}

	return data1
}

//use Stats package to produce the slope and intercepts
func produceAllResults() error {
	data1 := makeStatsData()

	_, _, err := produce
}

func main() {
	//See how long the code takes to run
	startTime := time.Now()
	err := produceAllResults()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	endTime := time.Now()
	timeElapsed := endTime.Sub(startTime)
	fmt.Printf("Time running regressions: %v\n", timeElapsed)
}
