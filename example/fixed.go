package main

import (
	"fmt"
	"io/ioutil"

	"github.com/brentnd/go-snowboy"
)

func main() {
	words := snowboy.Hotwords{}
	words.Add(snowboy.Hotword{
		Model: "./build/alexa.umdl",
		Sensitivity: 0.5,
		Keyword: snowboy.KeywordAlexa,
	})

	d := snowboy.NewDetector("./build/common.res", words)
	defer d.Close()

	wav, err := ioutil.ReadFile("./audio/alexa_request.wav")
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