package main

import (
	"fmt"
	"math"
)

func hash(x float64) int {
	if x < 0 {
		return int(-math.Floor(math.Abs(x)/10) * 10)
	}
	return int(math.Floor(x/10) * 10)
}

func main() {
	result := make(map[int][]float64)
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, x := range arr {
		result[hash(x)] = append(result[hash(x)], x)
	}

	for ind, x := range result {
		fmt.Println(ind, ":", x)
	}
}
