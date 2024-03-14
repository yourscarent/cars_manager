package postgres

import (
	_ "database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yourscarent/cars_manager/internal/config"
)

func MustConnect(cfg config.DB) *sqlx.DB {
	dts := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	conn, err := sqlx.Open("postgres", dts)
	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("failed to ping db: " + err.Error())
	}

	return conn
}
