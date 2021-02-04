package util

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

const (
	defaultMaxIdle     = 2
	defaultMaxActive   = 10
	defaultIdleTimeout = 60 // max idle time for each conn in seconds
)

type MysqlConfig struct {
	// etc. user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	Uri         string          `json:"uri"`
	LogLevel    logger.LogLevel `json:"log_level" mapstructure:"log_level"`
	MaxIdle     int             `json:"max_idle" mapstructure:"max_idle"`         // current default: 2
	MaxActive   int             `json:"max_active" mapstructure:"max_active"`     // larger than maxIdle
	IdleTimeout time.Duration   `json:"idle_timeout" mapstructure:"idle_timeout"` // suggest less than 8 * 3600
}

func DefaultConfig() *MysqlConfig {
	return &MysqlConfig{
		MaxIdle:     defaultMaxIdle,
		MaxActive:   defaultMaxActive,
		IdleTimeout: defaultIdleTimeout * time.Second,
	}
}

func NewMysql(cfg *MysqlConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(cfg.Uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		Logger.Log.Errorf("failed to new mysql, %v", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		Logger.Log.Errorf("failed to get mysql DB, %v", err)
		return
	}
	if cfg.MaxActive > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxActive)
	}
	if cfg.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdle)
	}
	if cfg.IdleTimeout > 0 {
		sqlDB.SetConnMaxIdleTime(cfg.IdleTimeout * time.Second)
	}
	return
}
