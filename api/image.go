package handler

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const IMAGE_SIZE = 600

func Handler(w http.ResponseWriter, r *http.Request) {
	background := getRandomColor()
	image := createImage(background)

	writeImage(w, &image)
}

func getRandNumberInRange(max int, min int) uint8 {
	return uint8(rand.Intn(max-min) + min)
}

func getRandomColor() color.RGBA {
	rand.Seed(time.Now().UnixNano())

	min := 0
	max := 255

	color := color.RGBA{
		getRandNumberInRange(max, min),
		getRandNumberInRange(max, min),
		getRandNumberInRange(max, min),
		1,
	}

	return color
}

// create rectangular image from a RGBA color
func createImage(background color.RGBA) image.Image {
	rect := image.Rect(0, 0, IMAGE_SIZE, IMAGE_SIZE)
	m := image.NewRGBA(rect)
	draw.Draw(m, m.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)
	var img image.Image = m

	return img
}

// writeImage encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)

	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
