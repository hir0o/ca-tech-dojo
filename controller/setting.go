package controller

import "database/sql"

type Connect struct {
	DB *sql.DB
}
