package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {
	var userAnswer string
	var correct int
	var wrong int

	fileName := flag.String("csv", "problems.csv", "Expected file should be a .csv file, the default is problem.csv")

	csvFile, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the provided file %s\n", *fileName))
	}

	fileContent := csv.NewReader(csvFile)

	csvLines, err := fileContent.ReadAll()
	if err != nil {
		exit("Failed to parse the provided file")
	}
	problems := parseLines(csvLines)

	for i, line := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, line.question)
		fmt.Scanf("%s\n", &userAnswer)
		if userAnswer == line.answer {
			correct++
		} else {
			wrong++
		}
	}
	fmt.Printf("You got %d out of %d correct!\n", correct, len(csvLines))
	fmt.Printf("Your grade: %g\n", grade(correct, len(csvLines)))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func grade(num1 int, num2 int) float64 {
	return math.Floor(float64(num1) / float64(num2) * float64(100))
}

//import the csv file

//read its content

//loop through the key values

//capture correct/incorrect answers

//give the user the results
