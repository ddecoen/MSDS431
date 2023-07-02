package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

func createDatasets(data [][]float64) []stats.Coordinate {
	datasets := make([]stats.Coordinate, len(data))
	for i, d := range data {
		datasets[i] = stats.Coordinate{d[0], d[1]}
	}
	return datasets
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
	data := [][]float64{
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

	datasets := createDatasets(data)
	produceSlopeIntercept(datasets)

	//See how long the code takes to run
	startTime := time.Now()
	endTime := time.Now()
	timeElapsed := endTime.Sub(startTime)
	fmt.Printf("Time running regressions: %v\n", timeElapsed)

}
