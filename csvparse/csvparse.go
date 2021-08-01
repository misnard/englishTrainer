package csvparse

import (
	"englishTrainer/words"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

//ParseCsv get a filePath as argument to open and parse csv file
func ParseCsv(filePath string) []*words.WordCombination {

	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("An error occured on file open: %v", err)
	}
	defer f.Close()

	words := []*words.WordCombination{}

	if err := gocsv.UnmarshalFile(f, &words); err != nil {
		fmt.Printf("An error occured on csv unmarshall: %v", err)
	}

	return words
}
