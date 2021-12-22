package main

import (
	CONF "zero_blog/conf"
	ROUTERS "zero_blog/routers"
)

func main() {
	// 从配置文件读取配置
	CONF.Init()

	// 装载路由
	r := ROUTERS.AllRouter()
	r.Run(":3000")
}
