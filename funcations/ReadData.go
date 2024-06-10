package funcations

import (
	"bufio"
	"fmt"
	"os"
)

func ReadData(FileName string) ([]int, error) {
	file, error := os.Open(FileName)
	if error != nil {
		return nil, error
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := 0
		_, error := fmt.Sscanf(scanner.Text(), "%d", &value)
		if error != nil {
			return nil, error
		}
		data = append(data, value)
	}
	if error := scanner.Err(); error != nil {
		return nil, error
	}
	return data, nil
}
