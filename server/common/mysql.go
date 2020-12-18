package common

import (
	"database/sql"
	"fmt"
)

func OpenDatabase(config Mysql) *sql.DB {
	logger := GetLogger()
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.Username, config.Password, config.Host, config.Port, config.Name, config.Option)
	conn, err := sql.Open("mysql", dbUrl)

	if err != nil {
		logger.Panicf("failed to connect database - %s", dbUrl)
	}

	if config.MaxIdle > 0 {
		conn.SetMaxIdleConns(config.MaxIdle)
	} else {
		conn.SetMaxIdleConns(10)
	}

	if config.MaxActive > 0 {
		conn.SetMaxOpenConns(config.MaxActive)
	} else {
		conn.SetMaxOpenConns(50)
	}

	err = conn.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	return conn
}
