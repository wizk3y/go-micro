package redis

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	DefaultMaxOpenConn  = 100
	DefaultMaxIdleConn  = 10
	DefaultConnLifetime = 300
)

type Config struct {
	Host         string
	Port         int
	Username     string
	Password     string
	Database     int
	MaxOpenConn  int
	MaxIdleConn  int
	ConnLifetime int
}

// GetDefaultConfig --
func GetDefaultConfig() *Config {
	return GetConfigWithPrefix("redis")
}

// GetConfigWithPrefix --
func GetConfigWithPrefix(prefix string) *Config {
	maxOpenConn := viper.GetInt(fmt.Sprintf("%s.max_open_conn", prefix))
	if maxOpenConn <= 0 {
		maxOpenConn = DefaultMaxOpenConn
	}

	maxIdleConn := viper.GetInt(fmt.Sprintf("%s.max_idle_conn", prefix))
	if maxIdleConn <= 0 {
		maxIdleConn = DefaultMaxIdleConn
	}

	connLifeTime := viper.GetInt(fmt.Sprintf("%s.conn_life_time", prefix))
	if connLifeTime <= 0 {
		connLifeTime = DefaultConnLifetime
	}

	return &Config{
		Host:         viper.GetString(fmt.Sprintf("%s.host", prefix)),
		Port:         viper.GetInt(fmt.Sprintf("%s.port", prefix)),
		Username:     viper.GetString(fmt.Sprintf("%s.username", prefix)),
		Password:     viper.GetString(fmt.Sprintf("%s.password", prefix)),
		Database:     viper.GetInt(fmt.Sprintf("%s.database", prefix)),
		MaxOpenConn:  maxOpenConn,
		MaxIdleConn:  maxIdleConn,
		ConnLifetime: connLifeTime,
	}
}

// BuildDataSource --
func (c *Config) BuildDataSource() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// GetMaxOpenConn --
func (c *Config) GetMaxOpenConn() int {
	return c.MaxOpenConn
}

// GetMaxIdleConn --
func (c *Config) GetMaxIdleConn() int {
	return c.MaxIdleConn
}

// GetConnMaxLifetime --
func (c *Config) GetConnMaxLifetime() time.Duration {
	return time.Duration(c.ConnLifetime) * time.Second
}
