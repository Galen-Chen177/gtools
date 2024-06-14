package config

import (
	_ "embed"
	"os"
	"path/filepath"

	"gtools-wails/backend/global"
)

//go:embed config_temp.yaml
var tempConfig []byte

func WriteDefaultConfig() {
	// 创建并写入
	pwd, _ := os.Getwd()
	os.MkdirAll(filepath.Join(pwd, "config"), 0755)
	if err := os.WriteFile(global.ConfigFullName, tempConfig, 0664); err != nil {
		panic(err)
	}
}
