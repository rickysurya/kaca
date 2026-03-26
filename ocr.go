package main

import (
	"fmt"
	"regexp"
	"strings"

	"os"

	"github.com/otiai10/gosseract/v2"
)

func ocr(imageBytes []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	tessdata := os.Getenv("TESSDATA_PREFIX")
	if tessdata == "" {
		tessdata = "C:/Program Files/Tesseract-OCR/tessdata"
	}
	client.SetTessdataPrefix(tessdata)

	client.SetImageFromBytes(imageBytes)
	text, err := client.Text()
	if err != nil {
		return "", err
	}
	text = strings.TrimSpace(text)
	if text == "" {
		return "", fmt.Errorf("no text found in image")
	}

	re := regexp.MustCompile(`[^\x20-\x7E\s]`)
	text = re.ReplaceAllString(text, "")

	return text, nil
}
