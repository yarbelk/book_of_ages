package parser

import (
	"encoding/xml"
	. "github.com/yarbelk/book_of_ages/db"
	"io"
	"bufio"
	"strings"
)


func Parse(world *World, inputXML io.Reader) {
	decoder := xml.NewDecoder(inputXML)
	for {
		token, err := decoder.Token()

		if (err == io.EOF) {
			break
		}

		switch startElement := token.(type) {
		case xml.StartElement: 
			if startElement.Name.Local == "regions" {
				parseRegion(decoder)
			}
			if startElement.Name.Local == "underground_regions" {
				parseUndergroundRegion(decoder)
			}
		}
	}
}


func parseRegion(decoder *xml.Decoder) {
	var err error
	var region *Region
	for {
		region = &Region{}
		err = region.Decode(decoder)
		if err != nil {
			if err.Error() == "Wrong Tag Type" {
				continue
			}
			return
		}
	}
}

func parseUndergroundRegion(decoder *xml.Decoder) {
	var err error
	var underground_region *UndergroundRegion
	for {
		underground_region = &UndergroundRegion{}
		err = underground_region.Decode(decoder)
		if err != nil {
			if err.Error() == "Wrong Tag Type" {
				continue
			}
			return
		}
	}
}

// I hate unstructured text.
func ParseWorldHistory(world *World, inputHistory io.Reader) (err error) {
	buf := bufio.NewReader(inputHistory)
	return parserWorldName(world, buf)
}


func parserWorldName(world *World, buf *bufio.Reader) (err error) {
	var name, translated_name string
	name, err = buf.ReadString('\n')
	if err != nil { return }
	translated_name, err = buf.ReadString('\n')
	if err != nil { return }
	world.Name = strings.TrimSpace(name)
	world.TranslatedName = strings.TrimSpace(translated_name)
	return
	}
