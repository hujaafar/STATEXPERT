package funcations

import "math"

func StdDev(data []int) int {
	return int(math.Round(math.Sqrt(float64(Variance(data)))))
}
