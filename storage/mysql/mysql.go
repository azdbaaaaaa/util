package mysql

import (
	"github.com/azdbaaaaaa/util/log"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

const (
	defaultUri         = "127.0.0.1:3306"
	defaultMaxIdle     = 2
	defaultMaxActive   = 10
	defaultIdleTimeout = 60 // max idle time for each conn in seconds
)

type Config struct {
	// etc. user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	Uri         string          `json:"uri"`
	LogLevel    logger.LogLevel `json:"log_level" mapstructure:"log_level"`
	MaxIdle     int             `json:"max_idle" mapstructure:"max_idle"`         // current default: 2
	MaxActive   int             `json:"max_active" mapstructure:"max_active"`     // larger than maxIdle
	IdleTimeout time.Duration   `json:"idle_timeout" mapstructure:"idle_timeout"` // suggest less than 8 * 3600
}

func New(cfg *Config) (db *gorm.DB, err error) {
	var logLevel logger.LogLevel
	if cfg.LogLevel > 0 {
		logLevel = cfg.LogLevel
	} else {
		logLevel = logger.Silent
	}
	db, err = gorm.Open(mysql.Open(cfg.Uri), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Errorf("failed to new mysql, %v", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("failed to get mysql DB, %v", err)
		return
	}
	if cfg.MaxActive > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxActive)
	} else {
		sqlDB.SetMaxOpenConns(defaultMaxActive)
	}
	if cfg.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdle)
	} else {
		sqlDB.SetMaxIdleConns(defaultMaxIdle)
	}
	if cfg.IdleTimeout > 0 {
		sqlDB.SetConnMaxIdleTime(cfg.IdleTimeout * time.Second)
	} else {
		sqlDB.SetConnMaxIdleTime(defaultIdleTimeout * time.Second)
	}
	return
}
