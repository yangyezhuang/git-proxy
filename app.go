package main

import (
	"context"
	"git-proxy/backend"
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
}

func (a *App) SwitchMode(mode string) string {
	address := backend.SwitchProxy(mode)
	return address
}

func (a *App) QueryConfig() map[string]string {
	config, _ := backend.QueryConfig()
	return config
}

func (a *App) ResetSettings() string {
	result := backend.ClearConfig()
	return result
}

func (a *App) SaveSettings(http string, socks string) string {
	result := backend.SaveConfig(http, socks)
	return result
}
