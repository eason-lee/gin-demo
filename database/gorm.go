package db

import (
	"gin-demo/global"
	"gin-demo/model"
	"gin-demo/server"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Gorm ...
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB
func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// MysqlTables ...
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.StatisticsMeta{},
		model.StatisticsData{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}

// GormMysql ...
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB
func GormMysql() *gorm.DB {
	m := global.CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}


//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config
func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.CONFIG.Mysql.LogZap {
	case "silent", "Silent":
		config.Logger = server.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = server.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = server.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = server.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = server.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = server.Default.LogMode(logger.Info)
			break
		}
		config.Logger = server.Default.LogMode(logger.Silent)
	}
	return config
}
