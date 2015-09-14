package main

import (
	"github.com/yarbelk/book_of_ages/boa"
	"github.com/yarbelk/book_of_ages/db"
	"fmt"
	"database/sql"
)

func init() {
	var err error
	db.Db, err = sql.Open("sqlite3", db.Database)
	defer func() {
		db.Db.Close()
	}()

	if err != nil {
		panic("Failed to open DB")
	}

	boa.SetupProgramDirs()
	db.Database = boa.ProgramDataPath + "/boa.db"
	db.RunMigrations()
}

func main() {
	fmt.Println("BOA\n")
}
