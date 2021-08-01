package main

import (
	"englishTrainer/csvparse"
	"englishTrainer/randwords"
	"englishTrainer/speech"
	"englishTrainer/trainer"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	words := csvparse.ParseCsv("data.csv")
	words = randwords.RandWords(words)
	speech.LoadWordsData(words)

	trainer.StartTrain(words)
}
