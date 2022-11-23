package reqiestmethods

import (
	"jun2/structs"

	"github.com/jmoiron/sqlx"
)

// Config docker postgres:14.5-alpine
var Config = structs.Config{
	Host:     "127.0.0.1",
	Port:     5432,
	User:     "test_db",
	Password: "test_db",
	DBName:   "test_db",
}

var DB *sqlx.DB
