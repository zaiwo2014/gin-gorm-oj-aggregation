package redis

import (
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var REDIS = InitRedis()

func InitRedis() *redis.Client {

	// 初始化 viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	// 从配置文件获取数据库 DSN
	rsn := viper.GetString("database.redis.rsn")

	rdb := redis.NewClient(&redis.Options{
		Addr:     rsn,
		Password: "",
		DB:       0,
		// 添加以下配置来避免使用不兼容的特性
		Protocol: 2, //  旧协议
	})

	return rdb
}
