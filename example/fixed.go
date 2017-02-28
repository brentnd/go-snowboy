package main

import (
	"fmt"
	"io/ioutil"

	"github.com/brentnd/go-snowboy"
)

func main() {
	d := snowboy.Detector{
		ResourceFilename: "./build/common.res",
		AudioGain: 1.0,
	}
	defer d.Close()

	d.AddHotword(snowboy.Hotword{
		Model: "./build/alexa.umdl",
		Sensitivity: 0.5,
		Keyword: snowboy.KeywordAlexa,
	})

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