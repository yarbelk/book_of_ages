-- +migrate Up
CREATE TABLE IF NOT EXISTS worlds (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
  translated_name TEXT
);

CREATE TABLE IF NOT EXISTS regions (
	id INTEGER PRIMARY KEY,
	world_id INTEGER,
	name TEXT,
	type TEXT,
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);

CREATE TABLE IF NOT EXISTS underground_regions (
	id INTEGER PRIMARY KEY,
	world_id INTEGER,
	name TEXT,
	type TEXT,
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);

CREATE TABLE IF NOT EXISTS sites (
	id INTEGER PRIMARY KEY,
	world_id INTEGER,
	name TEXT,
	x_coord INTEGER,
	y_coord INTEGER,
  type TEXT,
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);


CREATE TABLE IF NOT EXISTS artifacts (
	id INTEGER PRIMARY KEY,
	world_id INTEGER,
	name TEXT,
	item TEXT,
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);

CREATE TABLE IF NOT EXISTS historical_figures (
	id INTEGER PRIMARY KEY,
	world_id INTEGER,
	name TEXT,
	race TEXT,
	caste TEXT,
  appeared INTEGER,
  birth_year INTEGER,
  death_year INTEGER,
  associated_type TEXT,

	FOREIGN KEY(world_id) REFERENCES worlds(id)
);

-- +migrate Down

DROP TABLE IF EXISTS worlds;
DROP TABLE IF EXISTS regions;
DROP TABLE IF EXISTS underground_regions;
DROP TABLE IF EXISTS sites;
DROP TABLE IF EXISTS artifacts;
DROP TABLE IF EXISTS historical_figures;
