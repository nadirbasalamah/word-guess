package lib

import (
	"encoding/json"
	"guess-word/models"
	"io"
	"os"
)

func ReadWordbank() (models.WordBank, error) {
	jsonFile, err := os.Open("data/wordbank.json")

	if err != nil {
		return models.WordBank{}, err
	}

	jsonData, err := io.ReadAll(jsonFile)

	if err != nil {
		return models.WordBank{}, err
	}

	var wordbank models.WordBank

	if err := json.Unmarshal(jsonData, &wordbank); err != nil {
		return models.WordBank{}, err
	}

	return wordbank, nil
}
