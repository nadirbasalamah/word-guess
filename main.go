package main

import (
	"fmt"
	"guess-word/lib"
	"guess-word/models"
	"log"
)

var wordbank models.WordBank
var userChoice int
var wordDatum models.Datum
var userAnswers []string

func main() {
	readWordBank()

	mainMenu()

	//TODO: start the game to guess the words based on the given letters
	mainGame()

	//TODO: if all words guessed, exit the game
}

func readWordBank() {
	var err error

	wordbank, err = lib.ReadWordbank()

	if err != nil {
		log.Fatalf("error when reading word bank: %v\n", err)
	}
}

func mainMenu() {
	fmt.Println("Welcome to Word Guess Game")
	fmt.Println("Choose the letters bag")

	fmt.Println("=====")
	for idx := range wordbank.Data {
		fmt.Printf("%d) Letter Bag %d\n", idx+1, idx+1)
	}
	fmt.Println("=====")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&userChoice)

	isValid := lib.ValidateChoice(userChoice, len(wordbank.Data))

	if !isValid {
		fmt.Println("invalid input, please try again")
	}

	userChoice--

	wordDatum = wordbank.Data[userChoice]
}

func mainGame() {
	var wordInput string

	fmt.Println("Your Letters")
	for _, letter := range wordDatum.Letters {
		fmt.Printf("[%s] ", letter)
	}
	fmt.Println("")

	fmt.Print("Enter a word: ")
	fmt.Scan(&wordInput)

	isValid := lib.ValidateAnswerInput(wordInput)

	if !isValid {
		fmt.Println("invalid input, please try again")
	}

	word := lib.SanitizeInput(wordInput)

	userAnswers = append(userAnswers, word)

	isDuplicate := lib.IsDuplicate(userAnswers, word)
	isCorrect := lib.CheckAnswer(wordDatum.Answers, word)

	fmt.Println("is duplicate?: ", isDuplicate)

	if isCorrect && !isDuplicate {
		fmt.Println("your answers is: ", userAnswers)
	} else {
		fmt.Println("wrong answer, please try again")
	}
}
