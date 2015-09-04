package main

import (
	"github.com/yarbelk/book_of_ages/boa"
	"github.com/yarbelk/book_of_ages/db"
)


func main() {
	boa.SetupProgramDirs()
	db.CreateTables()
}
