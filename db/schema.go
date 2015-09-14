package db


import (
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
	"gopkg.in/gorp.v1"
	"database/sql"
	"fmt"
)

var Database string = "db/boa.db"
var Migrations string = "db/migrations"
var Db *sql.DB
var dbmap *gorp.DbMap

var Env string = ""

func RunMigrations() {
	if Env == "" {
		Env = "development"
	}
	migrations := &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations",
	}

	n, err := migrate.Exec(Db, "sqlite3", migrations, migrate.Up)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Migrated %d\n", n)
}


func InitDb() *gorp.DbMap {
	if dbmap != nil { return dbmap }
	var err error
	Db, err = sql.Open("sqlite3", Database)
	if err != nil {
		fmt.Println(err.Error())
	}

	RunMigrations()
	// construct a gorp DbMap
    dbmap = &gorp.DbMap{Db: Db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTableWithName(World{}, "worlds").SetKeys(true, "Id")
	dbmap.AddTableWithName(Region{}, "regions").SetKeys(true, "Id")
	dbmap.AddTableWithName(UndergroundRegion{}, "underground_regions").SetKeys(true, "Id")

	return dbmap
}

func CloseDb() {
	Db.Close()
	dbmap = nil
}
