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
	x2 = []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y2 = []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
)

//use Stats package to Coordinate data from the arrays
func StatsData() ([]stats.Coordinate, []stats.Coordinate) {
	data1 := make([]stats.Coordinate, len(x1))
	data2 := make([]stats.Coordinate, len(x2))

	for i := 0; i < len(x1); i++ {
		data1[i] = stats.Coordinate{X: x1[i], Y: y1[i]}
		data2[i] = stats.Coordinate{X: x2[i], Y: y2[i]}
	}

	return data1, data2
}

//calculate the slope and intercept in linear regression
func slopeIntercept(reg stats.Series) (float64, float64, error) {
	if len(reg) == 0 {
		return 0, 0, fmt.Errorf("Series is empty")
	}

	var x1, y1, x2, y2 float64

	x1 = reg[0].X
	y1 = reg[0].Y

	slope := (y2 - y1) / (x2 - x1)
	intercept := y1 - slope*x1

	return intercept, slope, nil

}

// produceSlopeIntercept produces the slope and intercept for a regression number.
func produceSlopeIntercept(regNum int, data []stats.Coordinate) (float64, float64, error) {
	reg, err := stats.LinearRegression(data)
	if err != nil {
		return 0, 0, err
	}

	intercept, slope, err := slopeIntercept(reg)
	if err != nil {
		return 0, 0, err
	}

	fmt.Printf("Regression %d, Int: %f, Slope: %f\n\n", regNum, intercept, slope)
	return intercept, slope, nil

}

//use Stats package to produce the slope and intercepts
func AllResults() error {
	data1, data2 := StatsData()

	_, _, err := produceSlopeIntercept(1, data1)
	if err != nil {
		return err
	}

	_, _, err = produceSlopeIntercept(2, data2)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	//See how long the code takes to run
	startTime := time.Now()
	err := AllResults()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	endTime := time.Now()
	timeElapsed := endTime.Sub(startTime)
	fmt.Printf("Time running regressions: %v\n", timeElapsed)
}
