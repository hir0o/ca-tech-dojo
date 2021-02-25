package controller

import "database/sql"

type ConnectDB struct {
	DB *sql.DB
}
