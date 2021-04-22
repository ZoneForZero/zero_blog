package conf

import (
	"os"
	CACHE "zero_blog/cache"
	MODEL "zero_blog/model"
	UTIL "zero_blog/util"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	UTIL.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		UTIL.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	MODEL.Database(os.Getenv("MYSQL_DSN"))
	CACHE.Redis()
}
