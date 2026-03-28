package main

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/hotkey"
)

func (a *App) registerHotkey() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyX)
	err := hk.Register()
	if err != nil {
		return
	}
	go func() {
		for range hk.Keydown() {
			if runtime.WindowIsNormal(a.ctx) {
				runtime.WindowHide(a.ctx)
			} else {
				runtime.WindowShow(a.ctx)
			}
		}
	}()
}
