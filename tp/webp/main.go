package main

import (
	"image/png"
	"log"
	"os"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func main() {
	file, err := os.Open("input.png")
	if err != nil {
		log.Fatalln(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create("output.webp")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	if err := webp.Encode(output, img, options); err != nil {
		log.Fatalln(err)
	}
}
