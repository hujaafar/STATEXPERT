package funcations

import "math"



func Median(data []int) int {
	sortedData := make([]int, len(data))
    copy(sortedData, data)
    sortInts(sortedData)

    mid := len(sortedData) / 2
    if len(sortedData)%2 == 0 {
        return int(math.Round(float64(sortedData[mid-1]+sortedData[mid]) / 2.0))
    }
    return sortedData[mid]
}
