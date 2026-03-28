package main

import (
	"golang.design/x/hotkey"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) registerHotkey() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyX)
	err := hk.Register()
	if err != nil {
		return
	}
	go func() {
		for range hk.Keydown() {
			runtime.WindowShow(a.ctx)
		}
	}()
}