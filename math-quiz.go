package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	var answer int
	var correct int
	var wrong int

	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContent := csv.NewReader(file)

	quiz, err := fileContent.ReadAll()

	for i := range quiz {
		fmt.Printf("Problem #%v: %v = ", i+1, quiz[i][0])
		_, err := fmt.Scanf("%d", &answer)
		if err != nil {
			log.Fatal(err)
		}
		if strconv.Itoa(answer) == quiz[i][1] {
			fmt.Println("Correct")
			correct++
		} else {
			fmt.Println("Wrong")
			wrong++
		}
	}
	fmt.Printf("You got %v out of %v.\n", strconv.Itoa(correct), strconv.Itoa(len(quiz)))
	fmt.Printf("Your Grade: %v\n", grade(correct, len(quiz)))
}

func grade(num1 int, num2 int) float64 {
	return math.Floor(float64(num1) / float64(num2) * float64(100))
}

//import the csv file

//read its content

//loop through the key values

//store values

//capture correct/incorrect answers

//give the user the results
