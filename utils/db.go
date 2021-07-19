package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // this for sql
)

// NewDbConnection is function to create New DB Connection
func NewDbConnection(host, port, username, password, databaseName string) (*sql.DB, error) {

	strConn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		databaseName,
	)

	db, err := sql.Open("mysql", strConn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
