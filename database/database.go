package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	migrate "github.com/rubenv/sql-migrate"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db != nil {
		return db
	}
	//Conexão com o mysql do container
	db, err := sql.Open("mysql", "hostgator:hostgator123@tcp(172.18.0.2:3306)/teste?parseTime=true")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//Carrega o arquivo .sql na pasta especificada e executa as queries (Migrations)
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	//Executa os migraions carregados
	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)

	return db
}
