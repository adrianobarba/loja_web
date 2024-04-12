package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
