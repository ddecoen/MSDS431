package main

//Import the packages for the program
import (
	"fmt"

	"github.com/montanaflynn/stats"
)

type Coordiante struct {
	X, Y float64
}

func main() {
	//Anscombe's Quarter Dataset I
	data := []stats.Coordinate{
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

	slope, intercept := stats.LinearRegression(data)

	fmt.Println("Regression results:")
	fmt.Printf("Intercept: %0.2f\n", intercept)
	fmt.Printf("Slope: %0.2f\n", slope)
	//r, _ := stats.LinearRegression(data)
	//fmt.Println(r)
}
