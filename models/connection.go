package models

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
)

func connect() (*genmai.DB, error) {
	return genmai.New(&genmai.SQLite3Dialect{}, "./development.db")
}
