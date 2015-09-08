package boa


import (
	"os/user"
	"os"
)

var ProgramDataPath string

func SetupProgramDirs() {
	
	user_data, _ := user.Current()

	ProgramDataPath = user_data.HomeDir+ "/.boa"


	os.Mkdir(ProgramDataPath, os.ModeDir | 0700 )
}
