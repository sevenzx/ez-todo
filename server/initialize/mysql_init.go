package initialize

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sevenzx/eztodo/conf"
	"github.com/sevenzx/eztodo/global"
	"github.com/sevenzx/eztodo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// GormMysql 初始化Gorm的Mysql
func GormMysql() *gorm.DB {
	m := conf.Config.Mysql
	if m.Dbname == "" {
		return nil
	}

	// [参考](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), getGormMysqlConfig())
	if err != nil {
		hlog.Error("gorm.Open err", err)
		return nil
	} else {
		// 设置数据库实例选项，指定表的存储引擎
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
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
func getGormMysqlConfig() *gorm.Config {
	m := conf.Config.Mysql
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.TablePrefix,
			SingularTable: m.SingularTable,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	switch m.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
