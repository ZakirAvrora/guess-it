package main

import (
	"math"
)

func Sort(arr []float64) {
	var minIndex int
	for i := 0; i < len(arr); i++ {
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func CalculateMed(arr []float64) float64 {
	Sort(arr)
	if len(arr)%2 == 1 {
		return arr[len(arr)/2]
	} else {
		return (arr[len(arr)/2] + arr[len(arr)/2-1]) / 2
	}
}

func CalculateAvg(arr []float64) float64 {
	var sum float64 = 0
	for _, num := range arr {
		sum += num
	}
	return sum / float64(len(arr))
}

func CalculateVar(arr []float64) float64 {
	var sum float64 = 0
	average := CalculateAvg(arr)
	for _, num := range arr {
		sum += math.Pow(num-average, 2)
	}
	return sum / float64(len(arr))
}

func CalculateStd(arr []float64) float64 {
	return math.Sqrt(CalculateVar(arr))
}

func GetLimits(arr []float64) (float64, float64) {
	average := CalculateAvg(arr)
	stDev := CalculateStd(arr)

	low := average - 0.9*stDev
	high := average + 1.3*stDev
	return low, high
}

func IsOutline(arr []float64, n float64) bool {
	Sort(arr)
	length := len(arr)
	q1 := arr[(length+1)/4]
	q3 := arr[3*((length+1)/4)]
	iqr := q3 - q1
	upperFence := q3 + (1.5 * iqr)
	lowerFence := q1 - (1.5 * iqr)

	if n > upperFence || n < lowerFence {
		return true
	}
	return false
}

func IsOutliner(n int, arr []float64) bool {
	return (n/int(CalculateAvg(arr)) > 2 || float64(n)/(CalculateAvg(arr)) < 0.5)
}
