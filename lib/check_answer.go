package lib

import "slices"

func CheckAnswer(words []string, wordGuess string) bool {
	wordIdx := slices.Index[[]string](words, wordGuess)

	return wordIdx != -1
}
