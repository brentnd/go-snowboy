package main

import (
	"os"
	"fmt"
	"strings"

	"github.com/brentnd/go-snowboy"
)

func handleDetection(result string) {
	fmt.Println("Keyword Detected:", result)
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("usage: %s <resource> <keyword.umdl> <audio file>\n", os.Args[0])
		return
	}

	d := snowboy.Detector{
		ResourceFile: os.Args[1],
		AudioGain: 1.0,
	}
	defer d.Close()

	// Parse keyword from .umdl file path.
	// If keyword is known (not command line), just use
	// snowboy.Keyword("i know it already")
	k := os.Args[2]
	k = strings.TrimRight(k, ".umdl")
	kParts := strings.Split(k, "/")

	d.HandleFunc(snowboy.Hotword{
		Model: os.Args[2],
		Sensitivity: 0.5,
		Name: kParts[len(kParts) - 1],
	}, handleDetection)

	f, err := os.Open(os.Args[3])
	if err != nil {
		panic(err)
	}
	d.ReadAndDetect(f)
}