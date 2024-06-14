package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"gtools-wails/backend"
	"gtools-wails/backend/dialog"
	"gtools-wails/backend/global"
	"gtools-wails/backend/jsonfunc"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	MainInit()

	time.Sleep(3600 * time.Second)

	// Create an instance of the app structure
	app := backend.NewApp()
	jsonFunc := jsonfunc.NewJsonFunc()
	dialog := &dialog.BackendDialog{}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "gtools-wails",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: dialog.Startup,
		Bind: []interface{}{
			app,
			jsonFunc,
			dialog,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func MainInit() {
	// 只允许同时执行一个该程序
	if checkProgramRunning(global.ApplicationName) {
		os.Exit(0)
	}
	// 创建配置文件夹，以及写入信息
}

func checkProgramRunning(program string) bool {
	pid := os.Getpid()
	cmd := exec.Command("ps", "-ef")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("执行 ps 命令出错:", err)
		return false
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, fmt.Sprintf(" %d ", pid)) {
			continue
		}
		if strings.Contains(line, program) {
			fmt.Println(line)
			return true
		}
	}
	return false
}
