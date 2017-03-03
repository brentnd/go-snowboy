package main

import (
	"os"
	"fmt"
	"time"

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

	d := snowboy.NewDetector(os.Args[1])
	defer d.Close()

	d.HandleFunc(snowboy.NewDefaultHotword(os.Args[2]), handleDetection)

	d.HandleSilenceFunc(500 * time.Millisecond, func(string) {
		fmt.Println("silence detected")
	})

	sampleRate, numChannels, bitDepth := d.AudioFormat()
	fmt.Printf("sample rate=%d, num channels=%d, bit depth=%d\n", sampleRate, numChannels, bitDepth)

	f, err := os.Open(os.Args[3])
	if err != nil {
		panic(err)
	}
	d.ReadAndDetect(f)
}