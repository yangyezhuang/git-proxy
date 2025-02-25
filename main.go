package main

import (
	"embed"
	"fmt"
	"git-proxy/service"
	"github.com/getlantern/systray"
	"github.com/go-ini/ini"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed favicon.ico
var appIcon []byte

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// 系统托盘
	go systray.Run(func() { onReady(app) }, onExit)

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "GitProxy",
		Width:         220,
		Height:        355,
		Frameless:     true,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			// 禁用缩放手势
			DisablePinchZoom:    true,
			Theme:               windows.SystemDefault,
			BackdropType:        windows.Acrylic,
			WindowIsTranslucent: true,
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

func onReady(a *App) {
	systray.SetIcon(appIcon)
	systray.SetTitle("GitProxy")
	systray.SetTooltip("GitProxy")

	// 定义菜单项
	mDefault := systray.AddMenuItem("No Proxy", "Switch Default")
	mProxy := systray.AddMenuItem("Proxy", "Switch Proxy")
	mHttp := mProxy.AddSubMenuItem("Http", "Switch Http")
	mSocks := mProxy.AddSubMenuItem("Socks", "Switch Socks")
	systray.AddSeparator()
	mShow := systray.AddMenuItem("Show", "Show the application")
	mHide := systray.AddMenuItem("Hide", "Hide the application")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	go func() {
		for {
			select {
			case <-mDefault.ClickedCh:
				service.SwitchProxy("default")
			case <-mHttp.ClickedCh:
				service.SwitchProxy("http")
			case <-mSocks.ClickedCh:
				service.SwitchProxy("socks")
			case <-mShow.ClickedCh:
				runtime.Show(a.ctx)
			case <-mHide.ClickedCh:
				runtime.Hide(a.ctx)
			case <-mQuit.ClickedCh:
				runtime.Quit(a.ctx)
			}
		}
	}()
}

func onExit() {
	// 清理工作
}
