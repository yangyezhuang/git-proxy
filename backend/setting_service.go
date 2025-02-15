package backend

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"

	"github.com/go-ini/ini"
)

// 获取用户主目录
func getUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user: %v", err)
	}
	return usr.HomeDir
}

// 获取配置文件路径
func getConfFilePath() string {
	// 获取用户主目录
	homeDir := getUserHomeDir()
	fmt.Printf("User Home Directory: %s\n", homeDir)

	// 定义 ini文件路径
	filePath := filepath.Join(homeDir, ".gitproxy", "config.ini")
	fmt.Printf("Config FullPath: %s\n", filePath)

	return filePath
}

// 读取 ini文件
func readConfig(mode string) (string, error) {
	filePath := getConfFilePath()
	addr := ""

	cfg, err := ini.Load(filePath)
	if err != nil {
		return addr, err
	}

	// 读取配置
	fmt.Printf("proxy mode: %s\n", mode)
	addr = cfg.Section("proxy").Key(mode).String()
	fmt.Printf("%s addr: %s\n", mode, addr)

	return addr, nil
}

// 查询配置
func QueryConfig() (map[string]string, error) {
	filePath := getConfFilePath()
	httpAddr := ""
	socksAddr := ""
	var confMap = make(map[string]string)

	cfg, err := ini.Load(filePath)
	if err != nil {
		return confMap, err
	}

	// 读取配置
	httpAddr = cfg.Section("proxy").Key("http").String()
	socksAddr = cfg.Section("proxy").Key("socks").String()
	confMap["httpAddr"] = httpAddr
	confMap["socksAddr"] = socksAddr

	for key, value := range confMap {
		fmt.Printf("Key: %s, Addr: %s\n", key, value)
	}

	return confMap, nil
}

func ClearConfig() string {
	filePath := getConfFilePath()
	cfg, err := ini.Load(filePath)
	if err != nil {
		print(err)
	}
	cfg.Section("proxy").Key("http").SetValue("")
	cfg.Section("proxy").Key("socks").SetValue("")
	cfg.SaveTo(filePath)
	return "ok"
}

func SaveConfig(http string, socks string) string {
	filePath := getConfFilePath()

	cfg, err := ini.Load(filePath)
	if err != nil {
		print(err)
	}
	cfg.Section("proxy").Key("http").SetValue(http)
	cfg.Section("proxy").Key("socks").SetValue(socks)
	cfg.SaveTo(filePath)
	return "ok"
}
