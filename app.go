package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
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
