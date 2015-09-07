package parser

import (
	"encoding/xml"
	. "github.com/yarbelk/book_of_ages/db"
	"io"
	"fmt"
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
	for {
		token, _ := decoder.Token()
		switch startElement := token.(type) {
			case xml.EndElement:
				if startElement.Name.Local == "regions" {
					return
				}
			case xml.StartElement:
				if startElement.Name.Local == "region" {
					r := Region{}
					decoder.DecodeElement(&r, &startElement)
					fmt.Printf("%v \n", r)
				}
			}
	}
}

func parseUndergroundRegion(decoder *xml.Decoder) {
	for {
		token, _ := decoder.Token()
		switch startElement := token.(type) {
			case xml.EndElement:
				if startElement.Name.Local == "underground_regions" {
					return
				}
			case xml.StartElement:
				if startElement.Name.Local == "underground_region" {
					ur := UndergroundRegion{}
					decoder.DecodeElement(&ur, &startElement)
					fmt.Printf("%v \n", ur)
				}
			}
	}
}
