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

	var correctAnswers int

	for i, record := range records {
		var answer string
		fmt.Printf("Problem #%d: %v = ", i, record[0])
		fmt.Scanln(&answer)
		if answer == record[1] {
			correctAnswers += 1
		}
	}

	fmt.Printf("Total correct answers: %d\n", correctAnswers)
}
