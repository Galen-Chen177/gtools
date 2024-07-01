package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"gtools-wails/backend"
	"gtools-wails/backend/config"
	"gtools-wails/backend/dialog"
	"gtools-wails/backend/global"
	"gtools-wails/backend/jsonfunc"
	"gtools-wails/backend/log"
	"gtools-wails/backend/myredis"
	"gtools-wails/backend/utils"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	MainInit()

	// Create an instance of the app structure
	app := backend.NewApp()
	jsonFunc := jsonfunc.NewJsonFunc()
	dialog := &dialog.BackendDialog{}
	configFunc := config.NewConfig()
	logFunc := log.NewLog()
	redisFunc := myredis.NewRedis()

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
			configFunc,
			logFunc,
			redisFunc,
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
	// 初始化log
	pwd, _ := os.Getwd()
	logf, err := os.OpenFile(filepath.Join(pwd, "gtools-wails.log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(logf)

	logrus.Infoln("begin")
	// 创建配置文件夹，以及写入信息
	// 判断文件在不在，在的话不管，不在的话创建，并写入默认的信息
	global.ConfigFullName = filepath.Join(pwd, "config", global.ConfigName)
	if ok, err := utils.PathExists(global.ConfigFullName); err != nil {
		logrus.Panic(err)
	} else if !ok {
		config.WriteDefaultConfig()
	} else {
		config.ReadConfig()
	}
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
