package main

import (
	"bytes"
	"image"
	"image/png"

	"github.com/kbinani/screenshot"
)

func captureRegion(x, y, width, height int) ([]byte, error) {
	rect := image.Rect(x, y, x+width, y+height)
	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		return nil, err
	}
	
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
