package main

import (
	"os"
	"fmt"

	"github.com/brentnd/go-snowboy"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 6 {
		fmt.Printf("usage: %s <token> <name> <wav1> <wav2> <wav3>\n", os.Args[0])
		return
	}
	t := snowboy.TrainRequest{
		Token: os.Args[1],
		Name: os.Args[2],
		Language: snowboy.LanguageEnglish,
		AgeGroup: snowboy.AgeGroup20s,
		Gender: snowboy.GenderMale,
		Microphone: "standard USB mic",
	}
	t.AddWave(os.Args[3])
	t.AddWave(os.Args[4])
	t.AddWave(os.Args[5])
	pmdl, err := t.Train()
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(t.Name + ".pmdl", pmdl, 0644)
}