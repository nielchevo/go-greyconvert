package main

import (
	"image"
	"log"
	"os"
)

func otsu(greyFile os.File) {

	readerFile, err := os.Open(greyFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer readerFile.Close()

	newImg, format, err := image.Decode(readerFile)
	if err != nil {
		log.Fatal(err)
	}
	if format != "jpeg" {
		log.Fatalf("only JPEG/ .jpg format is supported")
	}

	// init variables
	threshold := 0

}
