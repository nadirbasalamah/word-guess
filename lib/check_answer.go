package lib

import "slices"

func IsDuplicate(answers []string, wordGuess string) bool {
	var freqs map[string]int = make(map[string]int)

	for _, answer := range answers {
		freqs[answer]++
	}

	return freqs[wordGuess] > 1
}

func CheckAnswer(words []string, wordGuess string) bool {
	wordIdx := slices.Index[[]string](words, wordGuess)

	return wordIdx != -1
}
