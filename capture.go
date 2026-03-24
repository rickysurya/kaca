package main

import (
	"bytes"
	"image"
	"image/png"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)


func (a *App) CaptureRegion(x, y, width, height int) ([]byte, error) {
	runtime.WindowHide(a.ctx)
	time.Sleep(150 * time.Millisecond)

    rect := image.Rect(x, y, x+width, y+height)
    img, err := screenshot.CaptureRect(rect)
    if err != nil {
        return nil, err
    }
	// temporary buffer storage for the image
    buf := new(bytes.Buffer)
    err = png.Encode(buf, img)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}
