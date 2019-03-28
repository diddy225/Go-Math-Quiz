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
	var correct int
	var wrong int

	customTimer := flag.Int("timer", 30, "Expected an Int for time in seconds")
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
	timer := time.NewTimer(time.Duration(*customTimer) * time.Second)
problemLoop:
	for i, line := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, line.question)
		answerCh := make(chan string)

		go func() {
			var userAnswer string
			fmt.Scanf("%s\n", &userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if formatAnswer(answer) == line.answer {
				correct++
			} else {
				wrong++
			}
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
