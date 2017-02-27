package main

import (
	"os"
	"fmt"
	"io/ioutil"

	"github.com/brentnd/go-snowboy"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("usage: %s <resource> <keyword.umdl> <audio file>\n", os.Args[0])
		return
	}
	d := snowboy.NewSnowboyDetector(os.Args[1], os.Args[2], "0.5", 1.0)
	defer d.Close()

	wav, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		panic(err)
	}

	res := d.RunDetection(wav)
	fmt.Printf("detected %d\n", res)
}