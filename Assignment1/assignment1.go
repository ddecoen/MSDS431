package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {
	data := []stats.Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
	}

	r, _ := stats.LinearRegression(data)
	fmt.Println(r)
}
