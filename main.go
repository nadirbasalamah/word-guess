package main

import (
	"fmt"
	"guess-word/lib"
	"guess-word/models"
	"log"
	"os"
	"os/exec"
	"runtime"
	"slices"
)

var wordbank models.WordBank
var userChoice int
var wordDatum models.Datum
var tmpAnswers []string
var userAnswers []string
var isChoiceValid bool
var continuePlaying string

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	readWordBank()

	for {
		for !isChoiceValid {
			mainMenu()
			if !isChoiceValid {
				fmt.Println("invalid input, please try again")
			}
		}

		for len(userAnswers) != len(wordDatum.Answers) {
			mainGame()
		}

		fmt.Println("Congratulations! you win the game!")

		cleanup()

		fmt.Print("Continue playing? (y/n): ")
		fmt.Scan(&continuePlaying)

		if continuePlaying == "n" {
			fmt.Println("Bye...")
			break
		} else if continuePlaying != "y" && continuePlaying != "n" {
			fmt.Println("please enter n to quit the game")
		} else {
			clearTerminal()
		}
	}
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
		isChoiceValid = false
		clearTerminal()
		return
	}

	isChoiceValid = true

	userChoice--

	wordDatum = wordbank.Data[userChoice]

	clearTerminal()
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

	tmpAnswers = append(tmpAnswers, word)

	isDuplicate := lib.IsDuplicate(tmpAnswers, word)
	isCorrect := lib.CheckAnswer(wordDatum.Answers, word)

	clearTerminal()

	if isCorrect && !isDuplicate {
		userAnswers = append(userAnswers, word)
		showAnswers(userAnswers)
	} else {
		fmt.Println("wrong answer, please try again")
		showAnswers(userAnswers)
	}
}

func showAnswers(answers []string) {
	fmt.Println("your answers")
	for _, answer := range answers {
		fmt.Printf("[%s] ", answer)
	}
	fmt.Println("")
}

func clearTerminal() {
	clearTerminalFunc, ok := clear[runtime.GOOS]
	if ok {
		clearTerminalFunc()
	} else {
		panic("unsupported platform. clear terminal failed")
	}
}

func cleanup() {
	tmpAnswers = slices.Delete[[]string](tmpAnswers, 0, len(tmpAnswers))
	userAnswers = slices.Delete[[]string](userAnswers, 0, len(userAnswers))
	isChoiceValid = false
}
