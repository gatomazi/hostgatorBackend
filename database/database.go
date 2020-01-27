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

// var DockerIP = flag.String("dockerIp", "localhost", "Use for set the docker mysql image IP")

var DockerIP = "localhost"

func GetDB() *sql.DB {
	// flag.Parse()
	var err error

	if db != nil {
		return db
	}

	//Conexão com o mysql do container
	db, err := sql.Open("mysql", "hostgator:hostgator123@tcp("+DockerIP+":3306)/hostgator?parseTime=true")

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
