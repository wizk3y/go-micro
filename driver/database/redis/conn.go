package redis

import (
	goredis "github.com/redis/go-redis/v9"
)

// NewRedisConnection --
func NewRedisConnection(conn *Config) (*goredis.Client, error) {
	if conn == nil {
		conn = GetDefaultConfig()
	}

	opt := goredis.Options{
		Addr:     conn.BuildDataSource(),
		Username: conn.Username,
		Password: conn.Password,
		DB:       conn.Database,
	}

	if conn.GetMaxOpenConn() > 0 {
		opt.PoolSize = conn.GetMaxOpenConn()
	}

	if conn.GetMaxIdleConn() > 0 && conn.GetMaxIdleConn() < conn.GetMaxOpenConn() {
		opt.MaxIdleConns = conn.GetMaxIdleConn()
	}

	opt.ConnMaxLifetime = conn.GetConnMaxLifetime()

	c := goredis.NewClient(&opt)

	return c, nil
}
