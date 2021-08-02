package speech

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"englishTrainer/words"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

//LoadWordsData take in entry []*csvparse.WordCombination and convert csvparse.WordCombination to .mp3 file
//If the word does not exists write it to data folder else skip download
func LoadWordsData(w []*words.WordCombination) {

	c := make(chan bool)

	//Check if output dir exisits
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0755)
	}

	w = append(w, &words.WordCombination{FrenchWord: "sample", EnglishWord: "sample"})

	for _, word := range w {
		fwFileName := "data/" + word.FrenchWord + ".mp3"
		ewFileName := "data/" + word.EnglishWord + ".mp3"

		go convert(word.FrenchWord, fwFileName, "fr-FR", c)
		go convert(word.EnglishWord, ewFileName, "en-US", c)

		fwOk, ewOk := <-c, <-c

		if !fwOk || !ewOk {
			log.Fatalf("An Error occured on word conversion French Word: '%v' English Word: '%v'\n", fwOk, ewOk)
		}
	}

	fmt.Println("All words loaded successfully !")
}

func convert(sentence, filename, languageCode string, c chan bool) {

	// Already converted
	if _, err := os.Stat(filename); err == nil {
		c <- true
		return
	}

	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		c <- false
		log.Fatal(err)
	}
	defer client.Close()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: sentence},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: languageCode,
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		c <- false
		log.Fatal(err)
	}

	// The resp's AudioContent is binary.
	err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		c <- false
		log.Fatal(err)
	}
	fmt.Printf("Audio content written to file: %v\n", filename)
	c <- true
}
