package main

import (
	"embed"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "GitProxy",
		Width:         600,
		Height:        400,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			// 禁用缩放手势
			DisablePinchZoom: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	// 获取配置文件路径
	configFilePath := getConfigFilePath()

	// 检查配置文件是否存在
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// 文件不存在则创建配置文件
		fmt.Println("Config file does not exist, creating...")
		createFile(configFilePath)
	} else if err != nil {
		// 处理其他可能的错误
		fmt.Printf("Error checking config file: %v\n", err)
		// 根据需要处理错误，比如退出程序
		// os.Exit(1)
	}
}

// 获取配置文件路径
func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err) // 或者处理错误
	}
	return filepath.Join(homeDir, ".gitproxy", "config.ini")
}

// 创建配置文件
func createFile(filePath string) {
	// 确保配置文件所在的目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err) // 或者处理错误
	}

	// 创建并写入初始配置到 ini文件
	cfg := ini.Empty()
	section, err := cfg.NewSection("proxy")
	if err != nil {
		panic(err) // 或者处理错误
	}
	section.NewKey("http", "")
	section.NewKey("socks", "")

	if err := cfg.SaveTo(filePath); err != nil {
		panic(err) // 或者处理错误
	}
}
