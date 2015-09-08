package db

import (
	"encoding/xml"
	"errors"
	"fmt"
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

type Modeler interface {
	Decode(decoder xml.Decoder) error
	Save()
	Value() interface{}
}


func (region *Region) Decode(decoder *xml.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return err
	}
	switch startElement := token.(type) {
		case xml.EndElement:
			if startElement.Name.Local == "regions" {
				return errors.New("end of stream")
			}
		case xml.StartElement:
			if startElement.Name.Local == "region" {
				decoder.DecodeElement(region, &startElement)
				fmt.Printf("%v \n", region)
				return nil
			} else {
			return errors.New("Wrong Tag Type")
			}
		default:
			return errors.New("Wrong Tag Type")
	 }
	return nil
}

func (region *Region) Save() (err error) {
	return
}

func (region *Region) Value() (interface {}) {
	return region
}
