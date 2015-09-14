-- +migrate Up
CREATE TABLE IF NOT EXISTS worlds (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
  translated_name TEXT
);

CREATE TABLE IF NOT EXISTS regions (
	id INTEGER,
	world_id INTEGER,
	name TEXT,
	type TEXT,
	PRIMARY KEY(id, world_id),
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);

CREATE TABLE IF NOT EXISTS underground_regions (
	id INTEGER,
	world_id INTEGER,
	name TEXT,
	type TEXT,
	PRIMARY KEY(id, world_id),
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);

CREATE TABLE IF NOT EXISTS sites (
	id INTEGER,
	world_id INTEGER,
	name TEXT,
	x_coord INTEGER,
	y_coord INTEGER,
  type TEXT,
	PRIMARY KEY(id, world_id),
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);


CREATE TABLE IF NOT EXISTS artifacts (
	id INTEGER,
	world_id INTEGER,
	name TEXT,
	item TEXT,
	PRIMARY KEY(id, world_id),
	FOREIGN KEY(world_id) REFERENCES worlds(id)
);

-- +migrate Down

DROP TABLE IF EXISTS worlds;
DROP TABLE IF EXISTS regions;
DROP TABLE IF EXISTS underground_regions;
DROP TABLE IF EXISTS sites;
DROP TABLE IF EXISTS artifacts;
