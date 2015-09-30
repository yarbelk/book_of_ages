package db

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"strconv"
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
	Id int `xml:"id" db:"id, primarykey"`
	World int `db:"world_id"`
	Name string `xml:"name" db:"name"`
	Type string `xml:"type" db:"type"`
}

type UndergroundRegion struct {
	Id int `xml:"id" db:"id,primarykey"`
	World int `db:"world_id"`
	Type string `xml:"type"`
	Depth int `xml:"depth"`
}


type Site struct {
	Id int `xml:"id" db:"id, primarykey"`
	World int `db:"world_id"`
	Name string `xml:"name" db:"name"`
	Coords string `xml:"coords"`
	X_coord int `db:"x_coord"`
	Y_coord int `db:"y_coord"`
	Type string `xml:"type"`
}

type Artifact struct {
	Id int `xml:"id" db:"id, primarykey"`
	World int `db:"world_id"`
	Name string `xml:"name" db:"name"`
	Item string `xml:"item" db:"item"`
}

type World struct {
	Id int `db:"id"`
	Name string `db:"name"`
	TranslatedName string `db:"translated_name"`
	Regions []Region `db:"-"`
	UndergroundRegions []UndergroundRegion `db:"-"`
	SiteList []Site `xml:"sites" db:"-"`
	ArtifactList []Artifact `xml:"artifacts" db:"-"`
}

type HistoricalFigure struct {
	Id int `xml:"id" db:"id, primarykey"`
	Name string `xml:"name" db:"name"`
	Race string `xml:"race" db:"race"`
	Cast string `xml:"cast" db:"cast"`
	Appeared int `xml:"appeared" db:"appeared"`
	BirthYear int `xml:"birth_year" db:"birth_year"`
	DeathYear int `xml:"death_year" db:"death_year"`
	AssociatedType string `xml:"associated_type" db:"associated_type"`
}

type Modeler interface {
	Decode(decoder *xml.Decoder) error
	Value() interface{}
}

func decodeXml(decoder *xml.Decoder, model Modeler, parent_tag, tag string) error {
	fmt.Printf("%v\n", parent_tag)
	token, err := decoder.Token()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return err
	}
	fmt.Printf("got token: %v\n", token)
	switch startElement := token.(type) {
		case xml.EndElement:
			if startElement.Name.Local == parent_tag {
				return errors.New("end of stream")
			}
		case xml.StartElement:
			if startElement.Name.Local == tag {
			    fmt.Printf("%v\n", token.(xml.StartElement).Copy().Name)
				err = decoder.DecodeElement(model, &startElement)
				if err != nil {
					fmt.Printf("Decode failed: %v\n", err.Error())
					return err
				}
				return nil
			} else {
			return errors.New("Wrong Tag Type")
			}
		default:
			return errors.New("Wrong Tag Type")
	 }
	return nil
}


func (region *Region) Decode(decoder *xml.Decoder) error {
	return decodeXml(decoder, region, "regions", "region")
}

func (region *Region) Value() (interface {}) {
	return region
}

func (underground_region *UndergroundRegion) Decode(decoder *xml.Decoder) error {
	return decodeXml(decoder, underground_region, "underground_regions", "underground_region")
}

func (underground_region *UndergroundRegion) Value() (interface {}) {
	return underground_region
}

func (site *Site) Decode(decoder *xml.Decoder) error {
	err := decodeXml(decoder, site, "sites", "site")
	if err != nil {
		return err
	}

	coords := strings.Split(site.Coords, ",")
	site.X_coord, _ = strconv.Atoi(coords[0])
	site.Y_coord, _ = strconv.Atoi(coords[1])
	return nil
}

func (site *Site) Value() (interface {}) {
	return site
}

func (artifact *Artifact) Decode(decoder *xml.Decoder) error {
	return decodeXml(decoder, artifact, "artifacts", "artifact")
}

func (artifact *Artifact) Value() (interface {}) {
	return artifact
}

func (historical_figure *HistoricalFigure) Decode(decoder *xml.Decoder) error {
	return decodeXml(decoder, historical_figure, "historical_figures", "historical_figure")
}

func (historical_figure *HistoricalFigure) Value() (interface {}) {
	return historical_figure
}
