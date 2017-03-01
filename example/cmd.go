package main

import (
	"os"
	"fmt"

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

	d.HandleFunc(snowboy.NewHotword(os.Args[2], 0.5), handleDetection)

	d.HandleSilenceFunc(func(string) {
		fmt.Println("silence detected")
	})

	f, err := os.Open(os.Args[3])
	if err != nil {
		panic(err)
	}
	d.ReadAndDetect(f)
}