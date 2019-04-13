package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	readerFile, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer readerFile.Close()

	img, format, err := image.Decode(readerFile)
	if err != nil {
		log.Fatal(err)
	}
	if format != "jpeg" {
		log.Fatalf("only JPEG/ .jpg format is supported")
	}

	fmt.Printf("img size: %d x %d", img.Bounds().Dx(), img.Bounds().Dy())
	//create a new props image
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	//newImg := image.NewRGBA(rect)
	grayImg := image.NewGray(rect)

	// itterate every pixels
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			//extract img rgba value
			pixel := img.At(x, y)

			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			r := float64(originalColor.R) * 0.3
			g := float64(originalColor.G) * 0.59
			b := float64(originalColor.B) * 0.11
			grey := uint8(r + g + b/3)

			// set grey color to new img container
			fmt.Printf("R : %e, G: %e, B: %e, channel: %u\n", r, g, b, grey)
			imgColor := color.RGBA{R: grey, G: grey, B: grey, A: originalColor.A}

			grayImg.Set(x, y, imgColor)
		}
	}

	// Save as new .jpg img
	saveFile, err := os.Create("grayed.jpg")
	if err != nil {
		log.Fatal(err)
	}
	err = jpeg.Encode(saveFile, grayImg, nil)

	defer saveFile.Close()
}
