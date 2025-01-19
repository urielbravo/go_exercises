package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println("Error openinf file:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for i, record := range records {
		// fmt.Printf("Row %d: %v\n", i, record)
		fmt.Printf("Problem #%d: %v\n", i, record[0])
	}
}
