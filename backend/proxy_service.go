package backend

import (
	"fmt"
	"os/exec"
	"syscall"
)

func SwitchProxy(mode string) string {
	addr, err := readConfig(mode)
	if err != nil {
		return fmt.Sprintf("Failed to read config: %v", err)
	}

	switch mode {
	case "http":
		setProxy("http", "http://"+addr)
		setProxy("https", "http://"+addr)
	case "socks":
		setProxy("http", "socks5://"+addr)
		setProxy("https", "socks5://"+addr)
	default:
		unsetProxy("http")
		unsetProxy("https")
	}
	return addr
}

func unsetProxy(scheme string) {
	cmd := exec.Command("git", "config", "--global", "--unset", fmt.Sprintf("%s.proxy", scheme))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to unset %s proxy: %v\n", scheme, err)
	}
}

func setProxy(scheme, proxy string) {
	cmd := exec.Command("git", "config", "--global", fmt.Sprintf("%s.proxy", scheme), proxy)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to set %s proxy to %s: %v\n", scheme, proxy, err)
	}
}
