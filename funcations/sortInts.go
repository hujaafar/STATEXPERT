package funcations

func sortInts(slice []int) {
    for i := 0; i < len(slice); i++ {
        minIdx := i
        for j := i + 1; j < len(slice); j++ {
            if slice[j] < slice[minIdx] {
                minIdx = j
            }
        }
        slice[i], slice[minIdx] = slice[minIdx], slice[i]
    }
}