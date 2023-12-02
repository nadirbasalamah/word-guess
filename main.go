package main

import (
	"fmt"
	"guess-word/lib"
	"guess-word/models"
	"log"
)

var wordbank models.WordBank
var userChoice int

func main() {
	//TODO: print bag of letters
	readWordBank()

	mainMenu()

	//TODO: start the game to guess the words based on the given letters

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
}
