package config

import (
	_ "embed"
	"encoding/json"
	"os"
	"path/filepath"

	"gtools-wails/backend/global"
	"gtools-wails/backend/utils"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

//go:embed config_temp.yaml
var tempConfig []byte

func WriteDefaultConfig() {
	// 创建并写入
	pwd, _ := os.Getwd()
	os.MkdirAll(filepath.Join(pwd, "config"), 0755)
	if err := os.WriteFile(global.ConfigFullName, tempConfig, 0664); err != nil {
		logrus.Panic(err)
	}
	ReadConfig()
}

func ReadConfig() error {
	f, err := os.ReadFile(global.ConfigFullName)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if err := yaml.Unmarshal(f, &Cfg); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func WriteConfig(configJson string) error {
	if len(configJson) < 10 {
		return nil
	}
	// json--> config -->yamlfile
	if err := json.Unmarshal([]byte(configJson), &Cfg); err != nil {
		logrus.Error(err)
		return err
	}

	// TODO:可以适当增加一些校验

	data, err := yaml.Marshal(&Cfg)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if err := os.WriteFile(global.ConfigFullName, data, 0664); err != nil {
		logrus.Error(err)
		return err
	}

	// 由于连接信息有可能更改，关闭链接，需要时重新链接
	for _, v := range CloseList {
		v()
	}
	CloseList = []func(){}

	return nil
}

type IConfig interface {
	Read() string
	Write(string) string
}

type ConfigFrontend struct{}

func NewConfig() IConfig {
	return &ConfigFrontend{}
}

func (c *ConfigFrontend) Read() string {
	if err := ReadConfig(); err != nil {
		return utils.Resp("", err)
	}
	res := map[string]any{
		"redishost": Cfg.Redis.Host,
		"redisport": Cfg.Redis.Port,
		"redisdb":   Cfg.Redis.DB,
		"redispwd":  Cfg.Redis.Password,
	}

	return utils.Resp(res, nil)
}

func (c *ConfigFrontend) Write(in string) string {
	logrus.Info("Write config:", in)
	return utils.Resp("", WriteConfig(in))
}

var CloseList = []func(){}

func AddCloseList(f func()) {
	CloseList = append(CloseList, f)
}
