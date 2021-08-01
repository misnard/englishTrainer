package randwords

import (
	"englishTrainer/words"
	"math/rand"
	"time"
)

//RandWords receive an array of csvparse.WordCombination and return shuffeled csvparse.WordCombination
func RandWords(words []*words.WordCombination) []*words.WordCombination {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })

	return words
}
