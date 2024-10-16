package db

import (
	"log"

	"github.com/PandaPy/pginer/template/initialize/config"
	"github.com/PandaPy/pginer/template/models"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBs map[string]*gorm.DB

// 初始化所有 MySQL 数据库连接
func Init() {
	DBs = make(map[string]*gorm.DB)
	for name, dbConfig := range config.AppConfig.Databases {
		color.Green("初始化数据库连接(%s)  ", dbConfig.Name)

		mysqlConfig := mysql.Config{
			DSN:                       dbConfig.Dsn(), // DSN data source name
			DefaultStringSize:         191,            // string 类型字段的默认长度
			SkipInitializeWithVersion: false,          // 根据版本自动配置
		}

		db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
		if err != nil {
			zap.L().Fatal("failed to initialize database", zap.Error(err))
		}

		// 获取底层的 sql.DB 对象并设置连接池配置
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("获取 sql.DB 失败: %v", err)
		}

		// 设置最大空闲连接数和最大打开连接数
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)

		// 检查 AutoMigrate 配置并决定是否进行自动迁移
		if dbConfig.AutoMigrate {
			err = db.AutoMigrate(models.Models...)
			if err != nil {
				zap.L().Fatal("failed to migrate database", zap.Error(err))
			}
		}

		// 保存数据库连接到 map 中
		DBs[name] = db
	}
}

// 获取指定的 MySQL 数据库连接，默认返回 default 数据库
func DB(name ...string) *gorm.DB {
	if len(name) == 0 {
		name = append(name, "default")
	}
	db, exists := DBs[name[0]]
	if !exists {
		log.Fatalf("数据库 %v 不存在", name[0])
	}
	return db
}
