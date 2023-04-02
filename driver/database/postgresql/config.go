package postgresql

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	DefaultMaxOpenConn  = 100
	DefaultMaxIdleConn  = 10
	DefaultConnLifetime = 600
)

type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Others       string
	MaxOpenConn  int
	MaxIdleConn  int
	ConnLifetime int
}

// GetDefaultConfig --
func GetDefaultConfig() *Config {
	return GetConfigWithPrefix("postgre")
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
		User:         viper.GetString(fmt.Sprintf("%s.user", prefix)),
		Password:     viper.GetString(fmt.Sprintf("%s.password", prefix)),
		DatabaseName: viper.GetString(fmt.Sprintf("%s.database_name", prefix)),
		MaxOpenConn:  maxOpenConn,
		MaxIdleConn:  maxIdleConn,
		ConnLifetime: connLifeTime,
	}
}

// BuildDataSource --
func (c *Config) BuildDataSource() string {
	ds := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DatabaseName)
	if c.Others != "" {
		ds = fmt.Sprintf("%s&%s", ds, c.Others)
	}

	return ds
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
