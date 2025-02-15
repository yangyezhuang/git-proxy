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

	// 定义 INI 文件路径
	filePath := filepath.Join(homeDir, "gitproxy.ini")
	fmt.Printf("Config FullPath: %s\n", filePath)

	return filePath
}

// 写入 INI 文件
func writeConfig(filePath string) error {
	cfg := ini.Empty()
	// 添加配置节（section）及键值对
	section1 := cfg.Section("proxy")
	section1.Key("http").SetValue("127.0.0.1:7890")
	section1.Key("socks").SetValue("127.0.0.1:80")

	section2 := cfg.Section("Section2")
	section2.Key("keyA").SetValue("valueA")
	section2.Key("keyB").SetValue("valueB")

	// 写入到文件
	filename := "git_proxy_config.ini"
	if err := cfg.SaveTo(filename); err != nil {
		panic(err)
	}

	// 输出确认信息
	fmt.Printf("INI file '%s' created successfully.\n", filename)

	// // 创建一些示例节和键
	// if err := cfg.Section("User").Key("Name").SetValue("John Doe"); err != nil {
	// 	return err
	// }
	// if err := cfg.Section("User").Key("Age").SetValue("30"); err != nil {
	// 	return err
	// }
	// if err := cfg.Section("App").Key("Version").SetValue("1.0.0"); err != nil {
	// 	return err
	// }

	// // 将配置写入文件
	// if err := cfg.SaveTo(filePath); err != nil {
	// 	return err
	// }

	return nil
}

// 读取 INI 文件
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
