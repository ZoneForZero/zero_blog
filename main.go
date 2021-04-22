package main

import (
	"zero_blog/conf"
	"zero_blog/routers"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := routers.AllRouter()
	r.Run(":3000")
}
