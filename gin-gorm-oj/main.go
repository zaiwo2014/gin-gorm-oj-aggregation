package main

import "gin-gorm-oj/router"

func main() {
	// 创建默认的 Gin 路由引擎
	r := router.Router()

	// 启动服务器，默认监听 :8080 端口
	r.Run(":9981")
}

//func main() {
//	rdb := redis.InitRedis()
//	ctx := context.Background()
//	test := rdb.Get(ctx, "test")
//	fmt.Print(test)
//}
