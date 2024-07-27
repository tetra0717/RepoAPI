package database

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() (*sql.DB, error) {
	count := 15
	for count > 1 {
		db, err := sql.Open("mysql","user:password@tcp(DB)/api")
		if err != nil {
			time.Sleep(time.Second * 2)
			count--
			continue
		}

		
		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("failed to ping database: %w", err)
		}
		return db, nil
	}
	
	return nil, fmt.Errorf("failed to connect to database")
}
