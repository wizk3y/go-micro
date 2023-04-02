package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/wizk3y/go-micro/logger"
)

// NewPostgreConnection --
func NewPostgreConnection(conn *Config) (*sql.DB, error) {
	if conn == nil {
		conn = GetDefaultConfig()
	}

	dataSource := conn.BuildDataSource()

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		logger.Errorf("[psql] could not connect database, details: %v", err)
		return nil, err
	}

	if conn.GetMaxOpenConn() > 0 {
		db.SetMaxOpenConns(conn.GetMaxOpenConn())
	}

	if conn.GetMaxIdleConn() > 0 && conn.GetMaxIdleConn() < conn.GetMaxOpenConn() {
		db.SetMaxIdleConns(conn.GetMaxIdleConn())
	}

	db.SetConnMaxLifetime(conn.GetConnMaxLifetime())

	err = db.Ping()
	if err != nil {
		logger.Errorf("[psql] could not ping to database, details: %v", err)
		return nil, err
	}

	return db, nil
}
