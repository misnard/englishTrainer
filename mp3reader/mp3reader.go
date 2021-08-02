package mp3reader

import (
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

//PlayWord is a function from https://github.com/hajimehoshi/go-mp3/blob/master/example/main.go
func PlayWord(word string, c *oto.Context) error {

	filepath := "data/" + word + ".mp3"
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}
	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}


func InitContext() (c *oto.Context) {
	filepath := "data/sample.mp3"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		log.Fatal(err)
	}

	c, err = oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		log.Fatal(err)
	}

	return c
}