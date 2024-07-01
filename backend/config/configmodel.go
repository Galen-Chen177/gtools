package config

var Cfg Config

type Config struct {
	Redis RedisConfig `json:"redis"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DB       int    `json:"db"`
	Password string `json:"password"`
}
