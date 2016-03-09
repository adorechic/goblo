package models

import (
	"github.com/naoina/genmai"
	_ "github.com/mattn/go-sqlite3"
)

func connect() (*genmai.DB, error) {
	return genmai.New(&genmai.SQLite3Dialect{}, "./development.db")
}
