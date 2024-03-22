package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hackaton/internal/config"
)

func GetDBConnection(cfg *config.Config) (*sql.DB, error) {
	strConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", strConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
