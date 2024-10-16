package initialize

import (
	"log"

	"github.com/PandaPy/pginer/template/initialize/cmd"
	"github.com/PandaPy/pginer/template/initialize/config"
	"github.com/PandaPy/pginer/template/initialize/db"
)

func Cmd() {
	config.Init() // 初始化配置
	db.Init()     // 初始化数据库
	if err := cmd.Execute(); err != nil {
		log.Fatalf("命令执行失败: %v", err)
	}
}
