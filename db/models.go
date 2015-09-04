package db

import (
	"encoding/xml"
)


/*
Tables are:

    regions
    underground_regions
    sites
    world_constructions
    artifacts
    historical_figures
    entity_populations
    entities
    historical_events
    historical_event_collections
    historical_eras
*/

type Region struct {
	Id int `xml:"id" db:"id"`
	World int `db:"world_id"`
	Name string `xml:"name" db:"name"`
	Type string `xml:"type" db:"type"`
}

type UndergroundRegion struct {
	Id int `xml:"id" db:"id"`
	World int `db:"world_id"`
	Type string `xml:"type"`
	Depth int `xml:"depth"`
}


type Site struct {
	Id int `xml:"id" db:"id"`
	World int `db:"world_id"`
	Name string `xml:"name" db:"name"`
	Coords string `xml:"coords"`
	X_coord int `db:"x_coord"`
	Y_coord int `db:"y_coord"`
}

type Artifact struct {
	Id int `xml:"id" db:"id"`
	World int `db:"world_id"`
	Name string `xml:"name" db:"name"`
	Item string `xml:"item" db:"item"`
}

type RegionList struct {
	Regions []Region `xml:"region"`
}

type UndergroundRegionList struct {
	UndergroundRegions []UndergroundRegion `xml:"underground_region"`
}

type World struct {
	XMLName xml.Name `xml:"df_world"`
	Id int `db:"id"`
	Name string `db:"name"`
	RegionList RegionList `xml:"regions"`
	UndergroundRegionList UndergroundRegionList `xml:"underground_regions"`
	SiteList []Site `xml:"sites"`
	ArtifactList []Artifact `xml:"artifacts"`
}
