package lib

import "strings"

func ValidateChoice(input, numOfChoices int) bool {
	isInvalid := input <= 0 || input >= numOfChoices+1

	return !isInvalid
}

func ValidateAnswerInput(input string) bool {
	return input != ""
}

func SanitizeInput(input string) string {
	return strings.ToLower(input)
}
