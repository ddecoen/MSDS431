package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

type DataSet [][]float64

func createDataset(dataset DataSet) []stats.Coordinate {
	coords := make([]stats.Coordinate, len(dataset))
	for i, d := range dataset {
		coords[i] = stats.Coordinate{d[0], d[1]}
	}
	return coords
}

func calculateSlopeIntercept(reg stats.Series) (float64, float64, error) {
	if len(reg) == 0 {
		return 0, 0, fmt.Errorf("Series is empty")
	}

	var x1, y1, x2, y2 float64

	x1 = reg[0].X
	y1 = reg[0].Y

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
		return 0, 0, fmt.Errorf("No same value X found")
	}

	slope := (y2 - y1) / (x2 - x1)
	intercept := y1 - slope*x1

	return intercept, slope, nil
}

func produceSlopeIntercept(data []stats.Coordinate) (float64, float64, error) {
	reg, err := stats.LinearRegression(data)
	if err != nil {
		return 0, 0, err
	}

	intercept, slope, err := calculateSlopeIntercept(reg)
	if err != nil {
		return 0, 0, err
	}
	fmt.Printf("Regression Anscombe Quartet details - Intercept: %f, Slope: %f\n\n", intercept, slope)
	return intercept, slope, nil
}

func main() {
	dataset0 := DataSet{
		{10.0, 8.04},
		{8.0, 6.95},
		{13.0, 7.58},
		{9.0, 8.81},
		{11.0, 8.33},
		{14.0, 9.96},
		{6.0, 7.24},
		{4.0, 4.26},
		{12.0, 10.84},
		{7.0, 4.82},
		{5.0, 5.68},
	}
	dataset1 := DataSet{
		{10.0, 9.14},
		{8.0, 8.14},
		{13.0, 8.74},
		{9.0, 8.77},
		{11.0, 9.26},
		{14.0, 8.10},
		{6.0, 6.13},
		{4.0, 3.10},
		{12.0, 9.13},
		{7.0, 7.26},
		{5.0, 4.74},
	}
	dataset2 := DataSet{
		{10.0, 7.46},
		{8.0, 6.77},
		{13.0, 12.74},
		{9.0, 7.11},
		{11.0, 7.81},
		{14.0, 8.84},
		{6.0, 6.08},
		{4.0, 5.39},
		{12.0, 8.15},
		{7.0, 6.42},
		{5.0, 5.73},
	}
	dataset3 := DataSet{
		{8.0, 6.58},
		{8.0, 5.76},
		{8.0, 7.71},
		{8.0, 8.84},
		{8.0, 8.47},
		{8.0, 7.04},
		{8.0, 5.25},
		{19.0, 12.50},
		{8.0, 5.56},
		{8.0, 7.91},
		{8.0, 6.89},
	}

	d := []DataSet{dataset0, dataset1, dataset2, dataset3}

	datasets := make([][]stats.Coordinate, len(d))
	for i, dataset := range d {
		datasets[i] = createDataset(dataset)
	}
	startTime := time.Now()
	for _, coords := range datasets {
		produceSlopeIntercept(coords)
	}
	endTime := time.Now()
	timeElapsed := endTime.Sub(startTime)
	fmt.Printf("Time running regressions: %v\n", timeElapsed)

	//See how long the code takes to run
	//startTime := time.Now()
	//endTime := time.Now()
	//timeElapsed := endTime.Sub(startTime)
	//fmt.Printf("Time running regressions: %v\n", timeElapsed)

}
