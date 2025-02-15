package backend

import (
	"log"
	"os/exec"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func onReady() {
	// systray.SetTemplateIcon(iconData, iconData)
	systray.SetIcon(icon.Data)
	systray.SetTitle("GitProxy")
	about := systray.AddMenuItem("关于", "GitProxy")
	systray.AddSeparator()
	openHelp := systray.AddMenuItem("使用帮助", "打开使用帮助")
	openOrigination := systray.AddMenuItem("访问官网", "https://www.anqicms.com")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "GitProxy")

	go func() {
		for {
			select {
			case <-about.ClickedCh:
				openURL("https://www.anqicms.com/about.html")
			case <-openHelp.ClickedCh:
				openURL("https://www.anqicms.com/help")
			case <-openOrigination.ClickedCh:
				openURL("https://www.anqicms.com/")
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func openURL(url string) {
	cmd := exec.Command("cmd", "/c", "start", url) // Windows
	// For macOS, you can use: cmd := exec.Command("open", url)
	// For Linux with xdg-open: cmd := exec.Command("xdg-open", url)

	err := cmd.Start()
	if err != nil {
		log.Println("Error opening URL:", err)
	}
}

func onExit() {
	// clean up here
	log.Println("退出程序")
}
