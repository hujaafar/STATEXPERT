package funcations

import "math"

func Average(data []int) int {
	sum := 0
    for _, value := range data {
        sum += value
    }
    return int(math.Round(float64(sum) / float64(len(data))))
}