package main

import (
	"math"
)

func average(salary []int) float64 {
	sum, max, min := 0, math.MinInt32, math.MaxInt32
	for _, v := range salary {
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return float64(sum-min-max) / float64(len(salary)-2)
}
