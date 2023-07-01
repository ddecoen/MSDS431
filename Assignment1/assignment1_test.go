//This is the testing file for assignment1 go

package main

import (
	"math"

	"testing"

	"github.com/montanaflynn/stats"
)

//Create comparison for running the various regressions
func CreateComparisonAllRegressions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		createAllResults()
	}
}

//Test the regression results against the expected results for each dataset
func CompareRegressionResults(t *testing.T) {
	//List the expected values for slope and intercepts
	expectedIntercept := 3.0
	expectedSlope := 0.5
	data1, data2, data3, data4 := createStatsData()

	//Test regression on Anscombe Quartet I
	compareRegression(t, 1, data1, expectedIntercept, expectedSlope)

	//Test regression on Anscombe Quartet II
	compareRegression(t, 2, data2, expectedIntercept, expectedSlope)

	//Test regression on Anscombe Quartet III
	compareRegression(t, 3, data3, expectedIntercept, expectedSlope)

	//Test regression on Anscombe Quartet IV
	compareRegression(t, 4, data4, expectedIntercept, expectedSlope)
}

//compareRegression calculates the regression values and compares to expected
func compareRegression(t *testing.T, regNum int, data []stats.Coordinate, expectedIntercept, expectedSlope float64) {
	//Calculate regression
	intercept, slope, err := createSlopeIntercept(regNum, data)
	if err != nil {
		t.Errorf("Regression %d failed: %v", regNum, err)
	}

	//Round the values from regression for comparison purposes
	roundIntercept := math.Round(intercept*100) / 100
	roundSlope := math.Round(slope*100) / 100

	//Compare regression actual to expected
	if roundIntercept != expectedIntercept {
		t.Errorf("The value is not correct for Intercept in regression %d. Expected: %f, Actual: %f", regNum, expectedIntercept, roundIntercept)
	}

	if roundSlope != expectedSlope {
		t.Errorf("The value is not correct for Slope in regression %d. Expected: %f, Actual: %f", regNum, expectedSlope, roundSlope)
	}

}
