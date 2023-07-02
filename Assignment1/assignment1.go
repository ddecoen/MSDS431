//Code inspired by https://github.com/weswest/msds431wk2/blob/main/main.go
//The purpose of this code is to reproduce the results of Anscombe's Quartet in Go
//Results were compared to Python and R results

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
	x1 = []float64{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0}
	y1 = []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	//Anscombe's Quartet Dataset II
	x2 = []float64{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0}
	y2 = []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	//Anscombe's Quartet Dataset III
	x3 = []float64{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0}
	y3 = []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
	//Ascombe's Quartet Dataset IV
	x4 = []float64{8.0, 8.0, 8.0, 8.0, 8.0, 8.0, 8.0, 19.0, 8.0, 8.0, 8.0}
	y4 = []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}
)

//use Stats package to Coordinate data from the arrays
func createStatsData() ([]stats.Coordinate, []stats.Coordinate, []stats.Coordinate, []stats.Coordinate) {
	data1 := make([]stats.Coordinate, len(x1))
	data2 := make([]stats.Coordinate, len(x2))
	data3 := make([]stats.Coordinate, len(x3))
	data4 := make([]stats.Coordinate, len(x4))

	for i := 0; i < len(x1); i++ {
		data1[i] = stats.Coordinate{X: x1[i], Y: y1[i]}
		data2[i] = stats.Coordinate{X: x2[i], Y: y2[i]}
		data3[i] = stats.Coordinate{X: x3[i], Y: y3[i]}
		data4[i] = stats.Coordinate{X: x4[i], Y: y4[i]}
	}

	return data1, data2, data3, data4
}

//calculate the slope and intercept in linear regression
func calcSlopeIntercept(reg stats.Series) (float64, float64, error) {
	if len(reg) == 0 {
		return 0, 0, fmt.Errorf("Series is empty")
	}

	var x1, y1, x2, y2 float64

	x1 = reg[0].X
	y1 = reg[0].Y

	//To deal with the error of n-1 data elemets use run go help build
	//build code to go over the regression data to find a different x value
	//without this code block the results are wrong - intercept is approx 0
	sameX := false
	for i := 1; i < len(reg); i++ {
		if reg[i].X != x1 {
			sameX = true
			x2 = reg[i].X
			y2 = reg[i].Y
			break
		}
	}

	if !sameX {
		return 0, 0, fmt.Errorf("No same X values found")
	}

	// Calculate the slope and intercept
	slope := (y2 - y1) / (x2 - x1)
	intercept := y1 - slope*x1

	return intercept, slope, nil

}

// createSlopeIntercept produces the slope and intercept for a regression number.
func createSlopeIntercept(regNum int, data []stats.Coordinate) (float64, float64, error) {
	reg, err := stats.LinearRegression(data)
	if err != nil {
		return 0, 0, err
	}

	intercept, slope, err := calcSlopeIntercept(reg)
	if err != nil {
		return 0, 0, err
	}

	fmt.Printf("Regression %d, Int: %f, Slope: %f\n\n", regNum, intercept, slope)
	return intercept, slope, nil

}

//use Stats package to produce the slope and intercepts
func createAllResults() error {
	data1, data2, data3, data4 := createStatsData()

	_, _, err := createSlopeIntercept(1, data1)
	if err != nil {
		return err
	}

	_, _, err = createSlopeIntercept(2, data2)
	if err != nil {
		return err
	}

	_, _, err = createSlopeIntercept(3, data3)
	if err != nil {
		return err
	}

	_, _, err = createSlopeIntercept(4, data4)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//See how long the code takes to run
	startTime := time.Now()
	err := createAllResults()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	endTime := time.Now()
	timeElapsed := endTime.Sub(startTime)
	fmt.Printf("Time running regressions: %v\n", timeElapsed)
}
