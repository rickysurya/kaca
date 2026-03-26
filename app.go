package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
    targetLang string
}

func NewApp() *App {
	return &App{targetLang: "id"}
}

func (a *App) setLanguage(lang string){
    a.targetLang = lang
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowFullscreen(ctx)
}

func (a *App) Translate(text string, targetLang string) string {
	result, err := translate(text, targetLang)
	if err != nil {
		return "translation error"
	}
	return result
}

func (a *App) CaptureRegion(x, y, width, height int) ([]byte, error) {
	return captureRegion(x, y, width, height)
}

func (a *App) CaptureAndTranslate(x, y, width, height int, targetLang string) (string, error) {
    imageBytes, err := captureRegion(x, y, width, height)
    if err != nil {
        return "", err
    }

    text, err := ocr(imageBytes)
    if err != nil {
        return "", err
    }

    translated, err := translate(text, targetLang)
    if err != nil {
        return "", err
    }

    return translated, nil
}