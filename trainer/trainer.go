package trainer

import (
	"bufio"
	"englishTrainer/mp3reader"
	"englishTrainer/words"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

//StartTrain take in entry words []*words.WordCombination and is the principal function who start training
func StartTrain(words []*words.WordCombination) {
	fmt.Println("Welcome to the English / French trainer !")

	for _, w := range words {
		var question string
		var answer string
		// Generate rand true or false
		if rand.Intn(2) == 0 {
			question = w.FrenchWord
			answer = w.EnglishWord
		} else {
			question = w.EnglishWord
			answer = w.FrenchWord
		}

		guessWord(question, answer)
	}

	fmt.Println("Your training was finished if you want to continue relaunch the application or add words in to the source file :)")
}

func guessWord(question, answer string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("---------------------------------------")
		fmt.Println("What is the translation of the word you heard ?")
		fmt.Println("If you want to replay the sound type replay else answer.")

		mp3reader.PlayWord(question)

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == answer {
			fmt.Println("Good anwser, well done ! :)")
			return
		}
	}
}
