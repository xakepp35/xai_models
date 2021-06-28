package controllers

import (
	"image"
	"image/color"
	"image/png"
	"net/http"

	"github.com/xakepp35/xai_models/bresenham"
	"github.com/xakepp35/xai_models/models"
)

func genImage() *image.RGBA {
	width := 2400
	height := 1000

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	//cyan := color.RGBA{100, 200, 200, 0xff}
	red := color.RGBA{255, 0, 0, 0xff}

	input := 5.5
	m := models.NewIzhikModel()
	var n, n2 models.Neuron
	// Set color for each pixel.
	scale := float64(height / 400)
	yPrev := height / 2
	for x := 1; x < width; x++ {
		m.Step(&n, input)
		m.Step(&n2, n.V/10)
		val := n2.V
		y := int(float64(height)/2 - val*scale)
		bresenham.Bresenham(img, x-1, yPrev, x, y, red)
		yPrev = y
	}
	return img
}

func Graph() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderContentType, ImagePNG)
		w.WriteHeader(http.StatusOK)
		img := genImage()
		png.Encode(w, img)
	}
}
