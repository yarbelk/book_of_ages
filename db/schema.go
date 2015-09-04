package db


import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
	"log"
)


var initial_schema = `
CREATE TABLE IF NOT EXISTS worlds (
	id INT PRIMARY KEY,
	name TEXT
);

CREATE TABLE IF NOT EXISTS regions (
	id INT,
	world_id INT,
	name TEXT,
	type TEXT,
	PRIMARY KEY(id, world_id),
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);

CREATE TABLE IF NOT EXISTS underground_regions (
	id INT,
	world_id INT,
	name TEXT,
	type TEXT,
	PRIMARY KEY(id, world_id),
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);

CREATE TABLE IF NOT EXISTS sites (
	id INT,
	world_id INT,
	name TEXT,
	x_coord INT,
	y_coord INT,
	PRIMARY KEY(id, world_id),
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);


CREATE TABLE IF NOT EXISTS artifacts (
	id INT,
	world_id INT,
	name TEXT,
	item TEXT,
	PRIMARY KEY(id, world_id),
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);
`

func CreateTables() {
	connection, err := sqlx.Connect("sqlite3", "book_of_ages.sqlite3")
    if err != nil {
        log.Fatalln(err)
    }

	connection.MustExec(initial_schema)
}
