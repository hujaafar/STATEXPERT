package main

import (
	"fmt"
	"func/funcations"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("go run your-program.go data.txt")
		os.Exit(1)
	}

	FileName := os.Args[1]
	data, error := funcations.ReadData(FileName)
	if error != nil {
		fmt.Println("Wrong")
		os.Exit(1)
	}

	Average := funcations.Average(data)
	Median := funcations.Median(data)
	variance := funcations.Variance(data)
	Standerd := funcations.StdDev(data)

	fmt.Println("Average:", Average)
	fmt.Println("Median:", Median)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", Standerd)
}
