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
	var totalProblems int

	for i, record := range records {
		totalProblems += 1
		var answer string
		fmt.Printf("Problem #%d: %v = ", i, record[0])
		fmt.Scanln(&answer)
		if answer == record[1] {
			correctAnswers += 1
		}
	}

	fmt.Printf("You scored %d out of %d\n", correctAnswers, totalProblems)
}
