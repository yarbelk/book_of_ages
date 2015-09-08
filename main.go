package main

import (
	"github.com/yarbelk/book_of_ages/boa"
	"github.com/yarbelk/book_of_ages/db"
	"fmt"
)

func init() {
	boa.SetupProgramDirs()
	db.Database = boa.ProgramDataPath + "/boa.db"
	db.RunMigrations()
}

func main() {
	fmt.Println("BOA\n")
}
