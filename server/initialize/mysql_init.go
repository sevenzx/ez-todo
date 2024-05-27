package initialize

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sevenzx/eztodo/config"
	"github.com/sevenzx/eztodo/core/global"
	"github.com/sevenzx/eztodo/model"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// GormMysql 初始化Gorm的Mysql
func GormMysql(lumberjack *lumberjack.Logger) *gorm.DB {
	c := config.Config.Mysql
	if c.Dbname == "" {
		return nil
	}

	// [参考](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Dsn(), // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), getGormMysqlConfig(lumberjack))
	if err != nil {
		hlog.Error("gorm.Open err", err)
		return nil
	} else {
		// 设置数据库实例选项，指定表的存储引擎
		db.InstanceSet("gorm:table_options", "ENGINE="+c.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
		return db
	}
}

// RegisterTables 初始化表
func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		hlog.Error("register table failed", err)
		os.Exit(0)
	}
	hlog.Info("register table success")
}

// getGormMysqlConfig 获取GORM的Mysql配置
func getGormMysqlConfig(lumberjack *lumberjack.Logger) *gorm.Config {
	c := config.Config.Mysql
	conf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: c.SingularTable,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(log.New(lumberjack, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      false,
	})

	switch c.LogMode {
	case "silent", "Silent":
		conf.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		conf.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		conf.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		conf.Logger = _default.LogMode(logger.Info)
	default:
		conf.Logger = _default.LogMode(logger.Info)
	}
	return conf
}
