package funcations

import (
	"math"
)

func Variance(data []int) int {
	avg := float64(Average(data))
	sum := 0.0
	for _, value := range data {
		diff := float64(value) - avg
		sum += diff * diff
	}
	return int(math.Round(sum / float64(len(data))))
}
