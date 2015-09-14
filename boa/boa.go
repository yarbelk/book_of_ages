package boa


import (
	"os/user"
	"os"
	"github.com/yarbelk/book_of_ages/db"
	"github.com/yarbelk/book_of_ages/parser"
	"io"
	"log"
)

var ProgramDataPath string

func SetupProgramDirs() {
	user_data, _ := user.Current()

	ProgramDataPath = user_data.HomeDir+ "/.boa"

	os.Mkdir(ProgramDataPath, os.ModeDir | 0700 )
}

func IngestData(input_xml, input_world io.Reader) (err error) {
	var world *db.World
	world = &db.World{}
	parser.ParseWorldHistory(world, input_world)

	dbmap := db.InitDb()

	err = dbmap.Insert(world)
	if err != nil {
		log.Println("Failed to Ingest World, ", err.Error())
	}
	return
}
