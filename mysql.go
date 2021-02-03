package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type mysqlLog struct{}

func (l mysqlLog) Print(v ...interface{}) {
	Logger.Info(v)
}

type MysqlConfig struct {
	Uri         string        `json:"uri"` // etc. user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
	Log         bool          `json:"log"`
	MaxIdle     int           `json:"max_idle" mapstructure:"max_idle"`         // current default: 2
	MaxActive   int           `json:"max_active" mapstructure:"max_active"`     // larger than maxIdle
	IdleTimeout time.Duration `json:"idle_timeout" mapstructure:"idle_timeout"` // suggest less than 8 * 3600
}

func NewMysql(cfg *MysqlConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", cfg.Uri)
	if err != nil {
		Logger.Errorf("NewMysql error, %v", err)
		return
	}
	db.DB().SetMaxIdleConns(cfg.MaxIdle)
	db.DB().SetMaxOpenConns(cfg.MaxActive)
	if cfg.IdleTimeout > 0 {
		db.DB().SetConnMaxIdleTime(cfg.IdleTimeout)
	}
	//db.DB().SetConnMaxLifetime(time.Duration(cfg.MaxLifeTime) / time.Second)
	db.SetLogger(mysqlLog{})
	db.LogMode(cfg.Log)
	return
}
