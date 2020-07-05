package main

import (
	"log"
	"os"
	"image"
	"image/color"
	"image/png"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("[Usage] ./lcd-calibration-image-generator w h output")
	}

	w, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	h, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	for x := 0; x < w; x++ {
		isWhite := x % 2 == 1

		for y := 0; y < h; y++ {
			if isWhite {
				img.Set(x, y, white)
			} else {
				img.Set(x, y, black)
			}

			isWhite = !isWhite
		}
	}

	file, err := os.Create(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if err = png.Encode(file, img); err != nil {
		log.Fatal(err)
	}
}


