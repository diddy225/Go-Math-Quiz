package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	var userAnswer string
	var correct int
	var wrong int

	customTimer := flag.Int("userTime", 30, "Expected an Int for time in seconds")
	fileName := flag.String("csv", "problems.csv", "Expected file should be a .csv file, the default is problem.csv")
	flag.Parse()

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

	examStart(*customTimer)

	now := time.Now()
	after := now.Add(time.Duration(*customTimer) * time.Second)

	for i, line := range problems {
		now = time.Now()

		fmt.Printf("Problem #%d: %s = ", i+1, line.question)
		fmt.Scanf("%s\n", &userAnswer)

		if formatAnswer(userAnswer) == line.answer {
			correct++
		} else {
			wrong++
		}
		if now.After(after) {
			break
		}
	}

	fmt.Printf("You got %d out of %d correct!\n", correct, len(csvLines))
	fmt.Printf("Your grade: %g\n", grade(correct, len(csvLines)))

}

func examStart(time int) {
	fmt.Printf("You will have %v seconds to take this exam.\nPress Enter to begin math test.", time)
	fmt.Scanln()
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

func formatAnswer(answer string) string {
	return strings.TrimSpace(strings.ToLower(answer))
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
