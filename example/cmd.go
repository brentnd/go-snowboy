package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/brentnd/go-snowboy"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("usage: %s <resource> <keyword.umdl> <audio file>\n", os.Args[0])
		return
	}

	// Parse keyword from .umdl file path.
	// If keyword is known (not command line), just use
	// snowboy.Keyword("i know it already")
	k := os.Args[2]
	k = strings.TrimRight(k, ".umdl")
	kParts := strings.Split(k, "/")

	words := snowboy.Hotwords{}
	words.Add(snowboy.Hotword{
		Model: os.Args[2],
		Sensitivity: 0.5,
		Keyword: snowboy.Keyword(kParts[len(kParts) - 1]),
	})

	d := snowboy.NewDetector(os.Args[1], words)
	d.SetAudioGain(1.0)
	defer d.Close()

	wav, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		panic(err)
	}

	res, err := d.RunDetection(wav)
	if res != nil {
		fmt.Printf("detected keyword %v, error %v\n", *res, err)
	} else {
		fmt.Printf("detect no keyword, error %v\n", err)
	}
}