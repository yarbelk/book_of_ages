-- +migrate Up
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
  type TEXT,
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

-- +migrate Down

DROP TABLE IF EXISTS worlds;
DROP TABLE IF EXISTS regions;
DROP TABLE IF EXISTS underground_regions;
DROP TABLE IF EXISTS sites;
DROP TABLE IF EXISTS artifacts;
