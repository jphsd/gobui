package main

import (
	"flag"
	"image"
	"image/png"
	"os"

	"github.com/jphsd/gobui"
)

// Display a PNG file in a browser at localhost:8080
func main() {
	flag.Parse()
	img, _ := readImage(flag.Args()[0])

	disp := gobui.NewDisplay(8080, ".")
	disp.Load(img)

	for {}
}

// Read a PNG image
func readImage(name string) (image.Image, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}
