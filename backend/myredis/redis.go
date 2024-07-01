package myredis

import (
	"context"
	"fmt"
	"strings"

	"gtools-wails/backend/config"
	"gtools-wails/backend/utils"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type IRedis interface {
	Exec(string) string
}

var redisConn *redis.Client

var redisCxt = context.Background()

func GetCon() (*redis.Client, error) {
	if redisConn == nil {
		redisConn = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", config.Cfg.Redis.Host, config.Cfg.Redis.Port),
			Password: config.Cfg.Redis.Password,
			DB:       config.Cfg.Redis.DB,
		})
		if err := redisConn.Ping(redisCxt).Err(); err != nil {
			logrus.Error("InitRedis failed:", err)
			return nil, err
		}
		config.AddCloseList(Close)
	}
	return redisConn, nil
}

func Close() {
	redisConn.Close()
	redisConn = nil
}

type Redis struct{}

func NewRedis() IRedis {
	return &Redis{}
}

func (r *Redis) Exec(in string) string {
	if in == "" {
		return utils.Resp("", nil)
	}
	cli, err := GetCon()
	if err != nil {
		return utils.Resp("", err)
	}

	inCmd := strings.Fields(in)
	args := inCmd[1:]
	switch strings.ToLower(inCmd[0]) {
	case "set":
		if len(args) < 2 {
			return utils.Resp(nil, fmt.Errorf("SET command requires key and value"))
		}
		key := args[0]
		value := args[1]
		if len(args) > 2 {
			//TODO: int转time.during
			t := cast.ToDuration(args[2])
			return utils.Resp(cli.Set(redisCxt, key, value, t).Result())
		}
		return utils.Resp(cli.Set(redisCxt, key, value, 0).Result())
	case "get":
		if len(args) < 1 {
			return utils.Resp(nil, fmt.Errorf("GET command requires key"))
		}
		key := args[0]
		return utils.Resp(redisConn.Get(redisCxt, key).Result())
	case "del":
		if len(args) < 1 {
			return utils.Resp(nil, fmt.Errorf("DEL command requires at least one key"))
		}
		keys := make([]string, len(args))
		for i, arg := range args {
			keys[i] = arg
		}
		return utils.Resp(redisConn.Del(redisCxt, keys...).Result())
	case "keys":
		if len(args) > 1 {
			return utils.Resp(nil, fmt.Errorf("KEYS command takes zero or one pattern argument"))
		}
		var pattern string
		if len(args) == 1 {
			pattern = args[0]
		} else {
			pattern = "*"
		}
		return utils.Resp(redisConn.Keys(redisCxt, pattern).Result())
	case "ttl":
		return utils.Resp(redisConn.TTL(redisCxt, args[0]).Result())

	// TODO:添加更多 Redis 命令的处理逻辑
	default:
		return utils.Resp("", fmt.Errorf("Unknown command: %v", in))
	}
}
