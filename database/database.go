package database

import (
	"database/sql"
	"go-project/config"

	_ "github.com/go-sql-driver/mysql" //Driver MySQL
)

// Connect Realiza conexão com banco
func Connect() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.DBConnection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil

}
