package db


import (
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
	"database/sql"
	"fmt"
)

var Database string = "db/boa.db"
var Migrations string = "db/migrations"

var Env string = ""

func RunMigrations() {
	if Env == "" {
		Env = "development"
	}
	migrations := migrate.FileMigrationSource{
		Dir: Migrations,
	}

	db, err := sql.Open("sqlite3", Database)
	defer func() {
		db.Close()
	}()

	if err != nil {
		panic("Failed to open DB")
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)

	if err != nil {
		panic("Failed to open DB")
	}

	fmt.Printf("Migrated %d\n", n)
}
